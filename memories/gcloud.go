package dolores_memories

import (
	"log"
	"os"

	dolores_corecode "github.com/OpenChaos/dolores/corecode"
)

func GcloudComputeInstances() {
	log.Println("[memory] remember GCloud instances")
	working_dir, _ := os.Getwd()
	parameters := []string{
		dolores_corecode.OverrideFromEnvVar("DOLORES_SHELL_OUTPUT_DIR", working_dir),
	}
	output, err := dolores_corecode.Exec("gcloud-compute-instances", parameters[0:]...)
	log.Println(output)
	if err != nil {
		log.Println("[ERROR]", err.Error())
	}
}
