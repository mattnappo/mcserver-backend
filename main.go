package main

import (
	"flag"

	"github.com/xoreo/mcserver-backend/api"
	"github.com/xoreo/mcserver-backend/common"
)

// APIServerPortFlag is the flag for the port on which the API server should run.
var APIServerPortFlag = flag.Int("api-port", 8000, "The API server's port.")

func main() {
	logger := common.NewLogger("main")
	flag.Parse()

	// Start an api server
	err := api.StartAPIServer(*APIServerPortFlag)
	if err != nil {
		logger.Criticalf(err.Error())
	}
	logger.Infof("API server started")
}
