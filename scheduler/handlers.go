package scheduler

import (
	"net/http"
	"io/ioutil"
	"bytes"
)

func (s *TypeScheduler) HandleGET() error {

	s.Requester.URL = s.Config.URLGET
	s.Requester.Method = "GET"
	s.Requester.RadisDB = s.Config.RadisDB

	errReq := s.Requester.NewRequest()
	if errReq != nil {
		return errReq
	}

	errReadAll := s.Requester.ReadAll()
	if errReadAll != nil {
		return errReadAll
	}

	return nil
}

func (s *TypeScheduler) HandlePOST() error {

	s.Requester.URL = s.Config.URLPOST
	s.Requester.Method = "POST"
	s.Requester.RadisDB = s.Config.RadisDB

	errReq := s.Requester.NewRequest()
	if errReq != nil {
		return nil
	}

	return nil
}

func (r *TypeRequester) NewRequest() error {

	req, errReq := http.NewRequest(r.Method, r.URL, bytes.NewReader(r.Body))
	if errReq != nil {
		return nil
	}

	q := req.URL.Query()
	q.Add("namedb", r.RadisDB)
	req.URL.RawQuery = q.Encode()

	client := http.Client{}
	resp, errResp := client.Do(req)
	if errResp != nil {
		return nil
	}

	r.Response = resp

	return nil
}

func (r *TypeRequester) ReadAll() error {
	b, errReadAll := ioutil.ReadAll(r.Response.Body)
	if errReadAll != nil {
		return errReadAll
	}

	r.Body = b

	return nil
}