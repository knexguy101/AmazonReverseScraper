package amazon

import (
	"github.com/PuerkitoBio/goquery"
	"net/http"
)

type StoreProductResult struct {
	Name  string
	Link  string
	Image string
}

func ScrapeImagesFromStore(c *http.Client, storeLink string) ([]*StoreProductResult, error) {
	req, _ := http.NewRequest("GET", storeLink, nil)
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

	var results []*StoreProductResult
	doc.Find("span[data-component-type='s-product-image']").Each(func(i int, selection *goquery.Selection) {

		link, ok := selection.Find("a").Attr("href")
		if !ok {
			return
		}

		img, ok := selection.Find(".s-image").Attr("src")
		if !ok {
			return
		}

		title, ok := selection.Find(".s-image").Attr("alt")
		if !ok {
			return
		}

		results = append(results, &StoreProductResult{
			Link:  link,
			Image: img,
			Name:  title,
		})
	})

	return results, nil
}
