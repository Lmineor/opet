package etcd

import (
	"github.com/coreos/etcd/clientv3"
	"time"
	"github.com/coreos/etcd/pkg/transport"
)

type Client struct {
	*clientv3.Client
}

func NewClient(flag *Flags) (*Client, error) {
	tlsInfo := transport.TLSInfo{
		CertFile:      flag.CertFile,
		KeyFile:       flag.KeyFile,
		TrustedCAFile: flag.CAFile,
	}
	tlfConfig, err := tlsInfo.ClientConfig()
	if err != nil {
		return nil, err
	}
	if len(flag.CertFile) == 0 && len(flag.KeyFile) == 0 && len(flag.CAFile) == 0 {
		tlfConfig = nil
	}
	config := &clientv3.Config{
		DialTimeout: 5 * time.Second,
		Endpoints:   flag.EndPoints,
		TLS:         tlfConfig,
	}
	c, err := clientv3.New(*config)
	if err != nil {
		return nil, err
	}
	return &Client{Client: c}, nil
}
