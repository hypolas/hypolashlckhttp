package hypolashlckhttp

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	extractJSON "github.com/hypolas/readjsonfromflatpath"
)

// GetHTTP call URL and return result
func GetHTTP() (result string) {
	// Load and stransform environment variable
	taskLoadEnvironnement()

	clientHTTP := constructHTTPClient()
	logf.VarDebug(clientHTTP, "clientHTTP")

	logf.Warn.Printf("%s\n", healthcheckHTTPUrl)
	reqHTTP, err := http.NewRequest("GET", healthcheckHTTPUrl, nil)
	logf.VarDebug(reqHTTP, "reqHTTP")
	if err != nil {
		logf.Err.Println(err)
	}

	reqHTTP.Header.Add("Accept", `application/json`)

	additionnalHeaders := splitFlatten(healthcheckHTTPHeaders)
	for _, header := range additionnalHeaders {
		splitedHeader := strings.Split(header, ",")
		reqHTTP.Header.Add(splitedHeader[0], splitedHeader[1])
	}

	resp, err := clientHTTP.Do(reqHTTP)

	if err != nil {
		logf.Err.Fatalf("%s\n", err)
	}
	defer resp.Body.Close()

	bodyHTTP, err := ioutil.ReadAll(resp.Body)
	logf.VarDebug(bodyHTTP, "bodyHTTP")
	logf.Err.Println(err)

	/*
	*	If chek is on status html code the test stop here
	 */
	logf.VarDebug(healthcheckHTTPUseCode, "healthcheckHttpUseCode")
	if healthcheckHTTPUseCode {
		if intIsIn(resp.StatusCode, healthcheckHTTPResponse) {
			os.Exit(0)
		} else {
			os.Exit(1)
		}
	}

	/*
	*	If chek is on REST API, the json will be tested
	 */
	logf.VarDebug(healthcheckHTTPJsonPath, "healthcheckHttpJsonPath")
	if healthcheckHTTPJsonPath != "" {
		return extractJSON.ReadJSONFromFlatPath("", bodyHTTP)
	} else {
		return strings.Trim(string(bodyHTTP), "\"")
	}
}

/*
*	Construct client HTTP
 */
func constructHTTPClient() *http.Client {
	client := &http.Client{
		Transport: &http.Transport{},
		Timeout:   0,
	}

	if healthcheckHTTPProxy != "" {
		proxyURL, err := url.Parse(healthcheckHTTPProxy)
		logf.Err.Println(err)
		client.Transport = &http.Transport{
			Proxy: http.ProxyURL(proxyURL),
		}
	}

	if healthcheckHTTPTimeout != 0 {
		client.Timeout = healthcheckHTTPTimeout * time.Second
	}

	return client
}

func intIsIn(i int, arrayInt []int) bool {
	for _, inte := range arrayInt {
		if i == inte {
			return true
		}
	}
	return false
}

func splitFlatten(flatten string) []string {
	flatten = strings.TrimSpace(flatten)
	if !strings.Contains(flatten, "__") {
		return []string{}
	}
	return strings.Split(flatten, "__")
}
