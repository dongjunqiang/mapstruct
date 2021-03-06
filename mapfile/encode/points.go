package encode

import (
	"fmt"
	"github.com/geo-data/mapstruct/types"
)

func (enc *Encoder) EncodePoints(p types.Points) (err error) {
	if err = enc.StartDirective("POINTS"); err != nil {
		return
	}

	var points []fmt.Stringer
	for _, point := range p {
		points = append(points, point)
	}

	if err = enc.EncodeStringers(points...); err != nil {
		return
	}

	if err = enc.EndDirective(); err != nil {
		return
	}

	return
}
