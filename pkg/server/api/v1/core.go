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
	putResp, err := global.EtClient.Put(ctx, kv.Key, kv.Value)
	if err != nil {
		klog.Error(err)
		ErrorResp(c)
	}
	Success(putResp, "success", c)
}

func GetKey(c *gin.Context) {
	resKv := make(map[string]string)
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	key := c.Query("key")
	resp, _ := global.EtClient.Get(ctx, key)
	if len(resp.Kvs) > 0 {
		resKv["key"] = string(resp.Kvs[0].Key)
		resKv["value"] = string(resp.Kvs[0].Value)
		Success(resKv, "aaa", c)
	} else {
		FailedResp(c)
	}
}

func QueryKeyWithPrefix(c *gin.Context) {
	Success("a", "mmm", c)
}

func DeleteKey(c *gin.Context) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	key := c.Query("key")
	resp, err := global.EtClient.Delete(ctx, key)
	if err != nil {
		FailedResp(c)
	}
	Success(resp, "success", c)
}
