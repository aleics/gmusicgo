package main

import(
	"fmt"
	"os"
	"github.com/aleics/gmusicgo"
)

func main(){
	gmusic := gmusicgo.Init()
	path := "/home/aleix/gmusic/"

	var header [5]string
	_, err := fmt.Scan(&header[0])
	if err != nil {
		os.Exit(1)
	}
	_, err = fmt.Scan(&header[1])
	if err != nil {
		os.Exit(1)
	}
	_, err = fmt.Scan(&header[2])
	if err != nil {
		os.Exit(1)
	}
	_, err = fmt.Scan(&header[3])
	if err != nil {
		os.Exit(1)
	}
	_, err = fmt.Scan(&header[4])
	if err != nil {
		os.Exit(1)
	}

	err = gmusic.Connect(header[0],header[1],header[2],header[3],header[4], path)
	if err != nil {
		fmt.Println("Couldn't connect with Gmusic")
	}

	err = gmusic.GetSong("b7e5d43b-0293-3020-8277-a686bd394eec", path + "songs/")
	if err != nil {
		fmt.Println("Couldn't get the song")
	}
	err = gmusic.Import(path)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Couldn't import all the data. Check exporting path")
	}

	err = gmusic.CreatePlaylist("Test", "test google api", true)

	if err != nil {
		fmt.Println(err)
		fmt.Println("Couldn't create the playlist.")
	}

	_ = gmusic.LoadUserPlaylist()

	
	//_ = gmusic.DeletePlaylist("")
}

/*package main

import(
	"fmt"
	"os"
	"github.com/aleics/clientlogin"
	"github.com/aleics/tokens"
	"github.com/aleics/track"
	"github.com/aleics/playlist"
	"github.com/aleics/plentry"
	"github.com/aleics/stream"
)

func main(){

	path := "/home/aleix/gmusic/"
	client := clientlogin.Init()

	var header [5]string
	_, err := fmt.Scan(&header[0])
	if err != nil {
                os.Exit(1)
        }
	_, err = fmt.Scan(&header[1])
	if err != nil {
                os.Exit(1)
        }
	_, err = fmt.Scan(&header[2])
	if err != nil {
                os.Exit(1)
        }
	_, err = fmt.Scan(&header[3])
	if err != nil {
                os.Exit(1)
        }
	_, err = fmt.Scan(&header[4])
	if err != nil {
		os.Exit(1)
	}


	client.SetHeader(header[0],header[1],header[2],header[3],header[4])

	value := client.MakeRequest(client.Header())
	

	if value[0] == "200 OK"{
		fmt.Println("AUTHORIZATION")
		fmt.Println("Good Request")
		fmt.Println("Auth saved on: " + path)
		fmt.Println(client.Auth())
		save := client.SaveInfo(path)
		if save != true {
			fmt.Println("User info couldn't be saved")
		} else {
			fmt.Println("User info saved!")
		}
	}

	fmt.Println("")
	fmt.Println("----------------------------------------------------------------------------------------------------")
	fmt.Println("")
	
	tokenvalues := tokens.Init()
	
	value2 := tokenvalues.MakeRequest(client.Auth())

	if value2[0] == "200 OK"{
		fmt.Println("TOKENS")
		fmt.Println("Good Request")
		fmt.Println("Tokens saved on: " + path)
		fmt.Println(tokenvalues.Xt())
		fmt.Println(tokenvalues.Sjsaid())
		save2 := tokenvalues.SaveInfo(path)
		if save2 != true {
			fmt.Println("Tokens value couldn't be saved")
		} else {
			fmt.Println("Tokens saved!")
		}
	}

	fmt.Println("")
	fmt.Println("----------------------------------------------------------------------------------------------------")
        fmt.Println("")

	track := track.Init()
	
	tracks := track.TracksRequest(client.Auth(), path)
	if tracks[0].Id() != "" {
		fmt.Println("TRACKS")
		fmt.Println("Good Request")
		fmt.Println("Tracks saved on: " + path)
		length := len(tracks)
		if length != 0 {
                        fmt.Println("Good Response")
		} else { fmt.Println("Error saving Tracks")}
	}

	fmt.Println("")
        fmt.Println("----------------------------------------------------------------------------------------------------")
        fmt.Println("")


	playlist := playlist.Init()
	
	playlists := playlist.PlaylistsRequest(client.Auth(), path)
	if playlists[0].Id() != "" {
		fmt.Println("PLAYLISTS")
                fmt.Println("Good Request")
		fmt.Println("Playlists saved on: " + path)
                length := len(playlists)
		if length != 0 {
                        fmt.Println("Good Response")
		} else {fmt.Println("Error saving Playlists")}
	}

	fmt.Println("")
	fmt.Println("----------------------------------------------------------------------------------------------------")
	fmt.Println("")


	plentry := plentry.Init()

        plentries := plentry.PlentryRequest(client.Auth(), path)
        if plentries[0].Id() != "" {
                fmt.Println("PLENTRIES")
                fmt.Println("Good Request")
		fmt.Println("Plentries saved on: " + path)
                length := len(plentries)
                if length != 0 {
			fmt.Println("Good Response")
		} else {fmt.Println("Error saving Plentries")}
        }

	fmt.Println("")
        fmt.Println("----------------------------------------------------------------------------------------------------")
        fmt.Println("")

	stream := stream.Init()

	stream.StreamRequest(client.Auth(), tokenvalues.Sjsaid(), tokenvalues.Xt(), tracks[0].Id(), path)
	if stream.StreamUrl() != "" {
		fmt.Println("STREAM")
		fmt.Println("Good Request")
		fmt.Println("Stream saved on: " + path)
	} else {fmt.Println("Error saving Stream")}
	
	fmt.Println("")
        fmt.Println("-----------------------------------------------------------------------------------------------------")
}*/
