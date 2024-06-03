package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/lissdx/qAPI/internal/feeders"
	"github.com/lissdx/qAPI/internal/models"
)

func main() {

	// TODO: replace the hardcoded fileName param with conf param
	urlSrc, err := feeders.GetData("/Users/lissdx/dev/GO/go/src/github.com/lissdx/qAPI/doc/queries.txt")
	if err != nil {
		panic(err)
	}

	searchEntries := make([]*models.SearchEntry, 0, len(urlSrc))

	for _, path := range urlSrc {
		// TODO remove separator to config
		// there are a couple of ways to solve it
		// an the solution depends on our approach
		// 1. Create custom url normalizer ( for example let's talk about
		//    interface NormalizeURL(string) string)
		// 2. Delegate the normalization responsibility
		//    to the custom HTTPClient
		// 3. Combined approach (more flexible): custom HTTPClient
		//    will have an access to the urlNormalizer
		splitedUrl := strings.Split(path, "?q=")
		//resp, err := http.Get(url.PathEscape(path))
		requestUrl := fmt.Sprintf("%s?q=%s", splitedUrl[0], url.QueryEscape(splitedUrl[1]))
		resp, err := http.Get(requestUrl)

		if err != nil {
			log.Default().Println(err.Error())
			continue
		}

		holdSE := models.SearchEntry{}
		body, err := io.ReadAll(resp.Body)

		err = json.Unmarshal(body, &holdSE)
		if err != nil {
			log.Default().Println(err.Error())
			continue
		}

		searchEntries = append(searchEntries, &holdSE)
		_ = resp.Body.Close()
	}

	res := models.SearchEntryConcat(searchEntries...)

	log.Default().Println("search result total: ", res.TotalCount)
	log.Default().Println("search result entries: ", res.Items)
}
