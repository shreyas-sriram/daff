package daff

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateRequest(t *testing.T) {
	log.SetOutput(ioutil.Discard)

	type output struct {
		request *http.Request
		err     bool
	}

	challenge := Challenge{
		"http://localhost:5000/get", Request{"GET", []string{"Authorization:Bearer foobar"}, []string{"fuzz:buzz"}, ""}, Response{200},
	}

	expectedRequest, _ := http.NewRequest("GET", "http://localhost:5000/get", nil)
	expectedRequest.Header.Add("Authorization", "Bearer foobar")
	expectedRequest.AddCookie(&http.Cookie{Name: "fuzz", Value: "buzz"})

	tests := []struct {
		name string
		args Challenge
		want output
	}{
		{name: "request created successfully", args: challenge, want: output{expectedRequest, false}},
	}

	for _, test := range tests {
		got, err := createRequest(test.args)
		if test.want.err {
			assert.Error(t, err, "got: %+v, want: %+v", err, test.want.err)
		} else {
			assert.NoError(t, err, "got: %+v, want: %+v", err, test.want.err)
		}
		assert.Equal(t, got, test.want.request, "got: %+v, want %+v", got, test.want.request)
	}
}

func TestSendRequest(t *testing.T) {
	expectedRespStatus := 200

	testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(expectedRespStatus)
	}))
	defer func() { testServer.Close() }()

	req, err := http.NewRequest(http.MethodGet, testServer.URL, nil)
	assert.NoError(t, err)

	res, err := sendRequest(req)
	assert.NoError(t, err)
	assert.Equal(t, res.StatusCode, expectedRespStatus, "got: %+v, want %+v", res.StatusCode, expectedRespStatus)
}
