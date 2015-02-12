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
		{URI: "/sample.jpg", Token: "img"},
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
				<link rel="stylesheet" type="text/css" href="/foo.css" />
				<script type="text/javascript" async="async" src="/async.js"></script>
				<script type="text/javascript" src="/foo.js"></script>
				<script></script>
			</head>
			<body>
				<!-- comment for good measure -->
				<a href="/relative"></a>
				<a href="http://example.com/absolute"></a>
				<a href="//google.com/protocol-relative"></a>
				<a href="#yep"></a>
				<a name="yep">yep</a>
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

	need, have := len(reqs), len(refs)
	if need != have {
		t.Errorf("Wrong number of refs returned. need: %d , have: %d", need, have)
	}

	// loop through and verify URI and Token names
	for i, req := range reqs {
		ref := refs[i]

		if req.URI != ref.URI {
			t.Errorf("Mismatch URI detected. need: %s , have: %s", req.URI, ref.URI)
		}

		if req.Token != ref.Token {
			t.Errorf("Mismatch Token detected. need: %s , have: %s", req.Token, ref.Token)
		}
	}
}
