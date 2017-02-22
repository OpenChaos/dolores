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

	dbAccessAllotCommand = allot.New("(give) db access to <user:string> for <appName:string> in <appEnv:string>")
	dbAccessNlpSamples   = []string{
		"{verb} db access to {user} for {appName} in {appEnv}",
		"{verb} access to {user} for db of {appName} in {appEnv}",
		"{verb} database access to {user} for {appName} in {appEnv}",
	}

	serverListAllotCommand = allot.New("server list for <appName:string> in <appEnv:string>")
	serverListNlpSamples   = []string{}

	nslookupAllotCommand = allot.New("nslookup <searchFor:string>")
	nslookupNlpSamples   = []string{}
)
