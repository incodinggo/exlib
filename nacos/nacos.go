package nacos

import (
	"context"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
)

type Client struct {
	cli config_client.IConfigClient
}

type Opt struct {
	IpAddr               string
	Port                 uint64
	Group                string
	TimeoutMs            uint64
	BeatInterval         int64
	NamespaceId          string
	AppName              string
	Endpoint             string
	RegionId             string
	AccessKey            string
	SecretKey            string
	OpenKMS              bool
	CacheDir             string
	UpdateThreadNum      int
	NotLoadCacheAtStart  bool
	UpdateCacheWhenEmpty bool
	Username             string
	Password             string
	LogDir               string
	LogLevel             string
	ContextPath          string
}

func Init(opt Opt) (error, *Client) {
	cliConfig := constant.ClientConfig{
		TimeoutMs:            opt.TimeoutMs,
		BeatInterval:         opt.BeatInterval,
		NamespaceId:          opt.NamespaceId,
		AppName:              opt.AppName,
		Endpoint:             opt.Endpoint,
		RegionId:             opt.RegionId,
		AccessKey:            opt.AccessKey,
		SecretKey:            opt.SecretKey,
		OpenKMS:              opt.OpenKMS,
		CacheDir:             opt.CacheDir,
		UpdateThreadNum:      opt.UpdateThreadNum,
		NotLoadCacheAtStart:  opt.NotLoadCacheAtStart,
		UpdateCacheWhenEmpty: opt.UpdateCacheWhenEmpty,
		Username:             opt.Username,
		Password:             opt.Password,
		LogDir:               opt.LogDir,
		LogLevel:             opt.LogLevel,
		ContextPath:          opt.ContextPath,
	}
	srvConfigs := []constant.ServerConfig{
		*constant.NewServerConfig(
			opt.IpAddr,
			opt.Port,
			constant.WithScheme("http"),
			constant.WithContextPath("/nacos"),
		),
	}
	cli, err := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": srvConfigs,
		"clientConfig":  cliConfig,
	})
	if err != nil {
		return err, nil
	}
	return nil, &Client{cli}
}

// Publish 发布配置
// content:"connectTimeoutInMills=3000"
func (c *Client) Publish(dataId, group, content string) (bool, error) {
	success, err := c.cli.PublishConfig(vo.ConfigParam{
		DataId:  dataId,
		Group:   group,
		Content: content,
	})
	if err != nil {
		return false, err
	}
	return success, err
}

type Result struct {
	config string
	toMap  map[string]interface{}
}

// Get 获取配置
func (c *Client) Get(dataId, group string) (*Result, error) {
	config, err := c.cli.GetConfig(vo.ConfigParam{
		DataId: dataId,
		Group:  group,
	})
	if err != nil {
		return nil, err
	}
	ret := &Result{config: config}
	ret.setVal()
	return ret, nil
}

// Delete 删除配置
func (c *Client) Delete(dataId, group string) (bool, error) {
	success, err := c.cli.DeleteConfig(vo.ConfigParam{
		DataId: dataId,
		Group:  group,
	})
	if err != nil {
		return false, err
	}
	return success, nil
}

func (c *Client) Onchange(ctx context.Context, dataId, group string, onChange func(namespace, group, dataId, data string)) {
	go func() {
		err := c.cli.ListenConfig(vo.ConfigParam{
			DataId:   dataId,
			Group:    group,
			OnChange: onChange,
		})
		if err != nil {
			return
		}
		select {
		case <-ctx.Done():
			return
		}
	}()
}

func (r *Result) setVal() {
	r.toMap = r.setMap()
}

func (r *Result) setMap() map[string]interface{} {
	var m map[string]interface{}
	err := jsoniter.Unmarshal([]byte(r.config), &m)
	if err != nil {
		fmt.Println("config not json type")
	}
	return m
}

func (r *Result) String(k string) string {
	if r.toMap == nil {
		return ""
	}
	v, ok := r.toMap[k]
	if !ok {
		return ""
	}
	return v.(string)
}

func (r *Result) Slice(k string) []interface{} {
	if r.toMap == nil {
		return nil
	}
	v, ok := r.toMap[k]
	if !ok {
		return nil
	}
	return v.([]interface{})
}

func (r *Result) Int64(k string) int64 {
	if r.toMap == nil {
		return 0
	}
	v, ok := r.toMap[k]
	if !ok {
		return 0
	}
	return int64(v.(float64))
}

func (r *Result) Int(k string) int {
	if r.toMap == nil {
		return 0
	}
	v, ok := r.toMap[k]
	if !ok {
		return 0
	}
	return int(v.(float64))
}

func (r *Result) Bool(k string) bool {
	if r.toMap == nil {
		return false
	}
	v, ok := r.toMap[k]
	if !ok {
		return false
	}
	return v.(bool)
}

func (r *Result) Float64(k string) float64 {
	if r.toMap == nil {
		return 0
	}
	v, ok := r.toMap[k]
	if !ok {
		return 0
	}
	return v.(float64)
}

func (r *Result) Float32(k string) float32 {
	if r.toMap == nil {
		return 0
	}
	v, ok := r.toMap[k]
	if !ok {
		return 0
	}
	return v.(float32)
}

func (r *Result) Json(k string, i interface{}) {
	if r.toMap == nil {
		return
	}
	v, ok := r.toMap[k]
	if !ok {
		return
	}
	err := jsoniter.Unmarshal([]byte(v.(string)), i)
	if err != nil {
		return
	}
}
