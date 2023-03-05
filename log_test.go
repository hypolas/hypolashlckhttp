package hypolashlckhttp

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

// TestHTTPApi get API et test result
func TestHTTPApi(t *testing.T) {
	result := GetHTTP()

	logf.Info.Println("Result => ", result)
	if result != os.Getenv("HYPOLAS_HEALTHCHECK_HTTP_EXPECTED") {
		logf.Err.Fatalln("result != os.Getenv(\"HYPOLAS_HEALTHCHECK_HTTP_EXPECTED\")")
	} else {
		logf.Info.Println("Match ! It ok !!: ", result)
	}

	readFile, err := os.Open(logf.LogFile.Name())

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
