package sms

import (
	dysmsapi20170525 "github.com/alibabacloud-go/dysmsapi-20170525/v4/client"
)

type AliSmsService struct {
	AccessKeyId string
}

func NewAliSmsService(_result *dysmsapi20170525.Client) *AliSmsService {
	return &AliSmsService{}
}

func Send() {

}
