package main

/*	License: GPLv3
	Authors:
		Mirko Brombin <mirko@fabricators.ltd>
		Vanilla OS Contributors <https://github.com/vanilla-os/>
	Copyright: 2024
	Description:
		This program simply adds the environment variables needed to run a program with the NVIDIA GPU.
		NOTE: We are using Go for such a simple program because all the other vanilla tools are written in Go.
*/

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: nrun <program> [args...]")
		os.Exit(1)
	}

	program := os.Args[1]
	args := os.Args[2:]

	envVars := map[string]string{
		"__NV_PRIME_RENDER_OFFLOAD": "1",
		"__VK_LAYER_NV_optimus":     "NVIDIA_only",
		"__GLX_VENDOR_LIBRARY_NAME": "nvidia",
	}

	cmd := exec.Command(program, args...)

	cmd.Env = os.Environ()
	for key, value := range envVars {
		cmd.Env = append(cmd.Env, fmt.Sprintf("%s=%s", key, value))
	}

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Printf("Error running %s: %s\n", program, err)
		os.Exit(1)
	}
}
