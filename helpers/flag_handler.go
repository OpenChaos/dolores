package dolores_helpers

import (
	"flag"
	"fmt"

	"github.com/abhishekkr/gol/golconfig"
)

// flags
var (
	flagConfig = flag.String("config", "", "the path to overriding config file")

	flagSlackBotName     = flag.String("slack-bot-name", "dolores", "slack bot name")
	flagSlackBotAPIToken = flag.String("slack-bot-api-token", "dolores", "slack bot api token")
)

/* assignIfEmpty assigns val to *key only if it's empty */
func assignIfEmpty(mapper golconfig.FlatConfig, key string, val string) {
	if mapper[key] == "" {
		mapper[key] = val
	}
}

/*
ConfigFromFlags configs from values provided to flags.
*/
func ConfigFromFlags() golconfig.FlatConfig {
	flag.Parse()

	var config golconfig.FlatConfig
	config = make(golconfig.FlatConfig)
	if *flagConfig != "" {
		configFile := golconfig.GetConfigurator("json")
		configFile.ConfigFromFile(*flagConfig, &config)
	}

	assignIfEmpty(config, "slack-bot-name", *flagSlackBotName)
	assignIfEmpty(config, "slack-bot-api-token", *flagSlackBotAPIToken)

	fmt.Println("Dolores config:")
	for cfg, val := range config {
		fmt.Printf("[ %v : %v ]\n", cfg, val)
	}
	fmt.Println("***********************************************************")
	return config
}
