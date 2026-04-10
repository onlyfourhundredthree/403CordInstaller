$link = "https://github.com/403Cord/Installer/releases/latest/download/403CordInstallerCli.exe"

$outfile = "$env:TEMP\403CordInstallerCli.exe"

Write-Output "Downloading installer to $outfile"

Invoke-WebRequest -Uri "$link" -OutFile "$outfile"

Write-Output ""

Start-Process -Wait -NoNewWindow -FilePath "$outfile"

# Cleanup
Remove-Item -Force "$outfile"
