package htmlUtils

import (
	"strings"

	"golang.org/x/net/html"
)

func ExtractTextFromHtml(inputHtml string) (string, error) {
	doc, err := html.Parse(strings.NewReader(inputHtml))
	if err != nil {
		return "", err
	}

	var builder strings.Builder
	var traverse func(*html.Node)
	traverse = func(node *html.Node) {
		if node.Type == html.TextNode {
			builder.WriteString(node.Data)
		}
		// Traverse child nodes
		for child := node.FirstChild; child != nil; child = child.NextSibling {
			traverse(child)
		}
	}

	traverse(doc)
	return builder.String(), nil
}
