package daff

import (
	"io/ioutil"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseFile(t *testing.T) {
	log.SetOutput(ioutil.Discard)

	type output struct {
		config *Config
		err    bool
	}

	config := &Config{}
	expectedConfig := &Config{
		map[string]Challenge{
			"chall-get": {"http://localhost:5000/get", Request{"GET", []string{"Authorization:Bearer foobar"}, []string(nil), ""}, Response{200}},
		},
	}

	tests := []struct {
		name string
		args string
		want output
	}{
		{name: "config not found", args: "oops.yaml", want: output{&Config{nil}, true}},
		{name: "config parsed successfully", args: "../../example/test-config.yaml", want: output{expectedConfig, false}},
	}

	for _, test := range tests {
		err := config.parseFile(test.args)
		if test.want.err {
			assert.Error(t, err, "got: %+v, want: %+v", err, test.want.err)
		} else {
			assert.NoError(t, err, "got: %+v, want: %+v", err, test.want.err)
		}
		assert.Equal(t, config, test.want.config, "got: %+v, want %+v", config, test.want.config)
	}
}
