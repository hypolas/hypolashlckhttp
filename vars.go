package hypolashlckhttp

import (
	helpers "github.com/hypolas/hypolashlckhelpers"
	"os"
	"strconv"
	"strings"
	"time"
)

var (
	log    = helpers.NewLogger()
	result = helpers.NewResult()
)

func taskLoadEnvironnement() {
	help := helpers.InitHlckCustom{}
	/*
	*	Http check variable
	 */
	healthcheckHTTPExpected = help.NewEnvVars("HYPOLAS_HEALTHCHECK_HTTP_EXPECTED", "")
	healthcheckHTTPJsonPath = help.NewEnvVars("HYPOLAS_HEALTHCHECK_HTTP_JSON", "")
	log.Info.Println("Titi")
	log.Warn.Printf("%s\n", os.Getenv("HYPOLAS_HEALTHCHECK_HTTP_URL"))
	healthcheckHTTPUrl = help.NewEnvVars("HYPOLAS_HEALTHCHECK_HTTP_URL", "")
	log.Warn.Printf("%s\n", healthcheckHTTPUrl)
	healthcheckHTTPProxy = help.NewEnvVars("HYPOLAS_HEALTHCHECK_HTTP_PROXY", "")
	healthcheckHTTPHeaders = help.NewEnvVars("HYPOLAS_HEALTHCHECK_HTTP_HEADERS", "")

	healthcheckHTTPTimeout, err = time.ParseDuration(help.NewEnvVars("HYPOLAS_HEALTHCHECK_HTTP_TIMEOUT", "0") + "s")
	if err != nil {
		log.Err.Fatalln(err)
	}

	statusCode := strings.Split(help.NewEnvVars("HYPOLAS_HEALTHCHECK_HTTP_RESPONSES", ""), ",")
	if statusCode[0] != "" {
		healthcheckHTTPUseCode = true
		for _, status := range statusCode {
			code, err := strconv.Atoi(status)
			log.Err.Fatalln(err)
			healthcheckHTTPResponse = append(healthcheckHTTPResponse, code)
		}
	}
}

var (
	// Expected http value
	healthcheckHTTPExpected string

	// JsonPath Flatter with double _
	healthcheckHTTPJsonPath string

	// URL to check
	healthcheckHTTPUrl string

	// Proxy if needed
	healthcheckHTTPProxy string

	// Add header if needed
	healthcheckHTTPHeaders string

	// Use return code ?
	healthcheckHTTPUseCode bool

	// Define HTTP Timeout
	healthcheckHTTPTimeout time.Duration

	// Check HTTP Status Code
	healthcheckHTTPResponse []int

	// Logs folder
	healthcheckLogsFolder string

	separator   = "__"
	isJSONEntry = true

	err error
)
