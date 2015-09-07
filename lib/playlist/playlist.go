package playlist

import (
        "fmt"
        "net/http"
        "net/url"
        "io/ioutil"
        "strings"                                                                                                                     
        "encoding/json"
        "os"
	"strconv"
	"time"
	"bytes"
	"github.com/aleics/gmusicgo/lib/gmusicjson"
	"github.com/aleics/gmusicgo/lib/request"
)

type Playlist struct{
        Kind string `json:"kind"`
        Id string `json:"id"`
        CreationTimestamp string `json:"creationTimestamp"`
        LastModifiedTimestamp string `json:"lastModifiedTimestamp"`
        Deleted bool `json:"deleted"`
        Name string `json:"name"`
        TypePlaylist string `json:typePlaylist`
}
func Init() *Playlist{
        p := new(Playlist)
        return p
}
func ArrayInit() (p *[]Playlist){
        return p
}
func (p Playlist) GetKind() string {
        return p.Kind
}
func (p *Playlist) SetKind(kind string) {
        p.Kind = kind
}
func (p Playlist) GetId() string {
        return p.Id
}
func (p *Playlist) SetId(id string) {
	p.Id = id
}
func (p Playlist) GetCreationTimestamp() string {
        return p.CreationTimestamp
}
func (p *Playlist) SetCreationTimestamp(creationTimestamp string) {
        p.CreationTimestamp = creationTimestamp
}
func (p Playlist) GetLastModifiedTimestamp() string {
        return p.LastModifiedTimestamp
}
func (p *Playlist) SetLastModifiedTimestamp(lastModifiedTimestamp string) {
        p.LastModifiedTimestamp = lastModifiedTimestamp
}
func (p Playlist) GetDeleted() bool {
        return p.Deleted
}
func (p *Playlist) SetDeleted(deleted bool) {
        p.Deleted = deleted
}
func (p Playlist) GetName() string {
        return p.Name
}
func (p *Playlist) SetName(name string) {
        p.Name = name
}
func (p Playlist) GetType() string {
        return p.TypePlaylist
}
func (p *Playlist) SetType(typePlaylist string) {
        p.TypePlaylist = typePlaylist
}
func (p *Playlist) NewPlaylist(kind string, id string, creationTimestamp string, lastModifiedTimestamp string, deleted bool, name string, typePlaylist string){
	
	p.Kind = kind
	p.Id = id
	p.CreationTimestamp = creationTimestamp
	p.LastModifiedTimestamp = lastModifiedTimestamp
	p.Deleted = deleted
	p.Name = name
	p.TypePlaylist = typePlaylist
}
func (p Playlist) ToMap() map[string]string {

        ret := make(map[string]string)

        ret["kind"] = p.Kind
        ret["id"] = p.Id
        ret["creationTimestamp"] = p.CreationTimestamp
        ret["lastModifiedTimestamp"] = p.LastModifiedTimestamp
        ret["deleted"] = strconv.FormatBool(p.Deleted)
        ret["name"] = p.Name
        ret["typePlaylist"] = p.TypePlaylist
        
        return ret
}
func PlaylistsRequest(auth string, path string) []Playlist{

        hostname := "https://www.googleapis.com"
        resource := "/sj/v1beta1/playlists"

        u, _ := url.ParseRequestURI(hostname)
        u.Path = resource
        urlStr := fmt.Sprintf("%v",u)

        cl := &http.Client{}

        r, err := http.NewRequest("GET", urlStr, nil)
        if err != nil {
                os.Exit(1)
        }

        auth_header := "GoogleLogin auth=" + auth

        r.Header.Add("Authorization", auth_header)


        resp, err := cl.Do(r)
        if err != nil {
		os.Exit(1)
        }

        defer resp.Body.Close()

        b, err := ioutil.ReadAll(resp.Body) //Get the body of the response                                                               
        if err != nil { //Error management                                                                                               
                os.Exit(1)
        }

        var f interface{}
        json.Unmarshal(b, &f)

        m := f.(map[string]interface{})

        playlistsmap := m["data"]
        playlists := playlistsmap.(map[string]interface{})


        itemsmap := playlists["items"]
        items := itemsmap.([]interface{})

        var singleitem map[string]interface{}

        length := len(items)
        arrayplaylists := make([]Playlist,length)

        for i := 0; i < length; i++ {

                singleitem = items[i].(map[string]interface{})


                arrayplaylists[i].NewPlaylist(singleitem["kind"].(string), singleitem["id"].(string), singleitem["creationTimestamp"].(string), singleitem["lastModifiedTimestamp"].(string), singleitem["deleted"].(bool), singleitem["name"].(string), singleitem["type"].(string))

        }

	pa := []string{path,"playlists.json"}
        jsonpath := strings.Join(pa,"")

        _, err = gmusicjson.Export(arrayplaylists, jsonpath)
	if err != nil {
		fmt.Println("Error exporting Playlist: ")
		fmt.Println(err)
	}

        return arrayplaylists
}
func (p *Playlist) CreatePlaylist(auth string, xt string, name string, description string, public bool) bool{

	hostname := "https://play.google.com"
	resource := "/music/services/createplaylist" + "?u=0&xt=" + xt  + "&format=jsarray"

        u, _ := url.ParseRequestURI(hostname)
        u.Path = resource
        urlStr := hostname + resource

	body_r := `[["",1],[null,"` + name  + `", "` + description + `", []]]`

	req := request.Request{}
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"

	_, _, err := req.MakeCall("POST", auth, xt, urlStr, body_r, headers)
	if err != nil {
		os.Exit(1)
		return false
	}
	return true
}

