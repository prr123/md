// new renderer
// V2 add header display
// v3 add text

package mdjs

import (
	"fmt"
	"io"

	"github.com/gomarkdown/markdown/ast"
//	"github.com/gomarkdown/markdown/parser"

)



type Renderer struct {
	//Opts RendererOptions

	closeTag string // how to end singleton tags: either " />" or ">"

	// Track heading IDs to prevent ID collision in a single generation.
	headingIDs map[string]int

	lastOutputLen int

	// if > 0, will strip html tags in Out and Outs
	DisableTags int

	// IsSafeURLOverride allows overriding the default URL matcher. URL is
	// safe if the overriding function returns true. Can be used to extend
	// the default list of safe URLs.
//	IsSafeURLOverride func(url []byte) bool

//	sr *SPRenderer

	documentMatter ast.DocumentMatters // keep track of front/main/back matter.

	txtlev int

	list [10]string
	level int
	skipPar bool
	nest int
	nestPar [10]bool
	nestTyp [10]byte
}



func (r *Renderer) RenderNode(w io.Writer, node ast.Node, entering bool) ast.WalkStatus {
/*
	par := node.GetParent()
	if par != nil {
		children := par.GetChildren()
		fmt.Printf("children %d\n", len(children))
	} else {
		fmt.Println("no children")
	}
*/

	switch node := node.(type) {
	case *ast.Text:
		fmt.Fprintf(w,"{typ:\"txt\", txt:`%s`},\n", node.Literal);

	case *ast.Softbreak:
//		fmt.Fprintf(w,"\n")
//		r.CR(w)
		// TODO: make it configurable via out(renderer.softbreak)
	case *ast.Hardbreak:
		fmt.Fprintf(w,"{typ:\"br\", ch:[]},\n")
//		r.HardBreak(w, node)
	case *ast.NonBlockingSpace:
		fmt.Fprintf(w,"{typ:\"txt\", txt:\"&nbsp\"}");
//		r.NonBlockingSpace(w, node)
	case *ast.Emph:
		if entering {
			r.txtlev++
			if r.txtlev == 1 {
				fmt.Fprintf(w,"{typ:\"span\",style:{")
			}
			fmt.Fprintf(w,"fontStyle:\"italic\",")
			ch := node.GetChildren()
			if len(ch) == 1 {
				switch ch[0].(type) {
					case *ast.Text:
					fmt.Fprintf(w,"}, ch:[")
					default:
				}
			}
		} else {
			r.txtlev--
			if r.txtlev == 0 {
				fmt.Fprintf(w,"]},")
			}
		}
	case *ast.Strong:
		if entering {
			r.txtlev++
			if r.txtlev == 1 {
				fmt.Fprintf(w,"{typ:\"span\",style:{")
			}
			fmt.Fprintf(w,"fontWeight:\"bold\",")
			ch := node.GetChildren()
			if len(ch) == 1 {
				switch ch[0].(type) {
					case *ast.Text:
					fmt.Fprintf(w,"}, ch:[")
					default:
				}
			}
		} else {
			r.txtlev--
			if r.txtlev == 0 {
				fmt.Fprintf(w,"]},")
			}
		}
	case *ast.Del:
		if entering {
			r.txtlev++
			if r.txtlev == 1 {
				fmt.Fprintf(w,"{typ:\"span\",style:{")
			}
			fmt.Fprintf(w,"textDecoration:\"line-through\",")
			ch := node.GetChildren()
			if len(ch) == 1 {
				switch ch[0].(type) {
					case *ast.Text:
					fmt.Fprintf(w,"}, ch:[")
					default:
				}
			}
		} else {
			r.txtlev--
			if r.txtlev == 0 {
				fmt.Fprintf(w,"]},")
			}
		}
	case *ast.BlockQuote:
//		fmt.Fprintf(w,"{typ:\"br\"},\n")
		if entering {
			fmt.Fprintf(w,"{typ:\"BLOCKQUOTE\", style:{padding:\"50px 30px 50px 30px\"}, ch:[")
		} else {
			fmt.Fprintf(w,"]},\n")
		}
	case *ast.Aside:
//		tag := TagWithAttributes("<aside", BlockAttrs(node))
//		r.OutOneOfCr(w, entering, tag, "</aside>")
	case *ast.Link:
		if entering {
			r.level++
			r.list[r.level] = "lel"
			fmt.Fprintf(w,"{typ:\"a\",href:\"%s\",ch:[\n",node.Destination)
		} else {
			r.level--
			fmt.Fprintf(w,"]},\n")
		}

//		r.Link(w, node, entering)
	case *ast.CrossReference:
//		link := &ast.Link{Destination: append([]byte("#"), node.Destination...)}
//		r.Link(w, link, entering)
	case *ast.Citation:
//		r.Citation(w, node)
	case *ast.Image:
//		if r.Opts.Flags&SkipImages != 0 {
//			return ast.SkipChildren
//		}
//		r.Image(w, node, entering)
	case *ast.Code:
//		r.Code(w, node)
	case *ast.CodeBlock:
//		r.CodeBlock(w, node)
	case *ast.Caption:
//		r.Caption(w, node, entering)
	case *ast.CaptionFigure:
//		r.CaptionFigure(w, node, entering)
	case *ast.Document:
		// do nothing
	case *ast.Paragraph:
	fmt.Printf("dbg -- par: %t\n", r.skipPar)
		if entering {
			if !r.skipPar {
				r.level++
				r.list[r.level] = "pel"
				fmt.Fprintf(w,"{typ:\"p\",ch:[\n")
			}
		} else {
			if !r.skipPar {
				fmt.Fprintf(w,"{typ:\"txt\", txt:`\n`},");
				fmt.Fprintf(w,"]},\n")
				r.level--
			}
		}
//		r.Paragraph(w, node, entering)
	case *ast.HTMLSpan:
		fmt.Fprintf(w,"{typ:\"span\",")
//			fmt.Fprintf(w, "hel.id = \"%s\";\n", node.HeadingID)
		fmt.Fprintf(w, "txt:`%s`}\n,", node.Literal)

//		r.HTMLSpan(w, node)
	case *ast.HTMLBlock:
//		r.HTMLBlock(w, node)
	case *ast.Heading:
		if entering {
			r.level++
			r.list[r.level] = "hel"
			fmt.Fprintf(w,"{typ:\"h%d\",", node.Level)
			fmt.Fprintf(w, "id:\"%s\",", node.HeadingID)
			fmt.Fprintf(w, "ch:[\n")
		} else {
			fmt.Fprintf(w,"]},\n")
			r.level--
		}

	case *ast.HorizontalRule:
		fmt.Fprintf(w,"{typ:\"hr\", ch:[]},")

//		r.HorizontalRule(w, node)
	case *ast.List:

		if entering {
			r.level++
			t := node.ListFlags&ast.ListTypeOrdered >0
fmt.Printf("dbg -- list: %v %t tight: %t nest: %d\n", node.ListFlags, t, node.Tight, r.nest)
			r.nest++
			listStylDesc := fmt.Sprintf("marginLeft:\"%dpx\"",+(r.nest-1)*10)
			if t {
				r.list[r.level] = "ol"
				r.nestTyp[r.nest] = 'o'
				fmt.Fprintf(w,"{typ:\"ol\",style:{%s},ch:[\n", listStylDesc)
			} else {
				r.list[r.level] = "ul"
				r.nestTyp[r.nest] = 'u'
				fmt.Fprintf(w,"{typ:\"ul\", style:{%s},ch:[\n", listStylDesc)
			}
			r.nestPar[r.nest] = false
			if node.Tight {r.nestPar[r.nest] = true}
		} else {
			fmt.Fprintf(w,"]},\n")
			r.level--
			r.nestPar[r.nest] = false
			r.nestTyp[r.nest] = ' '
			r.nest--
		}
//		r.List(w, node, entering)
	case *ast.ListItem:
		bulStylDesc := ""
		if entering {
			if r.nestPar[r.nest] {r.skipPar = true}
			if r.nestTyp[r.nest] == 'o' {bulStylDesc = "listStyleType:\"decimal\""}
			if r.nestTyp[r.nest] == 'u' {
				bulStylDesc = "listStyleType:\"disc\""
				if r.nest%2 == 0 {bulStylDesc = "listStyleType:\"circle\""}
			}
			fmt.Fprintf(w,"{typ:\"li\", style:{%s}, ch:[\n", bulStylDesc)
		} else {
			r.skipPar = false
			fmt.Fprintf(w,"]},\n")
		}
//		r.ListItem(w, node, entering)
	case *ast.Table:
//		tag := TagWithAttributes("<table", BlockAttrs(node))
//		r.OutOneOfCr(w, entering, tag, "</table>")
	case *ast.TableCell:
//		r.TableCell(w, node, entering)
	case *ast.TableHeader:
//		r.OutOneOfCr(w, entering, "<thead>", "</thead>")
	case *ast.TableBody:
//		r.TableBody(w, node, entering)
	case *ast.TableRow:
//		r.OutOneOfCr(w, entering, "<tr>", "</tr>")
	case *ast.TableFooter:
//		r.OutOneOfCr(w, entering, "<tfoot>", "</tfoot>")
	case *ast.Math:
//		r.OutOneOf(w, true, `<span class="math inline">\(`, `\)</span>`)
//		EscapeHTML(w, node.Literal)
//		r.OutOneOf(w, false, `<span class="math inline">\(`, `\)</span>`)
	case *ast.MathBlock:
//		r.OutOneOf(w, entering, `<p><span class="math display">\[`, `\]</span></p>`)
//		if entering {
//			EscapeHTML(w, node.Literal)
//		}
	case *ast.DocumentMatter:
//		r.DocumentMatter(w, node, entering)
	case *ast.Callout:
//		r.Callout(w, node)
	case *ast.Index:
//		r.Index(w, node)
	case *ast.Subscript:
		if entering {
//			r.txtattr = "sub"
		} else {
//			r.txtattr = ""
		}
	case *ast.Superscript:
		if entering {
//			r.txtattr = "sup"
		} else {
//			r.txtattr = ""
		}
	case *ast.Footnotes:
		// nothing by default; just output the list.
	default:
		panic(fmt.Sprintf("Unknown node %T", node))
	}

	return ast.GoToNext
}

