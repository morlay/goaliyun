package tesladam

import (
	"github.com/morlay/goaliyun"
)

type ActionDiskRmaRequest struct {
	DiskName    string `position:"Query" name:"DiskName"`
	ExecutionId string `position:"Query" name:"ExecutionId"`
	DiskSlot    string `position:"Query" name:"DiskSlot"`
	Hostname    string `position:"Query" name:"Hostname"`
	DiskMount   string `position:"Query" name:"DiskMount"`
	DiskReason  string `position:"Query" name:"DiskReason"`
	DiskSn      string `position:"Query" name:"DiskSn"`
	RegionId    string `position:"Query" name:"RegionId"`
}

func (req *ActionDiskRmaRequest) Invoke(client goaliyun.Client) (*ActionDiskRmaResponse, error) {
	resp := &ActionDiskRmaResponse{}
	err := client.Request("tesladam", "ActionDiskRma", "2018-01-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ActionDiskRmaResponse struct {
	Status  bool
	Message goaliyun.String
	Result  goaliyun.String
}
