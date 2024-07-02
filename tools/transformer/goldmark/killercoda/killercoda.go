// Package killercoda implements parsers for Killercoda Markdown where it differs from CommonMark.
//
// The fenced code block parser is based on the Goldmark parser.
// https://github.com/yuin/goldmark/blob/15ade8aace9a9f269846fb83d36fc7bcec875cd5/parser/fcode_block.go
// MIT License

// Copyright (c) 2019 Yusuke Inuzuka

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.
package killercoda

import (
	"bytes"
	"regexp"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/text"
	"github.com/yuin/goldmark/util"
)

type fencedCodeBlock struct{}

// FencedCodeBlock is an extension that allow you to use Killercoda fenced code blocks.
// The code blocks are similar to ordinary fenced code blocks, but the closing fence may contain an action in paired curly braces.
// For example:
// ```bash
// echo "Hello, World!"
// ```{{exec}}
// In the previous example, the action is `exec`.
var FencedCodeBlock = &fencedCodeBlock{}

func (e *fencedCodeBlock) Extend(m goldmark.Markdown) {
	m.Parser().AddOptions(parser.WithBlockParsers(
		util.Prioritized(NewFencedCodeBlockParser(), 101),
	))
}

type fencedCodeBlockParser struct{}

var defaultFencedCodeBlockParser = &fencedCodeBlockParser{}

// NewFencedCodeBlockParser return a new parser.BlockParser can parse Killercoda fenced code blocks.
func NewFencedCodeBlockParser() parser.BlockParser {
	return defaultFencedCodeBlockParser
}

type fenceData struct {
	char   byte
	indent int
	length int
	node   ast.Node
}

var fencedCodeBlockInfoKey = parser.NewContextKey()

func (b *fencedCodeBlockParser) Trigger() []byte {
	return []byte{'~', '`'}
}

func (b *fencedCodeBlockParser) Open(_ ast.Node, reader text.Reader, pc parser.Context) (ast.Node, parser.State) {
	line, segment := reader.PeekLine()
	pos := pc.BlockOffset()

	if pos < 0 || (line[pos] != '`' && line[pos] != '~') {
		return nil, parser.NoChildren
	}

	findent := pos
	fenceChar := line[pos]
	i := pos
	for ; i < len(line) && line[i] == fenceChar; i++ {
	}

	oFenceLength := i - pos
	if minFenceLength := 3; oFenceLength < minFenceLength {
		return nil, parser.NoChildren
	}

	var info *ast.Text
	if i < len(line)-1 {
		rest := line[i:]
		left := util.TrimLeftSpaceLength(rest)
		right := util.TrimRightSpaceLength(rest)

		if left < len(rest)-right {
			infoStart, infoStop := segment.Start-segment.Padding+i+left, segment.Stop-right
			value := rest[left : len(rest)-right]

			if fenceChar == '`' && bytes.IndexByte(value, '`') > -1 {
				return nil, parser.NoChildren
			}

			if infoStart != infoStop {
				info = ast.NewTextSegment(text.NewSegment(infoStart, infoStop))
			}
		}
	}

	node := ast.NewFencedCodeBlock(info)

	pc.Set(fencedCodeBlockInfoKey, &fenceData{fenceChar, findent, oFenceLength, node})

	return node, parser.NoChildren
}

var actionRegexp = regexp.MustCompile(`{{(exec|copy)}}\s*$`)

func (b *fencedCodeBlockParser) Continue(node ast.Node, reader text.Reader, pc parser.Context) parser.State {
	line, segment := reader.PeekLine()
	fdata := pc.Get(fencedCodeBlockInfoKey).(*fenceData)

	w, pos := util.IndentWidth(line, reader.LineOffset())
	if w < 4 {
		i := pos
		for ; i < len(line) && line[i] == fdata.char; i++ {
		}

		length := i - pos
		if length >= fdata.length {
			matches := actionRegexp.FindStringSubmatch(string(line[i:]))
			if matches != nil {
				switch matches[1] {
				case "copy":
					node.SetAttributeString("data-killercoda-copy", "true")
				case "exec":
					node.SetAttributeString("data-killercoda-exec", "true")
				}
			}

			newline := 1

			if line[len(line)-1] != '\n' {
				newline = 0
			}

			reader.Advance(segment.Stop - segment.Start - newline + segment.Padding)

			return parser.Close
		}
	}

	pos, padding := util.IndentPositionPadding(line, reader.LineOffset(), segment.Padding, fdata.indent)
	if pos < 0 {
		pos = util.FirstNonSpacePosition(line)
		if pos < 0 {
			pos = 0
		}

		padding = 0
	}

	seg := text.NewSegmentPadding(segment.Start+pos, segment.Stop, padding)

	// If the code block line begins with a tab, keep a tab as it is.
	if padding != 0 {
		preserveLeadingTabInCodeBlock(&seg, reader, fdata.indent)
	}

	node.Lines().Append(seg)
	reader.AdvanceAndSetPadding(segment.Stop-segment.Start-pos-1, padding)

	return parser.Continue | parser.NoChildren
}

func (b *fencedCodeBlockParser) Close(node ast.Node, _ text.Reader, pc parser.Context) {
	fdata := pc.Get(fencedCodeBlockInfoKey).(*fenceData)
	if fdata.node == node {
		pc.Set(fencedCodeBlockInfoKey, nil)
	}
}

func (b *fencedCodeBlockParser) CanInterruptParagraph() bool {
	return true
}

func (b *fencedCodeBlockParser) CanAcceptIndentedLine() bool {
	return false
}

func preserveLeadingTabInCodeBlock(segment *text.Segment, reader text.Reader, indent int) {
	offsetWithPadding := reader.LineOffset() + indent
	line, ss := reader.Position()

	reader.SetPosition(line, text.NewSegment(ss.Start-1, ss.Stop))

	if offsetWithPadding == reader.LineOffset() {
		segment.Padding = 0
		segment.Start--
	}

	reader.SetPosition(line, ss)
}
