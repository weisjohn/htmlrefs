// Package htmlrefs returns a slice of `Reference{URI, Token string}`s from an `io.Reader`.
package htmlrefs

import (
	"io"

	"golang.org/x/net/html"
)

// `Reference` are simply two strings: a `URI` and a `Token`
type Reference struct{ URI, Token string }

// a map of which tokens :: attr names to examine
var attrMap = map[string]string{
	"link":   "href",
	"a":      "href",
	"script": "src",
	"img":    "src",
	"source": "srcset",
	"video":  "src",
}

// `All` takes a reader object (like the one returned from http.Get())
// It returns a slice of struct Reference{uri, nodetype}
// It does not close the reader passed to it.
func All(httpBody io.Reader) []Reference {
	refs := make([]Reference, 0)
	page := html.NewTokenizer(httpBody)
	for {
		tokenType := page.Next()
		if tokenType == html.ErrorToken {
			return refs
		}
		token := page.Token()

		if tokenType == html.StartTagToken || tokenType == html.SelfClosingTagToken {

			// read the name of the token
			tokenString := token.DataAtom.String()
			mappedAttr, ok := attrMap[tokenString]

			// if we don't have an attribute mapping, loop
			if !ok {
				continue
			}

			// loop over attributes
			for _, attr := range token.Attr {
				if attr.Key == mappedAttr {
					refs = append(refs, Reference{URI: attr.Val, Token: tokenString})
				}
			}

		}
	}
}
