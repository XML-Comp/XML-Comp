[![LICENSE](https://img.shields.io/badge/license-MIT-orange.svg)](LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/XML-Comp/XML-Comp)](https://goreportcard.com/report/github.com/XML-Comp/XML-Comp)
[![codebeat badge](https://codebeat.co/badges/1600adbb-27a3-4c3b-803e-818e1834b51a)](https://codebeat.co/projects/github-com-xml-comp-xml-comp)
[![GoDoc](https://godoc.org/github.com/XML-Comp/XML-Comp?status.png)](https://godoc.org/github.com/XML-Comp/XML-Comp)
[![Top Level Coverage](https://coveralls.io/repos/github/XML-Comp/XML-Comp/badge.svg?branch=master)](https://coveralls.io/github/XML-Comp/XML-Comp?branch=master)
[![Travis Build Status](https://api.travis-ci.org/XML-Comp/XML-Comp.svg?branch=master)](https://travis-ci.org/XML-Comp/XML-Comp)
[![DOI](https://zenodo.org/badge/71943139.svg)](https://zenodo.org/badge/latestdoi/71943139)
[![Donate](https://www.paypalobjects.com/pt_BR/BR/i/btn/btn_donateCC_LG.gif)](https://www.paypal.com/cgi-bin/webscr?cmd=_donations&business=arxdsilva%40gmail%2ecom&lc=BR&item_name=xml%2dcomp&currency_code=USD&bn=PP%2dDonationsBF%3abtn_donateCC_LG%2egif%3aNonHosted)


## Menu
* [GoDoc](https://godoc.org/github.com/XML-Comp/XML-Comp/comparer)
* [XML-Comparer](https://github.com/xml-comp/xml-comp#xml-comparer)
* [Installing](https://github.com/xml-comp/xml-comp#installing)
* [Running](https://github.com/xml-comp/xml-comp#running)
* [How this works?](https://github.com/xml-comp/xml-comp#how-this-works)
* [Comparing any kind of document](https://github.com/xml-comp/xml-comp#)
* [XML-Comp CLI Usage](https://github.com/xml-comp/xml-comp#xml-comp-cli-usage-needed-go-17)
* [Contributing](https://github.com/xml-comp/xml-comp#contributing)
* [To Do](https://github.com/xml-comp/xml-comp#to-do---check-our-issues--milestones)
* [Using only the comparer package](https://github.com/xml-comp/xml-comp#using-only-the-comparer-package)

## XML-Comparer
This package is a command line tool that provides the capability of comparing two directories adding to files the differences between them, creating the possible files or folders that are missing and adding the missing tags on each file. It was made to help [RimWorld](http://rimworldgame.com/)'s [community translators](https://github.com/ludeon)(1) to know what was modified on the last XML updates and to let them keep in track of what they need to add/remove from what has been done.

(1) and maybe other indie games that uses XML

### Installing
```
$ go get github.com/XML-Comp/XML-Comp
```

### Running
```shell
$ XML-Comp -translation /path/to/language/translation
```

### How this works?
You need two paths that we call "original" (optional) & "translation", which are described bellow:
- **`"original"`**: Full path directory of your RimWorld English folder (optional).
- **`"translation"`**: Full path directory of your RimWorld ~Language~ folder cloned from [GitHub](https://github.com/ludeon).

My "original" path (optional): **`/Users/arthur/Library/Application Support/Steam/steamapps/common/RimWorld/RimWorldMac.app/Mods/Core/Languages/English`**

My "translation" path: **`/Users/arthur/Github/RimWorld-PortugueseBrazilian`**

With these paths in hand, running them with `xml-comp` program will let you know in every file of your project what is missing by adding lines to it with what is needed to translate! That simple!

#### Comparing any kind of document:
To compare any kind of files, all you need is to use the flag `-doc <type name>`, eg `-doc html`. This will use the paths that you gave only to compare the specified type of document. Another example:

```shell
$ XML-Comp -doc html -original path/to/It -translation path/to/It
```

OBS: This is not required, by default It's comparing all `.xml` files that are encountered.

#### XML-Comp CLI Usage (needed Go 1.7+)
```shell
$ git clone github.com/XML-Comp/xml-comp
$ cd xml-comp
$ go install
$ xml-comp help
```
#### [Contributing](https://github.com/XML-Comp/XML-Comp/blob/master/Contributing.md)

#### [Join our Gitter](https://gitter.im/XML-Comparer/Lobby)
#### To Do - Check our [Issues](https://github.com/XML-Comp/XML-Comp/issues) & [Milestones]()

## Using only the comparer package
1- Import the package
```go
import "github.com/XML-Comp/XML-Comp/comparer"
```
2- Set document type variable to the desired document
```go
// without the "." | eg: "xml" or "html"
comparer.DocType = "desired docType"
```
3- Start the main function with the full paths to compare
```go
// the firstPath is always what will be used as model
comparer.Compare(firstPath, comparingPath)
```
