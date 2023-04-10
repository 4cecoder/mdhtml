package main

import (
	"fmt"
	"mdhtml/basic"
)

func main() {
	md := `# Header 1
**Bold text**
*Italic text*
[Link to Google](https://www.google.com)`

	fmt.Println(basic.MarkdownToHTML(md))
}
