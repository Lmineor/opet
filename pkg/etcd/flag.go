package etcd

import "github.com/spf13/pflag"

type Flags struct {
	EndPoints []string
	DialTimeout int
}

func NewFlags()*Flags{
	return &Flags{
		EndPoints: []string{"127.0.0.1:2379"},
		DialTimeout: 5,
	}
}


func (f *Flags)AddFlag(fs *pflag.FlagSet){
	fs.StringSliceVar(&f.EndPoints, "endpoint", f.EndPoints, "etcd endpoint")
	fs.IntVar(&f.DialTimeout, "dial_timeout", f.DialTimeout, "dial timeout")
}
