package cmd

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"github.com/urfave/cli/v2/altsrc"
	"os"
)

func InitInputSourceWithContext(flags []cli.Flag, flagFileName string, keyPrefix string) cli.BeforeFunc {
	isc := func(context *cli.Context) (altsrc.InputSourceContext, error) {
		filePath := context.String(flagFileName)
		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			return nil, nil
		}

		return altsrc.NewYamlSourceFromFile(filePath)
	}

	return func(context *cli.Context) error {
		inputSource, err := isc(context)
		if err != nil {
			return fmt.Errorf("Unable to create input source with context: inner error: \n'%v'", err.Error())
		}
		if inputSource == nil {
			return nil
		}

		return altsrc.ApplyInputSourceValues(
			context,
			NewPrefixedInputSourceContext(flags, inputSource, keyPrefix),
			flags)
	}
}
