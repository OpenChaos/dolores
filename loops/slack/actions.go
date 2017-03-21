package dolores_slack

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/nlopes/slack"
	"github.com/sbstjn/allot"

	dolores_corecode "github.com/OpenChaos/dolores/corecode"
	dolores_drives "github.com/OpenChaos/dolores/drives"
	dolores_gcp "github.com/OpenChaos/dolores/drives/gcp"
)

var (
	helpMessageHandler = MessageHandler{name: "help",
		allotCommand: helpAllotCommand,
		nlpSamples:   helpNlpSamples,
		msgFoo:       helpMessage}

	sshAccessMessageHandler = MessageHandler{name: "ssh-access",
		allotCommand: sshAccessAllotCommand,
		nlpSamples:   sshAccessNlpSamples,
		msgFoo:       sshAccess}

	dbAccessMessageHandler = MessageHandler{name: "db-access",
		allotCommand: dbAccessAllotCommand,
		nlpSamples:   dbAccessNlpSamples,
		msgFoo:       dbAccess}

	serverListMessageHandler = MessageHandler{name: "server-list",
		allotCommand: serverListAllotCommand,
		nlpSamples:   serverListNlpSamples,
		msgFoo:       serverList}

	nslookupMessageHandler = MessageHandler{name: "nslookup",
		allotCommand: nslookupAllotCommand,
		nlpSamples:   nslookupNlpSamples,
		msgFoo:       nslookup}

	bootLogMessageHandler = MessageHandler{name: "bootlog",
		allotCommand: bootLogAllotCommand,
		nlpSamples:   bootLogNlpSamples,
		msgFoo:       bootLog}
)

func helpMessage(ev *slack.MessageEvent, match allot.MatchInterface) (reply string, err error) {
	reply = replyHelpMessage
	return
}

func sshAccess(ev *slack.MessageEvent, match allot.MatchInterface) (reply string, err error) {
	axn, _ := match.Match(0)
	prep, _ := match.Match(1)
	user, _ := match.Match(2)
	machinePattern, _ := match.Match(3)

	requestedBy := SenderEmail(ev)
	isAdmin := "no"
	if IsAdmin(requestedBy) {
		isAdmin = "yes"
	}

	Reply(ev, accessReplyDeferMessage)
	if axn == "give" && (prep == "to" || prep == "for") {
		reply, err = dolores_drives.GiveSshAccess(machinePattern, user, isAdmin)
		if err != nil && sshAccessReplyFailureMessage != "" {
			reply = sshAccessReplyFailureMessage
		} else if err == nil && sshAccessReplySuccessMessage != "" {
			reply = sshAccessReplySuccessMessage
		}
	}
	return
}

