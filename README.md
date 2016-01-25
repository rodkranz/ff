# Find File #

Find File or text in file with go language.

-----------------------
### Source Find File ###

* Find text or file name 
* Version: 0.0.1
* License: ISC


-----------------------
## Parameters 
##### Helper: 
```
	Usage of ./main:
	  -f string
	    	the file name that I have to looking for.
	  -no-color
	    	Disable color output
	  -p string
	    	path string (default "./")
	  -r int
	    	range between start and end of the line (default 10)
	  -t string
	    	the word that I have to looking for.
```

-----------------------
## Compile 
	compile files from different platforms

#### Linux 

	```$ env GOOS=linux GOARCH=arm GOARM=7 go build -o ff main.go``` 

#### MacOs

	```$ env GOOS=darwin GOARCH=386 go build -o ff main.go```

#### Windows 
	
	```$ env GOOS=windows GOARCH=386 go build -o ff.exe main.go```




