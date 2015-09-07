package request

import (
	//"fmt"
	"net/http"
	//"net/url"
	//"encoding/json"
	"os"
	"time"
	//"bytes"
	"errors"
	"bytes"
)

type Request struct{
	url string
	response http.Response
	status string
}
func (r *Request) MakeCall(_type string, auth string, xt string, urlStr string, body string, suppHeaders map[string]string) (*http.Response, string, error) {

	r.url = urlStr

	parameters := bytes.NewBufferString(body)
	auth_header := "GoogleLogin auth=" + auth

	expire := time.Now().AddDate(0, 0, 1)
	cookie := http.Cookie{"xt", xt, "/", "music.google.com", expire, expire.Format(time.UnixDate), 0, true, true, "xt=xttoken", []string{"xt=xttoken"}}

	cl := &http.Client{}

	switch(_type){
		case "GET":{
			req, err := http.NewRequest("GET", urlStr, parameters)
			if err != nil {
				os.Exit(1)
				return &http.Response{}, "", errors.New("GET request error: error creating request")
			}
			req.Header.Add("Authorization", auth_header)

			for key, value := range(suppHeaders){
				req.Header.Add(key, value)
			}

			req.AddCookie(&cookie)
		
			resp, err := cl.Do(req)
			if err != nil {
				os.Exit(1)
				return &http.Response{}, "", errors.New("GET request error: error making request")
			}
			defer resp.Body.Close()		
			
			r.response = *resp
			r.status = resp.Status
			
			return resp, resp.Status, nil
			break
		}
		case "POST":{
			req, err := http.NewRequest("POST", urlStr, parameters)
			if err != nil {
				os.Exit(1)
				return &http.Response{}, "", errors.New("POST request error: error creating request")
			}
			req.Header.Add("Authorization", auth_header)

			for key, value := range(suppHeaders){
				req.Header.Add(key, value)
			}


			req.AddCookie(&cookie)

			resp, err := cl.Do(req)
			if err != nil {
				os.Exit(1)
				return &http.Response{}, "", errors.New("POST request error: error making request")
			}
			defer resp.Body.Close()

			r.response = *resp
			r.status = resp.Status

			return resp, resp.Status, nil
			break
		}
		case "HEAD":{
			req, err := http.NewRequest("HEAD", urlStr, parameters)
			if err != nil {
				os.Exit(1)
				return &http.Response{}, "", errors.New("HEAD request error: error creating request")
			}
			req.Header.Add("Authorization", auth_header)
			for key, value := range(suppHeaders){
				req.Header.Add(key, value)
			}

			req.AddCookie(&cookie)

			resp, err := cl.Do(req)
			if err != nil {
				os.Exit(1)
				return &http.Response{}, "", errors.New("HEAD request error: error making request")
			}
			defer resp.Body.Close()

			r.response = *resp
			r.status = resp.Status

			return resp, resp.Status, nil
			break
		}
		case "PUT":{
			req, err := http.NewRequest("PUT", urlStr, parameters)
			if err != nil {
				os.Exit(1)
				return &http.Response{}, "", errors.New("PUT request error: error creating request")
			}
			req.Header.Add("Authorization", auth_header)
			for key, value := range(suppHeaders){
				req.Header.Add(key, value)
			}

			req.AddCookie(&cookie)

			resp, err := cl.Do(req)
			if err != nil {
				os.Exit(1)
				return &http.Response{}, "", errors.New("PUT request error: error making request")
			}
			defer resp.Body.Close()
			
			r.response = *resp
			r.status = resp.Status

			return resp, resp.Status, nil
			break
		}
	}
	return &http.Response{}, "", errors.New("Error: no request selected")
}
