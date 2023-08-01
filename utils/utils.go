package utils

import "os"

func IsDebugEnabled() (bool, string) {
	debugValue, _ := os.LookupEnv("IR8_DEBUG")

	switch debugValue {
	case "false", "0", "no", "":
		return false, debugValue
	default:
		return true, debugValue
	}
}
