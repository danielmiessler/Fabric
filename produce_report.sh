#!/bin/bash
numb=1
read -p "Name?" Name
read -p "male/female?" Sex
read -p "Age?" Age
echo -e "The Subject's name is $Name, he/she is a $Age year old $Sex. The following are a collection of texts, some are conversations between the subject and various AI models, some are various writings by the subject \n"  >> PsychoData
for i in $(ls *.txt );
        do
                echo -e "Text number $numb:" >> PsychoData
                cat $i >> PsychoData
                echo -e \n  >> PsychoData
                numb=$(expr $numb + 1)
        done
echo "\n Raw Data  \n"  >> ~/reports/psycho_analysis_$Name.txt
cat PsychoData >> ~/reports/psycho_analysis_$Name.txt
echo "\n Analysis: \n" >> ~/reports/psycho_analysis_$Name.txt
cat PsychoData | ~/Fabric/fabric -sp analyze_personality | tee ~/reports/psycho_analysis_$Name.txt
echo -e "Report is saved in ~/reports/psycho_analysis_$Name.txt"
