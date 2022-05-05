package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"

	"gopkg.in/go-playground/colors.v1"
)

var validHEX = regexp.MustCompile(`(?i)#[0-9a-f]{6}`)
var validRGB = regexp.MustCompile(`rgb\((?:([0-9]{1,2}|1[0-9]{1,2}|2[0-4][0-9]|25[0-5]), ?)(?:([0-9]{1,2}|1[0-9]{1,2}|2[0-4][0-9]|25[0-5]), ?)(?:([0-9]{1,2}|1[0-9]{1,2}|2[0-4][0-9]|25[0-5]))\)`)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Printf("%s\n", swap(scanner.Bytes()))
	}
}

// swap reads a []byte representing a line of text and replaces any HEX color
// codes with RGB color codes or vice versa.
//
// A HEX color code is in the form #F8F8F8 (lowercase is accepted). An RGB code
// is in the form 235, 100, 0 with or without spaces, though no more than one
// space and optionally with 0 padded numbers.
func swap(src []byte) []byte {
	if validHEX.Match(src) {
		return validHEX.ReplaceAllFunc(src, swapHEX)
	}
	return validRGB.ReplaceAllFunc(src, swapRGB)
}

// swapHEX replaces all HEX color codes with RGB color codes in a []byte.
func swapHEX(src []byte) []byte {
	color, err := colors.ParseHEX(string(src))
	if err != nil {
		return src
	}
	return []byte(color.ToRGB().String())
}

// swapRGB replaces all RGB color codes with HEX color codes in a []byte.
func swapRGB(src []byte) []byte {
	color, err := colors.ParseRGB(string(src))
	if err != nil {
		return src
	}
	return []byte(color.ToHEX().String())
}
