package cmd

import (
	goflag "flag"
	"github.com/spf13/cobra"
	"k8s.io/klog/v2"
	"opet/global"
	"opet/pkg/etcd"
	"opet/pkg/flag"
	"opet/pkg/server"
	"os"
)

func NewOpEtCmd() *cobra.Command {
	flags := flag.NewFlags()
	cmd := &cobra.Command{
		Use:   "opet",
		Short: "A server to operate etcd.",
		Long:  `A easy test server to operate etcd which is a k, v based database.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			flags.PrintAllFlags(cmd.Flags())

			svr := initServer(flags)
			etcdClient := initEtcdClient(flags)
			opet := &OpEt{
				Server:     svr,
				EtcdClient: etcdClient,
				Flags:      flags,
			}
			if err := Run(opet); err != nil {
				klog.Fatal(err)
			}
			return nil
		},
	}
	addFlags(cmd, flags)
	validateFlags(flags)
	return cmd
}

// Execute executes the root command.
func Execute(cmd *cobra.Command) error {
	return cmd.Execute()
}

func addFlags(cmd *cobra.Command, of *flag.Flags) {
	flags := cmd.PersistentFlags()
	flags.SetNormalizeFunc(flag.WordSepNormalizeFunc)
	of.AddFlags(flags)
	flags.AddGoFlagSet(goflag.CommandLine)
}

func validateFlags(flags *flag.Flags) {
	if len(flags.EndPoints) == 0 {
		klog.Fatal("Endpoint cannot be empty.")
		os.Exit(1)
	}
}

type OpEt struct {
	Server     *server.Server
	EtcdClient *etcd.Client
	Flags      *flag.Flags
}

func initEtcdClient(flags *flag.Flags) *etcd.Client {
	klog.Info("initializing Etcd Client...")
	c, err := etcd.NewClient(flags.EtcdFlags.Flags)
	if err != nil {
		klog.Error("init Etcd Client Failed")
		os.Exit(1)
	}
	global.EtClient = c
	return c
}

func initServer(flags *flag.Flags) *server.Server {
	klog.Info("initializing Server...")
	return server.NewServer(flags.ServerFlags.Flags)
}

func Run(opet *OpEt) error {
	err := opet.Server.Run(opet.Flags.ServerFlags.Flags)
	return err
}
