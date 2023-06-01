package tasks

import (
	"amazonReverse/globals"
	"amazonReverse/lens"
	"fmt"
	"net/http"
)

func SearchByImage(imageUrl string) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("[FATAL ERROR]", err)
		}
	}()

	var (
		c = &http.Client{}
	)

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
