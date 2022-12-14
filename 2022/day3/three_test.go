package main

import (
	"testing"

	aoc2022 "github.com/pineman/aoc2022"
)

// This file was auto-generated using vscode "Go: Generate Unit Tests"

func Test_asciiToPriority(t *testing.T) {
	type args struct {
		item rune
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{'A'}, 27},
		{"", args{'B'}, 28},
		{"", args{'Z'}, 52},
		{"", args{'a'}, 1},
		{"", args{'b'}, 2},
		{"", args{'z'}, 26},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := priority(tt.args.item); got != tt.want {
				t.Errorf("asciiToPriority() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_partOne(t *testing.T) {
	type args struct {
		input []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{aoc2022.GetInput(3)}, 8105},
		{"", args{aoc2022.GetBigBoyInput(3)}, 91038443},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := partOne(tt.args.input); got != tt.want {
				t.Errorf("partOne() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_partTwo(t *testing.T) {
	type args struct {
		input []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{aoc2022.GetInput(3)}, 2363},
		{"", args{aoc2022.GetBigBoyInput(3)}, 27355441},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := partTwo(tt.args.input); got != tt.want {
				t.Errorf("partTwo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_partTwoBigBoy(b *testing.B) {
	input := aoc2022.GetBigBoyInput(3)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		partTwo(input)
	}
}

func Benchmark_partTwoBigBoy_set(b *testing.B) {
	input := aoc2022.GetBigBoyInput(3)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		partTwo_set(input)
	}
}
