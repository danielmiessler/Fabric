# How to Install the Web Interface and PDF-to-Markdown

If Fabric is already installed and you see fabric/web, go to step 3

If fabric is not installed, ensure Go is installed https://go.dev/doc/install and node / npm for web https://nodejs.org/en/download.

There are many ways to install fabric. Here's one approach that usually works well:

## Step 1: clone the repo
In terminal, from the parent directory where you want to install fabric:
git clone https://github.com/danielmiessler/fabric.git

## Step 2 : Install Fabric
cd fabric
go install github.com/danielmiessler/fabric@latest

## Step 3: Install GUI
Navigate to the web directory and install dependencies:

cd web

npm install

npx svelte-kit sync

## Step 4: Install PDF-to-Markdown
Install the PDF conversion components in the correct order:
cd web
# Install dependencies in this specific order

npm install -D patch-package

npm install -D pdfjs-dist@2.5.207

npm install -D github:jzillmann/pdf-to-markdown#modularize


No build step is required after installation.

## Step 5: Update Shell Configuration if not already done from your fabric installation
For Mac/Linux users:

Add environment variables to your ~/.bashrc (Linux) or ~/.zshrc (Mac) file:

# For Intel-based Macs or Linux
export GOROOT=/usr/local/go
export GOPATH=$HOME/go
export PATH=$GOPATH/bin:$GOROOT/bin:$HOME/.local/bin:$PATH

# For Apple Silicon Macs
export GOROOT=$(brew --prefix go)/libexec
export GOPATH=$HOME/go
export PATH=$GOPATH/bin:$GOROOT/bin:$HOME/.local/bin:$PATH

REFER TO OFFICIAL FABRIC README.MD FILE FOR OTHER OPERATING SYSTEMS

Step 5: Create Aliases for Patterns
Add the following to your .zshrc or .bashrc file to create shorter commands:

```bash

# The following three lines of code are path examples, replace with your actual path.

# Add fabric to PATH
export PATH="/Users/USERNAME/Documents/fabric:$PATH"

# Define the base directory for Obsidian notes
obsidian_base="/Users/USERNAME/Documents/fabric/web/myfiles/Fabric_obsidian"

# Define the patterns directory
patterns_dir="/Users/USERNAME/Documents/fabric/patterns"


# Loop through all files in the ~/.config/fabric/patterns directory

for pattern_file in ~/.config/fabric/patterns/*; do
    # Get the base name of the file
    pattern_name=$(basename "$pattern_file")

    # Unalias any existing alias with the same name
    unalias "$pattern_name" 2>/dev/null

    # Define a function dynamically for each pattern
    eval "
    $pattern_name() {
        local title=\$1
        local date_stamp=\$(date +'%Y-%m-%d')
        local output_path=\"\$obsidian_base/\${date_stamp}-\${title}.md\"

        # Check if a title was provided
        if [ -n \"\$title\" ]; then
            # If a title is provided, use the output path
            fabric --pattern \"$pattern_name\" -o \"\$output_path\"
        else
            # If no title is provided, use --stream
            fabric --pattern \"$pattern_name\" --stream
        fi
    }
    "
done

# YouTube shortcut function
yt() {
    local video_link="$1"
    fabric -y "$video_link" --transcript
}


After modifying your shell configuration file, apply the changes:

source ~/.zshrc  # or source ~/.bashrc for Linux

Step 6: Run Fabric Setup
Initialize fabric configuration:

fabric --setup

Step 7: Launch the Web Interface
Open two terminal windows and navigate to the web folder:

Terminal 1: Start the Fabric API Server
fabric --serve

Terminal 2: Start the Development Server
npm run dev


If you get an ** ERROR **.
It would be much appreciated that you copy /paste your error in your favorite LLM before opening a ticket, 90% of the time your llm will point you to the solution.

Also if you modify patterns, descriptions or tags in  Pattern_Descriptions/pattern_descriptions.json, make sure to copy the file over in  web/static/data/pattern_descriptions.json  

_____   ______   ______

OPTIONAL: Create Start/Stop Scripts 
You can create scripts to start/stop both servers at once.

### For Mac Users
When creating scripts on Mac using TextEdit:

1. Open TextEdit
2. **IMPORTANT:** Select "Format > Make Plain Text" from the menu BEFORE pasting any code
3. Paste the script content, follow instructions below ((Mac example)).


### For Windows Users
When creating scripts on Windows:

1. Use Notepad or a code editor like VS Code
2. Paste the script content
3. Save the file with the appropriate extension
4. Ensure line endings are set to LF (not CRLF) for bash scripts

ACTUAL SCRIPTS (Mac example)

Start Script 
1. Create a new file named start-fabric.command on your Desktop:

#!/bin/bash

# Change to the fabric web directory
cd "$HOME/Documents/Github/fabric/web"

# Start fabric serve in the background
osascript -e 'tell application "Terminal" to do script "cd '$HOME'/Documents/Github/fabric/web && fabric --serve; exit"'

# Wait a moment to ensure the fabric server starts
sleep 2

# Start npm development server in a new terminal
osascript -e 'tell application "Terminal" to do script "cd '$HOME'/Documents/Github/fabric/web && npm run dev; exit"'

# Close this script's terminal window after starting servers
echo "Fabric servers started!"
sleep 1
osascript -e 'tell application "Terminal" to close (every window whose name contains ".command")' &
exit

Stop Script

2. Create a new file named stop-fabric.command on your Desktop:

#!/bin/bash

# Kill the npm dev server
pkill -f "node.*dev"

# Kill the fabric server
pkill -f "fabric --serve"

# Force quit Terminal entirely and restart it
osascript <<EOD
tell application "Terminal" to quit
delay 1
tell application "Terminal" to activate
EOD

echo "Fabric servers stopped!"
sleep 1

# This script's terminal will already be closed by the quit command above
exit

3. Make both scripts executable:
chmod +x ~/Desktop/start-fabric.command
chmod +x ~/Desktop/stop-fabric.command

You can customize with icons by finding suitable .icns files, right-clicking each .command file, selecting "Get Info", and dragging your icon file onto the small icon in the top-left corner.

Note: You might need to allow the scripts to execute in your security settings by going to System Preferences → Security & Privacy after trying to run them the first time.



## 🎥 Demo Video
https://youtu.be/XMzjgqvdltM
