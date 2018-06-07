package crm

import (
	"github.com/morlay/goaliyun"
)

type DeleteContacterRequest struct {
	ContacterId int64  `position:"Query" name:"ContacterId"`
	RegionId    string `position:"Query" name:"RegionId"`
}

func (req *DeleteContacterRequest) Invoke(client goaliyun.Client) (*DeleteContacterResponse, error) {
	resp := &DeleteContacterResponse{}
	err := client.Request("crm", "DeleteContacter", "2015-03-24", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteContacterResponse struct {
	Success       bool
	ResultCode    goaliyun.String
	ResultMessage goaliyun.String
}
