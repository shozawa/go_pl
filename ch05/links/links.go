package links

import (
	"bytes"
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

	// [question] html.Parseとファイルの保存（io.Copy）で2回ストリームを読むので
	// バッファを巻き戻すかTeeで二方向に出力しなきゃいけない？ もっといい方法ありそう
	buf := bytes.NewBuffer(nil)
	tee := io.TeeReader(res.Body, buf)

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("getting %s: %s", url, res.Status)
	}
	doc, err := html.Parse(tee)
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}
	// ドキュメントの保存
	if cache {
		save(url, buf)
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

func save(url string, src io.Reader) error {
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
	_, err = io.Copy(f, src)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
