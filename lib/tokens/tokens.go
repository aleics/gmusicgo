package tokens

import (
	"fmt"
        "net/http"
        "net/url"
        "strings"
	"github.com/aleics/gmusicgo/lib/gmusicjson"
)
//Declaration of the Tokens struct
type Tokens struct {
	Xt string `json:"xt"`
}
func Init() *Tokens{
	t := new(Tokens)
	return t
}
func (t Tokens) GetXt() string{
	return t.Xt
}
func (t *Tokens) SetXt(s string){
	t.Xt = s
}
func (t *Tokens) MakeRequest(auth string) [2]string{

	hostname := "https://play.google.com"

	resource := "/music/listen"

	u, _ := url.ParseRequestURI(hostname)
	u.Path = resource
	urlStr := fmt.Sprintf("%v",u)
		
	cl := &http.Client{}

	r, err := http.NewRequest("GET", urlStr, nil) //GET Request
	if err != nil {
		return [2]string{"404 ERROR",""}
	}

	auth_header := "GoogleLogin auth=" + auth //Add the auth value on the Authorization header

	r.Header.Add("Authorization", auth_header)


	resp, err := cl.Do(r)
	if err != nil {
                return [2]string{"404 ERROR",""}
	}

	defer resp.Body.Close()

	responsehead := resp.Header

	xt := responsehead["Set-Cookie"][0][3:79] //Get the sjsaid token from the Response header


	t.Xt = xt
		
	response := [2]string{resp.Status, xt} //return the tokens and the status of the communication

	return response
}
func (t Tokens) SaveInfo(path string) bool{

	p := []string{path,"tokens.json"}
	jsonpath := strings.Join(p,"")
	
	_, err := gmusicjson.Export(t, jsonpath)
	if err != nil {
		fmt.Println(err)
		return false
	}

	return true
}
