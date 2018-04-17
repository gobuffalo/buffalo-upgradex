package upgradex

// go get -u -v -tags sqlite github.com/gobuffalo/pop/soda
func SodaCLI(opts Options) error {
	if opts.SkipPop {
		return nil
	}
	return getter("github.com/gobuffalo/pop/soda", opts)
}
