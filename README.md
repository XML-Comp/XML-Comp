[![Go Report Card](https://goreportcard.com/badge/github.com/arxdsilva/XML-Comp)](https://goreportcard.com/report/github.com/arxdsilva/XML-Comp)
[![codebeat badge](https://codebeat.co/badges/1600adbb-27a3-4c3b-803e-818e1834b51a)](https://codebeat.co/projects/github-com-arxdsilva-xml-comp)

## XML-Comparer
Made this software to help [RimWorld](http://rimworldgame.com/) [community translators](https://github.com/ludeon)(1) to know what was modified on the last XML updates and to let them keep in track of what they need to add/remove from what has been done.

(1) and maybe other indie games that uses XML

### How this works? - RimWorld translator
You need two paths that we call pathA & pathB, which are described bellow:
- pathA: Full path directory of your RimWorld English folder

My path - as an example: pathA = **/Users/arthur/Library/Application Support/Steam/steamapps/common/RimWorld/RimWorldMac.app/Mods/Core/Languages/English**
- pathB: Full path directory of your RimWorld ~Language~ folder cloned from [GitHub](https://github.com/ludeon)

Me again: pathB = **/Users/arthur/Github/RimWorld-PortugueseBrazilian**

With these paths in hand you run our program and It will let you know in a `missingSomethieng.txt` file what is missing and where in your translation! That simple!

#### Dev Status
Now It only compares containing folders in given pathA to pathB and creates a "missingFolder.txt" file into the pathB directory.

#### XML-Comp CLI
- Creating an independent binary:
```shell
$ go build -o $GOPATH/bin/xml-comp
$ xml-comp help
```

- With executable binary:
```shell
$ go build -o xml-comp
$ ./xml-comp help
```

#### To Do
- Enter folders on pathA & pathB
- Store files each path on separated variables
- Compare stored Files on 'same' folder
- Create Same folder 'missingFiles.txt'
- Read 'equal' files on same folder
- Create 'missingTags.txt' starting with <--Filename Missing Tags-->
- Compare fileB to fileA to get old tags
- Include in 'missingTags.txt' not matching tags starting with <--Filename Not Matching Tags-->
- Turn this project CLI viable
