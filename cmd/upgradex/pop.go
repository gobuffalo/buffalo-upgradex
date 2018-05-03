package upgradex

import "github.com/pkg/errors"

// go get -u -v -tags sqlite github.com/gobuffalo/pop/soda
func SodaCLI(opts Options) error {
	if opts.SkipPop {
		return nil
	}
	if err := getter("github.com/gobuffalo/pop", opts); err != nil {
		return errors.WithStack(err)
	}
	return getter("github.com/gobuffalo/pop/soda", opts)
}
