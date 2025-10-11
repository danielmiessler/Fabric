#!/bin/bash

# Script to install the word-generator, removing any existing installations first

echo "Removing existing word-generator extension..."
fabric --rmextension=word-generator

echo "Copying updated word-generator configuration..."
cp /Users/ourdecisions/devHelp/Fabric-fix/internal/plugins/template/Examples/word-generator.yaml /Users/ourdecisions/.config/fabric/extensions/configs/word-generator.yaml

echo "Copying updated word-generator script..."
cp /Users/ourdecisions/devHelp/Fabric-fix/internal/plugins/template/Examples/word-generator.py /Users/ourdecisions/.config/fabric/extensions/bin/word-generator.py

fabric --addextension /Users/ourdecisions/.config/fabric/extensions/configs/word-generator.yaml

# Update the text of path of /Users/ourdecisions/.config/fabric/extensions/configs/word-generator.yaml
# of executable: /usr/local/bin/word-generator.py to /Users/ourdecisions/.config/fabric/extensions/bin/word-generator.py
sed -i '' 's|/usr/local/bin/word-generator.py|/Users/ourdecisions/.config/fabric/extensions/bin/word-generator.py|g' /Users/ourdecisions/.config/fabric/extensions/configs/word-generator.yaml

echo "Word-generator extension has been updated successfully!"

echo "Generate with Word-generator and it works as expected:"

# Run (generate 3 random words) with debug logging
echo "{{ext:word-generator:generate:3}}" | fabric --debug=3

####### Now we will add word-counter extension #########
