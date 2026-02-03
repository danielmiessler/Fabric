#!/bin/bash
if [ $# < 1 ]
then
	printf "Please provide an agent name!\n"
	printf "./create_patterns_in_opencode.sh foobar"
	exit 0
fi
agentname=$1
mkdir ~/.config/opencode/commands
for pattern in ../../data/patterns/*;
do
if [[ -f $pattern ]]
then
	printf "$pattern is a file, skipping...\n"
	continue
fi
target=$(echo $pattern | awk -F'/' '{print $NF}')
desc="${target%.*}"
cat >> ~/.config/opencode/commands/$target.md << EOL
---
description: $desc
agent: $agentname
---
EOL
cat $pattern/system.md >> ~/.config/opencode/commands/$target.md
# Not sure if it is usefull to cat user.md if existing?
# if [ -f $pattern/user.md ]
# then
# 	cat $pattern/user.md >> ~/.config/opencode/commands/$target.md
# fi
printf "Wrote $target\n"
done
