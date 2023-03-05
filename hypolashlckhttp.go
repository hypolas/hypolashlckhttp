package hypolashlckhttp

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	helpers "github.com/hypolas/hypolashlckhelpers"
	extractJSON "github.com/hypolas/readjsonfromflatpath"
)

// Call call URL and return result
func Call() helpers.Result {
	// Load and stransform environment variable
	taskLoadEnvironnement()

	clientHTTP := constructHTTPClient()
	log.VarDebug(clientHTTP, "clientHTTP")

	log.Warn.Printf("%s\n", healthcheckHTTPUrl)
	reqHTTP, err := http.NewRequest("GET", healthcheckHTTPUrl, nil)
	log.VarDebug(reqHTTP, "reqHTTP")
	if err != nil {
		log.Err.Println(err)
	}

	reqHTTP.Header.Add("Accept", `application/json`)

	additionnalHeaders := splitFlatten(healthcheckHTTPHeaders)
	for _, header := range additionnalHeaders {
		splitedHeader := strings.Split(header, ",")
		reqHTTP.Header.Add(splitedHeader[0], splitedHeader[1])
	}

	resp, err := clientHTTP.Do(reqHTTP)

	if err != nil {
		log.Err.Fatalf("%s\n", err)
	}
	defer resp.Body.Close()

	bodyHTTP, err := ioutil.ReadAll(resp.Body)
	result.Output = string(bodyHTTP)

	log.VarDebug(bodyHTTP, "bodyHTTP")
	log.Err.Println(err)

	/*
	*	If chek is on status html code the test stop here
	 */
	log.VarDebug(healthcheckHTTPUseCode, "healthcheckHttpUseCode")
	if healthcheckHTTPUseCode {
		if intIsIn(resp.StatusCode, healthcheckHTTPResponse) {
			result.IsUP = true
			return result
		}
		return result
	}

	/*
	*	If chek is on REST API, the json will be tested
	 */
	log.VarDebug(healthcheckHTTPJsonPath, "healthcheckHttpJsonPath")
	if healthcheckHTTPJsonPath != "" {
		if healthcheckHTTPExpected == extractJSON.ReadJSONFromFlatPath("", bodyHTTP) {
			result.IsUP = true
			return result
		}
	} else {
		if healthcheckHTTPExpected == strings.Trim(string(bodyHTTP), "\"") {
			result.IsUP = true
			return result
		}
	}
	return result
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
		log.Err.Println(err)
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
