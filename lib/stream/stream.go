package stream

import(
	"fmt"
	"net/http"
	"net/url"
	"io/ioutil"
	"errors"
	"os"
	"time"
	"encoding/json"
)

type Stream struct{
		streamUrl string
		audioFile string
}
func Init() *Stream{
	s := new(Stream)
	return s
}
func (s Stream) StreamUrl() string{
	return s.streamUrl
}
func (s Stream) AudioFile() string{
	return s.audioFile
}
func (s *Stream) SetStreamUrl(streamUrl string) {
	s.streamUrl = streamUrl
}
func (s *Stream) SetAudioFile(audioFile string) {
	s.audioFile = audioFile 
}
func (s *Stream) StreamRequest(auth string, xt string, songid string, path string) error{
	b := s.StreamAudioRequest(s.StreamUrlRequest(auth, xt, songid, path + "streamurl_" + songid + ".json"), path + "file_" + songid + ".mp3")
	if b != true {
		return errors.New("Error stream request")
	}
	return nil
}
func (s *Stream) StreamAudioRequest(streamUrl string, path string) bool{
	
	out := true
	urlStr := streamUrl
	
	cl := &http.Client{}
	
	r, err := http.NewRequest("GET", urlStr, nil)
	if err != nil {
		panic(err)
	}
	r.Header.Add("Referer", "https://play.google.com/music/listen")
	resp, err := cl.Do(r)
	if err != nil {
		panic(err)
		out = false
	}
	
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body) //Get the body of the response

	if path != "" {
		f, err := os.Create(path)
		if err != nil {
			panic(err)
			out = false
		}

		defer func() {
			if err := f.Close(); err != nil {
				panic(err)
				out = false
			}
		}()
		
		_, err = f.Write([]byte(b))
		if err != nil {
			panic(err)
			out = false
		}
	}

	return out

}
func (s *Stream) StreamUrlRequest(auth string, xt string, songid string, path string) string{

	hostname := "https://play.google.com"
	resource := "/music/play?u=0&songid=" + songid + "&pt=e"
	
	u, _ := url.ParseRequestURI(hostname)
	u.Path = resource
	urlStr := fmt.Sprintf("%v",u)
	
	urlStr = hostname + resource
	
	cl := &http.Client{}
	
	r, err := http.NewRequest("GET", urlStr, nil)
	if err != nil {
		panic(err)
	}
	
	auth_header := "GoogleLogin auth=" + auth
	
	r.Header.Add("Authorization", auth_header)
	r.Header.Add("Content-Type", "application/json")


	expire := time.Now().AddDate(0, 0, 1)
	cookie := http.Cookie{"sjsaid", "", "/", "music.google.com", expire, expire.Format(time.UnixDate), 0, true, true, "sjsaid=sjsaidtoken", []string{"sjsaid=sjsaidtoken"}}
	r.AddCookie(&cookie)
	cookie = http.Cookie{"xt", xt, "/", "music.google.com", expire, expire.Format(time.UnixDate), 0, true, true, "xt=xttoken", []string{"xt=xttoken"}}
	r.AddCookie(&cookie)

	resp, err := cl.Do(r)
	if err != nil {
		panic(err)
        }

	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body) //Get the body of the response                                             
	
	result := make(map[string]interface{})
	json.Unmarshal(b, &result) //Decode the json body response

        if err != nil { //Error management                   
		panic(err)
        }
	
	s.SetStreamUrl(result["url"].(string))

	if path != "" {
		f, err := os.Create(path)
		if err != nil {
			panic(err)
		}

		defer func() {
			if err := f.Close(); err != nil {
				panic(err)
			}
		}()

		_, err = f.Write([]byte(b))
		if err != nil {
			panic(err)
		}
	}
	
	return s.streamUrl
}
