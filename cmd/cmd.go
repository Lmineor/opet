package cmd

import (
	"github.com/spf13/cobra"
	"k8s.io/klog/v2"
	"opet/global"
	"opet/pkg/etcd"
	"opet/pkg/flag"
	"opet/pkg/server"
	"os"
)

func NewOpEtCmd(flags *flag.Flags) *cobra.Command {
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
	//goflag.Parse()
	return cmd
}

// Execute executes the root command.
func Execute(cmd *cobra.Command) error {
	return cmd.Execute()
}

func addFlags(cmd *cobra.Command, of *flag.Flags) {
	flags := cmd.PersistentFlags()
	flags.SetNormalizeFunc(flag.WordSepNormalizeFunc)
	flags.StringVarP(&of.Port, "port", "p", of.Port, "server's port")
	flags.StringVar(&of.IP, "ip", of.IP, "server's ip")
	flags.StringVar(&of.GinMode, "gin-mode", of.GinMode, "gi-mode")
	//flags.StringVar(&of.LogDir, "lod-dir", of.LogDir, "the dir of logs.")
	flags.StringSliceVar(&of.EndPoints, "endpoints", of.EndPoints, "the endpoint of etcd.")
	flags.IntVar(&of.DialTimeout, "dial-timeout", of.DialTimeout, "dial-timeout")
	//cmd.PersistentFlags().AddGoFlag()
	//cmd.AddCommand(addCmd)
	//cmd.AddCommand(initCmd)
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
