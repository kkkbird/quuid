# QUUID

Generate uuid like redis stream style

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