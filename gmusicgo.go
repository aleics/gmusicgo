package gmusicgo

import (
	"errors"
	"github.com/aleics/gmusicgo/lib/clientlogin"
	"github.com/aleics/gmusicgo/lib/gmusicjson"
	"github.com/aleics/gmusicgo/lib/playlist"
	"github.com/aleics/gmusicgo/lib/plentry"
	"github.com/aleics/gmusicgo/lib/stream"
	"github.com/aleics/gmusicgo/lib/tokens"
	"github.com/aleics/gmusicgo/lib/track"
	"os"
)

type Gmusicgo struct {
	gclient    clientlogin.Clientlogin
	gtokens    tokens.Tokens
	gtracks    []track.Track
	gplaylists []playlist.Playlist
	gplentries []plentry.Plentry
	gstream    stream.Stream
}

func Init() *Gmusicgo {
	gmusic := new(Gmusicgo)
	gmusic.gclient = *clientlogin.Init()
	gmusic.gtokens = *tokens.Init()
	gmusic.gtracks = make([]track.Track, 0)
	gmusic.gplaylists = make([]playlist.Playlist, 0)
	gmusic.gplentries = make([]plentry.Plentry, 0)
	gmusic.gstream = *stream.Init()
	return gmusic
}
func (g Gmusicgo) GetGmusicgo() Gmusicgo {
	return g
}
func (g Gmusicgo) SetGmusicgo(gm Gmusicgo) {
	g = gm
}
func (g Gmusicgo) GetGclient() clientlogin.Clientlogin {
	return g.gclient
}
func (g *Gmusicgo) SetGclient(gc clientlogin.Clientlogin) {
	g.gclient = gc
}
func (g Gmusicgo) GetGtokens() tokens.Tokens {
	return g.gtokens
}
func (g *Gmusicgo) SetGtokens(gt tokens.Tokens) {
	g.gtokens = gt
}
func (g Gmusicgo) GetGtracks() []track.Track {
	return g.gtracks
}
func (g *Gmusicgo) SetGtracks(gtr []track.Track) {
	g.gtracks = gtr
}
func (g Gmusicgo) GetGplaylists() []playlist.Playlist {
	return g.gplaylists
}
func (g *Gmusicgo) SetGplaylists(gpla []playlist.Playlist) {
	g.gplaylists = gpla
}
func (g Gmusicgo) GetGplentries() []plentry.Plentry {
	return g.gplentries
}
func (g *Gmusicgo) SetGplentries(gple []plentry.Plentry) {
	g.gplentries = gple
}

func (g *Gmusicgo) Connect(accountType string, email string, passwd string, service string, source string, path string) error {

	header := [5]string{accountType, email, passwd, service, source}
	g.gclient.SetHeader(accountType, email, passwd, service, source)

	if g.gclient.MakeRequest(header)[0] != "200 OK" {
		return errors.New("Error clientlogin request")
		os.Exit(1)
	}

	if g.gtokens.MakeRequest(g.gclient.GetAuth())[0] != "200 OK" {
		return errors.New("Error tokens request")
		os.Exit(1)
	}

	if path != "" {
		g.gclient.SaveInfo(path)
		g.gtokens.SaveInfo(path)
	}

	g.gtracks = track.TracksRequest(g.gclient.GetAuth(), path)
	if g.gtracks[0].GetId() == "" {
		return errors.New("Error track request")
		os.Exit(1)
	}

	g.gplaylists = playlist.PlaylistsRequest(g.gclient.GetAuth(), path)
	if g.gplaylists[0].GetId() == "" {
		return errors.New("Error playlist request")
		os.Exit(1)
	}

	g.gplentries = plentry.PlentryRequest(g.gclient.GetAuth(), path)
	if g.gplentries[0].GetId() == "" {
		return errors.New("Error plentry request")
		os.Exit(1)
	}

	return nil
}
func (g *Gmusicgo) Update(path string) error {

	if path != "" {
		g.gclient.SaveInfo(path)
		g.gtokens.SaveInfo(path)
	}

	g.gtracks = track.TracksRequest(g.gclient.GetAuth(), path)
	if g.gtracks[0].GetId() == "" {
		return errors.New("Error track request")
		os.Exit(1)
	}

	g.gplaylists = playlist.PlaylistsRequest(g.gclient.GetAuth(), path)
	if g.gplaylists[0].GetId() == "" {
		return errors.New("Error playlist request")
		os.Exit(1)
	}

	g.gplentries = plentry.PlentryRequest(g.gclient.GetAuth(), path)
	if g.gplentries[0].GetId() == "" {
		return errors.New("Error plentry request")
		os.Exit(1)
	}

	return nil
}
func (g Gmusicgo) TracksToMap() ([]map[string]string, error) {
	thisMap := make([]map[string]string, 0)
	for i := 0; i < len(g.gtracks); i++ {
		thisMap = append(thisMap, g.gtracks[i].ToMap())
	}
	if len(thisMap) == 0 {
		return thisMap, errors.New("Tracks struct empty or error on the function: TracksToMap()")
	}
	return thisMap, nil
}
func (g Gmusicgo) PlaylistsToMap() ([]map[string]string, error) {
	thisMap := make([]map[string]string, 0)
	for i := 0; i < len(g.gplaylists); i++ {
		thisMap = append(thisMap, g.gplaylists[i].ToMap())
	}
	if len(thisMap) == 0 {
		return thisMap, errors.New("Playlists struct empty or error on the function: PlaylistsToMap()")
	}
	return thisMap, nil
}
func (g Gmusicgo) PlentriesToMap() ([]map[string]string, error) {
	thisMap := make([]map[string]string, 0)
	for i := 0; i < len(g.gplentries); i++ {
		thisMap = append(thisMap, g.gplentries[i].ToMap())
	}
	if len(thisMap) == 0 {
		return thisMap, errors.New("Plentries struct empty or error on the function: PlentriesToMap()")
	}
	return thisMap, nil
}

