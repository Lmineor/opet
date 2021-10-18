package flag

import (
	"github.com/spf13/pflag"
	"k8s.io/klog/v2"
	"opet/pkg/etcd"
	"opet/pkg/server"
	"strings"
)

type Flags struct {
	EtcdFlags
	ServerFlags
}

func WordSepNormalizeFunc(f *pflag.FlagSet, name string) pflag.NormalizedName {
	if strings.Contains(name, "_") {
		nname := strings.Replace(name, "_", "-", -1)
		klog.Warningf("Warning: %s is deprecated and will be removed in a future version, use %s instead.\n", name, nname)
		return pflag.NormalizedName(nname)
	}
	return pflag.NormalizedName(name)
}

func (f *Flags) PrintAllFlags(flags *pflag.FlagSet) {
	flags.VisitAll(func(flag *pflag.Flag) {
		klog.V(1).Infof("Flag: --%s=%q", flag.Name, flag.Value)
	})
}

func NewFlags() *Flags {
	return &Flags{
		EtcdFlags:   EtcdFlags{etcd.NewFlags()},
		ServerFlags: ServerFlags{server.NewFlags()},
	}
}

func (f *Flags) AddFlags(flags *pflag.FlagSet) {
	f.EtcdFlags.AddFlags(flags)
	f.ServerFlags.AddFlag(flags)
}

type EtcdFlags struct {
	*etcd.Flags
}

type ServerFlags struct {
	*server.Flags
}
