package main

type HTTPError struct {
	message string
	status  int
}

// Implements the error interface with a method - we'll cover interfaces and
// methods later
func (he *HTTPError) Error() {
	return he.message
}

func main() {
	r, err := httpResponse()
	// Runtime type check - we'll cover this later
	if he, ok := err.(*HTTPError); ok {
		sendHTTPError(err)
	}

	sendGoodResponse(r)
}

// START OMIT
func httpResponse() (string, error) {
	return "foo", nil
}

func sendHTTPError(e error) {}

func sendGoodResponse(r string) {}
