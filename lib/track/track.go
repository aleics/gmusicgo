package track

import (
	"fmt"
	"net/http"
	"net/url"
	"io/ioutil"
	"strings"  
	"encoding/json"
	"os"                                                                                                    
	"strconv"
	"github.com/aleics/gmusicgo/lib/gmusicjson"
)

type Track struct{
	Kind string `json:"kind"`
	Id string `json:"id"`
	ClientId string `json:"clientId"`
	CreationTimestamp string `json:"creationTimestamp"`
	LastModifiedTimestamp string `json:"lastModifiedTimestamp"`
	Deleted bool `json:"deleted"`
	Title string `json:"title"`
	Artist string `json:"artist"`
	Composer string `json:"composer"`
	Album string `json:"album"`
	AlbumArtist string `json:"albumArtist"`
	Year float64 `json:"year"`
	Comment string `json:"comment"`
	TrackNumber float64 `json:"trackNumber"`
	Genre string `json:"genre"`
	DurationMillis string `json:"durationMillis"`
	BeatsPerMinute float64 `json:"beatsPerMinute"`
	AlbumArtRefurl string `json:"albumArtRefurl"`
	PlayCount float64 `json:"playCount"`
	TotalTrackCount float64 `json:"totalTrackCount"`
	DiscNumber float64 `json:"discNumber"`
	TotalDiscCount float64 `json:"totalDiscount"`
	Rating string `json:"rating"`
	EstimatedSize string `json:"estimatedSize"`
}
func Init() *Track{
        t := new(Track)
        return t
}
func ArrayInit() (t *[]Track){
	return t
}
func (t Track) GetKind() string {
        return t.Kind
}
func (t *Track) SetKind(kind string) {
        t.Kind = kind
}
func (t Track) GetId() string {
	return t.Id
}
func (t *Track) SetId(id string) {
        t.Id = id
}
func (t Track) GetClientId() string {
        return t.ClientId
}
func (t *Track) SetClientId(clientId string) {
        t.ClientId = clientId
}
func (t Track) GetCreationTimestamp() string {
        return t.CreationTimestamp
}
func (t *Track) SetCreationTimestamp(creationTimestamp string) {
        t.CreationTimestamp = creationTimestamp
}
func (t Track) GetLastModifiedTimestamp() string {
        return t.LastModifiedTimestamp
}
func (t *Track) SetLastModifiedTimestamp(lastModifiedTimestamp string) {
        t.LastModifiedTimestamp = lastModifiedTimestamp
}
func (t Track) GetDeleted() bool {
        return t.Deleted
}
func (t *Track) SetDeleted(deleted bool) {
        t.Deleted = deleted
}
func (t Track) GetTitle() string {
        return t.Title
}
func (t *Track) SetTitle(title string) {
        t.Title = title
}
func (t Track) GetArtist() string {
        return t.Artist
}
func (t *Track) SetArtist(artist string) {
        t.Artist = artist
}
func (t Track) GetComposer() string {
        return t.Composer
}
func (t *Track) SetComposer(composer string) {
        t.Composer = composer
}
func (t Track) GetAlbum() string {
        return t.Album
}
func (t *Track) SetAlbum(album string) {
        t.Album = album
}
func (t Track) GetAlbumArtist() string {
        return t.AlbumArtist
}
func (t *Track) SetAlbumArtist(albumArtist string) {
        t.AlbumArtist = albumArtist
}
func (t Track) GetYear() float64 {
        return t.Year
}
func (t *Track) SetYear(year float64) {
        t.Year = year
}
func (t Track) GetComment() string {
        return t.Comment
}
func (t *Track) SetComment(comment string) {
        t.Comment = comment
}
func (t Track) GetTrackNumber() float64 {
        return t.TrackNumber
}
func (t *Track) SetTrackNumber(trackNumber float64) {
        t.TrackNumber = trackNumber
}
func (t Track) GetGenre() string {
        return t.Genre
}
func (t *Track) SetGenre(genre string) {
        t.Genre = genre
}
func (t Track) GetDurationMillis() string {
        return t.DurationMillis
}
func (t *Track) SetDurationMillis(durationMillis string) {
        t.DurationMillis = durationMillis
}
func (t Track) GetBeatsPerMinute() float64 {
        return t.BeatsPerMinute
}
func (t *Track) SetBeatsPerMinute(beatsPerMinute float64) {
        t.BeatsPerMinute = beatsPerMinute
}
func (t Track) GetAlbumArtRefurl() string {
        return t.AlbumArtRefurl
}
func (t *Track) SetAlbumArtRefurl(albumArtRefurl string) {
        t.AlbumArtRefurl = albumArtRefurl
}
func (t Track) GetPlayCount() float64 {
        return t.PlayCount
}
func (t *Track) SetPlayCount(playCount float64) {
        t.PlayCount = playCount
}
func (t Track) GetTotalTrackCount() float64 {
        return t.TotalTrackCount
}
func (t *Track) SetTotalTrackCount(totalTrackCount float64) {
        t.TotalTrackCount = totalTrackCount
}
func (t Track) GetDiscNumber() float64 {
        return t.DiscNumber
}
func (t *Track) SetDiscNumber(discNumber float64) {
        t.DiscNumber = discNumber
}
func (t Track) GetTotalDiscCount() float64 {
        return t.TotalDiscCount
}
func (t *Track) SetTotalDiscCount(totalDiscCount float64) {
        t.TotalDiscCount = totalDiscCount
}
func (t Track) GetRating() string {
        return t.Rating
}
func (t *Track) SetRating(rating string) {
        t.Rating = rating
}
func (t Track) GetEstimatedSize() string {
        return t.EstimatedSize
}
func (t *Track) SetEstimatedSize(estimatedSize string) {
        t.EstimatedSize = estimatedSize
}
func (t *Track) NewTrack(kind string, id string, clientId string, creationTimestamp string, lastModifiedTimestamp string, deleted bool, title string, artist string, composer string, album string, albumArtist string, year float64, comment string, trackNumber float64, genre string, durationMillis string, beatsPerMinute float64, albumArtRefurl string, playCount float64, totalTrackCount float64, discNumber float64, totalDiscCount float64,rating string, estimatedSize string){
	t.Kind = kind
	t.Id = id
	t.ClientId = clientId
	t.CreationTimestamp = creationTimestamp
	t.LastModifiedTimestamp = lastModifiedTimestamp
	t.Deleted = deleted
	t.Title = title
	t.Artist = artist
	t.Composer = composer
	t.Album = album
	t.AlbumArtist = albumArtist
	t.Year = year
	t.Comment = comment
	t.TrackNumber = trackNumber
	t.Genre = genre
	t.DurationMillis = durationMillis
	t.BeatsPerMinute = beatsPerMinute
	t.AlbumArtRefurl = albumArtRefurl
	t.PlayCount = playCount
	t.TotalTrackCount = totalTrackCount
	t.DiscNumber = discNumber
	t.TotalDiscCount = totalDiscCount
	t.Rating = rating
	t.EstimatedSize = estimatedSize
}
func (t Track) ToMap() map[string]string {
	
	ret := make(map[string]string)

	ret["kind"] = t.Kind 
	ret["id"] = t.Id
	ret["creationTimestamp"] = t.CreationTimestamp
	ret["lastModifiedTimestamp"] = t.LastModifiedTimestamp
	ret["deleted"] = strconv.FormatBool(t.Deleted)
	ret["title"] = t.Title
	ret["artist"] = t.Artist
	ret["composer"] = t.Composer
	ret["album"] = t.Album
	ret["albumArtist"] = t.AlbumArtist
	ret["year"] = strconv.FormatFloat(t.Year, 'f', 0, 64)
	ret["comment"] = t.Comment
	ret["trackNumber"] = strconv.FormatFloat(t.TrackNumber, 'f', 0, 64)
	ret["genre"] = t.Genre
	ret["durationMillis"] = t.DurationMillis
	ret["beatsPerMinute"] = strconv.FormatFloat(t.BeatsPerMinute, 'f', 0, 64)
	ret["albumArtRefurl"] = t.AlbumArtRefurl
	ret["playCount"] = strconv.FormatFloat(t.PlayCount, 'f', 0, 64)
	ret["totalTrackCount"] = strconv.FormatFloat(t.TotalTrackCount, 'f', 0, 64)
	ret["discNumber"] = strconv.FormatFloat(t.DiscNumber, 'f', 0, 64)
	ret["totalDiscCount"] = strconv.FormatFloat(t.TotalDiscCount, 'f', 0, 64)
	ret["rating"] = t.Rating
	ret["estimatedSize"] = t.EstimatedSize

	return ret
}

