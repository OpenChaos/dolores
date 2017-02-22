package dolores_drives

import (
	dolores_corecode "github.com/OpenChaos/dolores/corecode"
)

func Nslookup(args ...string) (reply string, err error) {
	reply, err = dolores_corecode.Exec("grep-gcloud-compute", args[0:]...)
	return
}
