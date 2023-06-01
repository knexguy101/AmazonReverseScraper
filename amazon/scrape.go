package amazon

import (
	"github.com/PuerkitoBio/goquery"
	"net/http"
)

func ScrapeProductLinks(c *http.Client, productLink string) ([]string, error) {
	req, _ := http.NewRequest("GET", productLink, nil)
	req.Header = map[string][]string{
		"Accept":        {"text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7"},
		"Cache-Control": {"max-age=0"},
		"User-Agent":    {"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36"},
	}
	res, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}

	var results []string
	doc.Find("#imgTagWrapperId > img").Each(func(i int, selection *goquery.Selection) {
		src, ok := selection.Attr("src")
		if !ok {
			return
		}

		results = append(results, src)
	})

	return results, nil
}
