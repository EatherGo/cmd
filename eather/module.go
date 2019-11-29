package main

// ModuleXML module.xml template
const ModuleXML = `
<?xml version="1.0" encoding="UTF-8"?>
<module>
    <name>{{name}}</name>
    <version>1.0.0</version>
</module>
`

// ModuleMain main.go template
const ModuleMain = `package main

import (
	"github.com/EatherGo/eather"
)

type module struct{}

// {{name}} to export in plugin
func {{name}}() (f eather.Module, err error) {
	f = module{}
	return
}
`

// ModuleMainConf template to add module to modules.xml
const ModuleMainConf = `	<module>
        <name>{{name}}</name>
        <enabled>true</enabled>
    </module>`
