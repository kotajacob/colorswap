package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	log.SetPrefix("")
	log.SetFlags(0)

	hexPtr := flag.Bool("hex", false, "convert to hex")
	rgbPtr := flag.Bool("rgb", false, "convert to rgb")
	rgbaPtr := flag.Bool("rgba", false, "convert to rgba")
	vec3Ptr := flag.Bool("vec3", false, "convert to vec3")
	vec4Ptr := flag.Bool("vec4", false, "convert to vec4")
	flag.Parse()

	format, err := getFormat(
		hexPtr,
		rgbPtr,
		rgbaPtr,
		vec3Ptr,
		vec4Ptr,
	)
	if err != nil {
		log.Fatalln(err)
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Printf("%s\n", Swap(scanner.Bytes(), format))
	}
}

func getFormat(opts ...*bool) (Format, error) {
	for i, op := range opts {
		if *op {
			return Format(i), nil
		}
	}
	return 0, fmt.Errorf("no output format selected")
}
