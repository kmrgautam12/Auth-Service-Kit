package main

import (
	utils "Pay-AI/financial-transaction-server/Utils"
	"Pay-AI/financial-transaction-server/restrouters"
)

var logger = utils.Logger

func main() {
	logger.Info("Starting handler")
	restrouters.Handler()
}
