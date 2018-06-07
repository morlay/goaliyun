package crm

import (
	"github.com/morlay/goaliyun"
)

type FindContacterRequest struct {
	ContacterId int64  `position:"Query" name:"ContacterId"`
	RegionId    string `position:"Query" name:"RegionId"`
}

func (req *FindContacterRequest) Invoke(client goaliyun.Client) (*FindContacterResponse, error) {
	resp := &FindContacterResponse{}
	err := client.Request("crm", "FindContacter", "2015-03-24", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type FindContacterResponse struct {
	Success       bool
	ResultCode    goaliyun.String
	ResultMessage goaliyun.String
	Data          FindContacterData
}

type FindContacterData struct {
	ContacterId       goaliyun.Integer
	KpId              goaliyun.Integer
	CustomerId        goaliyun.Integer
	ContacterName     goaliyun.String
	ContacterEmail    goaliyun.String
	ContacterMobile   goaliyun.String
	ContacterPosition goaliyun.String
}
