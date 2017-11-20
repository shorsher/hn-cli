package main 

import (
	"bytes"
	"fmt"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
)

// Host is the HN API endpoint 
const Host  = "https://hacker-news.firebaseio.com/v0"

type Story struct {
	Kids        []int  `json:kids`
	Descendents int    `json:descendents`
	Id          int    `json:id`
	Score       int    `json:score`
	Time        int    `json:time`
	By          string `json:by`
	Title       string `json:title`
	Url         string `json:url`
}

func main() {
	ids := getStoryID("topstories")
	stories := getStories(ids, 30)
	fmt.Println(stories)
}

func getStoryID(list string) []int {
	url := Host + "/" + list + ".json"
	b := get(url)

	var ids []int
	if parseErr := json.Unmarshal(b.Bytes(), &ids); parseErr != nil {
		log.Fatal(parseErr)
	}
	return ids
}

func get(url string) *bytes.Buffer {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	var b bytes.Buffer
	if _, err := io.Copy(&b, resp.Body); err != nil {
		log.Fatal(err)
	}

	return &b
}

func getStories(ids []int, limit int) *bytes.Buffer {
	var b bytes.Buffer

	if len(ids) > 0 {
		for _, id := range ids[:limit] {
			item := getStory(id)
			fmt.Fprintf(&b, "\t%s\n\t%s\n\n", item.Title, item.Url)
		}
	}

	return &b
}

func getStory(id int) Story {
	url := Host + "/item/" + strconv.Itoa(id) + ".json"
	b := get(url)

	story := Story{}

	if parseErr := json.Unmarshal(b.Bytes(), &story); parseErr != nil {
		log.Fatal(parseErr)
	}

	return story
}