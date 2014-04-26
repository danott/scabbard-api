package main

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
)

func PassageQuery(s string) (Passage, error) {
	requestUrl, err := url.Parse("http://www.esvapi.org/v2/rest/passageQuery")
	requestUrl.RawQuery = esvParams(s).Encode()

	resp, err := http.Get(requestUrl.String())
	if err != nil {
		return Passage{}, err
	}

	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Passage{}, err
	}

	body := string(respBody)

	if matched, _ := regexp.MatchString("^ERROR:", body); matched {
		return Passage{}, errors.New(body)
	}

	re := regexp.MustCompile("<h2>(.*)</h2>")
	title := re.FindStringSubmatch(body)[1]

	return Passage{title, body}, nil
}

func esvParams(s string) url.Values {
	v := make(url.Values)

	v.Set("include-passage-references", "true")
	v.Set("include-verse-numbers", "true")
	v.Set("include-footnotes", "false")
	v.Set("include-footnote-links", "false")
	v.Set("include-headings", "false")
	v.Set("include-subheadings", "false")
	v.Set("include-surrounding-chapters", "false")
	v.Set("include-word-ids", "false")
	v.Set("link-url", "http,//www.gnpcb.org/esv/search/")
	v.Set("include-audio-link", "false")
	v.Set("audio-format", "mp3")
	v.Set("audio-version", "hw")
	v.Set("include-short-copyright", "false")
	v.Set("include-copyright", "false")
	v.Set("output-format", "html")
	v.Set("key", apiKey)
	v.Set("passage", s)

	return v
}