func (r *Renderer) RenderHeader(w io.Writer, ast ast.Node) {
	fmt.Fprintf(w,"const frag= {typ:\"div\", style:{margin:\"20px\", listStylePosition: \"inside\",},ch:[\n")
	r.list[0] = "div"
	r.level = 0
//	fmt.Fprintf(w, "{\"doc\":[\n")
	return
}


func (r *Renderer) RenderFooter(w io.Writer, ast ast.Node) {
	
	fmt.Fprintf(w,"],};\n")
	return
}

func NewRenderer() *Renderer {
	// configure the rendering engine
	closeTag := ">"
//	if opts.Flags&UseXHTML != 0 {
//		closeTag = " />"
//	}

/*
	if opts.FootnoteReturnLinkContents == "" {
		opts.FootnoteReturnLinkContents = `<sup>[return]</sup>`
	}
	if opts.CitationFormatString == "" {
		opts.CitationFormatString = `<sup>[%s]</sup>`
	}
	if opts.Generator == "" {
		opts.Generator = `  <meta name="GENERATOR" content="github.com/gomarkdown/markdown markdown processor for Go`
	}
*/
	return &Renderer{
//		Opts: opts,

		closeTag:   closeTag,
		headingIDs: make(map[string]int),

//		sr: NewSmartypantsRenderer(opts.Flags),
	}
}

