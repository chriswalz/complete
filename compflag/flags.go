// Autogenerated by go run compflag/gen/main.go. DO NOT EDIT.

package compflag

import (
	"flag"
	"fmt"
	"time"

	"github.com/chriswalz/complete/v2/predict"
)

// String if a flag function for a flag of type string.
func String(name string, value string, usage string, options ...predict.Option) *string {
	return CommandLine.String(name, value, usage, options...)
}

// StringVar if a flag function for a flag of already exiting variable of type string.
func StringVar(v *string, name string, value string, usage string, options ...predict.Option) {
	CommandLine.StringVar(v, name, value, usage, options...)
}

// Bool if a flag function for a flag of type bool.
func Bool(name string, value bool, usage string, options ...predict.Option) *bool {
	return CommandLine.Bool(name, value, usage, options...)
}

// BoolVar if a flag function for a flag of already exiting variable of type bool.
func BoolVar(v *bool, name string, value bool, usage string, options ...predict.Option) {
	CommandLine.BoolVar(v, name, value, usage, options...)
}

// Int if a flag function for a flag of type int.
func Int(name string, value int, usage string, options ...predict.Option) *int {
	return CommandLine.Int(name, value, usage, options...)
}

// IntVar if a flag function for a flag of already exiting variable of type int.
func IntVar(v *int, name string, value int, usage string, options ...predict.Option) {
	CommandLine.IntVar(v, name, value, usage, options...)
}

// Duration if a flag function for a flag of type time.Duration.
func Duration(name string, value time.Duration, usage string, options ...predict.Option) *time.Duration {
	return CommandLine.Duration(name, value, usage, options...)
}

// DurationVar if a flag function for a flag of already exiting variable of type time.Duration.
func DurationVar(v *time.Duration, name string, value time.Duration, usage string, options ...predict.Option) {
	CommandLine.DurationVar(v, name, value, usage, options...)
}

// String if a flag function for a flag of type string.
func (fs *FlagSet) String(name string, value string, usage string, options ...predict.Option) *string {
	p := new(string)
	fs.StringVar(p, name, value, usage, options...)
	return p
}

// StringVar if a flag function for a flag of already exiting variable of type string.
func (fs *FlagSet) StringVar(p *string, name string, value string, usage string, options ...predict.Option) {
	(*flag.FlagSet)(fs).Var(newStringValue(value, p, predict.Options(options...)), name, usage)
}

// Bool if a flag function for a flag of type bool.
func (fs *FlagSet) Bool(name string, value bool, usage string, options ...predict.Option) *bool {
	p := new(bool)
	fs.BoolVar(p, name, value, usage, options...)
	return p
}

// BoolVar if a flag function for a flag of already exiting variable of type bool.
func (fs *FlagSet) BoolVar(p *bool, name string, value bool, usage string, options ...predict.Option) {
	(*flag.FlagSet)(fs).Var(newBoolValue(value, p, predict.Options(options...)), name, usage)
}

// Int if a flag function for a flag of type int.
func (fs *FlagSet) Int(name string, value int, usage string, options ...predict.Option) *int {
	p := new(int)
	fs.IntVar(p, name, value, usage, options...)
	return p
}

// IntVar if a flag function for a flag of already exiting variable of type int.
func (fs *FlagSet) IntVar(p *int, name string, value int, usage string, options ...predict.Option) {
	(*flag.FlagSet)(fs).Var(newIntValue(value, p, predict.Options(options...)), name, usage)
}

// Duration if a flag function for a flag of type time.Duration.
func (fs *FlagSet) Duration(name string, value time.Duration, usage string, options ...predict.Option) *time.Duration {
	p := new(time.Duration)
	fs.DurationVar(p, name, value, usage, options...)
	return p
}

// DurationVar if a flag function for a flag of already exiting variable of type time.Duration.
func (fs *FlagSet) DurationVar(p *time.Duration, name string, value time.Duration, usage string, options ...predict.Option) {
	(*flag.FlagSet)(fs).Var(newDurationValue(value, p, predict.Options(options...)), name, usage)
}

// ============================================================================================== //

type stringValue struct {
	v *string
	predict.Config
}

func newStringValue(val string, p *string, c predict.Config) *stringValue {
	*p = val
	return &stringValue{v: p, Config: c}
}

func (v *stringValue) Set(val string) error {
	var err error
	*v.v, err = parseString(val)
	if err != nil {
		return fmt.Errorf("bad value for String flag")
	}
	return v.Check(val)
}

func (v *stringValue) Get() interface{} {
	return *v.v
}

func (v *stringValue) String() string {
	if v == nil || v.v == nil {
		return ""
	}
	return formatString(*v.v)
}

func (v *stringValue) Predict(prefix string) []string {
	if v.Predictor != nil {
		return v.Predictor.Predict(prefix)
	}
	return []string{""}
}

// ============================================================================================== //

type boolValue struct {
	v *bool
	predict.Config
}

func newBoolValue(val bool, p *bool, c predict.Config) *boolValue {
	*p = val
	return &boolValue{v: p, Config: c}
}

func (v *boolValue) Set(val string) error {
	var err error
	*v.v, err = parseBool(val)
	if err != nil {
		return fmt.Errorf("bad value for Bool flag")
	}
	return v.Check(val)
}

func (v *boolValue) Get() interface{} {
	return *v.v
}

func (v *boolValue) String() string {
	if v == nil || v.v == nil {
		return ""
	}
	return formatBool(*v.v)
}

func (v *boolValue) IsBoolFlag() bool { return true }

func (v *boolValue) Predict(prefix string) []string {
	if v.Predictor != nil {
		return v.Predictor.Predict(prefix)
	}
	return predictBool(*v.v, prefix)
}

// ============================================================================================== //

type intValue struct {
	v *int
	predict.Config
}

func newIntValue(val int, p *int, c predict.Config) *intValue {
	*p = val
	return &intValue{v: p, Config: c}
}

func (v *intValue) Set(val string) error {
	var err error
	*v.v, err = parseInt(val)
	if err != nil {
		return fmt.Errorf("bad value for Int flag")
	}
	return v.Check(val)
}

func (v *intValue) Get() interface{} {
	return *v.v
}

func (v *intValue) String() string {
	if v == nil || v.v == nil {
		return ""
	}
	return formatInt(*v.v)
}

func (v *intValue) Predict(prefix string) []string {
	if v.Predictor != nil {
		return v.Predictor.Predict(prefix)
	}
	return []string{""}
}

// ============================================================================================== //

type durationValue struct {
	v *time.Duration
	predict.Config
}

func newDurationValue(val time.Duration, p *time.Duration, c predict.Config) *durationValue {
	*p = val
	return &durationValue{v: p, Config: c}
}

func (v *durationValue) Set(val string) error {
	var err error
	*v.v, err = parseDuration(val)
	if err != nil {
		return fmt.Errorf("bad value for Duration flag")
	}
	return v.Check(val)
}

func (v *durationValue) Get() interface{} {
	return *v.v
}

func (v *durationValue) String() string {
	if v == nil || v.v == nil {
		return ""
	}
	return formatDuration(*v.v)
}

func (v *durationValue) Predict(prefix string) []string {
	if v.Predictor != nil {
		return v.Predictor.Predict(prefix)
	}
	return []string{""}
}
