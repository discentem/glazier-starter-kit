Param(
    [string]config_root_path="http://glazier"
)

if ($config_root_path -eq "") {
  $config_root_path = "http://glazier"
}

$Host.UI.RawUI.WindowTitle = 'Glazier'
$env:LOCALAPPDATA = 'X:\'
$env:PYTHONPATH = 'X:\glazier\src'
Write-Output 'Starting Glazier imaging process...'

# For a full list of Glazier flags, execute `python autobuild.py --helpfull`
$py_args = @(
  "X:\glazier\src",
  "--config_root_path=$config_root_path",
  '--resource_path=X:\glazier\resources',
  '--glazier_spec_os=windows10-stable',
  '--preserve_tasks=true'
)

& X:\Python\python.exe $py_args