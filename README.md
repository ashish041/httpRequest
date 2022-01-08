# Purpose of the tool
---------------------
The purpose of this tool is to make parallel http requests,
to finish the http requests as soon as possible. At the same
time adjust the number of parallel processes to the system
resources to prevent the high CPU usages.

# Installation
--------------
```
Command to build the tool:
	go build -o myhttp
Command to test the tool:
 	go test
Command to run the tool: 
	./myhttp http://www.yahoo.com http://google.com
	./myhttp facebook.com
	./myhttp -parallel 5 yahoo.com google.com facebook.com yahoo.com yandex.com twitter.com reddit.com/r/funny reddit.com/r/notfunny baroquemusiclibrary.com
```