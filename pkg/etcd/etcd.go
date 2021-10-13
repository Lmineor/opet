package etcd

import (
	"github.com/coreos/etcd/clientv3"
	"time"
)

type Client struct{
	*clientv3.Client
}


func NewClient(flag *Flags)(*Client, error){
	config := &clientv3.Config{
		DialTimeout: time.Duration(flag.DialTimeout) * time.Second,
		Endpoints: flag.EndPoints,
	}
	c, err := clientv3.New(*config)
	if err != nil {
		return nil, err
	}
	return &Client{Client: c}, nil
}
