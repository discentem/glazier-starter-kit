Param(
    [string]$config_server="http://glazier"
)

if ($config_server -eq "") {
  $config_server = "http://glazier"
}

$Host.UI.RawUI.WindowTitle = 'Glazier'
$env:LOCALAPPDATA = 'X:\'
$env:PYTHONPATH = 'X:\glazier\'
Write-Output 'Starting Glazier imaging process...'

# For a full list of Glazier flags, execute `python autobuild.py --helpfull`
$py_args = @(
  "X:\glazier\glazier\autobuild.py",
  "--config_server=$config_server",
  '--resource_path=X:\glazier-resources',
  '--glazier_spec_os=windows10-stable',
  '--preserve_tasks=true'
)

& X:\Python\python.exe $py_args