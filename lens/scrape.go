package lens

import (
	"io"
	"net/http"
	"regexp"
	"strings"
)

const (
	LINK_TYPE   = 0
	ID_TYPE     = 1
	NAME_TYPE   = 2
	OTHER_TYPE  = 3
	GOOGLE_TYPE = 4
)

var (
	colorHexRegex, _   = regexp.Compile(`#[a-z0-9]{3,6}`)
	colorRGBRegex, _   = regexp.Compile(`rgba[(][0-9.,]*[)]`)
	colorPixelRegex, _ = regexp.Compile(`[0-9]{1,4}px`)
)

type ScrapeResult struct {
	Title string
	Link  string
}

func ScrapeUrl(c *http.Client, resultLink string) ([]*ScrapeResult, error) {
	req, _ := http.NewRequest("GET", resultLink, nil)
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

	b, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	s := string(b)

	var (
		results []*ScrapeResult

		currentLink  = ""
		currentTitle = ""
	)
	items := strings.Split(s, `null,null,null,"`)
	for x := 1; x < len(items); x++ {
		s := cleanString(items[x])
		switch identifyType(s) {
		case LINK_TYPE:
			currentLink = s
			break
		case NAME_TYPE:
			currentTitle = s
			break
		}
		if currentLink != "" && currentTitle != "" {
			results = append(results, &ScrapeResult{
				Link:  currentLink,
				Title: currentTitle,
			})
			currentLink = ""
			currentTitle = ""
		}
	}

	return results, nil
}

func cleanString(s string) string {
	return strings.Split(s, `"`)[0]
}

func identifyType(s string) int {
	s = cleanString(s)
	if strings.Contains(s, "https") {
		//im sorry, but just live with it, im lazy
		if strings.Contains(s, "gstatic.com") || strings.Contains(s, "google.com") || strings.Contains(s, "googleusercontent") {
			return GOOGLE_TYPE
		} else {
			return LINK_TYPE
		}
	} else if len(s) == 14 || len(s) < 25 { //the < 25 rule is just ballpark, could break shit, but for now its fine
		return ID_TYPE
	} else if s == "" || colorRGBRegex.MatchString(s) || colorHexRegex.MatchString(s) || colorPixelRegex.MatchString(s) {
		return OTHER_TYPE
	} else {
		return NAME_TYPE
	}
}
