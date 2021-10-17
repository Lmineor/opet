package api

import (
	"context"
	"github.com/gin-gonic/gin"
	"k8s.io/klog/v2"
	"opet/global"
	"opet/pkg/server/common"
	"time"
)

func PutKey(c *gin.Context) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	klog.Info("Put key")
	var kv common.KVBody
	c.ShouldBindJSON(&kv)
	global.EtClient.Put(ctx, kv.Key, kv.Value)
	common.Success("aaa", "aaa", c)
}

func GetKey(c *gin.Context) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	var kv common.KVBody
	c.ShouldBindJSON(&kv)
	global.EtClient.Put(ctx, kv.Key, kv.Value)
	common.Success("aaa", "aaa", c)
}

func QueryKeyWithPrefix(c *gin.Context) {
	common.Success("a", "mmm", c)
}
