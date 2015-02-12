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
    fmt.Println(refs)
}
```


### struct

`htmlrefs` returns a slice 

```
type Reference struct {
    URI, Token string
}
```