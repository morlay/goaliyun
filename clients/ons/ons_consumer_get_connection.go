package ons

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type OnsConsumerGetConnectionRequest struct {
	PreventCache int64  `position:"Query" name:"PreventCache"`
	OnsRegionId  string `position:"Query" name:"OnsRegionId"`
	OnsPlatform  string `position:"Query" name:"OnsPlatform"`
	ConsumerId   string `position:"Query" name:"ConsumerId"`
	RegionId     string `position:"Query" name:"RegionId"`
}

func (req *OnsConsumerGetConnectionRequest) Invoke(client goaliyun.Client) (*OnsConsumerGetConnectionResponse, error) {
	resp := &OnsConsumerGetConnectionResponse{}
	err := client.Request("ons", "OnsConsumerGetConnection", "2017-09-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type OnsConsumerGetConnectionResponse struct {
	RequestId goaliyun.String
	HelpUrl   goaliyun.String
	Data      OnsConsumerGetConnectionData
}

type OnsConsumerGetConnectionData struct {
	ConnectionList OnsConsumerGetConnectionConnectionDoList
}

type OnsConsumerGetConnectionConnectionDo struct {
	ClientId   goaliyun.String
	ClientAddr goaliyun.String
	Language   goaliyun.String
	Version    goaliyun.String
}

type OnsConsumerGetConnectionConnectionDoList []OnsConsumerGetConnectionConnectionDo

func (list *OnsConsumerGetConnectionConnectionDoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]OnsConsumerGetConnectionConnectionDo)
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	for _, v := range m {
		*list = v
		break
	}
	return nil
}
