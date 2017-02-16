package dolores_drives

import dolores_corecode "github.com/OpenChaos/dolores/corecode"

func GiveSshAccess(parameters ...string) (output string, err error) {
	output, err = dolores_corecode.Exec("gcloud-ssh-access", parameters[0:]...)
	return
}
