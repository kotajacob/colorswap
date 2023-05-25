package main

import (
	"image/color"
	"reflect"
	"testing"
)

func TestSwap(t *testing.T) {
	type test struct {
		input  string
		want   string
		format Format
	}

	tests := []test{
		{
			input:  "rgb(18, 52, 86)",
			want:   "#123456",
			format: Hex,
		},
		{
			input:  "#123456",
			want:   "rgb(18,52,86)",
			format: RGB,
		},
		{
			input:  "rgba(20,20,20)",
			want:   "rgba(20,20,20)",
			format: RGB,
		},
		{
			input:  "rgb(20,20,20)",
			want:   "rgba(20,20,20,255)",
			format: RGBA,
		},
		{
			input:  "rgba(20,20,20,20)",
			want:   "rgba(20,20,20,20)",
			format: RGBA,
		},
		{
			input:  "rgb(20,20,20)",
			want:   "vec3(0.078431,0.078431,0.078431)",
			format: Vec3,
		},
		{
			input:  "rgb(20,40,60)",
			want:   "vec3(0.078431,0.156863,0.235294)",
			format: Vec3,
		},
		{
			input:  "rgb(20,40,60)",
			want:   "vec4(0.078431,0.156863,0.235294,1.000000)",
			format: Vec4,
		},
	}

	for _, tc := range tests {
		got := Swap([]byte(tc.input), tc.format)
		if string(got) != tc.want {
			t.Fatalf("expected: %s, got: %s", tc.want, got)
		}
	}
}

func TestParse(t *testing.T) {
	type test struct {
		input  string
		want   color.NRGBA
		format Format
	}

	tests := []test{
		{
			input: "#123456",
			want:  color.NRGBA{18, 52, 86, 255},
		},
		{
			input: "#eee",
			want:  color.NRGBA{238, 238, 238, 255},
		},
		{
			input: "#EEE",
			want:  color.NRGBA{238, 238, 238, 255},
		},
		{
			input: "#eeeeee",
			want:  color.NRGBA{238, 238, 238, 255},
		},
		{
			input: "#EEEEEE",
			want:  color.NRGBA{238, 238, 238, 255},
		},
		{
			input: "rgb(18,52,86)",
			want:  color.NRGBA{18, 52, 86, 255},
		},
		{
			input: "rgb(0,0,0)",
			want:  color.NRGBA{0, 0, 0, 255},
		},
		{
			input: "rgb(255, 0, 255)",
			want:  color.NRGBA{255, 0, 255, 255},
		},
		{
			input: "rgb(10, 0, 255)",
			want:  color.NRGBA{10, 0, 255, 255},
		},
		{
			input: "rgba(10, 0, 255, 32)",
			want:  color.NRGBA{10, 0, 255, 32},
		},
		{
			input: "vec3(0.224, 0.0, 1.0)",
			want:  color.NRGBA{57, 0, 255, 255},
		},
		{
			input: "vec4(0.224, 0.0, 0.75, 0.60)",
			want:  color.NRGBA{57, 0, 191, 153},
		},
	}

	for _, tc := range tests {
		got := Parse([]byte(tc.input))
		if !reflect.DeepEqual(got, tc.want) {
			t.Fatalf("expected: %v, got: %v", tc.want, got)
		}
	}
}
