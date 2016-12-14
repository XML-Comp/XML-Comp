[![LICENSE](https://img.shields.io/badge/license-MIT-orange.svg)](LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/XML-Comp/XML-Comp)](https://goreportcard.com/report/github.com/XML-Comp/XML-Comp)
[![codebeat badge](https://codebeat.co/badges/1600adbb-27a3-4c3b-803e-818e1834b51a)](https://codebeat.co/projects/github-com-XML-Comp-xml-comp)
[![GoDoc](https://godoc.org/github.com/XML-Comp/xml-comp?status.png)](https://godoc.org/github.com/XML-Comp/xml-comp)
[![Top Level Coverage](https://coveralls.io/repos/github/XML-Comp/XML-Comp/badge.svg?branch=master)](https://coveralls.io/github/XML-Comp/XML-Comp?branch=master) / [![comparer coverage](https://gocover.io/_badge/github.com/XML-Comp/xml-comp/comparer?0 "comparer coverage")](http://gocover.io/github.com/XML-Comp/XML-Comp/comparer)
[![Travis Build Status](https://api.travis-ci.org/XML-Comp/XML-Comp.svg?branch=master)](https://travis-ci.com/XML-Comp/xml-comp)

## XML-Comparer
This package is a command line tool that provides the capability of comparing two directories and outputs files with differences between them, including: missing Files, missing Folders and missing Tags of .xml files. It was made to help [RimWorld](http://rimworldgame.com/)'s [community translators](https://github.com/ludeon)(1) to know what was modified on the last XML updates and to let them keep in track of what they need to add/remove from what has been done.

(1) and maybe other indie games that uses XML

### Installing
```
$ go get github.com/XML-Comp/XML-Comp
```

### Running
```shell
$ XML-Comp -comp -original /path/to/language/english -translation /path/to/language/translation
```

### How this works?
You need two paths that we call "original" & "translation", which are described bellow:
- **`"original"`**: Full path directory of your RimWorld English folder
- **`"translation"`**: Full path directory of your RimWorld ~Language~ folder cloned from [GitHub](https://github.com/ludeon)

My "original" path: **`/Users/arthur/Library/Application Support/Steam/steamapps/common/RimWorld/RimWorldMac.app/Mods/Core/Languages/English`**

My "translation" path: **`/Users/arthur/Github/RimWorld-PortugueseBrazilian`**

With these paths in hand you run our program and It will let you know in a `missingSomethieng.txt` file what is missing and where in your translation! That simple!

#### XML-Comp CLI Usage (needed Go 1.7+)
```shell
$ git clone github.com/XML-Comp/xml-comp
$ cd xml-comp
$ go install
$ xml-comp help
```
#### [Join our Gitter](https://gitter.im/XML-Comparer/Lobby)
#### To Do - Check our [Issues](https://github.com/XML-Comp/XML-Comp/issues) & [Milestones]()