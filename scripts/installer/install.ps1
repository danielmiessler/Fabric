# Fabric Windows Installer Script
# Usage: iwr -useb https://raw.githubusercontent.com/danielmiessler/fabric/main/scripts/installer/install.ps1 | iex
# Usage with custom directory: $env:INSTALL_DIR="C:\tools"; iwr -useb https://raw.githubusercontent.com/danielmiessler/fabric/main/scripts/installer/install.ps1 | iex

param(
    [string]$InstallDir = $env:INSTALL_DIR
)

$ErrorActionPreference = "Stop"

# Colors for output (Windows Console colors)
$Colors = @{
    Red    = "Red"
    Green  = "Green"
    Yellow = "Yellow"
    Blue   = "Cyan"
    White  = "White"
}

# Print functions
function Write-Info {
    param([string]$Message)
    Write-Host "[INFO] $Message" -ForegroundColor $Colors.Blue
}

function Write-Success {
    param([string]$Message)
    Write-Host "[SUCCESS] $Message" -ForegroundColor $Colors.Green
}

function Write-Warning {
    param([string]$Message)
    Write-Host "[WARNING] $Message" -ForegroundColor $Colors.Yellow
}

function Write-Error {
    param([string]$Message)
    Write-Host "[ERROR] $Message" -ForegroundColor $Colors.Red
}

# Detect Windows architecture
function Get-Architecture {
    $arch = $env:PROCESSOR_ARCHITECTURE
    $archAMD64 = $env:PROCESSOR_ARCHITEW6432

    # Check for ARM64
    if ($arch -eq "ARM64") {
        return "arm64"
    }

    # Check for x86_64/AMD64
    if ($arch -eq "AMD64" -or $archAMD64 -eq "AMD64") {
        return "x86_64"
    }

    # Check for x86 (32-bit)
    if ($arch -eq "X86") {
        return "i386"
    }

    Write-Error "Unsupported architecture: $arch"
    Write-Error "This installer supports x86_64, i386, and arm64"
    exit 1
}

# Test if running with appropriate permissions for directory
function Test-WritePermission {
    param([string]$Path)

    try {
        if (!(Test-Path $Path)) {
            New-Item -Path $Path -ItemType Directory -Force | Out-Null
        }

        $testFile = Join-Path $Path "fabric_write_test.tmp"
        "test" | Out-File -FilePath $testFile -Force
        Remove-Item $testFile -Force
        return $true
    }
    catch {
        return $false
    }
}

# Download and install Fabric
function Install-Fabric {
    param(
        [string]$Architecture,
        [string]$InstallDirectory
    )

    # Construct download URL
    $filename = "fabric_Windows_$Architecture.zip"
    $downloadUrl = "https://github.com/danielmiessler/fabric/releases/latest/download/$filename"

    Write-Info "Downloading Fabric for Windows $Architecture..."
    Write-Info "URL: $downloadUrl"

    # Create temporary directory
    $tempDir = Join-Path $env:TEMP "fabric_install_$(Get-Random)"
    New-Item -Path $tempDir -ItemType Directory -Force | Out-Null
    $tempFile = Join-Path $tempDir "fabric.zip"

    try {
        # Download the archive
        Write-Info "Downloading archive..."
        Invoke-WebRequest -Uri $downloadUrl -OutFile $tempFile -UseBasicParsing

        Write-Info "Extracting Fabric binary..."

        # Extract the zip file
        Add-Type -AssemblyName System.IO.Compression.FileSystem
        $zip = [System.IO.Compression.ZipFile]::OpenRead($tempFile)

        # Find and extract only fabric.exe
        $fabricEntry = $zip.Entries | Where-Object { $_.Name -eq "fabric.exe" }
        if (!$fabricEntry) {
            Write-Error "fabric.exe not found in the downloaded archive"
            exit 1
        }

        # Create install directory if it doesn't exist
        if (!(Test-Path $InstallDirectory)) {
            Write-Info "Creating install directory: $InstallDirectory"
            New-Item -Path $InstallDirectory -ItemType Directory -Force | Out-Null
        }

        # Extract fabric.exe to install directory
        $fabricPath = Join-Path $InstallDirectory "fabric.exe"
        Write-Info "Installing Fabric to $fabricPath..."

        [System.IO.Compression.ZipFileExtensions]::ExtractToFile($fabricEntry, $fabricPath, $true)
        $zip.Dispose()

        Write-Success "Fabric installed successfully to $fabricPath"
        return $fabricPath
    }
    catch {
        Write-Error "Failed to download or extract Fabric: $($_.Exception.Message)"
        exit 1
    }
    finally {
        # Clean up
        if (Test-Path $tempDir) {
            Remove-Item $tempDir -Recurse -Force -ErrorAction SilentlyContinue
        }
    }
}

