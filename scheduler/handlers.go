package scheduler

import (
	"net/http"
	"io/ioutil"
	"bytes"
)

func (s *TypeSchaduler) HandleGET() ([]byte, error) {

	req, errReq := http.NewRequest("GET", s.Config.URLGET, nil)
	if errReq != nil {
		return nil, errReq
	}

	client := http.Client{}
	resp, errResp := client.Do(req)
	if errResp != nil {
		return nil, errResp
	}

	b, errReadAll := ioutil.ReadAll(resp.Body)
	if errReadAll != nil {
		return nil, errReadAll
	}

	return b, nil
}

func (s *TypeSchaduler) HandlePOST(data []byte) (*http.Response, error) {

	req, errReq := http.NewRequest("POST", s.Config.URLPOST, bytes.NewReader(data))
	if errReq != nil {
		return nil, errReq
	}

	client := http.Client{}
	resp, errResp := client.Do(req)
	if errResp != nil {
		return nil, errResp
	}

	return resp, nil
}
