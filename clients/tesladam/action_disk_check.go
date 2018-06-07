package tesladam

import (
	"github.com/morlay/goaliyun"
)

type ActionDiskCheckRequest struct {
	DiskMount string `position:"Query" name:"DiskMount"`
	Ip        string `position:"Query" name:"Ip"`
	RegionId  string `position:"Query" name:"RegionId"`
}

func (req *ActionDiskCheckRequest) Invoke(client goaliyun.Client) (*ActionDiskCheckResponse, error) {
	resp := &ActionDiskCheckResponse{}
	err := client.Request("tesladam", "ActionDiskCheck", "2018-01-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ActionDiskCheckResponse struct {
	Status  bool
	Message goaliyun.String
	Result  goaliyun.String
}
