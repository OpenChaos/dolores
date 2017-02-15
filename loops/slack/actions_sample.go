package dolores_slack

import "github.com/sbstjn/allot"

var (
	helpAllotCommand = allot.New("(help|sos)")
	helpNlpSamples   = []string{
		"please help",
		"need help",
	}

	sshAccessAllotCommand = allot.New("(give|remove) access (to|for|from|of) <user:string> for <machinePattern:string>")
	sshAccessNlpSamples   = []string{
		"{verb} ssh access to {user} for {machinePattern}",
		"{verb} access to {user} for {machinePattern}",
		"{verb} login access to {user} for {machinePattern}",
	}
)
