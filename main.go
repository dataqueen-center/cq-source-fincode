package main

import (
	"github.com/cloudquery/plugin-sdk/v3/serve"
	"github.com/dataqueen-center/cq-source-fincode/plugin"
)

func main() {
	serve.Source(plugin.Plugin())
}
