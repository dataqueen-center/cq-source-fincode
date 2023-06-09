package plugin

import (
	"github.com/cloudquery/plugin-sdk/v3/plugins/source"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/dataqueen-center/cq-source-fincode/client"
	"github.com/dataqueen-center/cq-source-fincode/resources"
)

var Version = "Development"

func Plugin() *source.Plugin {
	return source.NewPlugin(
		"fincode",
		Version,
		schema.Tables{
			resources.Payments(),
		},
		client.New,
	)
}
