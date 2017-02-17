package dolores_drives

import (
	dolores_corecode "github.com/OpenChaos/dolores/corecode"
)

func ServerListFor(serverKeyword ...string) (body string, err error) {
	body, err = dolores_corecode.Exec("give-server-list", serverKeyword[0:]...)
	return
}
