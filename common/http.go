package common

import "net/http"

func Close(resp *http.Response) error {
	if resp != nil && resp.Body != nil {
		return resp.Body.Close()
	}
	return nil
}

func CheckClose(resp *http.Response, err *error) {
	if resp != nil && resp.Body != nil {
		cerr := resp.Body.Close()

		if *err == nil {
			*err = cerr
		}
	}
}
