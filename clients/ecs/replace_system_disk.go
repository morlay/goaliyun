package ecs

import (
	"github.com/morlay/goaliyun"
)

type ReplaceSystemDiskRequest struct {
	ResourceOwnerId             int64  `position:"Query" name:"ResourceOwnerId"`
	ImageId                     string `position:"Query" name:"ImageId"`
	ResourceOwnerAccount        string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken                 string `position:"Query" name:"ClientToken"`
	OwnerAccount                string `position:"Query" name:"OwnerAccount"`
	SecurityEnhancementStrategy string `position:"Query" name:"SecurityEnhancementStrategy"`
	KeyPairName                 string `position:"Query" name:"KeyPairName"`
	OwnerId                     int64  `position:"Query" name:"OwnerId"`
	Platform                    string `position:"Query" name:"Platform"`
	Password                    string `position:"Query" name:"Password"`
	InstanceId                  string `position:"Query" name:"InstanceId"`
	SystemDiskSize              int64  `position:"Query" name:"SystemDiskSize"`
	DiskId                      string `position:"Query" name:"DiskId"`
	UseAdditionalService        string `position:"Query" name:"UseAdditionalService"`
	Architecture                string `position:"Query" name:"Architecture"`
	RegionId                    string `position:"Query" name:"RegionId"`
}

func (req *ReplaceSystemDiskRequest) Invoke(client goaliyun.Client) (*ReplaceSystemDiskResponse, error) {
	resp := &ReplaceSystemDiskResponse{}
	err := client.Request("ecs", "ReplaceSystemDisk", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ReplaceSystemDiskResponse struct {
	RequestId goaliyun.String
	DiskId    goaliyun.String
}
