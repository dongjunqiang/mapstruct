package decode

import (
	"fmt"
	"github.com/geo-data/mapfile/types"
)

func (t *Decoder) Legend() (l *types.Legend, err error) {
	token := t.Value()
	if token != "LEGEND" {
		err = fmt.Errorf("expected token LEGEND, got: %s", token)
		return
	}
	t.Next()

	l = new(types.Legend)
	for t != nil {
		token := t.Value()
		switch token {
		case "IMAGECOLOR":
			if l.ImageColor, err = t.Color(); err != nil {
				return
			}
		case "END":
			return
		default:
			err = fmt.Errorf("unhandled mapfile token: %s", token)
			return
		}

		t = t.Next()
	}

	return
}