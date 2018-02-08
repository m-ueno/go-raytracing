package main

import (
	"flag"
	"log"
	"os"

	. "github.com/m-ueno/raytracing"
)

func main() {
	var antialiasing = flag.Bool("aa", false, "Enable antialiasing (slow)")
	var output = flag.String("output", "", "Output filepath")
	var size = flag.Int("size", 0, "Image size in pixels")

	flag.Parse()
	if *size == 0 {
		log.Fatalln("Please provide `-size <size>`")
		os.Exit(1)
	}

	if *output == "" {
		log.Fatalln("Please provide `-output <path>`")
		os.Exit(1)
	}

	log.Println("start")
	log.Println("antialiasing:", *antialiasing)
	log.Println("output:", *output)
	log.Println("size:", *size)

	f, err := os.OpenFile(*output, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}
	defer f.Close()

	scene := NewScene33(*size)
	scene.Render(*antialiasing, f)

	log.Println("end")
}
