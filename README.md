# File-Sorter
Python/Bash script to sort big data files such as Data-Breaches to be able to grep/find a line much quicker.


<ul>
<li>
<h2>What it does?</h2>
</li>
<p>
<pre>sudo make install</pre>
</p>

<li>
<h2>How to install it?</h2>
</li>
<p>
<pre>git clone https://github.com/bastien8060/file-sorter</pre>
</p>

<li>
<h1>How to import data?</h1>
</li>
<p><pre>cd file-sorter
cd ./python
./addbreach.py</pre>
Note: The data file has to be in <pre>python/inputbreach/breach.txt</pre>
<h3>
Options:
</h3>
<ul>
<li>-D: delete source file after completed.</li>
</ul>


<li>
<h1>How to query data?</h1>
</li>
<p>After you have finished importing the data file, you can query for it by typing:
E.g.<pre> ./query.sh name@example.com</pre> </p>


<li>
<h2>Information:</h2>
</li><ul>
<li><p>List of imported files are in the "imported.log" file. The script keeps track of imported file in this log with their SHA sums to prevent a file to be added twice.</p></li>
 <li><p>All data is in "data" folder.</p></li>
 <li><p>Use only for educational and penetration testing purposes.</p></li></ul>
</ul>






