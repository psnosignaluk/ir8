package utils

import (
	"os"
)

const CurrentVersion = "0.0.1"

func AppVersion() string {
	return CurrentVersion
}

func IsDebugEnabled() (bool, string) {
	debugValue, _ := os.LookupEnv("IR8_DEBUG")

	switch debugValue {
	case "false", "0", "no", "":
		return false, debugValue
	default:
		return true, debugValue
	}
}
