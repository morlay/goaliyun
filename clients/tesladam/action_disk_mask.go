package tesladam

import (
	"github.com/morlay/goaliyun"
)

type ActionDiskMaskRequest struct {
	Op        string `position:"Query" name:"Op"`
	DiskMount string `position:"Query" name:"DiskMount"`
	Ip        string `position:"Query" name:"Ip"`
	RegionId  string `position:"Query" name:"RegionId"`
}

func (req *ActionDiskMaskRequest) Invoke(client goaliyun.Client) (*ActionDiskMaskResponse, error) {
	resp := &ActionDiskMaskResponse{}
	err := client.Request("tesladam", "ActionDiskMask", "2018-01-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ActionDiskMaskResponse struct {
	Status  bool
	Message goaliyun.String
	Result  goaliyun.String
}
