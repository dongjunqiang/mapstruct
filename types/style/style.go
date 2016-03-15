package style

import (
	"github.com/geo-data/mapfile/mapfile/encode"
	"github.com/geo-data/mapfile/types"
	"github.com/geo-data/mapfile/types/color"
)

type Style struct {
	Color         *color.Color `json:",omitempty"`
	OutlineColor  *color.Color `json:",omitempty"`
	Symbol        types.Union  `json:",omitempty"`
	Size          types.Union  `json:",omitempty"`
	Width         types.Union  `json:",omitempty"`
	GeomTransform types.String `json:",omitempty"`
}

func (s *Style) Encode(enc *encode.MapfileEncoder) (err error) {
	if err = enc.TokenStart("STYLE"); err != nil {
		return
	}

	if s.Color != nil {
		if err = enc.TokenStringer("COLOR", s.Color); err != nil {
			return
		}
	}
	if s.OutlineColor != nil {
		if err = enc.TokenStringer("OUTLINECOLOR", s.OutlineColor); err != nil {
			return
		}
	}
	if err = enc.TokenUnion("SYMBOL", s.Symbol); err != nil {
		return
	}
	if err = enc.TokenUnion("SIZE", s.Size); err != nil {
		return
	}
	if err = enc.TokenUnion("WIDTH", s.Width); err != nil {
		return
	}
	if err = enc.TokenStringer("GEOMTRANSFORM", s.GeomTransform); err != nil {
		return
	}

	if err = enc.TokenEnd(); err != nil {
		return
	}

	return
}
