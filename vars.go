package hypolashlckhttp

import (
	helpers "github.com/hypolas/hypolashlckhelpers"
	"github.com/hypolas/hypolaslogger"
	"os"
	"strconv"
	"strings"
	"time"
)

func makeLogger(logPath string) hypolaslogger.HypolasLogger {
	l := hypolaslogger.NewLogger(logPath)
	return l
}

var (
	logf = makeLogger(os.Getenv("HYPOLAS_LOGS_FOLDER"))
)

func taskLoadEnvironnement() {
	help := helpers.InitHlckCustom{}

	/*
	*	Http check variable
	 */
	healthcheckHTTPExpected = help.InitEnvVars("HYPOLAS_HEALTHCHECK_HTTP_EXPECTED", "")
	healthcheckHTTPJsonPath = help.InitEnvVars("HYPOLAS_HEALTHCHECK_HTTP_JSON", "")
	logf.Warn.Printf("%s\n", os.Getenv("HYPOLAS_HEALTHCHECK_HTTP_URL"))
	healthcheckHTTPUrl = help.InitEnvVars("HYPOLAS_HEALTHCHECK_HTTP_URL", "")
	logf.Warn.Printf("%s\n", healthcheckHTTPUrl)
	healthcheckHTTPProxy = help.InitEnvVars("HYPOLAS_HEALTHCHECK_HTTP_PROXY", "")
	healthcheckHTTPHeaders = help.InitEnvVars("HYPOLAS_HEALTHCHECK_HTTP_HEADERS", "")

	healthcheckHTTPTimeout, err = time.ParseDuration(help.InitEnvVars("HYPOLAS_HEALTHCHECK_HTTP_TIMEOUT", "0") + "s")
	if err != nil {
		logf.Err.Fatalln(err)
	}

	statusCode := strings.Split(help.InitEnvVars("HYPOLAS_HEALTHCHECK_HTTP_RESPONSES", ""), ",")
	if statusCode[0] != "" {
		healthcheckHTTPUseCode = true
		for _, status := range statusCode {
			code, err := strconv.Atoi(status)
			logf.Err.Fatalln(err)
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

	returnedValue string
	separator     = "__"
	isJSONEntry   = true

	err error
)
