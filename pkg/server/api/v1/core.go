package api

import (
	"context"
	"github.com/gin-gonic/gin"
	"k8s.io/klog/v2"
	"opet/global"
	"time"
)

func PutKey(c *gin.Context) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	klog.Info("Put key")
	var kv KVBody
	c.ShouldBindJSON(&kv)
	global.EtClient.Put(ctx, kv.Key, kv.Value)
	Success("aaa", "aaa", c)
}

func GetKey(c *gin.Context) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	key := c.Query("key")
	resp, _ := global.EtClient.Get(ctx, key)
	klog.Info(resp)
	Success("aaa", "aaa", c)
}

func QueryKeyWithPrefix(c *gin.Context) {
	Success("a", "mmm", c)
}
