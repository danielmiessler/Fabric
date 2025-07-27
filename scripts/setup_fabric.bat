@echo off
setlocal enabledelayedexpansion

:: Check if running with administrator privileges
net session >nul 2>&1
if %errorlevel% neq 0 (
    echo Please run this script as an administrator.
    pause
    exit /b 1
)

:: Install Chocolatey (package manager for Windows)
if not exist "%ProgramData%\chocolatey\bin\choco.exe" (
    echo Installing Chocolatey...
    @"%SystemRoot%\System32\WindowsPowerShell\v1.0\powershell.exe" -NoProfile -InputFormat None -ExecutionPolicy Bypass -Command "[System.Net.ServicePointManager]::SecurityProtocol = 3072; iex ((New-Object System.Net.WebClient).DownloadString('https://chocolatey.org/install.ps1'))" && SET "PATH=%PATH%;%ALLUSERSPROFILE%\chocolatey\bin"
)

:: Install Go
where go >nul 2>&1
if %errorlevel% neq 0 (
    echo Installing Go...
    choco install golang -y
    set "PATH=%PATH%;C:\Program Files\Go\bin"
)

:: Install Git
where git >nul 2>&1
if %errorlevel% neq 0 (
    echo Installing Git...
    choco install git -y
)

:: Refresh environment variables
call refreshenv

:: Install Fabric
echo Installing Fabric...
go install github.com/danielmiessler/fabric@latest

:: Run Fabric setup
echo Running Fabric setup...
fabric --setup

:: Install yt helper
echo Installing yt helper...
go install github.com/danielmiessler/yt@latest

:: Prompt user for YouTube API Key
set /p YOUTUBE_API_KEY=Enter your YouTube API Key (press Enter to skip): 
if not "!YOUTUBE_API_KEY!"=="" (
    echo YOUTUBE_API_KEY=!YOUTUBE_API_KEY!>> %USERPROFILE%\.config\fabric\.env
)

:: Prompt user for OpenAI API Key
set /p OPENAI_API_KEY=Enter your OpenAI API Key (press Enter to skip): 
if not "!OPENAI_API_KEY!"=="" (
    echo OPENAI_API_KEY=!OPENAI_API_KEY!>> %USERPROFILE%\.config\fabric\.env
)

:: Run Fabric
:run_fabric
cls
echo Fabric is now installed and ready to use.
echo.
echo Available options:
echo 1. Run Fabric with custom options
echo 2. List patterns
echo 3. List models
echo 4. Update patterns
echo 5. Exit
echo.
set /p CHOICE=Enter your choice (1-5): 

if "%CHOICE%"=="1" (
    set /p PATTERN=Enter pattern (or press Enter to skip): 
    set /p CONTEXT=Enter context (or press Enter to skip): 
    set /p SESSION=Enter session (or press Enter to skip): 
    set /p MODEL=Enter model (or press Enter to skip): 
    set /p TEMPERATURE=Enter temperature (or press Enter for default): 
    set /p STREAM=Do you want to stream output? (Y/N): 

    set "FABRIC_CMD=fabric"
    if not "!PATTERN!"=="" set "FABRIC_CMD=!FABRIC_CMD! --pattern !PATTERN!"
    if not "!CONTEXT!"=="" set "FABRIC_CMD=!FABRIC_CMD! --context !CONTEXT!"
    if not "!SESSION!"=="" set "FABRIC_CMD=!FABRIC_CMD! --session !SESSION!"
    if not "!MODEL!"=="" set "FABRIC_CMD=!FABRIC_CMD! --model !MODEL!"
    if not "!TEMPERATURE!"=="" set "FABRIC_CMD=!FABRIC_CMD! --temperature !TEMPERATURE!"
    if /i "!STREAM!"=="Y" set "FABRIC_CMD=!FABRIC_CMD! --stream"

    echo Running Fabric with command: !FABRIC_CMD!
    !FABRIC_CMD!
    pause
    goto run_fabric
) else if "%CHOICE%"=="2" (
    fabric --listpatterns
    pause
    goto run_fabric
) else if "%CHOICE%"=="3" (
    fabric --listmodels
    pause
    goto run_fabric
) else if "%CHOICE%"=="4" (
    fabric --updatepatterns
    pause
    goto run_fabric
) else if "%CHOICE%"=="5" (
    exit /b 0
) else (
    echo Invalid choice. Please try again.
    pause
    goto run_fabric
)
