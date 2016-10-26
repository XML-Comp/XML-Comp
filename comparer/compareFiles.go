package comparer

func CompareFiles(pathA, pathB string) error {
	if (pathA == "") || (pathB == "") {
		return fmt.Errorf("Empty path")
	}
	missing, err := missingFiles(pathA, pathB)
}

func missingFiles(A, B string) error {

}