func TracksRequest(auth string, path string) []Track{

	hostname := "https://www.googleapis.com"
	resource := "/sj/v1beta1/tracks"

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

	tracksmap := m["data"]
	tracks := tracksmap.(map[string]interface{})


	
	itemsmap := tracks["items"]
	items := itemsmap.([]interface{})
	
	var singleitem map[string]interface{}

	var singlealbumrefs map[string]interface{}
	var albumref string

	length := len(items)
	arraytracks := make([]Track,length)

	for i := 0; i < length; i++ {
	
		singleitem = items[i].(map[string]interface{})
		if singleitem["albumArtRef"] == nil {
			albumref = ""
		} else {
		albumrefs := singleitem["albumArtRef"].([]interface{})
		singlealbumrefs = albumrefs[0].(map[string]interface{})
		albumref = singlealbumrefs["url"].(string)
		}

		arraytracks[i].NewTrack(singleitem["kind"].(string), singleitem["id"].(string), singleitem["clientId"].(string), singleitem["creationTimestamp"].(string), singleitem["lastModifiedTimestamp"].(string), singleitem["deleted"].(bool), singleitem["title"].(string), singleitem["artist"].(string), singleitem["composer"].(string), singleitem["album"].(string), singleitem["albumArtist"].(string), singleitem["year"].(float64), singleitem["comment"].(string), singleitem["trackNumber"].(float64), singleitem["genre"].(string), singleitem["durationMillis"].(string), singleitem["beatsPerMinute"].(float64), albumref, singleitem["playCount"].(float64), singleitem["totalTrackCount"].(float64), singleitem["discNumber"].(float64), singleitem["totalDiscCount"].(float64), singleitem["rating"].(string), singleitem["estimatedSize"].(string)) 

	}

	p := []string{path,"tracks.json"}
	jsonpath := strings.Join(p,"")
	

	_, err = gmusicjson.Export(arraytracks, jsonpath)
	if err != nil {
		fmt.Println("Error exporting Tracks")
		fmt.Println(err)
	}

	return arraytracks
}

