package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"

	"github.com/gorilla/feeds"
)

func main() {

	var (
		limit  int
		err    error
		search string
	)

	if len(os.Args) < 2 {
		log.Fatalf("call URL [LIMIT] [SEARCH]")
	}

	if len(os.Args) > 2 {
		limit, err = strconv.Atoi(os.Args[2])
		if err != nil {
			log.Fatalf("can not parse limit: %s", err.Error())
		}
	}

	if len(os.Args) > 3 {
		search = os.Args[3]
	}

	urlArg := os.Args[1]

	url, err := url.ParseRequestURI(urlArg)
	if err != nil {
		log.Fatalf("can not parse URL: %s", err.Error())
	}

	req, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		log.Fatalf("can not create request: %s", err.Error())
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("can not execute request: %s", err.Error())
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("can not read response: %s", err.Error())
	}

	if len(b) == 0 {
		log.Fatalf("RSS feed was empty")
	}

	var feed feeds.RssFeedXml
	err = xml.Unmarshal(b, &feed)
	if err != nil {
		log.Fatalf("can not parse response: %s", err.Error())
	}

	// shorten the results
	if limit > 0 && limit > len(feed.Channel.Items) {
		limit = len(feed.Channel.Items) - 1
		feed.Channel.Items = feed.Channel.Items[:limit]
	}

	if search != "" {
		newItems := make([]*feeds.RssItem, 0)
		for i := range feed.Channel.Items {
			item := feed.Channel.Items[i]
			if strings.Contains(item.Description, search) {
				newItems = append(newItems, item)
			}
		}
		feed.Channel.Items = newItems
	}

	if feed.Channel == nil {
		log.Fatalf("rss feed has no channel")
	}

	b, err = json.MarshalIndent(feed, "", "    ")
	if err != nil {
		log.Fatalf("can not parse item: %s", err.Error())
	}

	fmt.Printf("%v", string(b))

}
