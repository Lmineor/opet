package etcd

type Flags struct {
	EndPoints   []string
	DialTimeout int
}

func NewFlags() *Flags {
	return &Flags{
		EndPoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5,
	}
}
