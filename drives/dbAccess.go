package dolores_drives

import dolores_corecode "dolores/corecode"

func GiveDbAccess(parameters ...string) (output string, err error) {
	// usage: psql-readonly-account <SSH-USER> <DATABASE_MASTER_BOX> <DATABASE_NAME> <DATABASE_USERNAME> <DATABASE_PASSWORD>
	output, err = dolores_corecode.Exec("psql-readonly-account", parameters[0:]...)
	return
}
