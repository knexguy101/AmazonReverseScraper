package tasks

import (
	"amazonReverse/amazon"
	"amazonReverse/globals"
	"amazonReverse/lens"
	"fmt"
	"net/http"
)

func SearchByStore(storeUrl string, maxItems int) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("[FATAL ERROR]", err)
		}
	}()

	var (
		c = &http.Client{}
	)

	products, err := amazon.ScrapeImagesFromStore(c, storeUrl)
	if err != nil {
		panic(err)
	} else if len(products) <= 0 {
		panic("No products were found for this store link")
	}

	for x := 0; x < len(products) && x < maxItems; x++ {
		imageUrl := products[x].Image

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
		fmt.Println("[ORIGINAL PRODUCT DETAILS]")
		fmt.Println("Name:", products[x].Name)
		fmt.Println("Product Link:", products[x].Link)
		fmt.Println("Image:", products[x].Image)
		fmt.Println("[RESULTS]")
		fmt.Println("")
		for _, v := range results {
			fmt.Printf("[%s] %s\n", v.Title, v.Link)
		}
	}
}
