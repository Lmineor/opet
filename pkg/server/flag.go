package server

import (
	"github.com/spf13/pflag"
)

type Flags struct {
	IP      string
	Port    uint16
	GinMode string
}

func NewFlags() *Flags {
	return &Flags{
		IP:      "127.0.0.1",
		Port:    8080,
		GinMode: "release",
	}
}

func (f *Flags) AddFlag(flags *pflag.FlagSet) {
	if f.IP != "" {
		flags.StringVar(&f.IP, flagServerIP, f.IP, "The IP address to bind")
	}
	if f.Port != 0 {
		flags.Uint16VarP(&f.Port, flagServerPort, "p", f.Port, "The Port to bind")
	}
	if f.GinMode != "" {
		flags.StringVar(&f.GinMode, flagMode, f.GinMode, "The gin mode")
	}

}
