package main

import "strings"

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

type template interface {
	parseData(name string) string
}

type templater struct {
	template string
}

func (t templater) parseData(name string) string {
	return strings.Replace(t.template, "{{name}}", name, -1)
}
