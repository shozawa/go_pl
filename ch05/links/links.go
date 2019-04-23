package links

import(
	"fmt"
	"net/http"

	"golang.org/x/net/html"
)

func Exract(url string)([]string, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("getting %s: %s", url, res.Status)
	}
	doc, err := html.Parse(res.Body)
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}
	var links []string
	forEachNode(doc, func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key != "href" {
					continue
				}
				link, err := res.Request.URL.Parse(a.Val)
				if err != nil {
					continue
				}
				links = append(links, link.String())
			}
		}
	})

	return links, nil
}

func forEachNode(n *html.Node, proc func(*html.Node)) {
	if n == nil {
		return
	}
	proc(n)
	forEachNode(n.FirstChild, proc)
	forEachNode(n.NextSibling, proc)
}