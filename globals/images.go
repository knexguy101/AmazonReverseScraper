package globals

import (
	"io"
	"net/http"
)

func DownloadImage(c *http.Client, url string) ([]byte, error) {
	req, _ := http.NewRequest("GET", url, nil)
	req.Header = map[string][]string{
		"Accept": {"*/*"},
	}
	res, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	return io.ReadAll(res.Body)
}
