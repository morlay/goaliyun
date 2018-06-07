package cloudwf

import (
	"github.com/morlay/goaliyun"
)

type SaveApgroupConfigRequest struct {
	Country     string `position:"Query" name:"Country"`
	LogLevel    int64  `position:"Query" name:"LogLevel"`
	Name        string `position:"Query" name:"Name"`
	EchoInt     int64  `position:"Query" name:"EchoInt"`
	Scan        int64  `position:"Query" name:"Scan"`
	Description string `position:"Query" name:"Description"`
	Id          int64  `position:"Query" name:"Id"`
	Dai         string `position:"Query" name:"Dai"`
	LogIp       string `position:"Query" name:"LogIp"`
	RegionId    string `position:"Query" name:"RegionId"`
}

func (req *SaveApgroupConfigRequest) Invoke(client goaliyun.Client) (*SaveApgroupConfigResponse, error) {
	resp := &SaveApgroupConfigResponse{}
	err := client.Request("cloudwf", "SaveApgroupConfig", "2017-03-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SaveApgroupConfigResponse struct {
	RequestId goaliyun.String
	Success   bool
	Message   goaliyun.String
	Data      goaliyun.String
	ErrorCode goaliyun.Integer
	ErrorMsg  goaliyun.String
}
