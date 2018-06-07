package sas_api

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type RenewInstanceRequest struct {
	Duration     int64  `position:"Query" name:"Duration"`
	InstanceId   string `position:"Query" name:"InstanceId"`
	ClientToken  string `position:"Query" name:"ClientToken"`
	BuyNumber    int64  `position:"Query" name:"BuyNumber"`
	OwnerId      int64  `position:"Query" name:"OwnerId"`
	PricingCycle string `position:"Query" name:"PricingCycle"`
	RegionId     string `position:"Query" name:"RegionId"`
}

func (req *RenewInstanceRequest) Invoke(client goaliyun.Client) (*RenewInstanceResponse, error) {
	resp := &RenewInstanceResponse{}
	err := client.Request("sas-api", "RenewInstance", "2017-07-05", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type RenewInstanceResponse struct {
	Code        goaliyun.String
	Message     goaliyun.String
	OrderId     goaliyun.String
	InstanceId  goaliyun.String
	RequestId   goaliyun.String
	InstanceIds RenewInstanceInstanceIdList
}

type RenewInstanceInstanceIdList []goaliyun.String

func (list *RenewInstanceInstanceIdList) UnmarshalJSON(data []byte) error {
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
