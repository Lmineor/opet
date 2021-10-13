package cmd

import (
	"github.com/golang/glog"
	"github.com/spf13/cobra"
	"opet/global"
	"opet/pkg/etcd"
	"opet/pkg/flag"
	"opet/pkg/server"
	"os"
)

func NewOpEtCmd() *cobra.Command{
	opetFlags := flag.NewFlags()
	cmd := &cobra.Command{
		Use:   "opet",
		Short: "A server to operate etcd.",
		Long:  `A easy test server to operate etcd which is a k, v based database.`,
		RunE: func(cmd *cobra.Command, args []string) error{

			opetFlags.PrintAllFlags(cmd.Flags())
			svr := initServer(opetFlags)
			etcdClient := initEtcdClient(opetFlags)
			opet := &OpEt{
				Server:     svr,
				EtcdClient: etcdClient,
				Flags:      opetFlags,
			}
			if err := Run(opet);err != nil{
				glog.Fatal(err)
			}
			return nil
		},
	}
	opetFlags.AddFlag(cmd.Flags())
	//pflag.Parse()
	return cmd
}


type OpEt struct {
	Server *server.Server
	EtcdClient *etcd.Client
	Flags *flag.Flags
}


func initEtcdClient(flags *flag.Flags) *etcd.Client{
	glog.Info("initializing Etcd Client...")
	c, err := etcd.NewClient(flags.EtcdFlags.Flags)
	if err != nil{
		glog.Error("Init Etcd client Failed")
		os.Exit(1)
	}
	global.EtClient = c
	return c
}

func initServer(flags *flag.Flags) *server.Server{
	glog.Info("initializing Server...")
	return server.NewServer(flags.ServerFlags.Flags)
}

func Run(opet *OpEt)error{
	err := opet.Server.Run(opet.Flags.ServerFlags.Flags)
	return err
}