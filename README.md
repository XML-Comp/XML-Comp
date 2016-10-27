## XML-Comparer
Made this software to help [RimWorld](http://rimworldgame.com/) [community translators](https://github.com/ludeon)(1) to know what was modified on the last XML updates and to let them keep in track of what they need to add/remove from what has been done.

(1) and maybe other indie games that uses XML

### Installing
```go get github.com/ArxdSilva/XML-Comp```

### Running
```xml-comp -original /path/to/language/english -translation /path/to/language/translation```

### How this works? - RimWorld translator
You need two paths that we call "original" & "translation", which are described bellow:
- "original": Full path directory of your RimWorld English folder

My path - as an example: "original" = **/Users/arthur/Library/Application Support/Steam/steamapps/common/RimWorld/RimWorldMac.app/Mods/Core/Languages/English**
- "translation": Full path directory of your RimWorld ~Language~ folder cloned from [GitHub](https://github.com/ludeon)

Me again: "translation" = **/Users/arthur/Github/RimWorld-PortugueseBrazilian**

With these paths in hand you run our program and It will let you know in a `missingSomethieng.txt` file what is missing and where in your translation! That simple!

#### Dev Status
Now It only compares containing folders in given "original" to "translation" and creates a "missingFolder.txt" file into the "translation" directory.

#### To Do
- Enter folders on "original" & "translation"
- Store files each path on separated variables
- Compare stored Files on 'same' folder
- Create Same folder 'missingFiles.txt'
- Read 'equal' files on same folder
- Create 'missingTags.txt' starting with <--Filename Missing Tags-->
- Compare fileB to fileA to get old tags
- Include in 'missingTags.txt' not matching tags starting with <--Filename Not Matching Tags-->
- Turn this project CLI viable
