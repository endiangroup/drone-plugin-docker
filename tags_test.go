package docker

import (
	"reflect"
	"testing"
)

func Test_stripTagPrefix(t *testing.T) {
	var tests = []struct {
		Before string
		After  string
	}{
		{"refs/tags/1.0.0", "1.0.0"},
		{"refs/tags/v1.0.0", "1.0.0"},
		{"v1.0.0", "1.0.0"},
	}

	for _, test := range tests {
		got, want := stripTagPrefix(test.Before), test.After
		if got != want {
			t.Errorf("Got tag %s, want %s", got, want)
		}
	}
}

func TestDefaultTags(t *testing.T) {
	var tests = []struct {
		Before string
		After  []string
	}{
		{"", []string{"dev"}},
		{"refs/heads/master", []string{"latest"}},
		{"refs/heads/dev", []string{"dev"}},
		{"refs/heads/feature/new-thing", []string{"feature/new-thing"}},
		{"refs/tags/0.9.0", []string{"latest", "0.9", "0.9.0"}},
		{"refs/tags/1.0.0", []string{"latest", "1", "1.0", "1.0.0"}},
		{"refs/tags/v1.0.0", []string{"latest", "1", "1.0", "1.0.0"}},
		{"refs/tags/v1.0.0-alpha.1", []string{"1.0.0-alpha.1"}},

		// malformed or errors
		{"refs/tags/x1.0.0", []string{"dev"}},
		{"v1.0.0", []string{"dev"}},
	}

	for _, test := range tests {
		got, want := DefaultTags(test.Before), test.After
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Default tags for \"%v\" failed. Got tag %v, want %v", test.Before, got, want)
		}
	}
}

func Test_stripHeadPrefix(t *testing.T) {
	type args struct {
		ref string
	}
	tests := []struct {
		args args
		want string
	}{
		{
			args: args{
				ref: "refs/heads/master",
			},
			want: "master",
		},
	}
	for _, tt := range tests {
		if got := stripHeadPrefix(tt.args.ref); got != tt.want {
			t.Errorf("stripHeadPrefix() = %v, want %v", got, tt.want)
		}
	}
}
