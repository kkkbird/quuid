# QUUID

Generate uuid like redis stream entry ID, refer to https://redis.io/topics/streams-intro

## Example

``` golang
package main

import (
	"fmt"
	"github.com/kkkbird/quuid"
)

func main() {
	id := quuid.UUID()
	fmt.Println(id)
}
```

``` shell
1548237747771-0
```