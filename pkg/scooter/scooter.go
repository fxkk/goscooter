package scooter

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

type Response struct {
	Id        int    `json:"id"`
	Permalink string `json:"permalink"`
	Hash      string `json:"hash"`
	Quote     Quote  `json:"quote"`
}

type rawResponse struct {
	Id        flexInt `json:"id"`
	Permalink string  `json:"permalink"`
	Hash      string  `json:"hash"`
	Quote     Quote   `json:"quote"`
}

type Quote struct {
	Text             string `json:"text"`
	Track            string `json:"track"`
	Album            string `json:"album"`
	Year             string `json:"year"`
	AlbumInformation string `json:"album_information"`
	AlbumCover       string `json:"album_cover"`
	AlbumThumb       string `json:"album_thumb"`
	Releasedate      string `json:"releasedate"`

	//Track data is optional
	TrackCover  string `json:"trackcover"`
	TrackThumb  string `json:"trackthumb"`
	TrackMaster string `json:"trackmaster"`

	//Not all data have videos available
	Videos []Video `json:"videos"`
}

type Video struct {
	Title string `json:"title"`
	Link  string `json:"link"`
}

// the flexInt Type is needed, because the apis response id type inconsistent between int and string
//code from https://docs.bitnami.com/tutorials/dealing-with-json-with-non-homogeneous-types-in-go
type flexInt int

func (fi *flexInt) UnmarshalJSON(b []byte) error {
	if b[0] != '"' {
		return json.Unmarshal(b, (*int)(fi))
	}
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	i, err := strconv.Atoi(s)
	if err != nil {
		return err
	}
	*fi = flexInt(i)
	return nil
}

const (
	BaseURL = "https://howmuchisthe.fish/json"
)

func newClient() *http.Client {

	return &http.Client{
		Timeout: 30 * time.Second,
	}
}

func httpGet(path string) (*http.Response, error) {
	client := newClient()
	ctx := context.Background()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf("%s/%s", BaseURL, path), nil)
	if err != nil {
		return &http.Response{}, err
	}
	log.Printf("%s: %s", req.Method, req.URL)
	res, err := client.Do(req)

	return res, err
}

func getResponse(path string) (*Response, error) {
	res, err := httpGet(path)
	if err != nil {
		return &Response{}, err
	}
	defer res.Body.Close()

	log.Printf("Status: %d", res.StatusCode)

	if res.StatusCode != http.StatusOK {
		return &Response{}, errors.New(fmt.Sprintf("unexpected status: got %v", res.Status))
	}

	var rawResponse rawResponse
	//body, _ := ioutil.ReadAll(res.Body)
	//log.Println(string(body))
	err = json.NewDecoder(res.Body).Decode(&rawResponse)
	if err != nil {
		return &Response{}, err
	}

	response := Response{
		Id:        int(rawResponse.Id),
		Permalink: rawResponse.Permalink,
		Hash:      rawResponse.Hash,
		Quote:     rawResponse.Quote,
	}
	return &response, nil

}

func Random() (string, error) {
	response, err := RandomFull()
	if err != nil {
		return "", err
	}
	text := response.Quote.Text
	return text, nil
}

func RandomFull() (*Response, error) {
	response, err := getResponse("random")
	if err != nil {
		return &Response{}, err
	}
	return response, nil
}

func Daily() (string, error) {
	response, err := DailyFull()
	if err != nil {
		return "", err
	}
	text := response.Quote.Text
	return text, nil
}

func DailyFull() (*Response, error) {
	response, err := getResponse("daily")
	if err != nil {
		return &Response{}, err
	}
	return response, nil
}

func QuoteById(id int) (string, error) {
	response, err := QuoteByIdFull(id)
	if err != nil {
		return "", err
	}
	text := response.Quote.Text
	return text, nil
}

func QuoteByIdFull(id int) (*Response, error) {
	response, err := getResponse(fmt.Sprintf("perma/%d", id))
	if err != nil {
		return &Response{}, err
	}
	return response, nil
}
