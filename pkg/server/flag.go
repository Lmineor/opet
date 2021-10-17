package server

import (
	"github.com/spf13/pflag"
)

type Flags struct {
	IP      string
	Port    string
	GinMode string
}

func NewFlags() *Flags {
	return &Flags{
		IP:      "127.0.0.1",
		Port:    "8080",
		GinMode: "release",
	}
}

func (f *Flags) AddFlag(fs *pflag.FlagSet) {
	fs.StringVar(&f.IP, "server_ip", f.IP, "set server ip")
	fs.StringVar(&f.Port, "server_port", f.Port, "set server port")
	fs.StringVar(&f.GinMode, "mode", f.GinMode, "set the gin mode")
}
