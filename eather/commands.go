package main

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/joho/godotenv"
	"github.com/urfave/cli"
)

func createNew(c *cli.Context) error {
	name := c.String("name")

	if gitPath, err := exec.LookPath("git"); err == nil {
		cmd := exec.Command(gitPath, "clone", "--depth=1", "--branch=master", "https://github.com/EatherGo/starter-app", name)

		var errb bytes.Buffer
		cmd.Stderr = &errb

		err := cmd.Run()
		if err != nil {
			log.Fatal(errb.String())
		}

		cmd = exec.Command("rm", "-rf", name+"/.git")
		cmd.Stderr = &errb

		err = cmd.Run()
		if err != nil {
			log.Fatal(errb.String())
		}

		sourceFile := name + "/.env.example"
		destinationFile := name + "/.env"

		input, err := ioutil.ReadFile(sourceFile)
		if err != nil {
			fmt.Println(err)
		}

		err = ioutil.WriteFile(destinationFile, input, 0644)
		if err != nil {
			fmt.Println("Error creating", destinationFile)
			fmt.Println(err)
		}

		if err := os.MkdirAll(name+"/src/Modules", os.ModePerm); err != nil {
			return errors.New("cannot create modules path")
		}

		fmt.Println("Application " + name + " was created. Let's start:\n")
		fmt.Println("cd " + name)
		fmt.Println("go run main.go")

		return nil
	}

	return errors.New("git not found")
}

func createModule(c *cli.Context) error {
	name := c.String("name")

	err := godotenv.Load()
	if err != nil {
		return errors.New("Error loading .env file")
	}

	modulesDir := os.Getenv("CUSTOM_MODULES_DIR")
	if modulesDir == "" {
		return errors.New("Error loading CUSTOM_MODULE_DIR from env")
	}

	if _, err := os.Stat(modulesDir); os.IsNotExist(err) {
		return errors.New(modulesDir + " does not exists.")
	}

	newModule(modulesDir, name)

	fmt.Println(name)

	return nil
}

func newModule(dir string, name string) error {
	path := dir + "/" + name
	if err := os.MkdirAll(path, os.ModePerm); err != nil {
		return errors.New("cannot create module" + name)
	}

	if err := os.MkdirAll(path+"/etc", os.ModePerm); err != nil {
		return errors.New("cannot create module" + name)
	}

	moduleXML := strings.Replace(ModuleXML, "{{name}}", name, -1)

	err := ioutil.WriteFile(path+"/etc/module.xml", []byte(moduleXML), 0644)
	if err != nil {
		fmt.Println("Error creating")
		fmt.Println(err)
	}

	mainGo := strings.Replace(ModuleMain, "{{name}}", name, -1)

	err = ioutil.WriteFile(path+"/main.go", []byte(mainGo), 0644)
	if err != nil {
		fmt.Println("Error creating")
		fmt.Println(err)
	}

	dat, _ := ioutil.ReadFile("config/modules.xml")

	index := strings.Index(string(dat), "</modules>")

	mod := strings.Replace(ModuleMainConf, "{{name}}", name, -1)
	dats := string(dat[:index]) + mod + "\n" + string(dat[index:])

	f, err := os.OpenFile("config/modules.xml", os.O_RDWR, 0644)
	if err != nil {
		fmt.Println(err)
	}

	f.Truncate(0)

	f.Write([]byte(dats))

	return nil
}
