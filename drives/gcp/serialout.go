package dolores_gcp

import dolores_corecode "github.com/OpenChaos/dolores/corecode"

func GcloudSerialOut(args ...string) (reply string, err error) {
	reply, err = dolores_corecode.Exec("gcloud-serialout", args[0:]...)
	return
}
