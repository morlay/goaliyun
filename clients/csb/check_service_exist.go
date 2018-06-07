package csb

import (
	"github.com/morlay/goaliyun"
)

type CheckServiceExistRequest struct {
	CsbId       int64  `position:"Query" name:"CsbId"`
	ServiceName string `position:"Query" name:"ServiceName"`
	RegionId    string `position:"Query" name:"RegionId"`
}

func (req *CheckServiceExistRequest) Invoke(client goaliyun.Client) (*CheckServiceExistResponse, error) {
	resp := &CheckServiceExistResponse{}
	err := client.Request("csb", "CheckServiceExist", "2017-11-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CheckServiceExistResponse struct {
	Code      goaliyun.Integer
	Message   goaliyun.String
	RequestId goaliyun.String
	Data      CheckServiceExistData
}

type CheckServiceExistData struct {
	Exist bool
}
