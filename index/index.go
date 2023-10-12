package index

import (
	"encoding/gob"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
)

// Get all comics metadata once and create local index
func CreateIndex(filename string) error {
	// Open/Create file
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0o666)
	if err != nil {
		return err
	}

	comics := &Comics{}
	index := &Index{File: filename}

	// Loop for creating Index
	for i := 1; i <= lastComicsNum; i++ {
		var err error
		var statusCode int

		comics, statusCode, err = GetComicsJSON(i)
		if statusCode != http.StatusOK {
			continue
		}
		if err != nil {
			return err
		}
		index.Items = append(index.Items, comics)
	}

	// Serialize and write to file db
	en := gob.NewEncoder(file)

	if err := en.Encode(index); err != nil {
		return fmt.Errorf("xkcd: %v", err)
	}

	return nil
}

// Search comics
func SearchComics(term string, filename string) error {
	// Open/Create file
	file, err := os.OpenFile(filename, os.O_RDONLY, 0o666)
	if err != nil {
		return err
	}

	index := &Index{File: filename}

	// Desserialize and read to struct
	dec := gob.NewDecoder(file)

	if err := dec.Decode(index); err != nil {
		return fmt.Errorf("xkcd: %v", err)
	}

	// Search loop
	for _, comics := range index.Items {
		if strings.Contains(comics.Title, term) || strings.Contains(comics.Transcript, term) {
			fmt.Printf("URL: %q\nTranscript: %s\n\n", comics.Img, comics.Transcript)
		}
	}
	return nil
}

// Get metadata for commics #number
func GetComicsJSON(number int) (*Comics, int, error) {
	resp, err := http.Get(URL + "/" + strconv.Itoa(number) + "/info.0.json")
	if err != nil {
		return nil, 0, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, resp.StatusCode, nil
	}

	comics := &Comics{}

	if err = json.NewDecoder(resp.Body).Decode(comics); err != nil {
		return nil, 0, fmt.Errorf("xkcd: %v", err)
	}

	resp.Body.Close()
	return comics, resp.StatusCode, nil
}
