package links

import (
	"fmt"
	"io"
	"net/http"
	neturl "net/url"
	"os"
	pathpackage "path"

	"golang.org/x/net/html"
)

func Exract(url string, cache bool) ([]string, error) {
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
	// ドキュメントの保存
	if cache {
		save(url, res)
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

func save(url string, resp *http.Response) error {
	// FIXME: なんどもParseしてて無駄
	u, err := neturl.Parse(url)
	if err != nil {
		return err
	}
	// FIXME: index.html 以外の名前を使う
	dir, _ := pathpackage.Split(u.Path)
	dir = pathpackage.Join(".", u.Hostname(), dir)
	os.MkdirAll(dir, 0755)
	if err != nil {
		return err
	}
	f, err := os.Create(pathpackage.Join(dir, "index.html"))
	defer f.Close()
	// FXIME: コピーできてない...空ファイルができる
	_, err = io.Copy(f, resp.Body)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
