package lens

import (
	"amazonReverse/globals"
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"strings"
	"time"
)

type SearchResult struct {
	Link string
}

func Search(t *http.Client, imageBytes []byte, imageUrl string) (*SearchResult, error) {

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	fw, err := writer.CreateFormFile("encoded_image", "encoded_image.jpg")
	if err != nil {
		return nil, err
	}

	_, err = io.Copy(fw, bytes.NewReader(imageBytes))
	if err != nil {
		return nil, err
	}

	fw, err = writer.CreateFormField("image_url")
	if err != nil {
		return nil, err
	}

	_, err = io.Copy(fw, strings.NewReader(imageUrl))
	if err != nil {
		return nil, err
	}

	fw, err = writer.CreateFormField("sbisrc")
	if err != nil {
		return nil, err
	}

	_, err = io.Copy(fw, strings.NewReader("Google Chrome 113.0.5672.127 (Official) Windows"))
	if err != nil {
		return nil, err
	}

	writer.Close()

	req, _ := http.NewRequest("POST", fmt.Sprintf("https://lens.google.com/upload?ep=ccm&re=dcsp&s=4&st=%d&lm=1&sideimagesearch=1", time.Now().UnixMilli()), bytes.NewReader(body.Bytes()))
	req.Header = map[string][]string{
		"Content-Type":  {writer.FormDataContentType()},
		"Accept":        {"text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7"},
		"Cache-Control": {"max-age=0"},
		"User-Agent":    {"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36"},
	}
	res, err := t.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	b, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	link, err := globals.SafeSplit(string(b), "URL=", `"`)
	if err != nil {
		return nil, err
	}

	return &SearchResult{
		Link: link,
	}, nil
}
