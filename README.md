# File-Sorter
Golang/Python/Bash script to sort big data files such as Data-Breaches to be able to grep/find a line much quicker.


<ul>

<h2>What it does?</h2>

<p>
It is a Golang/Python/Bash script to sort big data files such as Data-Breaches to be able to grep/find a line much quicker.
</p>


<h2>Okay but... How?</h2>

<p>
It takes the input file (Has to be located under inputbreach/breach.txt) and splits it equally in files which are then orginised neatly where they can be found quickly. This way the system doesn't have to grep a huge file to find an occurence, it just needs to scan the right file which will be much smaller, resulting in faster speeds.
</p>


<h2>How to install it?</h2>

<p>Easy! In a one-liner:<br>
<pre>git clone https://github.com/bastien8060/file-sorter</pre>
</p>


<h2>How to import data?</h2>



<li><h3><u>Golang</u></h3></li>
<p><pre>
cd file-sorter
cd ./golang
./addbreach</pre>
Note: The data file has to be in <pre>golang/inputbreach/breach.txt</pre>
<h4>
Options:
</h4>
<ul>
<li>-D: delete source file after completed.</li>
</ul>


<li><h3><u>Python</u></h3></li>
<p><pre>cd file-sorter
cd ./python
./addbreach.py</pre>
Note: The data file has to be in <pre>python/inputbreach/breach.txt</pre>
<h4>
Options:
</h4>
<ul>
<li>-D: delete source file after completed.</li>
</ul>


<li><h3><u>Bash</u></h3></li>
<p><pre>cd file-sorter
cd ./bash
./sorter.sh</pre>
Note: The data file has to be in <pre>bash/inputbreach/breach.txt</pre>
<h4>
Options:
</h4>
<ul>
<li>-D: delete source file after completed.</li>
</ul>



<h2>How to query data?</h2>

<p>After you have finished importing the data file, you can query for it by typing:
E.g.<pre> ./query.sh name@example.com</pre> </p>


<li>
<h2><u>Information:</u></h2>
</li><ul>
<li><p>In term of speed, the Golang version is much faster to run, taking a couple second only to import 50mb.</p><p>In second place comes the Python3 version taking about nearly a minute.</p><p>I would avoid relying on the bash version as it is old and hasn't been maintained in over a year and an half. However, for information, it takes the bash version few minutes to import a 50mb file.</p></li>
<li><p>List of imported files are in the "imported.log" file. The script keeps track of imported file in this log with their SHA sums to prevent a file to be added twice. Each platforms are works differently and do not interfere with each others, therefore, it does not check duplication across platforms (e.g. from Python to Golang)</p></li>
 <li><p>All data is in "data" folder.</p></li>
 <li><p>Use only for educational and penetration testing purposes.</p></li></ul>
</ul>







