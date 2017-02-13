package dolores_slack

import dolores_corecode "dolores/corecode"

var (
	replyHelpMessage = `hey,
for help:
	what you just typed is just fine, viable commands are "help" and "sos"

to get ssh access to machines:
	'give access to userID@somedomain.com for machine-Pattern'
	or
	'give access to userID for machine-Pattern'
	it will try to fetch key for this from internal/scm portals, so make sure your keys are updated"
`

	sshAccessReplyDeferMessage = "sure, let me check if it's allowed as of now :)"

	sshAccessReplySuccessMessage, sshAccessReplyFailureMessage string

	doloresWrongCmdMessage = "say what, that got no meaning for me\n try asking for `help`"
)

func overrideMessagesFromEnv() {
	replyHelpMessage = dolores_corecode.OverrideFromEnvVar(
		"DOLORES_HELP_REPLY", replyHelpMessage)

	sshAccessReplyDeferMessage = dolores_corecode.OverrideFromEnvVar(
		"SSH_ACCESS_REPLY_DEFER", sshAccessReplyDeferMessage)

	sshAccessReplySuccessMessage = dolores_corecode.OverrideFromEnvVar(
		"SSH_ACCESS_REPLY_SUCCESS", sshAccessReplySuccessMessage)

	sshAccessReplyFailureMessage = dolores_corecode.OverrideFromEnvVar(
		"SSH_ACCESS_REPLY_FAILURE", sshAccessReplyFailureMessage)

	doloresWrongCmdMessage = dolores_corecode.OverrideFromEnvVar(
		"DOLORES_WRONG_CMD_REPLY", doloresWrongCmdMessage)
}
