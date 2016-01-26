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
$ ff -h
Usage of ./ff:
  -a int
    	Range around of the word that I found. (default 10)
  -d	Show Debug Mode
  -f string
    	Filter by name of file or the file name
  -no-color
    	Disable color output
  -p string
    	The directory path (default "./")
  -r	Search by regex
  -t string
    	what I am searching
```

-----------------------

## Usage

Without parameters will show everything in subfolders.

	$ ff
	----------------------------------------------------------------------------------------------------
	Path : ./
	----------------------------------------------------------------------------------------------------
	[File] ./ 
	[File] .git 
	[File] .git/COMMIT_EDITMSG 
	[File] .git/FETCH_HEAD 
	[File] .git/HEAD 
	.....


-----------------------

### Parameter -f

With parameter `-f` will filter by file names. 


	$ ff -f css
	----------------------------------------------------------------------------------------------------
	Path : ./
	File : css
	----------------------------------------------------------------------------------------------------
	[File] lookingFor/resources/css 
	[File] lookingFor/resources/css/angular-bootstrap-datepicker.css 
	[File] lookingFor/resources/css/bootstrap3.3.2.min.css 
	[File] lookingFor/resources/css/font-awesome.css 
	....


-----------------------
### Parameter -t

With parameter `-t` will looking for text inside of file.
	

	$ ff -t rlopes
	----------------------------------------------------------------------------------------------------
	Path : ./
	Text : rlopes
	----------------------------------------------------------------------------------------------------
	[File] lookingFor/resources/css/toastr.css 
		[144] .rlopes { lorem i
	----------------------------------------------------------------------------------------------------
	[File] lookingFor/services.dev.ini 
		[3] omain = ".rlopes.realestat
		[6] = 'http://rlopes.realestat
		[15] = 'http://rlopes.realestat
	----------------------------------------------------------------------------------------------------


-----------------------

### Parameter -r

With parameter `-r` you will find by regular expression


	$ ff -r -t "(rlopes)|(consectetur)"
	----------------------------------------------------------------------------------------------------
	Path : ./
	Regex: (rlopes)|(consectetur)
	----------------------------------------------------------------------------------------------------
	[File] lookingFor/resources/css/none.txt 
		[1] sit amet, consectetur adipisici
		[4] sit amet, consectetur adipisici
	----------------------------------------------------------------------------------------------------
	[File] lookingFor/resources/css/toastr.css 
		[144] .rlopes { lorem i
	----------------------------------------------------------------------------------------------------
	[File] lookingFor/services.dev.ini 
		[3] omain = ".rlopes.realestat
		[6] = 'http://rlopes.realestat
		[15] = 'http://rlopes.realestat
	----------------------------------------------------------------------------------------------------


-----------------------

### Parameter Combination

you can combinate both of parameters as `-t` and `-f`


	$ ff -t "rlopes" -f ini
	----------------------------------------------------------------------------------------------------
	Path : ./
	File : ini
	Text : rlopes
	----------------------------------------------------------------------------------------------------
	[File] lookingFor/services.dev.ini 
		[3] omain = ".rlopes.realestat
		[6] = 'http://rlopes.realestat
		[15] = 'http://rlopes.realestat
	----------------------------------------------------------------------------------------------------


-----------------------

### Parameter -d 

with parameter `-d` you active the debug mode, you can see the time of searching


	$ ff -t "rlopes" -f ini -d
	----------------------------------------------------------------------------------------------------
	Path : ./
	File : ini
	Text : rlopes
	----------------------------------------------------------------------------------------------------
	[File] lookingFor/services.dev.ini 
		[3] omain = ".rlopes.realestat
		[6] = 'http://rlopes.realestat
		[15] = 'http://rlopes.realestat
	----------------------------------------------------------------------------------------------------
	final Execution took 2.031182ms


-----------------------

### Parameter -v

Show the version

-----------------------

### Parameter -h

Show the helper


-----------------------
## Compile 

 compile files from different platforms


#### Linux 

	$ env GOOS=linux GOARCH=arm GOARM=7 go build -o ff main.go


#### MacOs

	$ env GOOS=darwin GOARCH=386 go build -o ff main.go


#### Windows 
	
	$ env GOOS=windows GOARCH=386 go build -o ff.exe main.go
