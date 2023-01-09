package configs

import (
	"os"
	model "portfolio-api/internal/models"
	"strconv"
)

// FiberConfig func for configuration Fiber app.
// See: https://docs.gofiber.io/api/fiber#config
func Init() (env model.Environment) {

	var model = new(model.Environment)

	serverName := os.Getenv("BASE_SERVER_NAME")
	if serverName == "" {
		serverName = "SDK Service"
	}
	model.BASE_SERVER_NAME = serverName

	readTimeout := os.Getenv("BASE_SERVER_READ_TIMEOUT")
	if readTimeout == "" {
		readTimeout = "60"
	}

	serverReadTimeoutInt, err := strconv.Atoi(readTimeout)
	if err != nil {
		Disconnect("Failed casting timeout context: ", err)
	}

	model.BASE_SERVER_READ_TIMEOUT = serverReadTimeoutInt

	serverUrl := os.Getenv("BASE_SERVER_URL")
	if serverUrl == "" {
		Disconnect("'BASE_SERVER_URL' is required, Add it in environment file to skip this error. Please read 'Readme.md' for more information")
	}
	model.BASE_SERVER_URL = serverUrl

	dbServerUrl := os.Getenv("BASE_DB_SERVER_URL")
	if dbServerUrl == "" {
		Disconnect("'BASE_DB_SERVER_URL' is required, Add it in environment file to skip this error. Please read 'Readme.md' for more information")
	}
	model.BASE_DB_SERVER_URL = dbServerUrl

	return *model
}
