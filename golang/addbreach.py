#!/bin/python
import glob, os, sys, subprocess, re, time
from tqdm import tqdm
def file_len(fname):
    with open(fname) as f:
        for i, l in enumerate(f):
            pass
    return i + 1

if len(sys.argv) >= 2:
  if (sys.argv[1] == "-h") or (sys.argv[1] == "--help"):
    print("[*] Usage ")
    print(" Add breach files in the directory inputbreach and run sorter.py")
    print(" The breach has to be in the format:")
    print("   adress1@domain.com:password1")
    print("   adress2@domain.com:password2")
    print("   ...")
    print(" Once finished, use ./query.py to search for lines")
    print(" ")
    print("[*] Parameters:")
    print(' "-D" : Deletes input file or files after being imported')
    exit(0) 
#print("no help needed take the L")
with open('./scripts/checksum.sh', 'rb') as file:
  script = file.read()
rc = subprocess.call("bash ./scripts/checksum.sh "+str(os.getpid()), shell=True)
breach = "inputbreach/breach.txt"
filetoindex = breach
direc = os.path.dirname(os.path.abspath(__file__))
print("[*] Checking number of lines for "+filetoindex+".")
num_lines = sum(1 for line in open('./inputbreach/breach.txt'))
print("  ")
print("  ")
print("[*] Sorting "+filetoindex+".")
with open(breach,'r') as f:
  for line in tqdm(f, total=num_lines, unit="lines"):
    if (line != ""):
      #print(line)
      letter1 = line[0].lower();
      if re.search(r"[a-zA-Z0-9]", letter1):
        if os.path.isfile("data/"+letter1):
          with open("data/"+letter1, 'a') as f2:
            print(line, file=f2)
        else:
          letter2 = line[1].lower();
          if re.search(r"[a-zA-Z0-9]", letter2):
            if os.path.isfile("data/"+letter1+"/"+letter2):
              with open("data/"+letter1+"/"+letter2, 'a') as f2:
                #print >>f2, line
                print(line, file=f2)
            else:
              letter3 = line[2].lower();
              if re.search(r"[a-zA-Z0-9]", letter3):
                if os.path.isfile("data/"+letter1+"/"+letter2+"/"+letter3):
                  with open("data/"+letter1+"/"+letter2+"/"+letter3, 'a') as f2:
                    print(line, file=f2)
              else:
                if os.path.isfile("data/"+letter1+"/"+letter2+"/symbols"):
                  with open("data/"+letter1+"/"+letter2+"/symbols", 'a') as f2:
                    print(line, file=f2)               
          else:
            if os.path.isfile("data/"+letter1+"/symbols"):
              with open("data/"+letter1+"/symbols", 'a') as f2:
                print(line, file=f2)
      else:
        if os.path.isfile("data/symbols"):
          with open("data/symbols", 'a') as f2:
            print(line, file=f2)
    else:
      print("incorrect format! ./"+sys.argv[0]+" -h to see help")
if len(sys.argv) >= 2:
  if (sys.argv[1] == "-D"):
    rc = subprocess.call("bash ./scripts/shalog.sh -d", shell=True)
  else:
    rc = subprocess.call("bash ./scripts/shalog.sh", shell=True)
else:
  rc = subprocess.call("bash ./scripts/shalog.sh", shell=True)
