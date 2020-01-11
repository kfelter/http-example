package cfg

import (
	"encoding/base64"
	"fmt"
	"os"
	"strconv"

	"github.com/pkg/errors"
)

type Var struct {
	Name  string
	Value string
}

// Getenv returns a Var object with the value being the environment var
func Getenv(Name string) Var {
	return Var{Name: Name, Value: os.Getenv(Name)}
}

// MustGetenv panics if the value of os.Getenv is nil
func MustGetenv(Name string) Var {
	value := os.Getenv(Name)
	if value == "" {
		panic(fmt.Errorf("Environment var %s must have a value", Name))
	}
	return Var{Name: Name, Value: value}
}

// GetenvWithDefault will return the value it finds in the environment unless non is found
// then it will return the default provided
func GetenvWithDefault(Name string, Default string) Var {
	value := os.Getenv(Name)
	if value == "" {
		return Var{Name: Name, Value: Default}
	}
	return Var{Name: Name, Value: value}
}

// Int converts the Var object to an int or panics if there is an error
func (v Var) Int() int {
	value, err := strconv.Atoi(v.Value)
	if err != nil {
		panic(errors.Wrap(err, v.Name))
	}
	return value
}

// Base64Decode decodes the var or panics
func (v Var) Base64Decode() []byte {
	data, err := base64.StdEncoding.DecodeString(v.Value)
	if err != nil {
		panic(errors.Wrap(err, v.Name))
	}
	return data
}
