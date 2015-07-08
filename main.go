package main

import (
	"bytes"
	"github.com/BurntSushi/toml"
	"log"
	"os"
	"os/exec"

	"github.com/codegangsta/cli"
)

type LocalConfig struct {
	Name  string `toml:name`
	EMail string `toml:email`
}

func setConfig(command string, target string) {
	cmd := exec.Command("git", "config", "--local", command, target)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()

	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	app := cli.NewApp()
	app.Name = "gitconfig-setter"
	app.Usage = "Set local user.name & user.email in .gitconfig"
	app.HideHelp = true
	app.Flags = []cli.Flag{
		cli.HelpFlag,
	}
	cli.AppHelpTemplate = `
NAME:
   {{.Name}} - {{.Usage}}

USAGE:
   {{.Name}} [repository] [config]

VERSION:
   {{.Version}}{{if or .Author .Email}}

AUTHOR:{{if .Author}}
  {{.Author}}{{if .Email}} - <{{.Email}}>{{end}}{{else}}
  {{.Email}}{{end}}{{end}}

OPTIONS:
   {{range .Flags}}{{.}}
   {{end}}
	`
	app.Author = "Layzie <HIRAKI Satoru>"
	app.Email = "saruko313@gmail.com"
	app.Version = "0.0.1"
	app.Action = func(c *cli.Context) {
		var (
			config     LocalConfig
			repository string
			configFile string
		)

		if len(c.Args().Get(0)) != 0 {
			repository = c.Args().Get(0)
		} else {
			repository = "./"
		}

		if len(c.Args().Get(1)) != 0 {
			configFile = c.Args().Get(1)
		} else {
			configFile = "./"
		}

		_, err := toml.DecodeFile(configFile+"config.toml", &config)

		if err != nil {
			log.Fatal(err)
		}

		localName := config.Name
		localMail := config.EMail

		os.Chdir(repository)

		setConfig("user.name", localName)
		log.Print(repository + "'s local git config name has changed! Using " + configFile + "config.toml")

		setConfig("user.email", localMail)
		log.Print(repository + "'s local git config email has changed! Using " + configFile + "config.toml")
	}

	app.Run(os.Args)
}
