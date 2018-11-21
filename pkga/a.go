package pkga // import github.com/owais/gomodtest/pkga

import "github.com/uniplaces/carbon"

func Get() string {
	return carbon.Now().DateTimeString()
}
