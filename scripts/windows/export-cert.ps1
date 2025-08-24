# Export signing certificate into base64 encoded contents
$cert= Get-ChildItem -Path Cert:\CurrentUser\My | where-object  {$_.Subject -like "*fabric*"}
if ($null -eq $cert) {
    Write-Error "Signing certificate '*fabric*' not found in CurrentUser\My store."
    exit 1
}
$password = Read-Host -AsSecureString -Prompt "Enter a password for the PFX file"
$pfxPath = [System.IO.Path]::GetTempFileName() + ".pfx"
Export-PfxCertificate -Cert $cert -FilePath $pfxPath -Password $password
$pfxContent = [System.IO.File]::ReadAllBytes($pfxPath)
[System.Convert]::ToBase64String($pfxContent)
Remove-Item $pfxPath
Write-Output "`n`nAdd these secrets to github secrets as WINDOWS_CERT_BASE64 and WINDOWS_CERT_PASSWORD`n"
