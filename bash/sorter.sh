#!/bin/bash

if [[ "$1" == "-h" || "$1" = "--help" || "$1" = "-help" || "$1" = "-H" ]]; then
echo "[*] Usage "
	echo '	Add breach files in the directory "inputbreach" and run ./sorter.sh'
	echo "	The breach has to be in the format:"
	echo "		adress1@domain.com:password1"
	echo "		adress2@domain.com:password2"
	echo "		..."
	echo "	Once finished, use ./query.sh to search for lines"
	echo " "
	echo "[*] Parameters:"
	echo '	"-D" : Deletes input file or files after being imported'
	exit 0
fi
ls ./inputbreach |sort| sed 's/^/.\/inputbreach\//' | while read inputfile;
do
echo "[*] Checking breach $inputfile checksum..." | tee -a debug
shasum="$(sha256sum "$inputfile" | cut -d' ' -f1)"
	if [ "$(cat imported.log | grep \"$shasum\" | wc -l)" == "0" ]; then
		cat imported.log | grep "$shasum" | tee -a debug


# data folder sorting
echo "[*] Sorting breaches ($inputfile)..." | tee -a debug
while IFS= read -r line
do
dir=$( cd "$( dirname ${BASH_SOURCE[0]} )" && pwd )
if [ "$line" != "" ]; then
	letter1="$(echo \"${line,,}\"||cut -b1)"
	if [[ "$letter1" == [a-zA-Z0-9] ]]; then
		if [ -f "$dir/data/$letter1" ]; then
			echo "$line" >>  "$dir/data/$letter1"
			#echo "added to \"$dir\"/data/\"$letter1\""
		else
			letter2=\"$(echo "${line,,}"|cut -b2)\"
			if [[ "$letter2" == [a-zA-Z0-9] ]]; then
				if [ -f "\"$dir\"/data/\"$letter1\"/\"$letter2\"" ]; then
					echo "$line" >>  "\"$dir\"/data/\"$letter1\"/\"$letter2\""
					#echo "added to $dir/data/$letter1/$letter2"
				else
					letter3="$(echo \"${line,,}\"|cut -b3)"
					if [[ "$letter3" == [a-zA-Z0-9] ]]; then
						if [ -f "\"$dir\"/data/$letter1/\"$letter2\"/\"$letter3\"" ]; then
							echo "$line" >>  "\"$dir\"/data/$letter1/\"$letter2\"/\"$letter3\""
							#echo "added to $dir/data/$letter1/$letter2/$letter3"
						fi
					else
						if [ -f "\"$dir\"/data/\"$letter1\"/\"$letter2\"/symbols" ]; then
							echo "$line" >> "\"$dir\"/data/\"$letter1\"/\"$letter2\"/symbols"
							#echo "added to $dir/data/$letter1/$letter2/symbols"
						fi
					fi
				fi
			else
				if [ -f "\"$dir\"/data/\"$letter1\"/symbols" ]; then
				echo "$line" >>  "\"$dir\"/data/\"$letter1\"/symbols"
				#echo "added to $dir/data/$letter1/symbols"
				fi
			fi
		fi
	else
		if [ -f "\"$dir\"/data/symbols" ]; then
			echo "$line" >>  "\"$dir\"/data/symbols"
			#echo "added to $dir/data/symbols"
		fi
	fi
else
	echo "[*] Incorrect format! Example:"
	echo "name@domain.com:12345678"
	echo "name@domain.com:12345678"
	echo "..."
	sleep 5
fi

done < "$inputfile"

		echo "[*] Logging sha256sum into imported.log..." | tee -a debug
		echo "$(date --rfc-3339=date): $(sha256sum "$inputfile")	$(du -h "$inputfile"|cut -d'	' -f1)" | tee -a imported.log | tee -a debug
		echo "------------------------------------------" | tee -a debug
		if [ "$1" == "-D" ]
		then
			echo "[*] Removing breach file $inputfile..." | tee -a debug
			rm "$inputfile" | tee -a debug
		fi
	else
		echo "[*] This breach is already imported." | tee -a debug
		cat imported.log | grep "$shasum" | tee -a debug
		echo "------------------------------------" | tee -a debug
	fi
done



