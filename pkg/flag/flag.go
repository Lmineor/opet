package flag

import (
	goflag "flag"
	"github.com/golang/glog"
	"github.com/spf13/pflag"
	"opet/pkg/etcd"
	"opet/pkg/server"
	"strings"
)

type Flags struct {
	LogDirector string
	EtcdFlags
	ServerFlags
}

func wordSepNormalizeFunc(f *pflag.FlagSet, name string) pflag.NormalizedName {
	if strings.Contains(name, "_") {
		nname := strings.Replace(name, "_", "-", -1)
		glog.Warningf("%s is deprecated and will be removed in a future version, use %s instead.", name, nname)
		return pflag.NormalizedName(nname)
	}
	return pflag.NormalizedName(name)
}

// InitFlags normalizes, parses, then logs the command line flags
func InitFlags() {
	pflag.CommandLine.SetNormalizeFunc(wordSepNormalizeFunc)
	pflag.CommandLine.AddGoFlagSet(goflag.CommandLine) // to support flags defined using Go's flag package
}

func (f *Flags)PrintAllFlags(flags *pflag.FlagSet){
	flags.VisitAll(func(flag *pflag.Flag) {
		glog.V(1).Infof("Flag: --%s=%q", flag.Name, flag.Value)
	})

}

func NewFlags()*Flags{
	return &Flags{
		EtcdFlags: EtcdFlags{etcd.NewFlags()},
		ServerFlags: ServerFlags{server.NewFlags()},
		LogDirector: "/Users/lex/code/coding/opet/log.log",
	}
}

type EtcdFlags struct {
	*etcd.Flags
}

type ServerFlags struct {
	*server.Flags
}

func (f *Flags)AddFlag(fs *pflag.FlagSet){
	f.EtcdFlags.AddFlag(fs)
	f.ServerFlags.AddFlag(fs)
}
