package main

import (
	"bytes"
	"fmt"
	"image/color"
	"math"
	"regexp"
)

type Format uint8

const (
	Hex Format = iota
	RGB
	RGBA
	Vec3
	Vec4
)

const (
	hexFormat      = "#%02x%02x%02x"
	hexShortFormat = "#%1x%1x%1x"
	rgbFormat      = "rgb(%d,%d,%d)"
	rgbaFormat     = "rgba(%d,%d,%d,%d)"
	vec3Format     = "vec3(%f,%f,%f)"
	vec4Format     = "vec4(%f,%f,%f,%f)"
)

var matchColor = regexp.MustCompile(`(?i)#(?:([0-9a-f]{1,2})([0-9a-f]{1,2})([0-9a-f]{1,2}))|rgba\((?:([0-9]{1,2}|1[0-9]{1,2}|2[0-4][0-9]|25[0-5]), ?)(?:([0-9]{1,2}|1[0-9]{1,2}|2[0-4][0-9]|25[0-5]), ?)(?:([0-9]{1,2}|1[0-9]{1,2}|2[0-4][0-9]|25[0-5]), ?)(?:([0-9]{1,2}|1[0-9]{1,2}|2[0-4][0-9]|25[0-5]))\)|rgb\((?:([0-9]{1,2}|1[0-9]{1,2}|2[0-4][0-9]|25[0-5]), ?)(?:([0-9]{1,2}|1[0-9]{1,2}|2[0-4][0-9]|25[0-5]), ?)(?:([0-9]{1,2}|1[0-9]{1,2}|2[0-4][0-9]|25[0-5]))\)|vec4\((?:([0-1]\.?\d*), ?([0-1]\.?\d*), ?([0-1]\.?\d*), ?([0-1]\.?\d*)\))|vec3\((?:([0-1]\.?\d*), ?([0-1]\.?\d*), ?([0-1]\.?\d*)\))`)

// Swap reads a []byte representing a block of text and replaces any color
// codes with color codes in the specified Format.
func Swap(src []byte, format Format) []byte {
	switch format {
	case Hex:
		return matchColor.ReplaceAllFunc(src, toHex)
	case RGB:
		return matchColor.ReplaceAllFunc(src, toRGB)
	case RGBA:
		return matchColor.ReplaceAllFunc(src, toRGBA)
	case Vec3:
		return matchColor.ReplaceAllFunc(src, toVec3)
	case Vec4:
		return matchColor.ReplaceAllFunc(src, toVec4)
	default:
		return matchColor.ReplaceAllFunc(src, toHex)
	}
}

// toHex replaces all color codes with Hex color codes in a []byte.
func toHex(src []byte) []byte {
	c := Parse(src)
	r, g, b, _ := c.RGBA()
	return []byte(fmt.Sprintf(hexFormat, uint8(r), uint8(g), uint8(b)))
}

// toRGB replaces all color codes with NRGB color codes in a []byte.
func toRGB(src []byte) []byte {
	c := Parse(src)
	r, g, b := c.R, c.G, c.B
	return []byte(fmt.Sprintf(rgbFormat, r, g, b))
}

// toRGBA replaces all color codes with NRGBA color codes in a []byte.
func toRGBA(src []byte) []byte {
	c := Parse(src)
	r, g, b, a := c.R, c.G, c.B, c.A
	return []byte(fmt.Sprintf(rgbaFormat, r, g, b, a))
}

// toVec3 replaces all color codes with Vec3 color codes in a []byte.
func toVec3(src []byte) []byte {
	c := Parse(src)
	r, g, b := c.R, c.G, c.B
	rf := float64(r) / 255
	gf := float64(g) / 255
	bf := float64(b) / 255
	return []byte(fmt.Sprintf(vec3Format, rf, gf, bf))
}

// toVec4 replaces all color codes with Vec4 color codes in a []byte.
func toVec4(src []byte) []byte {
	c := Parse(src)
	r, g, b, a := c.R, c.G, c.B, c.A
	rf := float64(r) / 255
	gf := float64(g) / 255
	bf := float64(b) / 255
	af := float64(a) / 255
	return []byte(fmt.Sprintf(vec4Format, rf, gf, bf, af))
}

func Parse(src []byte) color.NRGBA {
	if len(src) < 4 {
		// Invalid color!
		// If this happens the matchColor regex above is incorrect.
		panic("invalid color: less than 4 bytes")
	}

	s := string(bytes.ToLower(src))
	switch {
	case s[:1] == "#":
		return parseHex(s)
	case s[:4] == "rgba":
		return parseRGBA(s)
	case s[:3] == "rgb":
		return parseRGB(s)
	case s[:4] == "vec3":
		return parseVec3(s)
	case s[:4] == "vec4":
		return parseVec4(s)
	default:
		panic("invalid color: unknown color prefix")
	}
}

func parseHex(s string) color.NRGBA {
	var r, g, b uint8
	switch len(s) {
	case 4:
		fmt.Sscanf(s, hexShortFormat, &r, &g, &b)
		r *= 17
		g *= 17
		b *= 17
	case 7:
		fmt.Sscanf(s, hexFormat, &r, &g, &b)
	default:
		panic("invalid hex color: was not 4 or 7 bytes")
	}
	return color.NRGBA{R: r, G: g, B: b, A: 255}
}

func parseRGB(s string) color.NRGBA {
	var r, g, b uint8
	fmt.Sscanf(s, rgbFormat, &r, &g, &b)
	return color.NRGBA{R: r, G: g, B: b, A: 255}
}

func parseRGBA(s string) color.NRGBA {
	var r, g, b, a uint8
	fmt.Sscanf(s, rgbaFormat, &r, &g, &b, &a)
	return color.NRGBA{R: r, G: g, B: b, A: a}
}

func parseVec3(s string) color.NRGBA {
	fmt.Println(s)
	var r, g, b float64
	fmt.Sscanf(s, vec3Format, &r, &g, &b)
	r = math.Round(r * 255)
	g = math.Round(g * 255)
	b = math.Round(b * 255)
	return color.NRGBA{R: uint8(r), G: uint8(g), B: uint8(b), A: 255}
}

func parseVec4(s string) color.NRGBA {
	fmt.Println(s)
	var r, g, b, a float64
	fmt.Sscanf(s, vec4Format, &r, &g, &b, &a)
	r = math.Round(r * 255)
	g = math.Round(g * 255)
	b = math.Round(b * 255)
	a = math.Round(a * 255)
	return color.NRGBA{R: uint8(r), G: uint8(g), B: uint8(b), A: uint8(a)}
}
