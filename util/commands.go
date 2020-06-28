package main

import (
	"errors"
	"fmt"
)

type command struct {
	command        string
	usage          string
	description    string
	argCount       int
	commandHandler func([]string) error
}

// Help must be added in the main method to prevent
// circular reference in the initalization of commands
var helpCommand = command{
	"help",
	"help",
	"display all golf toolkit commands",
	0,
	func([]string) error {
		for _, command := range commands {
			if command.description == "" {
				continue
			}
			title := command.command
			for len(title) < 8 {
				title += " "
			}

			fmt.Println("   " + title + command.description)
		}
		return nil
	},
}

var commands = []command{
	command{
		"about",
		"about",
		"display information about the golf toolkit",
		0,
		func([]string) error {
			printBlockText("the golf toolkit is designed to help you make games with the golf engine. "+
				"The golf engine produces WASM files with all your sprite, and map data packed in. "+
				"This makes sharing games simpler, as there are no external dependancies, as well as making the resulting game files smaller. "+
				"This means, however, that all game assests must be converted to the correct golf format. "+
				"The golf toolkit provides you with a small set of tools to convert these files quickly and simply. ", 80)
			return nil
		},
	},

	command{
		"exit",
		"exit",
		"exit the golf toolkit",
		0,
		func([]string) error {
			fmt.Print("GoodByte!\n")
			cmdDone = true
			return nil
		},
	},

	command{
		"map",
		"map <map file> <sprite file> <output file>",
		"<map file> <sprite file> <output file> converts an image file into a golf map data",
		3,
		func(args []string) error {
			return convertMap(args[0], args[1], args[2])
		},
	},

	command{
		"csvmap",
		"csvmap <map file> <output file>",
		"<map file> <output file> converts a csv file into a golf map data",
		2,
		func(args []string) error {
			return convertCSVMap(args[0], args[1])
		},
	},

	command{
		"sprite",
		"sprite <sprite file> <output file>",
		"<sprite file> <output file>",
		2,
		func(args []string) error {
			return convertSpriteSheet(args[0], args[1])
		},
	},

	command{
		"flag",
		"flag <flag file> <output file>",
		"<flag file> <output file> convert a csv file into golf sprite flag data",
		2,
		func(args []string) error {
			return convertFlag(args[0], args[1])
		},
	},

	command{
		command: "",
	},

	command{
		"startserver",
		"startserver",
		"starts a server in the current directory that can be used to play your golf engine games",
		0,
		func(args []string) error {
			return errors.New("This function is not implemented yet")
		},
	},

	command{
		"build",
		"build <output file>",
		"<output file> builds the current golf project and creates a wasm file",
		1,
		func(args []string) error {
			return errors.New("this function is not implemented yet")
		},
	},
}
