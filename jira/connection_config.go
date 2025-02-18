package jira

import (
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/schema"
)

type jiraConfig struct {
	BaseUrl  *string `cty:"base_url"`
	Username *string `cty:"username"`
	Token    *string `cty:"token"`
	AuthType *string `cty:"auth_type"`
}

var ConfigSchema = map[string]*schema.Attribute{
	"base_url": {
		Type: schema.TypeString,
	},
	"username": {
		Type: schema.TypeString,
	},
	"token": {
		Type: schema.TypeString,
	},
	"auth_type": {
		Type: schema.TypeString,
	},
}

func ConfigInstance() interface{} {
	return &jiraConfig{}
}

// GetConfig :: retrieve and cast connection config from query data
func GetConfig(connection *plugin.Connection) jiraConfig {
	if connection == nil || connection.Config == nil {
		return jiraConfig{}
	}
	config, _ := connection.Config.(jiraConfig)
	return config
}
