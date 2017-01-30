package dolores_drives

import dolores_corecode "dolores/corecode"

func GiveSshAccess(machinePattern, user string) (err error) {
	err = dolores_corecode.Exec([]string{
		"gcloud-ssh-access",
		machinePattern,
		user,
	})
	return
}
