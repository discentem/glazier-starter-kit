#$MountMyWindowsImage | Dismount-MyWindowsImage

# $scriptPath is helpful for reference files relative to wherever this script is running
$scriptPath = split-path -parent $MyInvocation.MyCommand.Definition

# Read a value for config_root_path from $scriptPath\autobuild.json. Note tools\autobuild.json is in the .gitignore file
$abjson = "$scriptPath\autobuild.json"
$json = ""
$config_root_path = ""
if ((Test-Path $abjson) -eq $True) {
    $json = Get-Content $abjson | Out-String | ConvertFrom-Json
    $config_root_path = $json.config_root_path
}

# Set some variables for python metadata
$pyVersion = "3.9.5"
$pythonSavePath = "~\Downloads\python-$pyVersion-amd64.exe"
$pythonInstallHash = "53a354a15baed952ea9519a7f4d87c3f"
$pyEXEUrl = "https://www.python.org/ftp/python/$pyVersion/python-$pyVersion-amd64.exe"

# Install Chocoately. Borrowed from https://tseknet.com/blog/chocolatey
Set-ExecutionPolicy -ExecutionPolicy RemoteSigned -Force
iwr https://chocolatey.org/install.ps1 -UseBasicParsing | iex
# Install Windows ADK
choco install windows-adk -y
choco install windows-adk-winpe -y

# Install OSD Powershell Module. This isn't idempotent yet.
(Install-PackageProvider -Name NuGet -MinimumVersion 2.8.5.201 -Force) -and (Install-Module OSD -Force)
# Import the module. This isn't idempotent yet
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

# Mount our WIM. Borrowed from https://github.com/OSDeploy/OSD/blob/master/Public/OSDCloud/Edit-OSDCloud.winpe.ps1
$WorkspacePath = Get-OSDCloud.workspace -ErrorAction Stop
$MountMyWindowsImage = Mount-MyWindowsImage -ImagePath "$WorkspacePath\Media\Sources\boot.wim"
$MountPath = $MountMyWindowsImage.Path

# Variables we can use to install Python inside of the mounted WIM
$pyTargetDir = "$MountPath\Python"
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
$env:Path = [System.Environment]::GetEnvironmentVariable("Path","Machine") + ";" + [System.Environment]::GetEnvironmentVariable("Path","User") 
# Install Glazier requirements
& $pyEXE -m pip install -r C:\glazier\requirements.txt
Write-Host "Running '$pyEXE -m pip install pywin32'"
# Install an additional Glazier dependency that isn't included in requirements.txt.
& $pyEXE -m pip install pywin32


# Recursively copy Glazier source code into WIM
Write-Host "Copying Glazier source code inside the image"
robocopy "C:\glazier\" "$MountPath\glazier\" /E

Write-Host "Copying autobuild.ps1 to WIM"
# Copy autobuild.ps1 into WIM
robocopy "$scriptPath\" "$MountPath\Windows\System32\" autobuild.ps1

# Write out Startnet.cmd, which will call autobuild.ps1
$Startnet = @"
wpeinit
start PowerShell -Nol -W Mi
powershell -NoProfile -NoLogo -Command Set-ExecutionPolicy -ExecutionPolicy Bypass -Force
powershell -NoProfile -NoLogo -WindowStyle Maximized -NoExit -File "X:\Windows\System32\autobuild.ps1" -config_root_path {0}
"@ -f $config_root_path
Write-Host "Writing Startnet.cmd"
$Startnet | Out-File -FilePath "$MountPath\Windows\System32\Startnet.cmd" -Force -Encoding ascii
Write-Host "Saving WIM"
# Dismount and save WIM
$MountMyWindowsImage | Dismount-MyWindowsImage -Save
Write-Host "Creating ISO"
New-OSDCloud.iso