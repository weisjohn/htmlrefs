# htmlrefs

Package htmlrefs returns a slice containing the unique resource URIs from an `io.Reader`.

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
    links := htmlrefs.All(resp.Body)
    fmt.Println(links)
}
```