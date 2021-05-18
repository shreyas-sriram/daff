package daff

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsResponseValid(t *testing.T) {
	type input struct {
		received *http.Response
		expected *Response
	}

	tests := []struct {
		name string
		args input
		want bool
	}{
		{name: "simple match", args: input{&http.Response{Status: "200 OK", StatusCode: 200}, &Response{Status: 200}}, want: true},
	}

	for _, test := range tests {
		got := isResponseValid(test.args.received, test.args.expected)
		assert.Equal(t, got, test.want, "got: %+v, want %+v", got, test.want)
	}
}
