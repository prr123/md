package main

// example for https://blog.kowalczyk.info/article/cxn3/advanced-markdown-processing-in-go.html

import (
	"os"
	"fmt"
	"bytes"

//	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/ast"
//	"github.com/gomarkdown/markdown/html"
	mdjs "goDemo/md/mdjs"
	"github.com/gomarkdown/markdown/parser"

)

func RenderDom(doc ast.Node, renderer mdjs.Renderer) []byte {
    var buf bytes.Buffer
    renderer.RenderHeader(&buf, doc)
    ast.WalkFunc(doc, func(node ast.Node, entering bool)(walk ast.WalkStatus) {
        xy := renderer.RenderNode(&buf, node, entering)
//        fmt.Printf("walk status: %d %s\n", xy, buf)
        return xy
    })
    renderer.RenderFooter(&buf, doc)
    return buf.Bytes()
}


var mds = `# header

Sample text.

[link](http://example.com)
`

var printAst = true

func mdToJsDom(md []byte) []byte {
	// create markdown parser with extensions
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock
	p := parser.NewWithExtensions(extensions)
	doc := p.Parse(md)

	if printAst {
		fmt.Print("--- AST tree:\n")
		ast.Print(os.Stdout, doc)
		fmt.Print("\n")
	}

	// create HTML renderer with extensions
//	htmlFlags := html.CommonFlags | html.HrefTargetBlank
//	opts := html.RendererOptions{Flags: htmlFlags}
//	renderer := html.NewRenderer(opts)

	renderer := mdjs.NewRenderer()

	return RenderDom(doc,*renderer)
}

func main() {
	md := []byte(mds)
	script := mdToJsDom(md)

	fmt.Printf("\n\n--- Markdown:\n%s\n\n--- jsDom:\n%s\n", md, script)
}
