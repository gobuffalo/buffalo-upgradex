package upgradex

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/gobuffalo/buffalo/buffalo/cmd/updater"
	"github.com/pkg/errors"
)

var ErrNotBuffaloApp = errors.New("not a Buffalo app")

// go get -u -v -tags sqlite github.com/gobuffalo/buffalo/buffalo
func BuffaloCLI(opts Options) error {
	if opts.SkipBuffalo {
		return nil
	}
	if err := getter("github.com/gobuffalo/buffalo", opts); err != nil {
		return errors.WithStack(err)
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

	fmt.Println("about run the app updater")
	return updater.Run()
}

func isApp(root string) bool {
	_, err := os.Stat(filepath.Join(root, ".buffalo.dev.yml"))
	return err == nil
}
