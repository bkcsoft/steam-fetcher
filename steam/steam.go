package steam

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// New creates a new AppList from a byte-array (with json-data)
func New(bytes []byte) (*AppList, error) {
	alr := new(AppListResponse)
	alr.AppList = new(AppList)
	err := json.Unmarshal(bytes, alr)
	return alr.AppList, err
}

func NewFromFile(filename string) (*AppList, error) {
	_, err := os.Stat(filename)
	if err != nil {
		log.Println("Couldn't read file: ", filename)
		return nil, err
	}

	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	return New(bytes)
}

// Search returns a new AppList with less data
func (al *AppList) Search(query string) *AppList {
	nal := new(AppList)
	for _, v := range al.Apps {
		s1, s2 := strings.ToLower(v.Name), strings.ToLower(query)
		if strings.Contains(s1, s2) {
			nal.Apps = append(nal.Apps, v)
		}
	}
	return nal
}

// ToJSON converts the damn struct to json...
func (al *AppList) ToJSON() ([]byte, error) {
	return json.Marshal(al)
}
