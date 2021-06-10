Set-ExecutionPolicy Bypass -Scope Process -Force; [System.Net.ServicePointManager]::SecurityProtocol = [System.Net.ServicePointManager]::SecurityProtocol -bor 3072; iex ((New-Object System.Net.WebClient).DownloadString('https://chocolatey.org/install.ps1')) 
choco install windows-adk -y
choco install windows-adk-winpe -y
(Install-PackageProvider -Name NuGet -MinimumVersion 2.8.5.201 -Force) -and (Install-Module OSD -Force)
Import-Module OSD

if ((Get-OSDCloud.template) -eq "C:\ProgramData\OSDCloud") { 
    Write-Host "'NEW-OSDCloud.template -WinRE -Verbose' already completed"
} else { 
    New-OSDCloud.template -WinRE -Verbose
}

if ((Get-OSDCloud.workspace) -eq "C:\OSDCloud") {
    Write-Host "'New-OSDCloud.workspace -WorkspacePath C:\OSDCloud' already completed"
} else {
    New-OSDCloud.workspace -WorkspacePath C:\OSDCloud
}

