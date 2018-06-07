package yundun

import (
	"github.com/morlay/goaliyun"
)

type SetDdosQpsRequest struct {
	InstanceId   string `position:"Query" name:"InstanceId"`
	InstanceType string `position:"Query" name:"InstanceType"`
	QpsPosition  int64  `position:"Query" name:"QpsPosition"`
	Level        int64  `position:"Query" name:"Level"`
	RegionId     string `position:"Query" name:"RegionId"`
}

func (req *SetDdosQpsRequest) Invoke(client goaliyun.Client) (*SetDdosQpsResponse, error) {
	resp := &SetDdosQpsResponse{}
	err := client.Request("yundun", "SetDdosQps", "2015-04-16", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SetDdosQpsResponse struct {
	RequestId goaliyun.String
}
