package azure

import "github.com/Azure/go-autorest/autorest"

func isNotFoundError(err error) bool {
	if detailedError, ok := err.(autorest.DetailedError); ok {
		intSC, ok := detailedError.StatusCode.(int)
		return ok && intSC == 404
	}
	return false
}
