package dolores_corecode

import (
	"bytes"
	"errors"
	"log"
	"os/exec"
)

func Exec(commandTokens []string) (err error) {
	token_count := len(commandTokens)
	var cmd *exec.Cmd
	if token_count == 0 {
		log.Println("[ERROR] Exec got no command to run")
		err = errors.New("Exec got no command to run")
	} else if token_count == 1 {
		cmd = exec.Command(commandTokens[0])
	} else {
		cmd = exec.Command(commandTokens[0], commandTokens[1:]...)
	}

	var out bytes.Buffer
	cmd.Stdout = &out
	err = cmd.Run()
	return
}
