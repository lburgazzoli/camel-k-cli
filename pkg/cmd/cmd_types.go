package cmd

import (
	pluralize "github.com/gertd/go-pluralize"
	"github.com/urfave/cli/v2"
	"github.com/urfave/cli/v2/altsrc"
	"time"
)

func NewPrefixedInputSourceContext(flags []cli.Flag, delegate altsrc.InputSourceContext, prefix string) altsrc.InputSourceContext {
	return prefixedInputSourceContext{
		flags:    flags,
		delegate: delegate,
		prefix:   prefix,
	}
}

// prefixedInputSourceContext ---
type prefixedInputSourceContext struct {
	flags    []cli.Flag
	delegate altsrc.InputSourceContext
	prefix   string
}

func (d prefixedInputSourceContext) Source() string {
	return d.delegate.Source()
}
func (d prefixedInputSourceContext) Int(name string) (int, error) {
	return d.delegate.Int(d.prefix + d.computeKey(name))
}
func (d prefixedInputSourceContext) Duration(name string) (time.Duration, error) {
	return d.delegate.Duration(d.prefix + d.computeKey(name))
}
func (d prefixedInputSourceContext) Float64(name string) (float64, error) {
	return d.delegate.Float64(d.prefix + d.computeKey(name))
}
func (d prefixedInputSourceContext) String(name string) (string, error) {
	return d.delegate.String(d.prefix + d.computeKey(name))
}
func (d prefixedInputSourceContext) StringSlice(name string) ([]string, error) {
	return d.delegate.StringSlice(d.prefix + d.computeKey(name))
}
func (d prefixedInputSourceContext) IntSlice(name string) ([]int, error) {
	return d.delegate.IntSlice(d.prefix + d.computeKey(name))
}
func (d prefixedInputSourceContext) Generic(name string) (cli.Generic, error) {
	return d.delegate.Generic(d.prefix + d.computeKey(name))
}
func (d prefixedInputSourceContext) Bool(name string) (bool, error) {
	return d.delegate.Bool(d.prefix + d.computeKey(name))
}

func (d prefixedInputSourceContext) computeKey(name string) string {
	for i := range d.flags {
		for _, flagName := range d.flags[i].Names() {
			if flagName != name {
				continue
			}
			switch d.flags[i].(type) {
			case *cli.StringSliceFlag:
				return pluralize.NewClient().Plural(name)
			case *altsrc.StringSliceFlag:
				return pluralize.NewClient().Plural(name)
			}
		}
	}

	return name
}
