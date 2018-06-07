package crm

import (
	"github.com/morlay/goaliyun"
)

type ModifyCustomerInfoRequest struct {
	KpId     int64  `position:"Query" name:"KpId"`
	Website  string `position:"Query" name:"Website"`
	Biz      string `position:"Query" name:"Biz"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *ModifyCustomerInfoRequest) Invoke(client goaliyun.Client) (*ModifyCustomerInfoResponse, error) {
	resp := &ModifyCustomerInfoResponse{}
	err := client.Request("crm", "ModifyCustomerInfo", "2015-03-24", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifyCustomerInfoResponse struct {
	Success       bool
	ResultCode    goaliyun.String
	ResultMessage goaliyun.String
}
