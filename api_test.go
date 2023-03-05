package hypolashlckhttp

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

// TestHTTPApi get API et test result
func TestHTTPApi(t *testing.T) {
	test := Call()

	log.Info.Println("Result => ", result)
	if !test.IsUP {
		log.Err.Fatalln(test.Output)
	} else {
		log.Info.Println("Match ! It ok !!: ")
	}

	readFile, err := os.Open(log.LogFile.Name())

	if err != nil {
		panic(err)
	}

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		fmt.Println(fileScanner.Text())
	}

	readFile.Close()
}
