### How to use:
You will need to create an instance of Data with your Paths, so It will be something similar to this:

`instance := comparer.Data{
  PathA: "C://Appdata/RimWorld/Example/Model/English",
  PathB: "C://Documents/Github/arxdsilva/RimWorld-Portuguese",
  }`

Now with this instance you can perform your first comparison in this folder just by doing:

`instance.CompareContainingFoldersAndFiles()`

This will look into these paths and create two files named: `missingFolders.txt` and `missingFiles.txt` that will show you which files that are contained in pathA that pathB does not have.
