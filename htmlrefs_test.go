package htmlrefs

import (
	"strings"
	"testing"
)

func TestAll(t *testing.T) {

	// the required resources to be found in the example
	reqs := [...]Reference{
		{URI: "/favicon.ico", Token: "link"},
		{URI: "/foo.css", Token: "link"},
		{URI: "/async.js", Token: "script"},
		{URI: "/foo.js", Token: "script"},
		{URI: "/relative", Token: "a"},
		{URI: "http://example.com/absolute", Token: "a"},
		{URI: "//google.com/protocol-relative", Token: "a"},
		{URI: "#yep", Token: "a"},
		{URI: "/sample.jpg", Token: "a"},
		{URI: "/logo.svg", Token: "source"},
		{URI: "/logo.webp", Token: "source"},
		{URI: "/logo.png", Token: "source"},
		{URI: "/logo.gif", Token: "img"},
		{URI: "/vid.mp4", Token: "video"},
	}

	// example HTML reader
	reader := strings.NewReader(`
		<!DOCTYPE html>
		<html>
			<head>
				<title>example</title>
				<link rel="icon" href="/favicon.ico">
				<link rel="stylesheet" type="text/css" href="/foo.css">
				<script type="text/javascript" async="async" src="/async.js"></script>
				<script type="text/javascript" src="/foo.js"></script>
			</head>
			<body>
				<a href="/relative"></a>
				<a href="http://example.com/absolute"></a>
				<a href="//google.com/protocol-relative"></a>
				<a href="#yep"></a>
				<img src="/sample.jpg">
				<picture>
				   <source type="image/svg+xml" srcset="/logo.svg" />
				   <source type="image/webp" srcset="/logo.webp" />
				   <source type="image/png" srcset="/logo.png" />
				   <img src="/logo.gif" alt="Company logo" />
				</picture>
				<p>nope</p>
				<video src="/vid.mp4"></video>
			</body>
		</html>
	`)

	// get the refs from the implementation
	refs := All(reader)

	if len(refs) != len(reqs) {
		t.Error("Wrong number of refs returned")
	}

	for i := 0; i < len(reqs); i++ {
		uri := reqs[i].URI
		ref := refs[i]

		if uri != ref {
			t.Errorf("Mismatch URI detected. need: %s , have: %s", uri, ref)
		}
	}
}
