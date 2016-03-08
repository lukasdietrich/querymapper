# querymapper [![coverage](https://gocover.io/_badge/github.com/lukasdietrich/querymapper)](https://gocover.io/github.com/lukasdietrich/querymapper)

Mapper for [net/url](https://golang.org/pkg/net/url/) query values in Go.

```go
import (
  "net/url"
  
  "github.com/lukasdietrich/querymapper"
)

type MyStruct struct {
  SomeField int `query:"some-field"`
}

func doSomething(query url.Values) (MyStruct, error) {
  val s MyStruct
  return s, querymapper.MapQuery(query, &s)
}
```
