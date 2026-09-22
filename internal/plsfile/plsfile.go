package plsfile

import (
	"bytes"
	"errors"
	"os"
	"strings"
)

type PlsFile struct {
	Jobs map[string]PlsFileJob
}

type PlsFileJob struct {
	Commands []string
}

func ReadFile(path string) (PlsFile, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return PlsFile{}, err
	}
	res, err := parseData(data)
	if err != nil {
		return PlsFile{}, err
	}
	return res, nil
}

func removeCommentFromLine(line string) string {
	index := strings.Index(line, "//")
	if index == -1 {
		return line
	}
	line = line[:index]
	return line
}

func parseData(data []byte) (PlsFile, error) {
	lines := bytes.Split(bytes.ReplaceAll(data, []byte("\r\n"), []byte("\n")), []byte("\n"))

	result := PlsFile{
		Jobs: make(map[string]PlsFileJob),
	}

	ccmdName := ""
	ccmds := []string{}
	for _, line := range lines {
		l := strings.TrimSpace(string(line))

		// TODO: remove comments if there are any

		if l == "" {
			continue
		}

		// starts a new command
		if strings.HasPrefix(l, "#:") {
			cmdName := strings.TrimPrefix(l, "#:")
			cmdName = strings.TrimSpace(cmdName)

			if cmdName == "" {
				return PlsFile{}, errors.New("commands without name are not allowed!")
			}

			// add the current cmd to the result
			if ccmdName != "" {
				result.Jobs[ccmdName] = PlsFileJob{
					Commands: ccmds,
				}
				ccmdName = ""
				ccmds = []string{}
			}

			ccmdName = cmdName
			continue
		}

		// read the current command line and add it
		ccmds = append(ccmds, l)
	}
	if ccmdName != "" {
		result.Jobs[ccmdName] = PlsFileJob{
			Commands: ccmds,
		}
	}
	return result, nil
}
