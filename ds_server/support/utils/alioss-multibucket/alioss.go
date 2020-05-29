package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/go-test/alioss-multibucket/global"
)

func HandleError(err error) {
	fmt.Println("occurred error:", err)
	os.Exit(-1)
}

func CreateNewBucket( endPoint,accessKeyID,accessKeySecret,bucketName string)(*oss.Bucket, error){
	client, err := oss.New(endPoint, accessKeyID, accessKeySecret)
	if err != nil {
		return nil, err
	}
	err = client.CreateBucket(bucketName)
	if err != nil {
		return nil, err
	}
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		return nil, err
	}

	return bucket, nil
}
func GetExistBucket(endPoint,accessKeyID,accessKeySecret,bucketName string)(*oss.Bucket, error)  {
	client, err := oss.New(endPoint, accessKeyID, accessKeySecret)
	if err != nil {
		return nil, err
	}
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		return nil, err
	}
	return bucket, nil
}
func PutObjectFromFileWithUrl(bucketName,objectKey,filePath string)bool  {
	bucket, err := GetExistBucket(global.GloConfig.Sqlconfig.AliossConfig.Endpoint,global.GloConfig.Sqlconfig.AliossConfig.AccessKeyID,global.GloConfig.Sqlconfig.AliossConfig.AccessKeySecret,bucketName)
	if err != nil {
		HandleError(err)
	}
	options := []oss.Option{
		oss.Meta("key-main", "key-value"),
		oss.ContentType("multipart/form-data"),
	}
	var signedURL string
	signedURL, err = bucket.SignURL(objectKey, oss.HTTPPut, 60, options...)
	if err != nil {
		HandleError(err)
	}

	err = bucket.PutObjectFromFileWithURL(signedURL, filePath, options...)
	if err != nil {
		HandleError(err)
	}
	return true
}
func GetObjectToFileWithUrl(bucketName,objectKey,filePath string)  {
	bucket, err := GetExistBucket(global.GloConfig.Sqlconfig.AliossConfig.Endpoint,global.GloConfig.Sqlconfig.AliossConfig.AccessKeyID,global.GloConfig.Sqlconfig.AliossConfig.AccessKeySecret,bucketName)
	if err != nil {
		HandleError(err)
	}
	// Get object
	signedURL, err := bucket.SignURL(objectKey, oss.HTTPGet, 60)
	if err != nil {
		HandleError(err)
	}

	body, err := bucket.GetObjectWithURL(signedURL)
	if err != nil {
		HandleError(err)
	}
	// Read content
	data, err := ioutil.ReadAll(body)
	body.Close()
	data = data // use data
	fmt.Println(signedURL)

	err = bucket.GetObjectToFileWithURL(signedURL, filePath)
	if err != nil {
		HandleError(err)
	}
}