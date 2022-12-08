package main

/*
import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type UselessFact struct {
	Facts string `json:"text"`
	Lang  string `json:"language"`
}

func main() {
	var webSiteUrl = "https://uselessfacts.jsph.pl/random.json"
	u, _ := url.Parse(webSiteUrl)

	values, _ := url.ParseQuery(u.RawQuery)
	values.Add("language", "en")

	u.RawQuery = values.Encode()

	req, _ := http.NewRequest(
		http.MethodGet,
		u.String(),
		nil,
	)
	req.Header.Add("Accept", "application/json")
	res, _ := http.DefaultClient.Do(req)
	responseBytes, _ := io.ReadAll(res.Body)

	var uselessFact UselessFact
	json.Unmarshal(responseBytes, &uselessFact)

	fmt.Println(uselessFact)
}
*/
