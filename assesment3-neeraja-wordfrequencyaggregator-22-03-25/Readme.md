Problem Statement:= Word Frequency Aggregator

question:= Read a .txt file and prints frequency of each word 

input:=(data.txt)

expected output:=
reads 1
file 3
and 2
words 1
Hello 1
frequency 2
aggregator 1
word 1
This 1
Word 1
It 1
counts 1
in 2
the 5
each 1
there 1
is 1
of 2
no 1
text 1
it 1
prints 1
Word frequency data saved to wordfrequency.csv

Code Explanation:=

1) ReadFileChunks() -- it opens the file and reads
2) SplitWords() -- splits the words 
3) CountFrequency() -- it will counts the frequency of the words
4) WordFrequency() --  it is used to communicate between multiple goroutines
5) go build -o "wordfrequencyaggregator.exe" main.go -- creates binary executable file


