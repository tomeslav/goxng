goxng
=====

A basic xng encoding library

A sample main program : 

```
package main

import (
	"fmt"
	"goxng"
	"os"
)

func main() {
	
	imgs := os.Args
	if len(imgs) < 2 {
		fmt.Println("Please specify images filenames")
		return
	}
	
	xml, err := goxng.GetXng(imgs[1:], 60)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(xml)
}
```
