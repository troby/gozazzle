package gozazzle

import (
	"fmt"
	"errors"
	"net/url"
	"net/http"
	"io/ioutil"
	"crypto/md5"
	"encoding/xml"
)

func (session *Session) fetch(method string) error {
	respStruct := new(Response)
	url := session.BaseUrl + "method=" + method
	url = url + "&vendorid=" + session.Id
	url = url + "&hash=" + session.Hash

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if err := xml.Unmarshal(respBytes, respStruct); err != nil {
		return err
	}
	switch respStruct.Code {
	case "ERROR":
		msg := fmt.Sprintf("%s\n", respBytes)
		err := errors.New(msg)
		return err
	case "SUCCESS":
	default:
		msg := fmt.Sprintf("unexpected zazzle code: %s\n", respStruct.Code)
		err := errors.New(msg)
		return err
	}
	session.Response = respStruct
	return nil
}

func hash(data string) string {
	signature	:= []byte(data)
	hashByte	:= fmt.Sprintf("%x", md5.Sum(signature))
	return string(hashByte)
}

func urlEncode(raw string) (string, error) {
	encoded := url.QueryEscape(raw)
	if len(encoded) > 200 {
		err := errors.New("activity restricted to 200 characters after encoding")
		return "", err
	}
	return encoded, nil
}