# Check if directory is in PATH
function Test-InPath {
    param([string]$Directory)

    $pathDirs = $env:PATH -split ';'
    return $pathDirs -contains $Directory
}

# Provide PATH setup instructions
function Show-PathInstructions {
    param([string]$InstallDir)

    if (Test-InPath $InstallDir) {
        Write-Success "✅ $InstallDir is already in your PATH"
    }
    else {
        Write-Warning "⚠️  $InstallDir is not in your PATH"
        Write-Info "To use fabric from anywhere, you have a few options:"
        Write-Info ""
        Write-Info "Option 1 - Add to PATH for current user (recommended):"
        Write-Info "  `$currentPath = [Environment]::GetEnvironmentVariable('PATH', 'User')"
        Write-Info "  [Environment]::SetEnvironmentVariable('PATH', `"`$currentPath;$InstallDir`", 'User')"
        Write-Info ""
        Write-Info "Option 2 - Add to PATH for all users (requires admin):"
        Write-Info "  `$currentPath = [Environment]::GetEnvironmentVariable('PATH', 'Machine')"
        Write-Info "  [Environment]::SetEnvironmentVariable('PATH', `"`$currentPath;$InstallDir`", 'Machine')"
        Write-Info ""
        Write-Info "Option 3 - Add to current session only:"
        Write-Info "  `$env:PATH += `";$InstallDir`""
        Write-Info ""
        Write-Info "After updating PATH, restart your terminal or run: refreshenv"
    }
}

# Verify installation
function Test-Installation {
    param([string]$FabricPath)

    if (Test-Path $FabricPath) {
        Write-Info "Verifying installation..."
        try {
            $version = & $FabricPath --version 2>$null
            if ($LASTEXITCODE -eq 0) {
                Write-Success "Fabric $version is working correctly!"
            }
            else {
                Write-Warning "Fabric binary exists but --version failed"
            }
        }
        catch {
            Write-Warning "Fabric binary exists but could not run --version"
        }
    }
    else {
        Write-Error "Fabric binary not found at $FabricPath"
        exit 1
    }
}

# Main installation function
function Main {
    Write-Info "🚀 Starting Fabric installation..."

    # Detect architecture
    $arch = Get-Architecture
    Write-Info "Detected architecture: $arch"

    # Determine install directory
    if (!$InstallDir) {
        $InstallDir = Join-Path $env:USERPROFILE ".local\bin"
    }

    Write-Info "Install directory: $InstallDir"

    # Check permissions
    if (!(Test-WritePermission $InstallDir)) {
        Write-Error "Cannot write to $InstallDir"
        Write-Error "Try running as Administrator or choose a different directory"
        Write-Info "Example with custom directory: `$env:INSTALL_DIR=`"C:\tools`"; iwr -useb ... | iex"
        exit 1
    }

    # Install Fabric
    $fabricPath = Install-Fabric -Architecture $arch -InstallDirectory $InstallDir

    # Verify installation
    Test-Installation -FabricPath $fabricPath

    # Check PATH and provide instructions
    Show-PathInstructions -InstallDir $InstallDir

    Write-Info ""
    Write-Success "🎉 Installation complete!"
    Write-Info ""
    Write-Info "Next steps:"
    Write-Info "  1. Run 'fabric --setup' to configure Fabric"
    Write-Info "  2. Add your API keys and preferences"
    Write-Info "  3. Start using Fabric with 'fabric --help'"
    Write-Info ""
    Write-Info "Documentation: https://github.com/danielmiessler/fabric"
}

# Run main function
Main
