package main

import (
	"io/ioutil"
	"log"
	"os"
	"path"

	"github.com/urfave/cli"
	"gopkg.in/AlecAivazis/survey.v1"
)

const tempPath string = "./templates"

var tempMap = map[string]map[string]string{
	"gitignore": {
		"input":  "gitignore.txt",
		"output": ".gitignore",
	},
	"npmignore": {
		"input":  "npmignore.txt",
		"output": ".npmignore",
	},
}

func makeFile(cwd string, input string, output string) {
	temp, err := ioutil.ReadFile(path.Join(tempPath, input))

	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile(path.Join(cwd, output), temp, 0644)

	if err != nil {
		log.Fatal(err)
	}
}

func makeFiles(fileNames *[]string) {
	cwd, _ := os.Getwd()

	for _, value := range *fileNames {
		temp := tempMap[value]
		makeFile(cwd, temp["input"], temp["output"])
	}
}

func appAction(c *cli.Context) error {
	tempFiles := []string{}
	keys := []string{}

	for key := range tempMap {
		keys = append(keys, key)
	}

	prompt := &survey.MultiSelect{
		Message: "Please select the required file",
		Options: keys,
	}

	err := survey.AskOne(prompt, &tempFiles, nil)

	if err != nil {
		return err
	}

	makeFiles(&tempFiles)

	return nil
}

func main() {
	app := cli.NewApp()
	app.Name = "gig"

	app.Action = appAction

	err := app.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}
}
