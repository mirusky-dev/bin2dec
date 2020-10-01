package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/urfave/cli/v2"
)

var input string

func main() {
	app := &cli.App{
		Name:   "Bin2Dec",
		Action: Bin2DecAction,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "input",
				Aliases:     []string{"i"},
				Usage:       "input value to be converted",
				EnvVars:     []string{"BIN2DEC_INPUT"},
				Destination: &input,
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
	}
}

// Bin2DecAction ...
func Bin2DecAction(c *cli.Context) error {

	if c.NArg() > 0 {
		input = c.Args().Get(0)
	}

	if len(input) == 0 {
		return errors.New("please use:\n bin2dec <number> or\n bin2dec -i <number> or\n export: BIN2DEC_INPUT=<number> && bin2dec")
	}

	if len(input) > 8 && len(input) < 1 {
		return errors.New("the length must between 1 and 8")
	}

	n, err := strconv.ParseInt(input, 2, 0)
	if err != nil {
		return errors.New("must be only binary digits")
	}

	fmt.Println(n)

	return nil
}
