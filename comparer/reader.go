package comparer

// read recieves a path and a file and returns It's tags & an error
func read(file, path string) ([]string, error) {
	var tags []string
	// do your magic
	return tags, nil
}

func readFiles(files []string, pathA, pathB string) error {
	for _, v := range files {
		tagsA, err := read(v, pathA)
		if err != nil {
			return err
		}
		tagsB, err := read(v, pathA)
		if err != nil {
			return err
		}
		missingTags, _ := findMissing(tagsA, tagsB)
		mTags := MissingFile{
			name:   "MissingTags.txt",
			path:   pathB,
			prefix: v,
		}
		err = mTags.fileCreator(missingTags)
		if err != nil {
			return err
		}
	}
	return nil
}
