package main

import (
	"log"
	"time"
	"fmt"

	"github.com/micro/go-grpc"
	"github.com/micro/go-micro"
	go_config "github.com/micro/go-config"

	"github.com/noahzaozao/alisms_service/coinfig"
	"github.com/noahzaozao/alisms_service/cache"
	"context"
	"github.com/noahzaozao/alisms_service/proto/alisms"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"

)

type AliSMSService struct {
	Config config.SettingConfig
}

func (aliSmsService *AliSMSService) SMSVerficationCode(
	ctx context.Context, in *alisms.SMSVerficationCodeData, out *alisms.SMSVerficationResponseData) error {

	client, err := sdk.NewClientWithAccessKey(
		"default",
		aliSmsService.Config.SMSConfig.ACCESS_KEY_ID,
		aliSmsService.Config.SMSConfig.ACCESS_KEY_SECRET)
	if err != nil {
		panic(err)
	}

	request := requests.NewCommonRequest()
	request.Method = "POST"
	request.Scheme = "https"
	request.Domain = "dysmsapi.aliyuncs.com"
	request.Version = "2017-05-25"
	request.ApiName = "SendSms"

	request.QueryParams["SignName"] = in.SignName
	request.QueryParams["PhoneNumbers"] = in.PhoneNumbers
	request.QueryParams["TemplateCode"] = in.TemplateCode
	request.QueryParams["TemplateParam"] = in.TemplateParam
	request.QueryParams["SmsUpExtendCode"] = in.SmsUpExtendCode
	request.QueryParams["OutId"] = in.OutId

	response, err := client.ProcessCommonRequest(request)
	if err != nil {
		panic(err)
	}
	fmt.Print(response.GetHttpContentString())

	return nil
}

func (aliSmsService *AliSMSService) SMSVerficationCodeCheck(
	ctx context.Context, in *alisms.SMSVerficationCodeCheckData, out *alisms.SMSVerficationResponseData) error {

	return nil
}

func (aliSmsService *AliSMSService) SMSVerficationQuery(
	ctx context.Context, in *alisms.SMSVerficationQueryData, out *alisms.SMSVerficationQueryResponseData) error {

	client, err := sdk.NewClientWithAccessKey(
		"default",
		aliSmsService.Config.SMSConfig.ACCESS_KEY_ID,
		aliSmsService.Config.SMSConfig.ACCESS_KEY_SECRET)
	if err != nil {
		panic(err)
	}

	request := requests.NewCommonRequest()
	request.Method = "POST"
	request.Scheme = "https"
	request.Domain = "dysmsapi.aliyuncs.com"
	request.Version = "2017-05-25"
	request.ApiName = "QuerySendDetails"
	request.QueryParams["PhoneNumber"] = in.PhoneNumbers
	request.QueryParams["SendDate"] = in.SendDate
	request.QueryParams["PageSize"] = in.PageSize
	request.QueryParams["CurrentPage"] = in.CurrentPage
	request.QueryParams["BizId"] = in.BizId

	response, err := client.ProcessCommonRequest(request)
	if err != nil {
		panic(err)
	}
	fmt.Print(response.GetHttpContentString())

	return nil
}

func main() {

	// Load json config file
	if err := go_config.LoadFile("./config.yaml"); err != nil {
		log.Println(err.Error())
		return
	}

	var settingsConfig config.SettingConfig

	if err := go_config.Get("config").Scan(&settingsConfig); err != nil {
		log.Println(err.Error())
		return
	}

	log.Println("DEBUG: " + settingsConfig.DEBUG)
	log.Println("CHARSET: " + settingsConfig.DEFAULT_CHARSET)

	if len(settingsConfig.CACHES) < 1 {
		log.Println("CACHES config not exist")
		return
	}

	if err := cache.CacheMgr().Init(settingsConfig.CACHES[0]); err != nil {
		log.Println(err.Error())
		return
	}
	log.Println("Init CACHE...")

	service := grpc.NewService(
		micro.Name("alisms.srv"),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
	)
	service.Init()

	alismsService := &AliSMSService{
		Config: settingsConfig,
	}

	if err := alisms.RegisterAuthServiceHandler(service.Server(), alismsService); err != nil {
		log.Println(err.Error())
		return
	}

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
