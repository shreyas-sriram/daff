package daff

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

// isResponseValid validates the received response against the config
//
// This is a separate function to accommodate enhancements like comparing response
// length, looking for keywords etc.
func isResponseValid(received *http.Response, expected *Response) bool {
	return received.StatusCode == expected.Status
}

// parseKeyValue parses and returns a map of the headers or cookies
func parseKeyValue(s []string) (map[string]string, error) {
	m := make(map[string]string)

	for _, _s := range s {
		parts := strings.Split(_s, headerDelimiter)
		if len(parts) != 2 {
			log.Printf("Error finding key and value\n")
			return nil, fmt.Errorf("error finding key and value")
		}
		m[parts[0]] = parts[1]
	}

	return m, nil
}
