package encode

import (
	"fmt"
	"github.com/geo-data/mapfile/types/point"
)

func (enc *Encoder) EncodePoints(p point.Points) (err error) {
	if err = enc.TokenStart("POINTS"); err != nil {
		return
	}

	var points []fmt.Stringer
	for _, point := range p {
		points = append(points, point)
	}

	if err = enc.EncodeStringers(points...); err != nil {
		return
	}

	if err = enc.TokenEnd(); err != nil {
		return
	}

	return
}
