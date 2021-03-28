package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Entry struct {
	Title     string
	Author    string
	URL       string
	Permalink string
}

type Feed struct {
	Data struct {
		Children []struct {
			Data Entry
		}
	}
}

func main() {
	fmt.Println("-------------------")
	fmt.Println("Starting API client")
	fmt.Println("--------------------")
	fmt.Println("")

	// url of REST endpoint we are grabbing data from
	url := "https://www.reddit.com/r/golang/new.json"

	// fetch url
	client := &http.Client{}

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalf("Error making request object - %v", err)
	}
	request.Header.Add("Accept", "application/json")
	request.Header.Add("User-Agent", "postman")

	resp, err := client.Do(request)
	if err != nil {
		log.Fatalf("Error making GET request to URL [%v] - %v", url, err)
	}

	defer resp.Body.Close()

	// confirm we received an OK status
	if resp.StatusCode != http.StatusOK {
		log.Println(resp.Body)
		log.Fatalln("Error Status not OK:", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response body data - %v", err)
	}

	// create an empty instance of Feed struct
	// this is what gets filled in when unmarshaling JSON
	var entries Feed
	if err := json.Unmarshal(body, &entries); err != nil {
		log.Fatalf("Error unmarshalling response into struct - %v", err)
	}

	for _, entry := range entries.Data.Children {
		post := entry.Data
		log.Println(">>>")
		log.Println("Title   :", post.Title)
		// log.Println("Author  :", post.Author)
		// log.Println("URL     :", post.URL)
		log.Printf("Comments: http://reddit.com%s \n", post.Permalink)
	}
}
