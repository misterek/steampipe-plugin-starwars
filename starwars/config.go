package starwars

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/schema"
)

// Don't really need anything here
type starwarsConfig struct {
}

var ConfigSchema = map[string]*schema.Attribute{}

func ConfigInstance() interface{} {
	return &starwarsConfig{}
}

func GetConfig(connection *plugin.Connection) starwarsConfig {
	if connection == nil || connection.Config == nil {
		return starwarsConfig{}
	}
	config, _ := connection.Config.(starwarsConfig)
	return config
}
