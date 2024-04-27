package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jfgrea27/hnsnap/internal/client"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "top-stories",
		Usage: "Top stories from https://news.ycombinator.com",
		Action: func(*cli.Context) error {

			hn := client.HNClient{}

			ts := hn.TopStories()

			for _, t := range ts {
				if len(t.Title) > 0 {
					fmt.Println(fmt.Sprintf("Title: %s", t.Title))
				}

			}

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
