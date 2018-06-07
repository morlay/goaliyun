package alidns

import (
	"github.com/morlay/goaliyun"
)

type UpdateDNSSLBWeightRequest struct {
	Lang         string `position:"Query" name:"Lang"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	RecordId     string `position:"Query" name:"RecordId"`
	Weight       int64  `position:"Query" name:"Weight"`
	RegionId     string `position:"Query" name:"RegionId"`
}

func (req *UpdateDNSSLBWeightRequest) Invoke(client goaliyun.Client) (*UpdateDNSSLBWeightResponse, error) {
	resp := &UpdateDNSSLBWeightResponse{}
	err := client.Request("alidns", "UpdateDNSSLBWeight", "2015-01-09", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type UpdateDNSSLBWeightResponse struct {
	RequestId goaliyun.String
	RecordId  goaliyun.String
	Weight    goaliyun.Integer
}
