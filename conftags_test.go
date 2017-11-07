package conftags

import (
	"os"
	"reflect"
	"testing"

	"github.com/alecthomas/assert"
)

func TestParse(t *testing.T) {
	var err error
	type Mystruct struct {
		Myfield int `env:"MYFIELD" default:"10"`
	}
	in := Mystruct{}
	// Test the fallback to the default field tag
	err = Parse(&in)
	assert.NoError(t, err, "error should be nil")
	assert.Equal(t, in.Myfield, 10)

	// Test the environment variable has an impact on the struct
	os.Setenv("MYFIELD", "11")
	err = Parse(&in)
	assert.NoError(t, err, "error should be nil")
	assert.Equal(t, in.Myfield, 11)

	// Test that an error is returned when the type is not compatible
	os.Setenv("MYFIELD", "random")
	err = Parse(&in)
	assert.Error(t, err)

	type erroneous struct {
		Myfield int `default:"x"`
	}
	// Test the fallback to the default field tag
	err = Parse(&erroneous{})
	assert.Error(t, err, "should throw back an error")
}

func Test_isZero(t *testing.T) {
	var nilslice []string
	var nilmap map[string]string
	type args struct {
		v reflect.Value
	}
	tests := []struct {
		name string
		args reflect.Value
		want bool
	}{
		{"nil slice should return true", reflect.ValueOf(nilslice), true},
		{"nil map should return true", reflect.ValueOf(nilmap), true},
		{"slice should be false", reflect.ValueOf([]string{}), false},
		{"map should be false", reflect.ValueOf(new(map[bool]bool)), false},
		{"struct instance should be false", reflect.ValueOf(struct{ One bool }{true}), false},
		{"empty struct should be true", reflect.ValueOf(struct{ One bool }{}), true},
		{"empty array should be true", reflect.ValueOf([0]string{}), true},
		{"array should be false", reflect.ValueOf([1]string{"hi"}), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isZero(tt.args); got != tt.want {
				t.Errorf("isZero() = %v, want %v", got, tt.want)
			}
		})
	}
}
