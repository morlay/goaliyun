package sas_api

import (
	"github.com/morlay/goaliyun"
)

type RefundInstanceRequest struct {
	InstanceId string `position:"Query" name:"InstanceId"`
	OwnerId    int64  `position:"Query" name:"OwnerId"`
	RegionId   string `position:"Query" name:"RegionId"`
}

func (req *RefundInstanceRequest) Invoke(client goaliyun.Client) (*RefundInstanceResponse, error) {
	resp := &RefundInstanceResponse{}
	err := client.Request("sas-api", "RefundInstance", "2017-07-05", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type RefundInstanceResponse struct {
	Code      goaliyun.String
	Message   goaliyun.String
	RequestId goaliyun.String
}
