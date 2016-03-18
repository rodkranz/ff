# Find File #

Find File or text in file with go language.


-----------------------

### Source Find File ###

* Find text or file name 
* Version: 1.1.2
* License: ISC

-----------------------

## Download

This is the latest version of FF.

* [Linux](https://mega.nz/#!5QZj3Rab!tVUxSvhxVPkfWhVdhm4x_h1Cx0q3RttW6IFajXCF-oo)
* [MacOSx](https://mega.nz/#!MIYWWL5T!kBwAp0vFcN67A_fcC2cd4JR5KDKaACqfjhqIQr1M1dg)
* [Windows](https://mega.nz/#!NdJU2ALK!eqhiuCq3EHmH59AUO5wxRNbDLzsUavrKa9Tm7MjcZjY)

Linux or MacOSx you need to put the file in your `bin`'s folder and Windows You have to put the file in `windows`'s folder

-----------------------

## How to Compile it

 Compile those files from different platforms

### Requirements

* [GO Language](https://golang.org/doc/install)

#### Compiling to *Linux*

	$ env GOOS=linux GOARCH=arm GOARM=7 go build -o ff main.go


#### Compiling to *MacOSx*

	$ env GOOS=darwin GOARCH=386 go build -o ff main.go


#### Compiling to *Windows*

	$ env GOOS=windows GOARCH=386 go build -o ff.exe main.go


-----------------------

## Parameters

##### Helper: 

```
        Usage of ff:
          --exclude-dir string
                Exclude dir from reader (default ".bzr,CVS,.git,.hg,.svn")
          --no-color
                Disable color output
          --version
                Show the version
          -alsologtostderr
                log to standard error as well as files
          -cpu int
                Number of CPU you have 4 available (default 4)
          -d string
                Directory searching (default "./")
          -f string
                Filter by file name
          -log_backtrace_at value
                when logging hits line file:N, emit a stack trace (default :0)
          -log_dir string
                If non-empty, write log files in this directory
          -logtostderr
                log to standard error instead of files
          -r int
                Range around of the word (default 10)
          -regex
                Search by this Regex
          -stderrthreshold value
                logs at or above this threshold go to stderr
          -t string
                Text searching
          -u	Use case sensitive (default true)
          -up
                Check update
          -v value
                log level for V logs
          -vmodule value
                comma-separated list of pattern=N settings for file-filtered logging
                       
```


-----------------------

## Usage

Without parameters will show everything in subfolders.

```
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
```


-----------------------

### Parameter -f

Parameter `-f` of file, this parameter you will define the name of file  that you should like to filter

```
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
```


-----------------------

### Parameter -t

Parameter `-t` of text, this text that you will search inside of file.
	
```
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
```


-----------------------

### Parameter --cpu

Parameter `-cpu`, this parameter you will define the number of cpu the program can use.

Normal use:
```
	$ ff -t log
	ff -t "log"  0.01s user 0.01s system 70% cpu 0.034 total
```

Four CPU use:
```
	$ ff -t log -cpu 4
    ff -t "log" -cpu 4  0.01s user 0.01s system *159%* cpu 0.024 total
```


-----------------------

### Parameter -r

With parameter `-r` as regex will defined that your `-t` is a regular expression

```
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
```


-----------------------

### Parameter Combination

you can match both of the parameters as `-t` and `-f` to create more complex search

```
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
```


-----------------------

### Parameter -d 

The `-d` parameter will active the debug mode, you can see the time of searching (more things coming soon)

```
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
```


-----------------------

### Parameter -v

Show the version of application
    
```
    $ ff -version                                                                                                                                                master  ✭
    ----------------------------------------------------------------------------------------------------
            This program has written by Rodrigo Lopes <dev.rodrigo.lopes@gmail.com>.
            Only for academic purposes
    ----------------------------------------------------------------------------------------------------
      Version : 1.1.2
      Language: GO Language
      License : ISC
      Project : https://github.com/rodkranz/ff
      Contact : dev.rodrigo.lopes@gmail.com
    ----------------------------------------------------------------------------------------------------
```


-----------------------

### Parameter -h

Show the helper and parameters available


-----------------------

### Parameter -up

Check if exist an update available of sistem.

```
    $ ff -up
        ----------------------------------------------------------------------------------------------------
        You have the latest version
        Current version is .
        ----------------------------------------------------------------------------------------------------
```


-----------------------

## Time Execution
	
The all tests was made in the same computer
the folder has `1434` files


-----------------------

The `ff` running looking for word `rlopes`

```
	$ time ff -t rlopes
	----------------------------------------------------------------------------------------------------
	Path : ./
	Text : rlopes
	----------------------------------------------------------------------------------------------------
	[File] resources/css/toastr.css 
		[144] .rlopes { lorem i
	----------------------------------------------------------------------------------------------------
	[File] services.dev.ini 
		[3] omain = ".rlopes.realestat
		[6] = 'http://rlopes.realestat
		[15] = 'http://rlopes.realestat
	----------------------------------------------------------------------------------------------------
	ff rlopes  0,55s user 0,04s system 75% cpu 0,199 total
```


-----------------------

The `ack` running looking for word `rlopes`
	
```
	$ time ack rlopes
	resources/css/toastr.css
	144:.rlopes { lorem ipsum }

	services.dev.ini
	3:domain = ".rlopes.realestateid.fixe"
	6:baseUri        = 'http://rlopes.realestateid.fixe'
	15:baseUri        = 'http://rlopes.realestateid.fixe'
    ack rlopes  0,29s user 0,03s system 99% cpu 0,322 total
```


-----------------------

The `grep` running looking for word `rlopes`

```
    $ time grep -i -r 'rlopes' ./*                                                                                                                                                                                            master  ✗ ✭
    ./jaws/jaws/configuration.dev.ini:site_url    = "http://rlopes.realestateid.fixe"
    ./jaws/jaws/configuration.dev.ini:base_uri    = "rlopes.realestateid.fixe/pixelandia/"
    ./jaws/jaws/services.dev.ini:domain = ".rlopes.realestateid.fixe"
    ./jaws/jaws/services.dev.ini:baseUri        = 'http://rlopes.realestateid.fixe'
    ./jaws/jaws/services.dev.ini:baseUri        = 'http://rlopes.realestateid.fixe'
    ./jaws/jaws/modules_customization/crm/config.dev.ini:base_uri = "http://rlopes.realestateid.fixe/images/"
    ./jaws/jaws/realestate/service/attribute/AttributeService.php: * User: rlopes
    grep --color=auto --exclude-dir={.bzr,CVS,.git,.hg,.svn} -i -r 'rlopes' ./  4,26s user 0,05s system 99% cpu 4,328 total
```


-----------------------
