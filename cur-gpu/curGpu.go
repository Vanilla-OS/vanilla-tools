package main

/*	License: GPLv3
	Authors:
		Mirko Brombin <mirko@fabricators.ltd>
		Vanilla OS Contributors <https://github.com/vanilla-os/>
	Copyright: 2023
	Description:
		This program is a simple OpenGL program that prints the GPU name and vendor.
*/

import (
	"flag"
	"fmt"
	"log"

	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
)

func main() {
	allFlag := flag.Bool("all", false, "Print all OpenGL extensions")
	flag.Parse()

	if err := glfw.Init(); err != nil {
		log.Fatalln("failed to initialize glfw:", err)
	}
	defer glfw.Terminate()

	glfw.WindowHint(glfw.Visible, glfw.False)
	window, err := glfw.CreateWindow(1, 1, "", nil, nil)
	if err != nil {
		log.Fatalln("failed to create glfw window:", err)
	}
	defer window.Destroy()

	window.MakeContextCurrent()
	if err := gl.Init(); err != nil {
		log.Fatalln("failed to initialize gl:", err)
	}

	var vendor, renderer, version string

	vendor = gl.GoStr(gl.GetString(gl.VENDOR))
	renderer = gl.GoStr(gl.GetString(gl.RENDERER))
	version = gl.GoStr(gl.GetString(gl.VERSION))
	fmt.Printf("OpenGL version: %s\n", version)
	fmt.Printf("OpenGL vendor: %s\n", vendor)
	fmt.Printf("OpenGL renderer: %s\n", renderer)

	if *allFlag {
		var numExtensions int32
		gl.GetIntegerv(gl.NUM_EXTENSIONS, &numExtensions)
		fmt.Printf("OpenGL extensions (%d):\n", numExtensions)
		for i := int32(0); i < numExtensions; i++ {
			fmt.Printf("\t%s\n", gl.GoStr(gl.GetStringi(gl.EXTENSIONS, uint32(i))))
		}
	}
}
