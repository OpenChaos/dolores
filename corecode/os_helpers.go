package dolores_corecode

import (
	"log"
	"os"
	"strings"
)

func OverrideFromEnvVar(envVar string, defaultValue string) string {
	if HasEnv(envVar) {
		envVarValue := os.Getenv(envVar)
		log.Printf("[env] %s: %s", envVar, envVarValue)
		return envVarValue
	}
	log.Printf("[env] %s: %s", envVar, defaultValue)
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
