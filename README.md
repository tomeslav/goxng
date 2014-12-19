goxng
=====

A basic xng encoding library.

Xng is a simple image format based on svg for moving images. Unlike all the clever gif repleacement formats, xng can be displayed in any modern browser.

Xng is an idea by Jasper St. Pierre. See his blog entry for more details :
http://blog.mecheye.net/2014/10/xng-gifs-but-better-and-also-magical/

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
