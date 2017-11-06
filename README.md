[![LICENSE](https://img.shields.io/badge/license-MIT-orange.svg)](LICENSE)
[![GoDoc](https://godoc.org/github.com/XML-Comp/XML-Comp?status.png)](https://godoc.org/github.com/XML-Comp/XML-Comp)
[![Go Report Card](https://goreportcard.com/badge/github.com/XML-Comp/XML-Comp)](https://goreportcard.com/report/github.com/XML-Comp/XML-Comp)
[![codebeat badge](https://codebeat.co/badges/1600adbb-27a3-4c3b-803e-818e1834b51a)](https://codebeat.co/projects/github-com-xml-comp-xml-comp)


[![Top Level Coverage](https://coveralls.io/repos/github/XML-Comp/XML-Comp/badge.svg?branch=master)](https://coveralls.io/github/XML-Comp/XML-Comp?branch=master)
[![Travis Build Status](https://api.travis-ci.org/XML-Comp/XML-Comp.svg?branch=master)](https://travis-ci.org/XML-Comp/XML-Comp)
[![DOI](https://zenodo.org/badge/71943139.svg)](https://zenodo.org/badge/latestdoi/71943139)


[![Donate](https://www.paypalobjects.com/pt_BR/BR/i/btn/btn_donateCC_LG.gif)](https://www.paypal.com/cgi-bin/webscr?cmd=_donations&business=arxdsilva%40gmail%2ecom&lc=BR&item_name=xml%2dcomp&currency_code=USD&bn=PP%2dDonationsBF%3abtn_donateCC_LG%2egif%3aNonHosted)


## Menu
* [What is XML-Comp?](https://github.com/xml-comp/xml-comp#what-is-XML-Comp)
* [Features](https://github.com/xml-comp/xml-comp#features)
* [Installing](https://github.com/xml-comp/xml-comp#installing)
* [Running](https://github.com/xml-comp/xml-comp#running)
* [How this works?](https://github.com/xml-comp/xml-comp#how-this-works)
* [Comparing any kind of document](https://github.com/xml-comp/xml-comp#comparing-any-kind-of-document)
* [Contributing](https://github.com/xml-comp/xml-comp#contributing)
* [To Do](https://github.com/xml-comp/xml-comp#to-do---check-our-issues--milestones)
* [Using only the comparer package](https://github.com/xml-comp/xml-comp#using-only-the-comparer-package)

## What is XML-Comp
This is a command line tool and a package that together they provide the capability of comparing two directories and appending to files the differences between the directories, also creates possible files or folders that are missing. It was made to help [RimWorld](http://rimworldgame.com/)'s [community translators](https://github.com/ludeon)(1) to know what was modified on the last XML updates and to let them keep in track of what they need to add/remove from what has been done.

(1) and maybe other indie games that uses XML

### Features
- [x] [Compare two given directories](https://github.com/XML-Comp/XML-Comp/issues/7)
- [x] [Append missing tags to the respective file](https://github.com/XML-Comp/XML-Comp/issues/8)
- [x] [Creates missing files on the compared directory](https://github.com/XML-Comp/XML-Comp/issues/9)
- [x] [Creates missing folders on the compared directory](https://github.com/XML-Comp/XML-Comp/issues/32)
- [x] [Compare one directory with multiple directories](https://github.com/XML-Comp/XML-Comp/issues/48)
- [x] [Compare automatically the game's english version with the given translation](https://github.com/XML-Comp/XML-Comp/issues?q=is%3Aissue+is%3Aclosed)
- [ ] Translate tag content
- [ ] Expose untracked files
- [ ] Expose untracked directories
- [ ] Expose untracked tags

### Installing
```
$ go get github.com/XML-Comp/XML-Comp
```

### Running
```shell
$ XML-Comp -translation /path/to/language/translation
```

### How this works?
You need the path that is called "translation", which are described bellow:
- **`"translation"`**: Full path directory of your RimWorld ~Language~ folder cloned from [GitHub](https://github.com/ludeon).

My "translation" path: **`/Users/arthur/Github/RimWorld-PortugueseBrazilian`**

With this path in hand, running `xml-comp -translation your/path/to/translation` will let you know in every file of your project what is missing by adding lines to it with what is needed to translate! That simple!

#### [RIMWORLD not installed in standard path]
If by any reason you did not install the game on Steam's standard path or want to use a different one, It's possible to use the `original` flag that exposes your customized path to the game as shows: 

- **`"original"`**: Full path directory of your RimWorld English folder (optional).
My "original" path (optional): **`/Users/arthur/Library/Application Support/Steam/steamapps/common/RimWorld/RimWorldMac.app/Mods/Core/Languages/English`**

In order to use this flag you'll need to parse it with the `translation` flag as in: `xml-comp -original path/... -translation path/...`

#### Comparing any kind of document
To compare any kind of files, all you need is to use the flag `-doc <type name>`, eg `-doc html`. This will use the paths that you gave only to compare the specified type of document. Another example:

```shell
$ XML-Comp -doc html -original path/to/It -translation path/to/It
```

OBS: This is not required, by default It's comparing all `.xml` files that are encountered.

#### [Contributing](https://github.com/XML-Comp/XML-Comp/blob/master/Contributing.md)

#### [Join our Gitter](https://gitter.im/XML-Comparer/Lobby)
#### To Do - Check our [Issues](https://github.com/XML-Comp/XML-Comp/issues) & [Milestones](https://github.com/XML-Comp/XML-Comp/milestones)

## Using only the comparer package
```go
// Import the package
import "github.com/XML-Comp/XML-Comp/comparer"
// Set document type variable to the desired document
// without the "." | eg: "xml" or "html"
comparer.DocType = "html"
// Start the main function with the full paths to compare
// the firstPath is always what will be used as model
func main() {
    err := comparer.Compare(firstPath, comparingPath)
    if err != nil {
        log.Fatal(err)
    }
}
```
