package main

import (
	"code/pkg/du"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name:  "hexlet-path-size",
		Usage: "print size of a file or directory",
		Arguments: []cli.Argument{
			&cli.StringArg{
				Name:      "path",
				Value:     ".",
				UsageText: "path to file or directory",
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			path := cmd.StringArg("path")
			if path == "" {
				path = "."
			}

			size, err := du.GetSize(path)
			if err != nil {
				return err
			}

			fmt.Printf("%d\t%s\n", size, path)
			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatalf("%v", err)
	}
}
