package main

// example for https://blog.kowalczyk.info/article/cxn3/advanced-markdown-processing-in-go.html

import (
	"os"
	"fmt"
	"log"
	"bytes"

//	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/ast"
//	"github.com/gomarkdown/markdown/html"
	mdjson "goDemo/md/mdjson"
	"github.com/gomarkdown/markdown/parser"
	util "github.com/prr123/utility/utilLib"
)

func RenderDom(doc ast.Node, renderer mdjson.Renderer) []byte {
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

	renderer := mdjson.NewRenderer()

	return RenderDom(doc,*renderer)
}

func main() {

    numArgs := len(os.Args)
	var md []byte
	var mds =`
# header

Sample text.

[link](http://example.com)
`
	md = []byte(mds)

	flags:=[]string{"dbg","md", "script"}

    useStr := "[/md=<markdown file>] [/script=<script file>] [/dbg]"
    helpStr := fmt.Sprintf("help: The program cmvert md files to a js Dom script\n")

    if numArgs > len(flags)+1 {
        fmt.Println("too many arguments in cl!")
        fmt.Println("usage: %s %s\n", os.Args[0], useStr)
        os.Exit(1)
    }


    if numArgs == 2 {
        if os.Args[1] == "help" {
            fmt.Printf("usage is: %s %s\n", os.Args[0], useStr)
            fmt.Printf("%s\n", helpStr)
            os.Exit(1)
        }
    }

    flagMap, err := util.ParseFlags(os.Args, flags)
    if err != nil {log.Fatalf("util.ParseFlags: %v\n", err)}

    dbg := false
    _, ok := flagMap["dbg"]
    if ok {dbg = true}

    mdFilnam := ""
    mdval, ok := flagMap["md"]
    if ok {
        if mdval.(string) == "none" {log.Fatalf("error -- no markdown file provided with /md flag!")}
        mdFilnam = mdval.(string)
//	idx := bytes.Index(mdFilnam,".md")
		mdFilnam += ".md"
		md, err = os.ReadFile(mdFilnam)
		if err != nil {log.Fatalf("error -- cannot read md: %v", err)}
//		mds = string(md)
	} 

    outFilnam := "test.json"
    oval, ok := flagMap["script"]
    if ok {
        if oval.(string) == "none" {log.Fatalf("error -- no script file provided with /script flag!")}
        outFilnam = oval.(string)
//	idx := bytes.Index(mdFilnam,".md")
		outFilnam += ".js"
	}
	outFilnam = "json/" + outFilnam

	outfil, err := os.Create(outFilnam)
	if err != nil {log.Fatalf("error -- cannot create script file: %v", err)}
	defer outfil.Close()

    if dbg {
		if len(mdFilnam)  == 0 {
			log.Printf("debug -- no md file!\n")
		} else {
        	log.Printf("debug -- md file: %s\n",mdFilnam)
		}
		log.Printf("debug -- script file: %s\n",outFilnam)
    }



	script := mdToJsDom(md)

	fmt.Printf("\n\n--- Markdown:\n%s\n\n--- jsDom:\n%s\n", md, script)

	if outfil != nil {
		_, err = outfil.Write(script)
		if err !=nil {log.Fatalf("error -- writing to script file! %v\n",err)}
	}
}
