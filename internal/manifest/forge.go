package manifest

import (
	"bytes"
	"fmt"
	"strings"
)

type Exec func(command string) error

type ForgeryName string

func (an ForgeryName) Validate() error {
	if strings.HasPrefix(string(an), "-") {
		return fmt.Errorf("action: \"%s\": name should not begin with \"-\"", an)
	}
	if strings.ContainsAny(string(an), "\t \n\r") {
		return fmt.Errorf("action: \"%s\": name should not contain whitespaces", an)
	}
	return nil
}

type Forgeries []Forgery

func (as Forgeries) Validate() error {
	var errs forgeErrors
	for _, v := range as {
		if err := v.Validate(); err != nil {
			errs = append(errs, err)
		}
	}
	if len(errs) > 0 {
		return errs
	}
	return nil

}

type forgeErrors []error

func (e forgeErrors) Error() string {
	sb := bytes.NewBufferString("errors occurred while reading forge file: \n")
	for i, v := range e {
		_, _ = fmt.Fprintf(sb, "%d: \n\t%v\n", i, v)
	}
	return sb.String()
}

type Forge struct {
	Include   []string  `yaml:"include"`
	Forgeries Forgeries `yaml:"forgeries"`
}

// Forgery defines forge action. A forge action consists of:
// name - the name of the action. Must be unique across Forge definition
// description - description of the action
type Forgery struct {
	Name        ForgeryName   `yaml:"name"`
	Description string        `yaml:"description"`
	Pre         []ForgeryName `yaml:"pre"`
	Cmd         string        `yaml:"cmd"`
}

func (a Forgery) Validate() error {
	var errs forgeErrors
	if err := a.Name.Validate(); err != nil {
		errs = append(errs, err)
	}
	for _, v := range a.Pre {
		if err := v.Validate(); err != nil {
			errs = append(errs, err)
		}
	}
	if len(errs) > 0 {
		return errs
	}
	return nil
}

func (a Forgery) Execute(execFunc Exec) (err error) {
	return execFunc(a.Cmd)
}
