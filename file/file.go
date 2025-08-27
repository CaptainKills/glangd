package file

import (
	"bufio"
	"log"
	"os"

	cmd "github.com/CaptainKills/glangd/cmd"
)

func ReadFile(fileName string) []string {
	inputFile, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("Could not read file %s! %q\n", fileName, err)
	}
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)
	if err := scanner.Err(); err != nil {
		log.Fatalf("Could not scan file! %q\n", err)
	}

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

func WriteFile(fileName string, commands []cmd.CompileCommand) {
	outputFile, err := os.Create(fileName)
	if err != nil {
		log.Fatalf("Could not create file! %q\n", err)
	}
	defer outputFile.Close()

	outputFile.WriteString("[\n")

	for index, cmd := range commands {
		if index > 0 {
			outputFile.WriteString(",\n")
		}
		outputFile.WriteString(cmd.ToJson())
	}

	outputFile.WriteString("\n]\n")
}
