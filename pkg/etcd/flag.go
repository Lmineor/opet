package etcd

import "github.com/spf13/pflag"

type Flags struct {
	EndPoints []string
	CertFile  string
	KeyFile   string
	CAFile    string
}

func NewFlags() *Flags {
	return &Flags{
		EndPoints: []string{"192.168.65.4:2379"},
		CertFile:  "/run/config/pki/etcd/server.crt",
		KeyFile:   "/run/config/pki/etcd/server.key",
		CAFile:    "/run/config/pki/etcd/ca.crt",
	}
}

func (f *Flags) AddFlags(flags *pflag.FlagSet) {
	if f.CAFile != "" {
		flags.StringVar(&f.CAFile, flagCAFile, f.CAFile, "Path to a cert file for the certificate authority")
	}

	if f.EndPoints != nil {
		flags.StringSliceVar(&f.EndPoints, flagEndPoints, f.EndPoints, "The endpoint of etcd")
	}
	if f.KeyFile != "" {
		flags.StringVar(&f.KeyFile, flagKeyFile, f.KeyFile, "Path to a client key file for TLS")
	}
	if f.CertFile != "" {
		flags.StringVar(&f.CertFile, flagCertFile, f.CertFile, "Path to a client certificate file for TLS")
	}
}
