package day06

import (
	"github.com/garethjevans/aoc-2022/pkg/file"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindMarker(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "one",
			args: args{input: "mjqjpqmgbljsphdztnvjfqwrcgsmlb"},
			want: 7,
		},
		{
			name: "two",
			args: args{input: "bvwbjplbgvbhsrlpgdmjqwftvncz"},
			want: 5,
		},
		{
			name: "three",
			args: args{input: "nppdvjthqldpwncqszvftbrmjlhg"},
			want: 6,
		},
		{
			name: "four",
			args: args{input: "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg"},
			want: 10,
		},
		{
			name: "five",
			args: args{input: "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw"},
			want: 11,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FindMarker(tt.args.input)
			if got != tt.want {
				t.Errorf("FindMarker() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindMessage(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "one",
			args: args{input: "mjqjpqmgbljsphdztnvjfqwrcgsmlb"},
			want: 19,
		},
		{
			name: "two",
			args: args{input: "bvwbjplbgvbhsrlpgdmjqwftvncz"},
			want: 23,
		},
		{
			name: "three",
			args: args{input: "nppdvjthqldpwncqszvftbrmjlhg"},
			want: 23,
		},
		{
			name: "four",
			args: args{input: "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg"},
			want: 29,
		},
		{
			name: "five",
			args: args{input: "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw"},
			want: 26,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FindMessage(tt.args.input)
			if got != tt.want {
				t.Errorf("FindMarker() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindMarkerForPart1(t *testing.T) {
	lines, err := file.AsLines("testdata/input.txt")
	assert.NoError(t, err)

	got := FindMarker(lines[0])
	t.Logf("Part1 > %d", got)
}

func TestFindMarkerForPart2(t *testing.T) {
	lines, err := file.AsLines("testdata/input.txt")
	assert.NoError(t, err)

	got := FindMessage(lines[0])
	t.Logf("Part2 > %d", got)
}
