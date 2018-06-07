package ecs

import (
	"github.com/morlay/goaliyun"
)

type ReInitDiskRequest struct {
	ResourceOwnerId             int64  `position:"Query" name:"ResourceOwnerId"`
	Password                    string `position:"Query" name:"Password"`
	ResourceOwnerAccount        string `position:"Query" name:"ResourceOwnerAccount"`
	AutoStartInstance           string `position:"Query" name:"AutoStartInstance"`
	OwnerAccount                string `position:"Query" name:"OwnerAccount"`
	DiskId                      string `position:"Query" name:"DiskId"`
	SecurityEnhancementStrategy string `position:"Query" name:"SecurityEnhancementStrategy"`
	KeyPairName                 string `position:"Query" name:"KeyPairName"`
	OwnerId                     int64  `position:"Query" name:"OwnerId"`
	RegionId                    string `position:"Query" name:"RegionId"`
}

func (req *ReInitDiskRequest) Invoke(client goaliyun.Client) (*ReInitDiskResponse, error) {
	resp := &ReInitDiskResponse{}
	err := client.Request("ecs", "ReInitDisk", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ReInitDiskResponse struct {
	RequestId goaliyun.String
}
