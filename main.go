package main

import (
	"github.com/misterek/steampipe-plugin-starwars/starwars"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		PluginFunc: starwars.Plugin})
}
