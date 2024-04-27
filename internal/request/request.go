package request

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

func Get(u string) (string, error) {
	resp, err := http.Get(u)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return "", err
	}

	if resp.StatusCode != 200 {
		return "", errors.New(fmt.Sprintf("Non-200 response from api for %s", u))
	}
	return string(body), nil

}
