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
const ModuleController = `package main

import (
	"net/http"

	"github.com/EatherGo/eather"
)

// Index
func Index(w http.ResponseWriter, r *http.Request) {
	eather.SendJSONResponse(w, eather.Response{Message: "Running", Status: true})
}
`

// ModuleMainMapRouter template for MapRoutes function
const ModuleMainMapRouter = `
func (m module) MapRoutes() {
	router := eather.GetRouter()

	router.HandleFunc("/index", Index).Methods("GET")
}
`

// ModuleEvents template for events
const ModuleEvents = `package main

import (
	"github.com/EatherGo/eather"
)

// GetEventFuncs will return slice of events
func (m module) GetEventFuncs() []eather.Fire {
	return eventFuncs
}

var eventFuncs = []eather.Fire{
	eather.Fire{Call: "added", Func: added},
}

var added = func(data ...interface{}) {
	// do stuff here
}
`

// ModuleEventsXML add events to xml
const ModuleEventsXML = `
	<events>
		<listener for="product_added" call="added" name="add_some_stuff"></listener>
	</events>
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

func newModule(dir string, name string) error {
	path := dir + "/" + name

	if err := os.MkdirAll(path+"/etc", os.ModePerm); err != nil {
		return errors.New("cannot create module" + name)
	}

	createFile(path+"/etc/module.xml", templater{template: ModuleXML, name: name})

	createFile(path+"/main.go", templater{template: ModuleMain, name: name})

	writeToFileBefore("config/modules.xml", "</modules>", templater{template: ModuleMainConf, name: name})

	return nil
}

func initModController(dir string, name string) error {
	path := dir + "/" + name

	createFile(path+"/controller.go", templater{template: ModuleController})

	writeToFileBefore(path+"/main.go", "// "+name, templater{template: ModuleMainMapRouter})

	return nil
}

func initModEvents(dir string, name string) error {
	path := dir + "/" + name

	createFile(path+"/events.go", templater{template: ModuleEvents})

	writeToFileBefore(path+"etc/module.xml", "</module>", templater{template: ModuleEventsXML})

	return nil
}

func writeToFileBefore(file string, needle string, template templater) {
	dat, _ := ioutil.ReadFile(file)

	index := strings.Index(string(dat), needle)

	dats := string(dat[:index]) + template.parseData() + "\n" + string(dat[index:])

	f, err := os.OpenFile(file, os.O_RDWR, 0644)
	if err != nil {
		fmt.Println(err)
	}

	f.Truncate(0)

	f.Write([]byte(dats))
}
