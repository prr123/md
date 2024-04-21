// new renderer

package mdjs

import (
	"fmt"
	"io"

	"github.com/gomarkdown/markdown/ast"
//	"github.com/gomarkdown/markdown/parser"

	)

/*
type Renderer interface {
	// RenderNode renders markdown node to w.
	// It's called once for a leaf node.
	// It's called twice for non-leaf nodes:
	// * first with entering=true
	// * then with entering=false
	//
	// Return value is a way to tell the calling walker to adjust its walk
	// pattern: e.g. it can terminate the traversal by returning Terminate. Or it
	// can ask the walker to skip a subtree of this node by returning SkipChildren.
	// The typical behavior is to return GoToNext, which asks for the usual
	// traversal to the next node.
	RenderNode(w io.Writer, node ast.Node, entering bool) ast.WalkStatus

	// RenderHeader is a method that allows the renderer to produce some
	// content preceding the main body of the output document. The header is
	// understood in the broad sense here. For example, the default HTML
	// renderer will write not only the HTML document preamble, but also the
	// table of contents if it was requested.
	//
	// The method will be passed an entire document tree, in case a particular
	// implementation needs to inspect it to produce output.
	//
	// The output should be written to the supplied writer w. If your
	// implementation has no header to write, supply an empty implementation.
	RenderHeader(w io.Writer, ast ast.Node)

	// RenderFooter is a symmetric counterpart of RenderHeader.
	RenderFooter(w io.Writer, ast ast.Node)
}
*/

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
}



func (r *Renderer) RenderNode(w io.Writer, node ast.Node, entering bool) ast.WalkStatus {

	switch node := node.(type) {
	case *ast.Text:
		fmt.Println(" -- text")
//		r.Text(w, node)
	case *ast.Softbreak:
//		r.CR(w)
		// TODO: make it configurable via out(renderer.softbreak)
	case *ast.Hardbreak:
//		r.HardBreak(w, node)
	case *ast.NonBlockingSpace:
//		r.NonBlockingSpace(w, node)
	case *ast.Emph:
//		r.OutOneOf(w, entering, "<em>", "</em>")
	case *ast.Strong:
//		r.OutOneOf(w, entering, "<strong>", "</strong>")
	case *ast.Del:
//		r.OutOneOf(w, entering, "<del>", "</del>")
	case *ast.BlockQuote:
//		tag := TagWithAttributes("<blockquote", BlockAttrs(node))
//		r.OutOneOfCr(w, entering, tag, "</blockquote>")
	case *ast.Aside:
//		tag := TagWithAttributes("<aside", BlockAttrs(node))
//		r.OutOneOfCr(w, entering, tag, "</aside>")
	case *ast.Link:
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
		if entering {
			fmt.Println("-- par start")
		} else {
			fmt.Println("-- par end")
		}
//		r.Paragraph(w, node, entering)
	case *ast.HTMLSpan:
//		r.HTMLSpan(w, node)
	case *ast.HTMLBlock:
//		r.HTMLBlock(w, node)
	case *ast.Heading:
		if entering {
			fmt.Println("-- heading start")
		} else {
			fmt.Println("-- heading end")
		}
//		r.Heading(w, node, entering)
	case *ast.HorizontalRule:
//		r.HorizontalRule(w, node)
	case *ast.List:
//		r.List(w, node, entering)
	case *ast.ListItem:
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
//		r.OutOneOf(w, true, "<sub>", "</sub>")
//		if entering {
//			Escape(w, node.Literal)
//		}
//		r.OutOneOf(w, false, "<sub>", "</sub>")
	case *ast.Superscript:
//		r.OutOneOf(w, true, "<sup>", "</sup>")
//		if entering {
//			Escape(w, node.Literal)
//		}
//		r.OutOneOf(w, false, "<sup>", "</sup>")
	case *ast.Footnotes:
		// nothing by default; just output the list.
	default:
		panic(fmt.Sprintf("Unknown node %T", node))
	}

	return ast.GoToNext
}

func (r *Renderer) RenderHeader(w io.Writer, ast ast.Node) {

	return
}


func (r *Renderer) RenderFooter(w io.Writer, ast ast.Node) {

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

