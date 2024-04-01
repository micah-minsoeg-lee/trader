package flag

import "flag"

var configFlag = flag.String("config", "", "")

type Flag struct {
	configFlag string
}

func NewFlag() *Flag {
	flag.Parse()

	return &Flag{
		configFlag: *configFlag,
	}
}

func (f *Flag) ConfigFlag() string {
	return f.configFlag
}