func (t *Track) Print(){
	fmt.Print("kind: ")
	fmt.Println(t.Kind)                                                                                                      
	fmt.Print("id: ")                                                                                                          
	fmt.Println(t.Id)                                                                                                          
	fmt.Print("clientId: ")                                                                                                      
	fmt.Println(t.ClientId)                                                                                                    
	fmt.Print("creationTimestamp: ")                                                                                              
	fmt.Println(t.CreationTimestamp)                                                                                           
	fmt.Print("lastModifiedTimestamp: ")                                                                                          
	fmt.Println(t.LastModifiedTimestamp)                                                                                       
	fmt.Print("deleted: ")                                                                                                       
	fmt.Println(t.Deleted)                                                                                                     
	fmt.Print("title: ")                                                                                                          
	fmt.Println(t.Title)                                                                                                       
	fmt.Print("artist: ")                                                                                                        
	fmt.Println(t.Artist)                                                                                                      
	fmt.Print("composer: ")                                                                                                      
	fmt.Println(t.Composer)                                                                                                    
	fmt.Print("album: ")                                                                                                          
	fmt.Println(t.Album)                                                                                                       
	fmt.Print("albumArtist: ") 
	fmt.Println(t.AlbumArtist)                                                                                                 
	fmt.Print("year: ")                                                                                                          
	fmt.Println(t.Year)                                                                                                        
	fmt.Print("comment: ")
	fmt.Println(t.Comment)                                                                                                     
	fmt.Print("trackNumber: ")
	fmt.Println(t.TrackNumber)                                                                                                 
	fmt.Print("genre: ")                                                                                                          
	fmt.Println(t.Genre)                                                                                                       
	fmt.Print("durationMillis: ")
	fmt.Println(t.DurationMillis)                                                                                              
	fmt.Print("beatsPerMinute: ")      
	fmt.Println(t.BeatsPerMinute)                                                                                              
	fmt.Print("playCount: ")   
	fmt.Println(t.PlayCount)
	fmt.Print("albumArtRefurl: ")
	fmt.Println(t.AlbumArtRefurl)
	fmt.Print("totalTrackCount: ")
	fmt.Println(t.TotalTrackCount)                                                                                             
	fmt.Print("discNumber: ")
        fmt.Println(t.DiscNumber)
	fmt.Print("totalDiscCount: ")
        fmt.Println(t.TotalDiscCount)
	fmt.Print("rating: ")      
	fmt.Println(t.Rating)                                                                                                      
	fmt.Print("estimatedSize: ")                                                                                                 
	fmt.Println(t.EstimatedSize)                                                                                               

}
