$pyVersion = "3.9.5"
$pythonSavePath = "~\Downloads\python-$pyVersion-amd64.exe"
$pythonInstallHash = "53a354a15baed952ea9519a7f4d87c3f"
$pyEXEUrl = "https://www.python.org/ftp/python/$pyVersion/python-$pyVersion-amd64.exe"
$pyTargetDir = "C:\OSDCloud\Autopilot\Python39\"
$pyEXE = "$pyTargetDir\python.exe"

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

# This is not idempotent yet
Write-Host "Installing Python $pyVersion"
& $pythonSavePath TargetDir= Include_launcher=0 /passive

choco install git -y

git clone https://github.com/google/glazier.git C:\OSDCloud\Autopilot\glazier
& $pyEXE -m pip install pywin32