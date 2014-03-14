package main

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"regexp"
)

func reportPassageQueryError(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func PassageQuery(s string) Passage {
	params := make(url.Values)

	apiKey := os.Getenv("ESV_API_KEY")
	if len(apiKey) < 1 {
		apiKey = "IP"
	}

	params.Set("include-passage-references", "true")
	params.Set("include-verse-numbers", "true")
	params.Set("include-footnotes", "false")
	params.Set("include-footnote-links", "false")
	params.Set("include-headings", "false")
	params.Set("include-subheadings", "false")
	params.Set("include-surrounding-chapters", "false")
	params.Set("include-word-ids", "false")
	params.Set("link-url", "http,//www.gnpcb.org/esv/search/")
	params.Set("include-audio-link", "false")
	params.Set("audio-format", "mp3")
	params.Set("audio-version", "hw")
	params.Set("include-short-copyright", "false")
	params.Set("include-copyright", "false")
	params.Set("output-format", "html")
	params.Set("key", apiKey)
	params.Set("passage", s)

	requestUrl, err := url.Parse("http://www.esvapi.org/v2/rest/passageQuery")
	requestUrl.RawQuery = params.Encode()

	resp, err := http.Get(requestUrl.String())
	reportPassageQueryError(err)

	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	reportPassageQueryError(err)

	body := string(respBody)

	if matched, _ := regexp.MatchString("^ERROR:", body); matched {
		reportPassageQueryError(errors.New(body))
	}

	re := regexp.MustCompile("<h2>(.*)</h2>")
	title := re.FindStringSubmatch(body)[1]

	return Passage{title, body}
}
