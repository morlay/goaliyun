package ons

import (
	"github.com/morlay/goaliyun"
)

type OnsTopicCreateRequest struct {
	PreventCache int64  `position:"Query" name:"PreventCache"`
	Cluster      string `position:"Query" name:"Cluster"`
	QueueNum     int64  `position:"Query" name:"QueueNum"`
	OnsRegionId  string `position:"Query" name:"OnsRegionId"`
	OnsPlatform  string `position:"Query" name:"OnsPlatform"`
	AppName      string `position:"Query" name:"AppName"`
	Qps          int64  `position:"Query" name:"Qps"`
	Topic        string `position:"Query" name:"Topic"`
	Remark       string `position:"Query" name:"Remark"`
	Appkey       string `position:"Query" name:"Appkey"`
	Order        string `position:"Query" name:"Order"`
	Status       int64  `position:"Query" name:"Status"`
	RegionId     string `position:"Query" name:"RegionId"`
}

func (req *OnsTopicCreateRequest) Invoke(client goaliyun.Client) (*OnsTopicCreateResponse, error) {
	resp := &OnsTopicCreateResponse{}
	err := client.Request("ons", "OnsTopicCreate", "2017-09-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type OnsTopicCreateResponse struct {
	RequestId goaliyun.String
	HelpUrl   goaliyun.String
}
