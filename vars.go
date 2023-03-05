package hypolashlckhttp

import (
	"os"
	"strconv"
	"strings"
	"time"

	helpers "github.com/hypolas/hypolashlckhelpers"

)

var (
	log    = helpers.NewLogger()
	result = helpers.NewResult()
)

func taskLoadEnvironnement() {
	/*
	*	Http check variable
	 */
	healthcheckHTTPExpected = helpers.NewEnvVars("HYPOLAS_HEALTHCHECK_HTTP_EXPECTED", "")
	healthcheckHTTPJsonPath = helpers.NewEnvVars("HYPOLAS_HEALTHCHECK_HTTP_JSON", "")
	healthcheckHTTPUrl = helpers.NewEnvVars("HYPOLAS_HEALTHCHECK_HTTP_URL", "")
	healthcheckHTTPProxy = helpers.NewEnvVars("HYPOLAS_HEALTHCHECK_HTTP_PROXY", "")
	healthcheckHTTPHeaders = helpers.NewEnvVars("HYPOLAS_HEALTHCHECK_HTTP_HEADERS", "")

	healthcheckHTTPTimeout, err = time.ParseDuration(helpers.NewEnvVars("HYPOLAS_HEALTHCHECK_HTTP_TIMEOUT", "0") + "s")
	if err != nil {
		log.Err.Fatalln(err)
	}

	statusCode := strings.Split(helpers.NewEnvVars("HYPOLAS_HEALTHCHECK_HTTP_RESPONSES", ""), ",")
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
