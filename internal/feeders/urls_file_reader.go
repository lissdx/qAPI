package feeders

import (
	"bufio"
	"os"
)

type UrlSlice []string

// GetData TODO change the feeder func to: GetData(fileName string) (<-chan string, error)
func GetData(fileName string) ([]string, error) {
	result := make([]string, 0, 100)
	readFile, err := os.Open(fileName)

	if err != nil {
		return nil, err
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		result = append(result, fileScanner.Text())
	}

	return result, nil
}
