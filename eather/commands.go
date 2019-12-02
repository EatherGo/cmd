package main

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"

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
	modulesDir := c.String("path")

	err := godotenv.Load()
	if err != nil {
		return errors.New("Error loading .env file")
	}

	if modulesDir == "" {
		modulesDir = os.Getenv("CUSTOM_MODULES_DIR")
		if modulesDir == "" {
			return errors.New("Error loading CUSTOM_MODULE_DIR from env")
		}
	}

	if _, err := os.Stat(modulesDir); os.IsNotExist(err) {
		return errors.New(modulesDir + " does not exists.")
	}

	if _, err := os.Stat(modulesDir + "/" + name); !os.IsNotExist(err) {
		return errors.New("module " + name + " already exists")
	}

	newModule(modulesDir, name)

	if c.Bool("controller") {
		initModController(modulesDir, name)
	}

	if c.Bool("events") {
		initModEvents(modulesDir, name)
	}

	if model := c.String("model"); model != "" {
		initModModel(modulesDir, name, model)
	}

	if c.Bool("upgrade") {
		initModUpgrade(modulesDir, name)
	}

	if c.Bool("cron") {
		initModCron(modulesDir, name)
	}

	if c.Bool("call") {
		initModCallable(modulesDir, name)
	}

	fmt.Println("Module " + name + " was crafted")

	return nil
}
