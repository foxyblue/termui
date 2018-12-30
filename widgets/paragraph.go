// Copyright 2017 Zack Guo <zack.y.guo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license that can
// be found in the LICENSE file.

package widgets

import (
	"image"

	. "github.com/gizak/termui"
)

type Paragraph struct {
	Block
	Text      string
	TextAttrs AttrPair
}

func NewParagraph(s string) *Paragraph {
	return &Paragraph{
		Block:     *NewBlock(),
		Text:      s,
		TextAttrs: Theme.Paragraph.Text,
	}
}

func (p *Paragraph) Draw(buf *Buffer) {
	p.Block.Draw(buf)

	point := p.Min.Add(image.Pt(1, 1))
	cells := WrapText(ParseText(p.Text, p.TextAttrs), p.Dx()-2)

	for i := 0; i < len(cells) && point.Y < p.Max.Y-1; i++ {
		if cells[i].Rune == '\n' {
			point = image.Pt(p.Min.X+1, point.Y+1)
		} else {
			buf.SetCell(cells[i], point)
			point = point.Add(image.Pt(1, 0))
		}
	}
}
