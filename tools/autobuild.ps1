$Host.UI.RawUI.WindowTitle = 'Glazier'
$env:LOCALAPPDATA = 'X:\'
$env:PYTHONPATH = 'X:\glazier\src'
Write-Output 'Starting Glazier imaging process...'

# For a full list of Glazier flags, execute `python autobuild.py --helpfull`
$py_args = @(
  '--config_root_path=',
  '--resource_path=X:\glazier\resources',
  '--glazier_spec_os=windows10-stable',
  '--preserve_tasks=true'
)

& X:\Python\python.exe "$env:PYTHON\src\autobuild.py" $py_args