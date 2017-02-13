package dolores_corecode

import (
	"os"
	"strings"
)

func OverrideFromEnvVar(envVar string, defaultValue string) string {
	if HasEnv(envVar) {
		return os.Getenv(envVar)
	}
	return defaultValue
}

func HasEnv(envVar string) bool {
	for _, envKeyVal := range os.Environ() {
		if strings.Split(envKeyVal, "=")[0] == envVar {
			return true
		}
	}
	return false
}
