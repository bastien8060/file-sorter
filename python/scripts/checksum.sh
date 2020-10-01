#!/bin/bash
touch imported.log
inputfile="inputbreach/breach.txt"
echo "[*] Checking breach $inputfile checksum..." | tee -a .debug
shasum="$(sha256sum "./inputbreach/breach.txt" | cut -d' ' -f1)"
	if [ "$(cat imported.log | grep "$shasum" | wc -l)" == "0" ]; then
		cat imported.log | grep "$shasum" | tee -a .debug

else
echo "[*] This breach is already imported." | tee -a .debug
		cat imported.log | grep "$shasum" | tee -a .debug
		echo "------------------------------------" | tee -a .debug
		echo " "
kill $1
fi

echo " "
echo " "
