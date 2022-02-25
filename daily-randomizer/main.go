package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/urfave/cli/v2"
)

// Application constants
const (
	name   = "daily-randomizer"
	usage  = "Super simple Go CLI tool for randomly picking people in Dailys"
	author = "Rene Schach @rescDev"
)

// Flags
var (
	peopleFlag = &cli.StringFlag{
		Name:    "people",
		Aliases: []string{"p"},
	}
)

func main() {
	app := &cli.App{
		Name:  name,
		Usage: usage,
		Authors: []*cli.Author{
			{
				Name: author,
			},
		},
		Flags: []cli.Flag{
			peopleFlag,
		},
		Action: func(c *cli.Context) error {
			people := c.String("people")
			names := strings.Split(people, ",")

			// Skip execution if no names provided or invalid format
			if names[0] == "" {
				fmt.Println("Hey, what are you doing? No people no party, please submit some names :)")
				return nil
			}

			fmt.Println("Starting daily with ...")
			fmt.Println()

			for len(names) > 0 {
				time.Sleep(500 * time.Millisecond)
				fmt.Print(". ")
				time.Sleep(500 * time.Millisecond)
				fmt.Print(". ")
				time.Sleep(500 * time.Millisecond)
				fmt.Print(". ")

				rand.Seed(time.Now().UnixNano())
				index := rand.Intn(len(names))

				fmt.Printf("%s !!!\n", names[index])

				names[index] = names[len(names)-1]

				names = names[:len(names)-1]

				buf := bufio.NewReader(os.Stdin)
				_, err := buf.ReadBytes('\n')
				if err != nil {
					fmt.Println(err)
					return err
				}
			}

			fmt.Println("No one left :)")
			return nil
		},
	}

	// Hide the 'help' command as it is already available as global option '--help'
	app.HideHelpCommand = true

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
