package crm

import (
	"github.com/morlay/goaliyun"
)

type FindCustomerInfoRequest struct {
	KpId     int64  `position:"Query" name:"KpId"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *FindCustomerInfoRequest) Invoke(client goaliyun.Client) (*FindCustomerInfoResponse, error) {
	resp := &FindCustomerInfoResponse{}
	err := client.Request("crm", "FindCustomerInfo", "2015-03-24", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type FindCustomerInfoResponse struct {
	Success       bool
	ResultCode    goaliyun.String
	ResultMessage goaliyun.String
	Data          FindCustomerInfoData
}

type FindCustomerInfoData struct {
	Website goaliyun.String
	Biz     goaliyun.String
}
