package comparer

type Data struct {
	PathA, PathB  string
	dirA          []string
	dirB          []string
	fileA         []string
	fileB         []string
	missingFolder []string
	missingFiles  []string
}
