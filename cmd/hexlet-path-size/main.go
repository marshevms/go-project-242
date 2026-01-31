package main

import (
	"code/pkg/du"
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

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
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "human",
				Aliases: []string{"H"},
				Usage:   "human-readable sizes (auto-select unit)",
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

			var strSize string
			if cmd.Bool("human") {
				strSize = FormatSize(size)
			} else {
				strSize = fmt.Sprintf("%d", size)
			}

			fmt.Printf("%s\t%s\n", strSize, path)
			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatalf("%v", err)
	}
}

type format struct {
	value int64
	unit  string
}

// FormatSize formats size in bytes to human-readable format
func FormatSize(size int64) string {
	const (
		KB int64 = 1 << 13
		MB int64 = 1 << 23
		GB int64 = 1 << 33
		TB int64 = 1 << 43
		PB int64 = 1 << 53
		// EB int64 = 1 << 63 TODO: разобраться
	)

	var formats = []format{
		// {value: EB, unit: "EB"},
		{value: PB, unit: "PB"},
		{value: TB, unit: "TB"},
		{value: GB, unit: "GB"},
		{value: MB, unit: "MB"},
		{value: KB, unit: "KB"},
	}

	var f format
	for _, f = range formats {
		if size >= f.value {
			strSize := strconv.FormatFloat(float64(size)/float64(f.value), 'f', 1, 64)
			strSize = strings.TrimRight(strSize, "0")
			strSize = strings.TrimRight(strSize, ".")
			return fmt.Sprintf("%s%s", strSize, f.unit)
		}
	}

	return fmt.Sprintf("%dB", size)
}