func dbAccess(ev *slack.MessageEvent, match allot.MatchInterface) (reply string, err error) {
	axn, _ := match.Match(0)
	dbUsername, _ := match.Match(1)
	appName, _ := match.Match(2)
	appEnv, _ := match.Match(3)
	appNameUpper := strings.ToUpper(appName)
	appEnvUpper := strings.ToUpper(appEnv)

	requestedBy := SenderEmail(ev)
	if !IsAdmin(requestedBy) && !IsDbAdmin(requestedBy) {
		reply = dbAccessReplyNotAdmin
		return
	}

	if !IsPersonalMessage(ev) {
		reply = notAllowedInChannel
		return
	}
	Reply(ev, accessReplyDeferMessage)

	whitelist_dbs := strings.Fields(os.Getenv("DATABASE_READONLY_WHITELIST_DBS"))
	whitelist_dbs_count := len(whitelist_dbs)
	for _token_idx, _token := range whitelist_dbs {
		if _token == appNameUpper {
			break
		}
		if _token_idx == whitelist_dbs_count {
			log.Printf("[ERROR] %s asked db-access for non-whitelisted %s", requestedBy, appNameUpper)
			reply = dbAccessReplyNotWhitelisted
			return
		}
	}
	dbSshUser := os.Getenv(fmt.Sprintf("DATABASE_READONLY_%s_%s_MASTER_SSH_USER", appNameUpper, appEnvUpper))
	dbMaster := os.Getenv(fmt.Sprintf("DATABASE_READONLY_%s_%s_MASTER", appNameUpper, appEnvUpper))
	dbSlave := os.Getenv(fmt.Sprintf("DATABASE_READONLY_%s_%s_SLAVE", appNameUpper, appEnvUpper))
	dbName := os.Getenv(fmt.Sprintf("DATABASE_READONLY_%s_%s_DBNAME", appNameUpper, appEnvUpper))
	dbPassword := dolores_corecode.GeneratePassword(25, true)

	if dbSshUser == "" || dbMaster == "" || dbSlave == "" || dbName == "" || dbUsername == "" || dbPassword == "" {
		reply = "sorry but `dbAccess` task might not be available for this environment, let `systems team` know"
		return
	}
	log.Printf("[info] dbaccess with master:%s, dbslave: %s, dbname: %s, dbuser: %s, dbpassword: %s",
		dbMaster, dbSlave, dbName, dbUsername, dbPassword)

	if axn == "give" {
		reply, err = dolores_drives.GiveDbAccess(dbSshUser, dbMaster, dbName, dbUsername, dbPassword)
		if err != nil && dbAccessReplyFailureMessage != "" {
			reply = dbAccessReplyFailureMessage
		} else if err == nil && dbAccessReplySuccessMessage != "" {
			reply = dbAccessReplySuccessMessage
		} else if reply == "" {
			reply = fmt.Sprintf(
				"host: `%s`\ndb: `%s`\nuser: `ro_%s`\npassword: `%s`\ncontact systems team in case of issues",
				dbSlave, dbName, dbUsername, dbPassword,
			)
		}
	}
	return
}

func serverList(ev *slack.MessageEvent, match allot.MatchInterface) (reply string, err error) {
	appName, _ := match.Match(0)
	appEnv, _ := match.Match(1)
	appNameUpper := strings.ToUpper(appName)
	appEnvUpper := strings.ToUpper(appEnv)

	AddReaction(ev, "+1")
	appKeywordEnvVar := fmt.Sprintf("SERVER_LIST_%s_%s", appEnvUpper, appNameUpper)
	serverKeyword := os.Getenv(appKeywordEnvVar)
	serverListPath := os.Getenv("GCLOUD_COMPUTE_LIST")
	log.Printf("[info] server list for env: %s\nkeyword: %s\nlist-path: %s", appKeywordEnvVar, serverKeyword, serverListPath)
	if serverKeyword == "" {
		reply = serverListAppNotFound
		return
	}
	reply, err = dolores_drives.ServerListFor(serverKeyword, serverListPath)
	if err != nil {
		reply = fmt.Sprintf("[ERROR] server list failed for %s app in %s", appName, appEnv)
	}
	reply = fmt.Sprintf("```\n%s\n```", reply)
	return
}

func nslookup(ev *slack.MessageEvent, match allot.MatchInterface) (reply string, err error) {
	searchFor, _ := match.Match(0)

	AddReaction(ev, "+1")
	serverListPath := os.Getenv("GCLOUD_COMPUTE_LIST")
	reply, err = dolores_drives.Nslookup(searchFor, serverListPath)
	if err != nil {
		reply = fmt.Sprintf("[ERROR] nslookup failed for %s", searchFor)
	}
	reply = fmt.Sprintf("```\n%s\n```", reply)
	return
}

func bootLog(ev *slack.MessageEvent, match allot.MatchInterface) (reply string, err error) {
	serialOutFor, _ := match.Match(0)
	serialOutLineCount, serialOutLineCountErr := match.Match(1)

	_, notInt := strconv.Atoi(serialOutLineCount)
	if serialOutLineCountErr != nil || notInt != nil {
		serialOutLineCount = "100" // default line count for logs
	}

	AddReaction(ev, "+1")
	serverListPath := os.Getenv("GCLOUD_COMPUTE_LIST")
	reply, err = dolores_gcp.GcloudSerialOut(serialOutFor,
		serialOutLineCount,
		serverListPath)
	if err != nil {
		reply = fmt.Sprintf("[ERROR] bootlog failed for %s", serialOutFor)
	}
	reply = fmt.Sprintf("```\n%s\n```", reply)
	return
}
