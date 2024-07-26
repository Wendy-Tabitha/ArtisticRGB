package check

import (
	"ascii/types"
	"reflect"
	"testing"
)

func TestUsage(t *testing.T) {
	type args struct {
		args []string
	}
	tests := []struct {
		name string
		args args
		want types.Data
	}{
		{
			name: "Valid input with one argument",
			args: args{args: []string{"textOnly"}},
			want: types.Data{
				Text: "textOnly",
			},
		},
		{
			name: "Valid input with two arguments and valid --color flag",
			args: args{args: []string{"--color=blue", "textOnly"}},
			want: types.Data{
				Color: "blue",
				Text:  "textOnly",
			},
		},
		{
			name: "Valid input with three arguments and valid --color flag and color substring",
			args: args{args: []string{"--color=green", "substring", "textOnly"}},
			want: types.Data{
				Color:       "green",
				ColorSubstr: "substring",
				Text:        "textOnly",
			},
		},
		{
			name: "Valid input with four arguments and valid --color flag, color substring, and banner",
			args: args{args: []string{"--color=red", "substring", "textOnly", "bannerText"}},
			want: types.Data{
				Color:       "red",
				ColorSubstr: "substring",
				Text:        "textOnly",
				Banner:      "bannerText",
			},
		},
		{
			name: "Invalid input with more than four arguments",
			args: args{args: []string{"--color=blue", "substring", "textOnly", "bannerText", "extraArg"}},
			want: types.Data{}, // Expecting no valid Data
		},
		{
			name: "Invalid input with zero arguments",
			args: args{args: []string{}},
			want: types.Data{}, // Expecting no valid Data
		},
		{
			name: "Invalid input where the first argument is not a valid --color flag",
			args: args{args: []string{"--invalidFlag=blue", "textOnly"}},
			want: types.Data{}, // Expecting no valid Data
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Usage(tt.args.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Usage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExpressions(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name  string
		args  args
		want  bool
		want1 string
	}{
		{
			name:  "Valid color value",
			args:  args{s: "--color=blue"},
			want:  true,
			want1: "blue",
		},
		{
			name:  "Another valid color value",
			args:  args{s: "--color=red"},
			want:  true,
			want1: "red",
		},
		{
			name:  "Missing --color=",
			args:  args{s: "--no-color=blue"},
			want:  false,
			want1: "",
		},
		{
			name:  "Invalid pattern",
			args:  args{s: "color=blue"},
			want:  false,
			want1: "",
		},
		{
			name:  "Leading whitespace",
			args:  args{s: "   --color=green"},
			want:  false,
			want1: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := Expressions(tt.args.s)
			if got != tt.want {
				t.Errorf("Expressions() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Expressions() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
