package plsfile

type PlsFile struct {
	Jobs map[string]PlsFileJob
}

type PlsFileJob struct {
	Commands []string
}
