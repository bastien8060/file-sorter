#!/bin/bash
echo " "
echo " "
touch imported.log
inputfile="inputbreach/breach.txt"
echo "[*] Checking breach $inputfile checksum..." | tee -a .debug
shasum="$(sha256sum "./inputbreach/breach.txt" | cut -d' ' -f1)"
	if [ "$(cat imported.log | grep "$shasum" | wc -l)" == "0" ]; then
		cat imported.log | grep "$shasum" | tee -a .debug
		echo "[*] Logging sha256sum into imported.log..." | tee -a .debug
		echo "$(date --rfc-3339=date): $(sha256sum "$inputfile")	$(du -h "$inputfile"|cut -d'	' -f1)" | tee -a imported.log | tee -a .debug
		echo "------------------------------------------" | tee -a .debug
		echo " "

else
if [[ $1 == "-d" ]]; then
rm $inputfile -f
rm "../inputbreach/breach.txt" -f
fi
echo "[*] This breach is already imported." | tee -a .debug
		cat imported.log | grep "$shasum" | tee -a .debug
		echo "------------------------------------" | tee -a .debug
		echo " "
killall python
fi
