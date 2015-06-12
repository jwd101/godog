package godog

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/DATA-DOG/godog/gherkin"
)

var cfg config

type config struct {
	featuresPath  string
	formatterName string
}

func init() {
	// @TODO: colorize flag help output
	flag.StringVar(&cfg.featuresPath, "features", "features", "Path to feature files")
	flag.StringVar(&cfg.formatterName, "formatter", "pretty", "Formatter name")
}

func (c config) validate() error {
	inf, err := os.Stat(c.featuresPath)
	if err != nil {
		return err
	}
	if !inf.IsDir() {
		return fmt.Errorf("feature path \"%s\" is not a directory.", c.featuresPath)
	}
	switch c.formatterName {
	case "pretty":
		// ok
	default:
		return fmt.Errorf("Unsupported formatter name: %s", c.formatterName)
	}
	return nil
}

func (c config) features() (lst []*gherkin.Feature, err error) {
	return lst, filepath.Walk(cfg.featuresPath, func(p string, f os.FileInfo, err error) error {
		if err == nil && !f.IsDir() && strings.HasSuffix(p, ".feature") {
			ft, err := gherkin.Parse(p)
			if err != nil {
				return err
			}
			lst = append(lst, ft)
		}
		return err
	})
}
