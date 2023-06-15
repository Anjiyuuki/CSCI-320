# Assignment 9 - Go 
# Angie Bui

This program accepts a filename in the command line which contains a list of URL's, then goes into a string slice and start a goroutine for each element in the string slice. It waits 5 seconds for the routine to complete. URL and the length of the web page body is returned. It then goes to the second part which is regular expressions, where the program is extended to extract the web page title and prints it. 


Part 1 Goroutines and Part 2 regular expressions are together. 

## Running
Open file in Visual Studio. 
From the Visual Studio terminal:
```
go build asn9.go
./asn9 Sample_A9.txt

```
