package dolores_corecode

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"os/exec"
)

func Exec(command string, commandArgs ...string) (out_string string, err error) {
	var cmd *exec.Cmd
	if command == "" {
		log.Println("[ERROR] Exec got no command to run")
		err = errors.New("Exec got no command to run")
	} else if len(commandArgs) == 0 {
		cmd = exec.Command(command)
	} else {
		cmd = exec.Command(command, commandArgs[:]...)
	}

	var out bytes.Buffer
	cmd.Stdout = &out
	err = cmd.Run()
	out_string = fmt.Sprintf("%s", out)
	return
}
