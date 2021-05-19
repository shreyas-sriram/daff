package daff

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsResponseValid(t *testing.T) {
	log.SetOutput(ioutil.Discard)

	type input struct {
		received *http.Response
		expected *Response
	}

	tests := []struct {
		name string
		args input
		want bool
	}{
		{name: "returns true", args: input{&http.Response{Status: "200 OK", StatusCode: 200}, &Response{Status: 200}}, want: true},
		{name: "returns false", args: input{&http.Response{Status: "200 OK", StatusCode: 200}, &Response{Status: 403}}, want: false},
	}

	for _, test := range tests {
		got := isResponseValid(test.args.received, test.args.expected)
		assert.Equal(t, got, test.want, "got: %+v, want: %+v", got, test.want)
	}
}

func TestParseKeyValue(t *testing.T) {
	log.SetOutput(ioutil.Discard)

	type output struct {
		result map[string]string
		err    error
	}

	tests := []struct {
		name string
		args []string
		want output
	}{
		{name: "single element", args: []string{"a:b"}, want: output{map[string]string{"a": "b"}, nil}},
		{name: "multiple elements", args: []string{"a:b", "c:d"}, want: output{map[string]string{"a": "b", "c": "d"}, nil}},
		{name: "no elements", args: []string{}, want: output{map[string]string{}, nil}},
		{name: "error parsing", args: []string{"foo bar"}, want: output{nil, fmt.Errorf("error finding key and value")}},
	}

	for _, test := range tests {
		got, err := parseKeyValue(test.args)
		assert.Equal(t, err, test.want.err, "got: %+v, want: %+v", err, test.want.err)
		assert.Equal(t, got, test.want.result, "got: %+v, want: %+v", got, test.want.result)
	}
}
