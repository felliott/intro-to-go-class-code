package main

type HTTPError struct {
	message string
	status  int
}

// Implements the error interface with a method - we'll cover interfaces and
// methods later
func (he *HTTPError) String() {
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
