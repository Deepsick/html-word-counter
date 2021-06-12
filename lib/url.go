package lib

import (
	"errors"
	"io"
	"net/http"
	"net/url"
	"os"
)

func ValidateUrl(link *string) error {
	_, err := url.ParseRequestURI(*link)
	if err != nil {
		return err
	}

	return nil
}

func GetLink(args *[]string) (*string, error) {
	if len(*args) < ARGS_COUNT {
		return nil, errors.New("please, provide url as cli argument")
	}

	return &os.Args[1], nil
}

func GetHtml(link *string) (*io.ReadCloser, error) {
	response, err := http.Get(*link)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != STATUS_OK {
		return nil, errors.New("something went wrong during request")
	}

	return &response.Body, nil
}
