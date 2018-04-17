package upgradex

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/pkg/errors"
)

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
	args := []string{"get", "-u", "-v"}
	if opts.withTests {
		args = append(args, "-t")
	}
	if opts.WithSQLite {
		args = append(args, "-tags", "sqlite")
	}
	args = append(args, pkg)
	return execr("go", args...)
}

func execr(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	fmt.Println(strings.Join(cmd.Args, " "))
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	return cmd.Run()
}
