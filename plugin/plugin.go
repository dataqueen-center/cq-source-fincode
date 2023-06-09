package plugin

import (
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/cloudquery/plugin-sdk/v3/plugins/source"
	"github.com/dataqueen-center/cq-source-fincode/resources"
	"github.com/dataqueen-center/cq-source-fincode/client"
)

var Version = "Development"

func Plugin() *source.Plugin {
	return source.NewPlugin(
		"googleanalytics",
		Version,
		schema.Tables{
			resources.Payments(),
		},
		client.New,
	)
}
