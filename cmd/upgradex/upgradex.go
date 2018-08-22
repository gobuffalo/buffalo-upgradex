package upgradex

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"

	"github.com/gobuffalo/envy"
	"github.com/pkg/errors"
)

var modsOn = (strings.TrimSpace(envy.Get("GO111MODULE", "off")) == "on")
var runners = []func(Options) error{BuffaloCLI, SodaCLI, BuffaloApp}

func All(opts Options) error {
	for _, r := range runners {
		if err := r(opts); err != nil {
			return errors.WithStack(err)
		}
	}
	return nil
}

func getter(pkg string, opts Options) error {
	if modsOn {
		return modGetter(pkg, opts)
	}
	args := []string{"get", "-u", "-v"}
	if opts.withTests {
		args = append(args, "-t")
	}
	if opts.WithSQLite {
		args = append(args, "-tags", "sqlite")
	}
	args = append(args, pkg)
	return execr(envy.Get("GO_BIN", "go"), args...)
}

func modGetter(pkg string, opts Options) error {
	if _, err := os.Stat("go.mod"); err != nil {
		pwd, _ := os.Getwd()
		dir, _ := ioutil.TempDir("", "buffalo-go-modules")
		defer os.Chdir(pwd)
		os.Chdir(dir)
		if err := execr(envy.Get("GO_BIN", "go"), "mod", "init", "temp"); err != nil {
			return errors.WithStack(err)
		}
	}
	args := []string{"get", "-u", "-v"}
	if opts.WithSQLite {
		args = append(args, "-tags", "sqlite")
	}
	args = append(args, pkg)
	return execr(envy.Get("GO_BIN", "go"), args...)
}

func execr(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	fmt.Println(strings.Join(cmd.Args, " "))
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	return cmd.Run()
}
