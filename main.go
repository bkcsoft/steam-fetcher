package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

const (
	// SteamAPIURL holds the API-Endpoint that we need (this might change in the future (and I hate hardcoded shit), hence a const)
	SteamAPIURL = `https://api.steampowered.com/ISteamApps/GetAppList/v2`
	// TempFile is the file in which we store the temporary steam-apps data
	TempFile = `steam-apps.json`

	// RefreshHours is the amount of HOURS to wait before re-fetching the data
	RefreshHours = 24
)

var (
	// HTTPPort to bind on
	HTTPPort int
)

func init() {
	flag.IntVar(&HTTPPort, "port", 8080, "HTTP Port to bind to")
	flag.Parse()
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/search", Search)

	host := fmt.Sprintf(`localhost:%d`, HTTPPort)
	log.Fatal(http.ListenAndServe(host, router))
}

func fetchSteamData(r *http.Request) bool {
	res, err := http.Get(SteamAPIURL)
	if err != nil {
		log.Printf("Shit we can has error! ABANDON SHIP! %s\n", err.Error())
		return false
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Printf("Shit we can has error! ABANDON SHIP! %s\n", err.Error())
		return false
	}
	err = ioutil.WriteFile(TempFile, body, 0664)
	if err != nil {
		log.Printf("Shit we can has error! ABANDON SHIP! %s\n", err.Error())
		return false
	}
	return true
}

// Search for a game in the DB...
func Search(w http.ResponseWriter, r *http.Request) {
	// We always wanna return json-data...
	w.Header().Set("content-type", "application/json; charset: utf-8")
	file, err := os.Stat(TempFile)
	if err != nil {
		log.Print("steam-apps.json not found, fetching...")
		if !fetchSteamData(r) {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, `{"error": "An error occured while fetching data form steam, please inform the server-admin..."}`)
			return
		}
		file, err = os.Stat(TempFile)
		if err != nil {
			log.Print("steam-apps.json not found after fetch. this is bad -.-...")
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, `{"error": "An error occured while fetching data form steam, please inform the server-admin..."}`)
			return
		}
	}
	lastUpdate := file.ModTime()
	if time.Since(lastUpdate).Hours() > RefreshHours { // Refetch...
		log.Printf("%s older than %d hours! refreshing!\n", TempFile, RefreshHours)
		if !fetchSteamData(r) {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, `{"error": "An error occured while fetching data form steam, please inform the server-admin..."}`)
			return
		}
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, `{"apps": []}`)
}
