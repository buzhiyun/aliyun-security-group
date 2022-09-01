package aliyun

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/auth/credentials"
	ecs "github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"github.com/buzhiyun/go-utils/cfg"
	"github.com/kataras/golog"
)

var aliyunClient *ecs.Client

func init() {
	key ,ok := cfg.Config().GetString("aliyun.key")
	if !ok {
		golog.Fatal("获取阿里云配置失败 aliyun.key")
	}
	secret , ok := cfg.Config().GetString("aliyun.secret")
	if !ok {
		golog.Fatal("获取阿里云配置失败 aliyun.secret")
	}

	config := sdk.NewConfig()
	// 是否开启重试机制
	config.WithAutoRetry(true);
	// 最大重试次数
	config.WithMaxRetryTime(3);

	credential := credentials.NewAccessKeyCredential(key, secret)
	/* use STS Token
	credential := credentials.NewStsTokenCredential("<your-access-key-id>", "<your-access-key-secret>", "<your-sts-token>")
	*/
	client, err := ecs.NewClientWithOptions("cn-hangzhou", config, credential)
	if err != nil {
		golog.Fatalf("初始化阿里云客户端异常 , %s",err.Error())
	}

	aliyunClient = client

}