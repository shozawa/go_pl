package main

import (
	"testing"

	"golang.org/x/net/html"
)

func TestElementById(t *testing.T) {
	/*
	<doc>
	  <div1></div1>
	  <div2><h1 id="title"></h1></div2>
	</doc>
	*/
	id := html.Attribute{Key: "id", Val: "title"}
	attrs := []html.Attribute{id}
	h1 := &(html.Node{Data: "h1", Attr: attrs})
	div2 := &(html.Node{FirstChild: h1})
	div1 := &(html.Node{NextSibling: div2})
	doc := &(html.Node{FirstChild: div1})

	if got := elementById(doc, "title"); got.Data != "h1" {
		t.Errorf("expect: %v, but got %v\n", h1, got)
	}
}
