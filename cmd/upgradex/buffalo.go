package upgradex

import (
	"os"
	"path/filepath"

	"github.com/gobuffalo/buffalo/meta"
	"github.com/pkg/errors"
)

var ErrNotBuffaloApp = errors.New("not a Buffalo app")

// go get -u -v -tags sqlite github.com/gobuffalo/buffalo/buffalo
func BuffaloCLI(opts Options) error {
	if opts.SkipBuffalo {
		return nil
	}
	return getter("github.com/gobuffalo/buffalo/buffalo", opts)
}

func BuffaloApp(opts Options) error {
	if opts.SkipApp {
		return nil
	}
	if !isApp(opts.Path) {
		return errors.Wrapf(ErrNotBuffaloApp, "path: %s", opts.Path)
	}

	app := meta.New(opts.Path)
	if app.WithDep {
		return withDep(app, opts)
	}
	return withoutDep(app, opts)
}

func withDep(app meta.App, opts Options) error {
	return execr("dep", "ensure", "-v", "-update")
}

func withoutDep(app meta.App, opts Options) error {
	opts.withTests = true
	return getter("./...", opts)
}

func isApp(root string) bool {
	_, err := os.Stat(filepath.Join(root, ".buffalo.dev.yml"))
	return err == nil
}
