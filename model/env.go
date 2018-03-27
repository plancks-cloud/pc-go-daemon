package model

import (
	"os"
)

//LogFormatEnv is the environment variable used to define the log format
const LogFormatEnv = "LOGGING_FORMAT"

//WalletEnv is the environment variable used to define the wallet to use
const WalletEnv = "WALLET"

//GetEnvLogFormat returns the environment variable set for logging format
func GetEnvLogFormat() string {
	return os.Getenv(LogFormatEnv)
}

//GetEnvWallet returns the environment variable set for the wallet
func GetEnvWallet() string {
	return os.Getenv(WalletEnv)
}
