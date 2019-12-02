package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

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

// ModuleController template to create controller for module
const ModuleController = `package controller

import (
	"net/http"

	"github.com/EatherGo/eather"
)

// Index
func Index(w http.ResponseWriter, r *http.Request) {
	eather.SendJSONResponse(w, eather.Response{Message: "Running", Status: true})
}
`

type template interface {
	parseData(name string) string
}

type templater struct {
	template string
	name     string
}

func (t templater) parseData() string {
	if t.name != "" {
		return strings.Replace(t.template, "{{name}}", t.name, -1)
	}

	return t.template
}

func createFile(fpath string, template templater) {
	err := ioutil.WriteFile(fpath, []byte(template.parseData()), 0644)
	if err != nil {
		fmt.Println("Error creating")
		fmt.Println(err)
	}
}

func addToModulesConf(template templater) {
	dat, _ := ioutil.ReadFile("config/modules.xml")

	index := strings.Index(string(dat), "</modules>")

	dats := string(dat[:index]) + template.parseData() + "\n" + string(dat[index:])

	f, err := os.OpenFile("config/modules.xml", os.O_RDWR, 0644)
	if err != nil {
		fmt.Println(err)
	}

	f.Truncate(0)

	f.Write([]byte(dats))
}

func newModule(dir string, name string) error {
	path := dir + "/" + name
	if err := os.MkdirAll(path+"/etc", os.ModePerm); err != nil {
		return errors.New("cannot create module" + name)
	}

	createFile(path+"/etc/module.xml", templater{template: ModuleXML, name: name})

	createFile(path+"/main.go", templater{template: ModuleMain, name: name})

	addToModulesConf(templater{template: ModuleMainConf, name: name})

	return nil
}

func initModController(dir string, name string) error {
	path := dir + "/" + name
	if err := os.MkdirAll(path+"/controller", os.ModePerm); err != nil {
		return errors.New("cannot create controller folder for module" + name)
	}

	createFile(path+"/controller/"+name, templater{template: ModuleController})

	return nil
}
