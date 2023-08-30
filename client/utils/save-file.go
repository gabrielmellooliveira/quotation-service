package utils

import (
	"bufio"
	"io"
	"os"
)

func WriteInFile(path string, content string) {
    _, err := os.Stat(path);

    if err != nil {
        _, err := os.Create(path)

        if err != nil {
            panic(err)
        }
    }

	err = insertStringToFile(path, content)

	if err != nil {
		panic(err)
	}
}

func insertStringToFile(path, str string) error {
	lines, err := file2lines(path)

	if err != nil {
		return err
	}

	fileContent := ""
	for _, line := range lines {
		fileContent += line
		fileContent += "\n"
	}

    fileContent += str
    fileContent += "\n"

	return os.WriteFile(path, []byte(fileContent), 0644)
}

func file2lines(filePath string) ([]string, error) {
	f, err := os.Open(filePath)

	if err != nil {
		return nil, err
	}

	defer f.Close()

	return linesFromReader(f)
}

func linesFromReader(r io.Reader) ([]string, error) {
	var lines []string
	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}