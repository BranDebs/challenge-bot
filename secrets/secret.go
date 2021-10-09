package secrets

import (
	"bufio"
	"os"
)

func APIToken(filepath string) (string, error) {
	file, err := os.Open(filepath) // the file is inside the local directory
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if !scanner.Scan() {
		return "", scanner.Err()
	}

	return scanner.Text(), nil
}
