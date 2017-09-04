package main

import (
	"os"
	"github.com/codegangsta/cli"
	"github.com/biezhi/agon/utils"
)

const (
	VERSION      = "0.0.1"
)

func main() {

	app := cli.NewApp()
	app.Name = "bar"
	app.Usage = "A simple static site generator"
	app.Author = "https://github.com/biezhi"
	app.Email = "biezhi.me@gmail.com"
	app.Version = VERSION
	app.Commands = []cli.Command{
		{
			Name:   "init",
			Usage:  "Create a new Bar folder",
			Action: Init,
		},
		{
			Name:      "new",
			ShortName: "n",
			Usage:     "Create a new article",
			Action:    NewArticle,
		},
		{
			Name:      "generate",
			ShortName: "g",
			Usage:     "Generate blog to public folder",
			Action:    Generate,
		},
		{
			Name:      "clean",
			ShortName: "c",
			Usage:     "Clean publish folder",
			Action:    Clean,
		},
		{
			Name:      "serve",
			ShortName: "s",
			Usage:     "Run in server mode to serve blog",
			Action:    Serve,
		},
		{
			Name:      "deploy",
			ShortName: "d",
			Usage:     "Generate blog to public folder and publish",
			Action:    Deploy,
		},
	}
	app.Run(os.Args)
	os.Exit(0)
}

// init bar config
func Init(ctx *cli.Context) {
	if len(ctx.Args()) > 0 {
		utils.MkDir(ctx.Args().First())
		utils.CopyDir("blog", ctx.Args().First())
	}
}

// create a article
func NewArticle(ctx *cli.Context) {

}

// generator site
func Generate(ctx *cli.Context) {

}

// clean publish folder
func Clean(ctx *cli.Context) {

}

// local preview
func Serve(ctx *cli.Context) {

}

// publish site
func Deploy(ctx *cli.Context) {

}

// parse config.yml
func ParseConfig()  {

}