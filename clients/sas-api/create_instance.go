package sas_api

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type CreateInstanceRequest struct {
	Duration          int64  `position:"Query" name:"Duration"`
	IsAutoRenew       string `position:"Query" name:"IsAutoRenew"`
	ClientToken       string `position:"Query" name:"ClientToken"`
	BuyNumber         int64  `position:"Query" name:"BuyNumber"`
	OwnerId           int64  `position:"Query" name:"OwnerId"`
	VersionCode       string `position:"Query" name:"VersionCode"`
	PricingCycle      string `position:"Query" name:"PricingCycle"`
	AutoRenewDuration int64  `position:"Query" name:"AutoRenewDuration"`
	RegionId          string `position:"Query" name:"RegionId"`
}

func (req *CreateInstanceRequest) Invoke(client goaliyun.Client) (*CreateInstanceResponse, error) {
	resp := &CreateInstanceResponse{}
	err := client.Request("sas-api", "CreateInstance", "2017-07-05", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CreateInstanceResponse struct {
	Code        goaliyun.String
	Message     goaliyun.String
	OrderId     goaliyun.String
	InstanceId  goaliyun.String
	RequestId   goaliyun.String
	InstanceIds CreateInstanceInstanceIdList
}

type CreateInstanceInstanceIdList []goaliyun.String

func (list *CreateInstanceInstanceIdList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]goaliyun.String)
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
