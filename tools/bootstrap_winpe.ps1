Param(
    [string]$config_server = "",
    [string]$make_usb="",
    [string]$driver_path = "C:\drivers",
    [bool]$export_drivers = $false
)

# $scriptPath is the absolute path where this script is executing from
$scriptPath = split-path -parent $MyInvocation.MyCommand.Definition

if ((Test-Path "$scriptPath\autobuild.ps1") -eq $False) {
    Write-Host "$scriptPath\autobuild.ps1 not found"
    exit(1)
}

# Make sure glazier-resources is available to copy
if ((Test-Path "$scriptPath\..\glazier-resources\") -eq $False) {
    Write-Host "$scriptPath\..\glazier-resources\ not found"
    exit(1)
}

# if $scriptPath\autobuild.json exists, get a value for $config_root_path from it. Note tools\autobuild.json is in the .gitignore file.
$abjson = "$scriptPath\autobuild.json"
if ((Test-Path $abjson) -eq $True) {
    $json = Get-Content $abjson | Out-String | ConvertFrom-Json
    $config_server = $json.config_server
# get a value for $config_server from the CLI flags. If it's empty, exit 1 (bail)
} else {
    if ($config_server -eq "") {
        Write-Host "config_server cannot be an empty string"
        exit(1)
    }
}

function isURIWeb($address) {
	$uri = $address -as [System.URI]
	$uri.AbsoluteURI -ne $null -and $uri.Scheme -match '[http|https]'
}

# ensure $config_server is a valid url. This avoids avoid frustration, trust me.
if (!(isURIWeb($config_server))) {
    Write-Host "$config_server is not a valid url"
    exit(1)
}

if($export_drivers) {
    New-Item -ItemType Directory -Force -Path $driver_path	
    dism /online /export-driver /destination:$driver_path
}

# $pyVersion is the Python version that will be downloaded
$pyVersion = "3.9.5"
# $pythonSavePath is place where the Python installer will be downloaded on disk
$pythonSavePath = "~\Downloads\python-$pyVersion-amd64.exe"
# $pythonInstallHash is the hash used to verify the Python installer download
$pythonInstallHash = "53a354a15baed952ea9519a7f4d87c3f"
# $pyEXEUrl is the url where the Python installer will be obtained
$pyEXEUrl = "https://www.python.org/ftp/python/$pyVersion/python-$pyVersion-amd64.exe"

# Install Chocoately. Borrowed from https://tseknet.com/blog/chocolatey.
Set-ExecutionPolicy -ExecutionPolicy RemoteSigned -Force
iwr https://chocolatey.org/install.ps1 -UseBasicParsing | iex
# Install Windows ADK
choco install windows-adk -y
choco install windows-adk-winpe -y

# Install OSD Powershell Module. This isn't idempotent yet.
(Install-PackageProvider -Name NuGet -MinimumVersion 2.8.5.201 -Force) -and (Install-Module OSD -Force)
# Import the module. This isn't idempotent yet.
Import-Module OSD

# Create OSDCloud template if not created yet.
if ((Get-OSDCloud.template) -eq "C:\ProgramData\OSDCloud") { 
    Write-Host "'NEW-OSDCloud.template -WinRE -Verbose' already completed"
} else {
    Write-Host "Running 'NEW-OSDCloud.template -WinRE -Verbose'"
    New-OSDCloud.template -WinRE -Verbose
}

# Create OSDCloud workspace if not created yet.
if ((Get-OSDCloud.workspace) -eq "C:\OSDCloud") {
    Write-Host "'New-OSDCloud.workspace -WorkspacePath C:\OSDCloud' already completed"
} else {
    Write-Host "Running 'New-OSDCloud.workspace -WorkspacePath C:\OSDCloud'"
    New-OSDCloud.workspace -WorkspacePath C:\OSDCloud
}

# Download Python if needed
if ((Test-Path $pythonSavePath) -eq $False) {
    Write-Host "Downloading Python $pyVersion"
    curl $pyEXEUrl -UseBasicParsing -OutFile $pythonSavePath
} else {
    Write-Host "Python $pyVersion was already downloaded"
}

# Verify hash of Python installer
while ((Get-FileHash ($pythonSavePath) -Algorithm MD5).Hash -ne $pythonInstallHash) {
    Write-Host "Python hashes did not match. Removing and redownloading"
    rm $pythonSavePath
    curl $pyEXEUrl -UseBasicParsing -OutFile $pythonSavePath
}

# Edit-OSDCloud is needed to make StartWinRE-Wifi work
Edit-OSDCloud.winpe -DriverPath $driver_path

# Mount our WIM. Borrowed from https://github.com/OSDeploy/OSD/blob/master/Public/OSDCloud/Edit-OSDCloud.winpe.ps1
$WorkspacePath = Get-OSDCloud.workspace -ErrorAction Stop
$MountMyWindowsImage = Mount-MyWindowsImage -ImagePath "$WorkspacePath\Media\Sources\boot.wim"
$MountPath = $MountMyWindowsImage.Path

# $pyTargetDir is the directory where Python will get installed
$pyTargetDir = "$MountPath\Python"
# $pyEXE is shorthand for referring to python.exe. This variable will be used to install pip modules later.
$pyEXE = "$pyTargetDir\python.exe"

# This is not idempotent yet
Write-Host "Installing Python $pyVersion"
mkdir $MountPath\Python
# Install Python in the mounted WIM
& $pythonSavePath TargetDir=$pyTargetDir Include_launcher=0 /passive

# Install git
choco install git -y

# Clone Glazier repository.
Write-Host "Cloning and pulling Glazier github repo..."
& 'C:\Program Files\Git\bin\git.exe' clone https://github.com/google/glazier.git C:\glazier
& 'C:\Program Files\Git\bin\git.exe' -C C:\glazier\ pull
Write-Host "Install Glazier pip requirements"
# Reload the path: https://stackoverflow.com/questions/17794507/reload-the-path-in-powershell
# This is required so pip can install from git repositories
$env:Path = [System.Environment]::GetEnvironmentVariable("Path","Machine") + ";" + [System.Environment]::GetEnvironmentVariable("Path","User") 
# Install Glazier requirements
& $pyEXE -m pip install -r C:\glazier\requirements.txt
Write-Host "Running '$pyEXE -m pip install pywin32'"
# Install an additional Glazier dependency that isn't included in requirements.txt.
& $pyEXE -m pip install pywin32

Write-Host "Copying Glazier source code inside the image"
robocopy "C:\glazier\" "$MountPath\glazier\" /E /PURGE

Write-Host "Copying glazier-resources to WIM"
mkdir "$MountPath\glazier-resources"
robocopy "$scriptPath\..\glazier-resources\" "$MountPath\glazier-resources\" *.* /E /PURGE

Write-Host "Copying autobuild.ps1 to WIM"
# Copy autobuild.ps1 into WIM
robocopy "$scriptPath\" "$MountPath\Windows\System32\" autobuild.ps1 /PURGE

# Write out Startnet.cmd, which will run when WinPE boots. This will launch the CLI Wifi menu and then autobuild.ps1.
$Startnet = @"
wpeinit
start powershell
start /wait PowerShell -NoL -C Start-WinREWiFi
powershell -NoProfile -NoLogo -Command Set-ExecutionPolicy -ExecutionPolicy Bypass -Force
powershell -NoProfile -NoLogo -WindowStyle Maximized -NoExit -File "X:\Windows\System32\autobuild.ps1" -config_server {0}
"@ -f $config_server
Write-Host "Writing Startnet.cmd"
# Save our changes to Startnet.cmd
$Startnet | Out-File -FilePath "$MountPath\Windows\System32\Startnet.cmd" -Force -Encoding ascii
Write-Host "Saving WIM"
# Dismount and save WIM
#Save-WindowsImage -Path $MountPath
$MountMyWindowsImage | Dismount-MyWindowsImage -Save
Write-Host "Creating ISO"
New-OSDCloud.iso
if ($make_usb -eq 'true') {
    Write-Host "Creating USB"
    New-OSDCloud.usb
}
