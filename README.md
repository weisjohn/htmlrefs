# htmlrefs

Package htmlrefs returns a slice of `Reference{URI, Token string}`s from an `io.Reader`.

### usage

```go
package main

import (
    "fmt"
    "net/http"

    "github.com/weisjohn/htmlrefs"
)

func main() {
    resp, _ := http.Get("http://johnweis.com")
    refs := htmlrefs.All(resp.Body)

    for _, ref := range refs {
        fmt.Println(ref.Token, ":", ref.URI)
    }
}
```

### output

```
$ go run example-htmlrefs.go
link : /favicon.ico
link : /css/main.css
a : /
a : /talks
a : /projects
a : http://www.clevertech.biz/
a : /talks
a : /projects
script : /components/jquery/dist/jquery.min.js
script : /components/bootstrap/dist/js/bootstrap.min.js
```

### struct

`htmlrefs` returns a slice of `Reference`s

```
type Reference struct {
    URI, Token string
}
```

### credits

Inspired by https://github.com/JackDanger/collectlinks