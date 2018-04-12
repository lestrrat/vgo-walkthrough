package foo // import "github.com/lestrrat/vgo-walkthrough/foo"

import (
	"github.com/lestrrat/vgo-walkthrough/bar"
	"github.com/pkg/errors"
)

func Main() error {
	if err := bar.Bar(); err != nil {
		return errors.Wrap(err, `failed to call bar.Bar()`)
	}
	return nil
}
