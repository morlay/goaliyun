package crm

import (
	"github.com/morlay/goaliyun"
)

type FindContacterTestRequest struct {
	ContacterId int64  `position:"Query" name:"ContacterId"`
	RegionId    string `position:"Query" name:"RegionId"`
}

func (req *FindContacterTestRequest) Invoke(client goaliyun.Client) (*FindContacterTestResponse, error) {
	resp := &FindContacterTestResponse{}
	err := client.Request("crm", "FindContacterTest", "2015-03-24", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type FindContacterTestResponse struct {
	Success       bool
	ResultCode    goaliyun.String
	ResultMessage goaliyun.String
	Data          FindContacterTestData
}

type FindContacterTestData struct {
	ContacterId       goaliyun.Integer
	KpId              goaliyun.Integer
	CustomerId        goaliyun.Integer
	ContacterName     goaliyun.String
	ContacterEmail    goaliyun.String
	ContacterMobile   goaliyun.String
	ContacterPosition goaliyun.String
}
