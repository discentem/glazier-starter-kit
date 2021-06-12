#$MountMyWindowsImage | Dismount-MyWindowsImage

$scriptPath = split-path -parent $MyInvocation.MyCommand.Definition

$abjson = "$scriptPath\autobuild.json"
$json = ""
$config_root_path = ""
if ((Test-Path $abjson) -eq $True) {
    $json = Get-Content $abjson | Out-String | ConvertFrom-Json
    $config_root_path = $json.config_root_path
}

$pyVersion = "3.9.5"
$pythonSavePath = "~\Downloads\python-$pyVersion-amd64.exe"
$pythonInstallHash = "53a354a15baed952ea9519a7f4d87c3f"
$pyEXEUrl = "https://www.python.org/ftp/python/$pyVersion/python-$pyVersion-amd64.exe"

Set-ExecutionPolicy Bypass -Scope Process -Force; [System.Net.ServicePointManager]::SecurityProtocol = [System.Net.ServicePointManager]::SecurityProtocol -bor 3072; iex ((New-Object System.Net.WebClient).DownloadString('https://chocolatey.org/install.ps1')) 
choco install windows-adk -y
choco install windows-adk-winpe -y

# This isn't idempotent yet
(Install-PackageProvider -Name NuGet -MinimumVersion 2.8.5.201 -Force) -and (Install-Module OSD -Force)
# This isn't idempotent yet
Import-Module OSD

if ((Get-OSDCloud.template) -eq "C:\ProgramData\OSDCloud") { 
    Write-Host "'NEW-OSDCloud.template -WinRE -Verbose' already completed"
} else {
    Write-Host "Running 'NEW-OSDCloud.template -WinRE -Verbose'"
    New-OSDCloud.template -WinRE -Verbose
}

if ((Get-OSDCloud.workspace) -eq "C:\OSDCloud") {
    Write-Host "'New-OSDCloud.workspace -WorkspacePath C:\OSDCloud' already completed"
} else {
    Write-Host "Running 'New-OSDCloud.workspace -WorkspacePath C:\OSDCloud'"
    New-OSDCloud.workspace -WorkspacePath C:\OSDCloud
}

if ((Test-Path $pythonSavePath) -eq $False) {
    Write-Host "Downloading Python $pyVersion"
    curl $pyEXEUrl -UseBasicParsing -OutFile $pythonSavePath
} else {
    Write-Host "Python $pyVersion was already downloaded"
}

if ((Get-FileHash ($pythonSavePath) -Algorithm MD5).Hash -ne $pythonInstallHash) {
    Write-Host "Python hashes did not match. Removing and redownloading"
    rm $pythonSavePath
    curl $pyEXEUrl -UseBasicParsing -OutFile $pythonSavePath
}

# Borrowed from https://github.com/OSDeploy/OSD/blob/master/Public/OSDCloud/Edit-OSDCloud.winpe.ps1
$WorkspacePath = Get-OSDCloud.workspace -ErrorAction Stop
$MountMyWindowsImage = Mount-MyWindowsImage -ImagePath "$WorkspacePath\Media\Sources\boot.wim"
$MountPath = $MountMyWindowsImage.Path

$pyTargetDir = "$MountPath\Python"
$pyEXE = "$pyTargetDir\python.exe"

# This is not idempotent yet
Write-Host "Installing Python $pyVersion"
mkdir $MountPath\Python
& $pythonSavePath TargetDir=$pyTargetDir Include_launcher=0 /passive

choco install git -y

# Not idempotent yet
Write-Host "Cloning and pulling Glazier github repo..."
& 'C:\Program Files\Git\bin\git.exe' clone https://github.com/google/glazier.git C:\glazier
& 'C:\Program Files\Git\bin\git.exe' -C C:\glazier\ pull
Write-Host "Install Glazier pip requirements"
& $pyEXE -m pip install -r C:\glazier\requirements.txt
Write-Host "Running '$pyEXE -m pip install pywin32'"
& $pyEXE -m pip install pywin32



Write-Host "Copying Glazier source code inside the image"
robocopy "C:\glazier\" "$MountPath\glazier\" /E

Write-Host "Copying autobuild.ps1 to WIM"
robocopy "$scriptPath\" "$MountPath\Windows\System32\" autobuild.ps1
$Startnet = @"
wpeinit
start PowerShell -Nol -W Mi
powershell -NoProfile -NoLogo -Command Set-ExecutionPolicy -ExecutionPolicy Bypass -Force
powershell -NoProfile -NoLogo -WindowStyle Maximized -NoExit -File "X:\Windows\System32\autobuild.ps1" -config_root_path {0}
"@ -f $config_root_path
Write-Host "Writing Startnet.cmd"
$Startnet | Out-File -FilePath "$MountPath\Windows\System32\Startnet.cmd" -Force -Encoding ascii
Write-Host "Saving WIM"
$MountMyWindowsImage | Dismount-MyWindowsImage -Save
Write-Host "Creating ISO"
New-OSDCloud.iso