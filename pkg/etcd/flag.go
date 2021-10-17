package etcd

type Flags struct {
	EndPoints []string
	CertFile  string
	KeyFile   string
	CaFile    string
}

func NewFlags() *Flags {
	return &Flags{
		EndPoints: []string{"192.168.65.4:2379"},
		CertFile:  "/run/config/pki/etcd/server.crt",
		KeyFile:   "/run/config/pki/etcd/server.key",
		CaFile:    "/run/config/pki/etcd/ca.crt",
	}
}
