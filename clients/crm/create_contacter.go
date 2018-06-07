package crm

import (
	"github.com/morlay/goaliyun"
)

type CreateContacterRequest struct {
	KpId              int64  `position:"Query" name:"KpId"`
	ContacterName     string `position:"Query" name:"ContacterName"`
	ContacterEmail    string `position:"Query" name:"ContacterEmail"`
	ContacterMobile   string `position:"Query" name:"ContacterMobile"`
	ContacterPosition string `position:"Query" name:"ContacterPosition"`
	RegionId          string `position:"Query" name:"RegionId"`
}

func (req *CreateContacterRequest) Invoke(client goaliyun.Client) (*CreateContacterResponse, error) {
	resp := &CreateContacterResponse{}
	err := client.Request("crm", "CreateContacter", "2015-03-24", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CreateContacterResponse struct {
	Success       bool
	ResultCode    goaliyun.String
	ResultMessage goaliyun.String
	Data          CreateContacterData
}

type CreateContacterData struct {
	ContacterId goaliyun.Integer
}
