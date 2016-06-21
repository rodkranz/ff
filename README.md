# Find File #

Find File or text in file with go language.

-----------------------
### Source Find File ###

* Find text or file name 
* Version: 1.1.2
* License: ISC


-----------------------

## How to Compile it

 Compile those files from different platforms

### Requirements

* [GO Language](https://golang.org/doc/install)

#### Compiling to *Linux*

	$ env GOOS=linux GOARCH=arm GOARM=7 go build -o ff main.go


#### Compiling to *MAcOSX*

	$ env GOOS=darwin GOARCH=386 go build -o ff main.go


#### Compiling to *Windows*

	$ env GOOS=windows GOARCH=386 go build -o ff.exe main.go


-----------------------

## Parameters

##### Helper: 

```
    Usage of ./ff:
      --exclude-dir string
            Exclude dir from reader (default ".bzr,CVS,.git,.hg,.svn")
      --no-color
            Disable color output
      -cis
            Search text case insensitive
      -cpu int
            Number of CPU you have 4 available (default 4)
      -d string
            Directory ffConfig (default "./")
      -f string
            Filter by file name
      -force
            Replace all result without ask.
      -reg
            Search by this Regex
      -replace string
            Replace result to text
      -t string
            Text ffConfig
      -up
            Check update
      -ver
            Show the version
```

-----------------------

## Usage

Without parameters will show everything in subfolders.

```
    $ ff
    [dir] ./g please wait... 
    [dir] admin
    [file] admin/config.dev.ini
    [file] admin/config.ini
    [dir] admin/controller
    [file] admin/controller/AjaxAdminController.php
    [file] admin/controller/AjaxAttributesController.php
    [file] admin/controller/AjaxAuthenticationController.php

	.....
```

With one parameter will find for text.
```
	$ ff rodrigo 
    [file] admin/controller/AjaxAdminController.php (lines: 1)
    [7] 	 * @author    Rodrigo Lopes <rodrigo.lopes@fixeads.com> 
    
    [file] admin/controller/AjaxAttributesController.php (lines: 1)
    [7] 	 * @author    Rodrigo Lopes <rodrigo.lopes@fixeads.com> 
    
    [file] admin/controller/AjaxCompaniesController.php (lines: 1)
    [7] 	 * @author Rodrigo Lopes <rodrigo.lopes@fixeds.com> 
    .....
```

-----------------------

### Parameter -f

Parameter `-f` of file, this parameter you will define the name of file  that you should like to filter

```
	$ ff -f css                                                                                                                                                                                                                   develop  ✭
    [dir] admin/resources/css
    [file] admin/resources/css/adminLogin.css
    [file] admin/resources/css/adminTheme.css
    [file] admin/resources/css/angular-bootstrap-datepicker.css
    [file] admin/resources/css/animate.css
	....
```

-----------------------
### Parameter -t

Parameter `-t` of text, this text that you will search inside of file.
	
```
	$ ff -t rodrigo
    [file] admin/controller/AjaxAdminController.php (lines: 1)
    [7] 	 * @author    Rodrigo Lopes <rodrigo.lopes@fixeads.com> 
    
    [file] admin/controller/AjaxAttributesController.php (lines: 1)
    [7] 	 * @author    Rodrigo Lopes <rodrigo.lopes@fixeads.com> 
    
    [file] admin/controller/AjaxCompaniesController.php (lines: 1)
    [7] 	 * @author Rodrigo Lopes <rodrigo.lopes@fixeds.com> 
    ...
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

### Parameter -update

Check if exist an update available of sistem.

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
