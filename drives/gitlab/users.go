package dolores_gitlab

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	dolores_corecode "github.com/OpenChaos/dolores/corecode"
)

type gitlabUser struct {
	Id               int    `json:"id"`
	Username         string `json:"username"`
	Name             string `json:"name"`
	State            string `json:"state"`
	AvatarUrl        string `json:"avatar_url"`
	WebUrl           string `json:"web_url"`
	Email            string `json:"email"`
	IsAdmin          bool   `json:"is_admin"`
	CanCreateGroup   bool   `json:"can_create_group"`
	CanCreateProject bool   `json:"can_create_project"`
	External         bool   `json:"external"`
}

func ExternalUserIds() (users []gitlabUser) {
	gitlabBaseURL := os.Getenv("DOLORES_GITLAB_BASE_URL")
	gitlabPrivateToken := os.Getenv("DOLORES_GITLAB_PRIVATE_TOKEN")

	if gitlabBaseURL == "" || gitlabPrivateToken == "" {
		log.Println("[error] Gitlab Environment Config missing")
		return
	}

	usersAPIUrl := fmt.Sprintf("%s/api/v3/users", gitlabBaseURL)
	getParams := map[string]string{
		"private_token": gitlabPrivateToken,
		"external":      "true",
		"per_page":      "1000",
	}
	httpHeaders := map[string]string{}

	reply, err := dolores_corecode.HttpGet(usersAPIUrl, getParams, httpHeaders)
	if err == nil {
		err = json.Unmarshal([]byte(reply), &users)
		log.Println("[error]", err)
	} else {
		log.Println("[error]", err)
	}
	return
}

func UserDetails(userId int) (user gitlabUser) {
	gitlabBaseURL := os.Getenv("DOLORES_GITLAB_BASE_URL")
	gitlabPrivateToken := os.Getenv("DOLORES_GITLAB_PRIVATE_TOKEN")

	if gitlabBaseURL == "" || gitlabPrivateToken == "" {
		log.Println("[error] Gitlab Environment Config missing")
		return
	}

	usersAPIUrl := fmt.Sprintf("%s/api/v3/users/%d", gitlabBaseURL, userId)
	getParams := map[string]string{
		"private_token": gitlabPrivateToken,
	}
	httpHeaders := map[string]string{}

	reply, err := dolores_corecode.HttpGet(usersAPIUrl, getParams, httpHeaders)
	if err == nil {
		err = json.Unmarshal([]byte(reply), &user)
	} else {
		log.Println("[error]", err)
	}
	return
}

func MarkUsersInternal() {
	log.Println("[info] marking all users internal for whitelisted domain")
	gitlabBaseURL := os.Getenv("DOLORES_GITLAB_BASE_URL")
	gitlabPrivateToken := os.Getenv("DOLORES_GITLAB_PRIVATE_TOKEN")
	gitlabWhitelistDomain := os.Getenv("DOLORES_GITLAB_WHITELIST_DOMAIN")

	if gitlabBaseURL == "" || gitlabPrivateToken == "" {
		log.Println("[error] Gitlab Environment Config missing")
		return
	}

	external_users := ExternalUserIds()
	for _, user := range external_users {
		userDetails := UserDetails(user.Id)
		if strings.HasSuffix(userDetails.Email, gitlabWhitelistDomain) {
			usersAPIUrl := fmt.Sprintf("%s/api/v3/users/%d", gitlabBaseURL, user.Id)
			getParams := map[string]string{
				"private_token": gitlabPrivateToken,
				"external":      "false",
			}
			httpHeaders := map[string]string{}

			_, err := dolores_corecode.HttpPut(usersAPIUrl, getParams, httpHeaders)
			if err != nil {
				log.Println("[error]", err)
			}
		} else {
			log.Println("[warn] untrusted domain users", userDetails.Username, userDetails.Email)
		}
	}
	return
}
