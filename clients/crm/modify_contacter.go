package crm

import (
	"github.com/morlay/goaliyun"
)

type ModifyContacterRequest struct {
	ContacterId       int64  `position:"Query" name:"ContacterId"`
	ContacterName     string `position:"Query" name:"ContacterName"`
	ContacterEmail    string `position:"Query" name:"ContacterEmail"`
	ContacterMobile   string `position:"Query" name:"ContacterMobile"`
	ContacterPosition string `position:"Query" name:"ContacterPosition"`
	RegionId          string `position:"Query" name:"RegionId"`
}

func (req *ModifyContacterRequest) Invoke(client goaliyun.Client) (*ModifyContacterResponse, error) {
	resp := &ModifyContacterResponse{}
	err := client.Request("crm", "ModifyContacter", "2015-03-24", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifyContacterResponse struct {
	Success       bool
	ResultCode    goaliyun.String
	ResultMessage goaliyun.String
}