func (g *Gmusicgo) GetSong(songid string, path string) error {
	err := g.gstream.StreamRequest(g.gclient.GetAuth(), g.gtokens.GetXt(), songid, path)
	if err != nil {
		return errors.New("Error stream request")
	}
	return nil
}
func (g *Gmusicgo) CreatePlaylist(name string, description string, public bool) error {
	err := g.gplaylists[0].CreatePlaylist(g.gclient.GetAuth(), g.gtokens.GetXt(), name, description, public)
	if err != true {
		return errors.New("Error creating playlist: " + name)
	}
	return nil
}

func (g *Gmusicgo) LoadUserPlaylist() error {
	err := g.gplaylists[0].LoadUserPlaylist(g.gclient.GetAuth(), g.gtokens.GetXt())
	if err != true {
		return errors.New("Error loading playlists")
	}
	return nil
}

func (g *Gmusicgo) DeletePlaylist(id string) error {
	err := g.gplaylists[0].DeletePlaylist(g.gclient.GetAuth(), g.gtokens.GetXt(), id)
	if err != true {
		return errors.New("Error deleting playlist: " + id)
	}
	return nil
}

func (g Gmusicgo) GetIdBySongTitle(songname string) (string, error) {
	for i := 0; i < len(g.gtracks); i++ {
		if g.gtracks[i].GetTitle() == songname {
			return g.gtracks[i].GetId(), nil
		}
	}
	return "", errors.New("Song name not found")

}
func (g Gmusicgo) GetIdsByArtist(artistname string) ([]string, error) {
	out := make([]string, 0)
	for i := 0; i < len(g.gtracks); i++ {
		if g.gtracks[i].GetArtist() == artistname {
			out = append(out, g.gtracks[i].GetId())
		}
	}

	if len(out) == 0 {
		return out, errors.New("Artist name not found")
	}
	return out, nil
}
func (g Gmusicgo) GetIdsByAlbum(albumname string) ([]string, error) {
	out := make([]string, 0)
	for i := 0; i < len(g.gtracks); i++ {
		if g.gtracks[i].GetAlbum() == albumname {
			out = append(out, g.gtracks[i].GetId())
		}
	}

	if len(out) == 0 {
		return out, errors.New("Album name not found")
	}

	return out, nil
}
func (g Gmusicgo) GetIdsByPlaylist(playlistname string) ([]string, error) {
	pla_id := ""
	for i := 0; i < len(g.gplaylists); i++ {
		if g.gplaylists[i].GetName() == playlistname {
			pla_id = g.gplaylists[i].GetId()
		}
	}

	out := make([]string, 0)
	for i := 0; i < len(g.gplentries); i++ {
		if g.gplentries[i].GetPlaylistId() == pla_id {
			out = append(out, g.gplentries[i].GetTrackId())
		}
	}

	if len(out) == 0 {
		return out, errors.New("Playlist name not found")
	}

	return out, nil
}

func (g Gmusicgo) Export(path string) error {
	_, err := gmusicjson.Export(g.gclient, path+"userinfo.json")
	if err != nil {
		return errors.New("Error exporting User Info")
	}
	_, err = gmusicjson.Export(g.gtokens, path+"tokens.json")
	if err != nil {
		return errors.New("Error exporting Tokens")
	}
	_, err = gmusicjson.Export(g.gtracks, path+"tracks.json")
	if err != nil {
		return errors.New("Error exporting Tracks")
	}
	_, err = gmusicjson.Export(g.gplaylists, path+"playlists.json")
	if err != nil {
		return errors.New("Error exporting Playlists")
	}
	_, err = gmusicjson.Export(g.gplentries, path+"plentries.json")
	if err != nil {
		return errors.New("Error exporting Plentries")
	}
	return nil
}

func (g *Gmusicgo) Import(path string) error {
	err := gmusicjson.Import(path+"userinfo.json", &g.gclient)
	if err != nil {
		return errors.New("Error importing User Info")
	}
	err = gmusicjson.Import(path+"tokens.json", &g.gtokens)
	if err != nil {
		return errors.New("Error importing Tokens")
	}
	err = gmusicjson.Import(path+"tracks.json", &g.gtracks)
	if err != nil {
		return errors.New("Error importing Tracks")
	}
	err = gmusicjson.Import(path+"playlists.json", &g.gplaylists)
	if err != nil {
		return errors.New("Error importing Playlists")
	}
	err = gmusicjson.Import(path+"plentries.json", &g.gplentries)
	if err != nil {
		return errors.New("Error importing Plentries")
	}

	return nil
}
