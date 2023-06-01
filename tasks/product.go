package tasks

import (
	"amazonReverse/amazon"
	"amazonReverse/globals"
	"amazonReverse/lens"
	"fmt"
	"net/http"
)

func SearchByProduct(productUrl string) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("[FATAL ERROR]", err)
		}
	}()

	var (
		c = &http.Client{}
	)

	links, err := amazon.ScrapeProductLinks(c, productUrl)
	if err != nil {
		panic(err)
	} else if len(links) <= 0 {
		panic("No images were found for this product link")
	}

	imageUrl := links[0]

	imageBytes, err := globals.DownloadImage(c, imageUrl)
	if err != nil {
		panic(err)
	}

	se, err := lens.Search(c, imageBytes, imageUrl)
	if err != nil {
		panic(err)
	}

	results, err := lens.ScrapeUrl(c, se.Link)
	if err != nil {
		panic(err)
	}

	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("[RESULTS]")
	fmt.Println("")
	for _, v := range results {
		fmt.Printf("[%s] %s\n", v.Title, v.Link)
	}
}
