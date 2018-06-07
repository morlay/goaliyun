package cloudwf

import (
	"github.com/morlay/goaliyun"
)

type SaveApConfigRequest struct {
	Country     string `position:"Query" name:"Country"`
	ApAssetId   int64  `position:"Query" name:"ApAssetId"`
	LogLevel    int64  `position:"Query" name:"LogLevel"`
	Name        string `position:"Query" name:"Name"`
	EchoInt     int64  `position:"Query" name:"EchoInt"`
	Scan        int64  `position:"Query" name:"Scan"`
	Description string `position:"Query" name:"Description"`
	Id          int64  `position:"Query" name:"Id"`
	Dai         string `position:"Query" name:"Dai"`
	LogIp       string `position:"Query" name:"LogIp"`
	Mac         string `position:"Query" name:"Mac"`
	RegionId    string `position:"Query" name:"RegionId"`
}

func (req *SaveApConfigRequest) Invoke(client goaliyun.Client) (*SaveApConfigResponse, error) {
	resp := &SaveApConfigResponse{}
	err := client.Request("cloudwf", "SaveApConfig", "2017-03-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SaveApConfigResponse struct {
	RequestId goaliyun.String
	Success   bool
	Message   goaliyun.String
	Data      goaliyun.String
	ErrorCode goaliyun.Integer
	ErrorMsg  goaliyun.String
}
