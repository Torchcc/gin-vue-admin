package utils

import (
	"context"
	"fmt"
	"mime/multipart"
	"net/http"
	"net/url"
	"time"

	"gin-vue-admin/global"
	"github.com/tencentyun/cos-go-sdk-v5"
)

const CommonBucketUrl = "https://common-1302104842.cos.ap-guangzhou.myqcloud.com"

func NewCosClient(bucketUrl string) *cos.Client {
	u, _ := url.Parse(bucketUrl)
	b := &cos.BaseURL{BucketURL: u}
	// 1.永久密钥
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  global.GVA_CONFIG.Cos.SecretID,
			SecretKey: global.GVA_CONFIG.Cos.SecretKey,
		},
	})
	return client
}

// 接收两个参数 一个文件流 一个 bucket 你的七牛云标准空间的名字
func Upload(file *multipart.FileHeader) (err error, path string, key string) {
	cli := NewCosClient(CommonBucketUrl)
	f, err := file.Open()
	if err != nil {
		fmt.Printf("failed to open mulipart.FileHeader, err: [%s]", err.Error())
		return
	}

	fileKey := fmt.Sprintf("%d%s", time.Now().Unix(), file.Filename) // 文件名格式 自己可以改 建议保证唯一性
	_, err = cli.Object.Put(context.Background(), fileKey, f, nil)
	if err != nil {
		fmt.Printf("fail to upload to tx cos, err: [%s]", err.Error())
		return
	}
	return err, CommonBucketUrl + "/" + fileKey, fileKey
}

func DeleteFile(key string) error {
	return nil

	// mac := qbox.NewMac(global.GVA_CONFIG.Qiniu.AccessKey, global.GVA_CONFIG.Qiniu.SecretKey)
	// cfg := storage.Config{
	// 	// 是否使用https域名进行资源管理
	// 	UseHTTPS: false,
	// }
	// // 指定空间所在的区域，如果不指定将自动探测
	// // 如果没有特殊需求，默认不需要指定
	// // cfg.Zone=&storage.ZoneHuabei
	// bucketManager := storage.NewBucketManager(mac, &cfg)
	// err := bucketManager.Delete(global.GVA_CONFIG.Qiniu.Bucket, key)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return err
	// }
	// return nil
}
