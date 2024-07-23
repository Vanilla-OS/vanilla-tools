package main

/*	License: GPLv3
	Authors:
		Mirko Brombin <mirko@fabricators.ltd>
		Vanilla OS Contributors <https://github.com/vanilla-os/>
	Copyright: 2024
	Description:
		This tool is used to lock apt and dpkg binaries to prevent
		usage of the package manager in an immutable system.
		Part of the vanilla-utils.
*/

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const binDir = "/usr/bin/"

func lockBins(path string, verbose bool) error {
	if path == "" {
		path = binDir
	}

	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if strings.HasPrefix(info.Name(), "apt") || strings.HasPrefix(info.Name(), "dpkg") {
			oldPath := path
			newPath := fmt.Sprintf("%sprivate.%s", filepath.Dir(path)+"/", info.Name())

			if verbose {
				fmt.Printf("Locking %s\n", oldPath)
			}

			if err := os.Rename(oldPath, newPath); err != nil {
				return err
			}
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

func unlockBins(path string, verbose bool) error {
	if path == "" {
		path = binDir
	}

	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if strings.HasPrefix(info.Name(), "private.apt") || strings.HasPrefix(info.Name(), "private.dpkg") {
			oldPath := path
			newPath := fmt.Sprintf("%s%s", filepath.Dir(path)+"/", strings.TrimPrefix(info.Name(), "private."))

			if verbose {
				fmt.Printf("Unlocking %s\n", oldPath)
			}

			if err := os.Rename(oldPath, newPath); err != nil {
				return err
			}
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options]\n\nOptions:\n", os.Args[0])
		flag.PrintDefaults()
	}

	lockFlag := flag.Bool("lock", false, "Locks apt and dpkg binaries")
	unlockFlag := flag.Bool("unlock", false, "Unlocks apt and dpkg binaries")
	dirFlag := flag.String("dpath", "", "Specify a custom path to search for apt and dpkg binaries")
	verboseFlag := flag.Bool("verbose", false, "Enables verbose output")

	flag.Parse()

	if *lockFlag && *unlockFlag {
		fmt.Fprintln(os.Stderr, "Error: --lock and --unlock cannot be used together")
		os.Exit(1)
	}

	if *lockFlag {
		if err := lockBins(*dirFlag, *verboseFlag); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %s\n", err)
			os.Exit(1)
		}

		os.Exit(0)
	}

	if *unlockFlag {
		if err := unlockBins(*dirFlag, *verboseFlag); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %s\n", err)
			os.Exit(1)
		}

		os.Exit(0)
	}

	flag.Usage()
	os.Exit(1)
}
