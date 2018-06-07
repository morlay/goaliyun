package emr

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ResizeDiskRequest struct {
	ResourceOwnerId int64                     `position:"Query" name:"ResourceOwnerId"`
	ClusterId       string                    `position:"Query" name:"ClusterId"`
	DiskConfigs     *ResizeDiskDiskConfigList `position:"Query" type:"Repeated" name:"DiskConfig"`
	RegionId        string                    `position:"Query" name:"RegionId"`
}

func (req *ResizeDiskRequest) Invoke(client goaliyun.Client) (*ResizeDiskResponse, error) {
	resp := &ResizeDiskResponse{}
	err := client.Request("emr", "ResizeDisk", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ResizeDiskDiskConfig struct {
	HostGroupId string `name:"HostGroupId"`
	NewDiskSize int64  `name:"NewDiskSize"`
}

type ResizeDiskResponse struct {
	RequestId goaliyun.String
	ClusterId goaliyun.String
}

type ResizeDiskDiskConfigList []ResizeDiskDiskConfig

func (list *ResizeDiskDiskConfigList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ResizeDiskDiskConfig)
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	for _, v := range m {
		*list = v
		break
	}
	return nil
}
