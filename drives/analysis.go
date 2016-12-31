package dolores_drives

import (
	dolores_corecode "dolores/corecode"
)

func HttpEcho(url string, instruction string, basicAuth string) (body string) {
	getParams := map[string]string{
		"instruction": instruction,
	}
	httpHeaders := map[string]string{
		"basicAuth": basicAuth,
	}

	body, _ = dolores_corecode.HttpGet(url, getParams, httpHeaders)
	return

}
