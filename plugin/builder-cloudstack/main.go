package main

import (
	"github.com/ddoc/packer/builder/cloudstack"
	"github.com/ddoc/packer/packer/plugin"
)

func main() {
	plugin.ServeBuilder(new(cloudstack.Builder))
}