func (p *Playlist) LoadUserPlaylist(auth string, xt string) bool{
	hostname := "https://play.google.com"
        resource := "/music/services/loaduserplaylist" + "?u=0&xt=" + xt  + "&format=jsarray"

        u, _ := url.ParseRequestURI(hostname)
        u.Path = resource
        urlStr := hostname + resource

        cl := &http.Client{}

        body_r := `[["", 1, ""],[""]]`
        parameters := bytes.NewBufferString(body_r)

        r, err := http.NewRequest("POST", urlStr, parameters)
        if err != nil {
                os.Exit(1)
                return false
        }

        auth_header := "GoogleLogin auth=" + auth

        r.Header.Add("Authorization", auth_header)
        r.Header.Add("Content-Type", "application/json")


        expire := time.Now().AddDate(0, 0, 1)
        cookie := http.Cookie{"xt", xt, "/", "music.google.com", expire, expire.Format(time.UnixDate), 0, true, true, "xt=xttoken", []string{"xt=xttoken"}}
        r.AddCookie(&cookie)

        resp, err := cl.Do(r)
        if err != nil {
                panic(err)
                return false
        }

	defer resp.Body.Close()

	if resp.Status != "200 OK" {
                return false
        }

	b, err := ioutil.ReadAll(resp.Body) //Get the body of the response                                                                                                                      
	if err != nil { //Error management                                                                                                                                                      
		return false
	}
	result := make(map[string]interface{})
        json.Unmarshal(b, &result)

	for i, el := range result {
		fmt.Println(i)
		fmt.Println(el.(string))
	}


	return true
}

func (p *Playlist) DeletePlaylist(auth string, xt string, id string) bool{
	
	Url, err := url.Parse("https://play.google.com")
	Url.Path += "/music/services/deleteplaylist"
	
	parameters := url.Values{}

	parameters.Add("u", "0")
	parameters.Add("xt", xt)
	
	Url.RawQuery = parameters.Encode()

	cl := &http.Client{}

	body_r := []byte(`{"id": "` + id  + `", "requestCause": 1, "requestType": 1, "sessionId": ""}`)
	
	req, err := http.NewRequest("POST", Url.String(), bytes.NewBuffer(body_r))
	if err != nil {
		os.Exit(1)
		return false
	}
	
	auth_header := "GoogleLogin auth=" + auth

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded;charset=UTF-8")
	req.Header.Add("Authorization", auth_header)
	req.Header.Add("referer", "https://play.google.com/music/listen")
	
	expire := time.Now().AddDate(0, 0, 1)
        cookie := http.Cookie{"xt", xt, "/", "music.google.com", expire, expire.Format(time.UnixDate), 0, true, true, "xt=xttoken", []string{"xt=xttoken"}}
        req.AddCookie(&cookie)

	resp, err := cl.Do(req)
	if err != nil {
		panic(err)
		return false
	}
	fmt.Println(req)
	fmt.Println(resp)
	defer resp.Body.Close()
	
	if resp.Status != "200 OK" {
		return false
	}
	return true
}

func (p *Playlist) Print(){
	fmt.Print("kind: ")
        fmt.Println(p.Kind)
	fmt.Print("id: ")
        fmt.Println(p.Id)
        fmt.Print("creationTimestamp: ")
        fmt.Println(p.CreationTimestamp)
        fmt.Print("lastModifiedTimestamp: ")
	fmt.Println(p.LastModifiedTimestamp)
        fmt.Print("deleted: ")
        fmt.Println(p.Deleted)
	fmt.Print("name: ")
        fmt.Println(p.Name)
	fmt.Print("type: ")
        fmt.Println(p.TypePlaylist)
}
