package ecs

import (
	"github.com/morlay/goaliyun"
)

type CopyImageRequest struct {
	Tag4Value              string `position:"Query" name:"Tag.4.Value"`
	ResourceOwnerId        int64  `position:"Query" name:"ResourceOwnerId"`
	ImageId                string `position:"Query" name:"ImageId"`
	Tag2Key                string `position:"Query" name:"Tag.2.Key"`
	Tag5Key                string `position:"Query" name:"Tag.5.Key"`
	ResourceOwnerAccount   string `position:"Query" name:"ResourceOwnerAccount"`
	DestinationImageName   string `position:"Query" name:"DestinationImageName"`
	DestinationRegionId    string `position:"Query" name:"DestinationRegionId"`
	OwnerAccount           string `position:"Query" name:"OwnerAccount"`
	Tag3Key                string `position:"Query" name:"Tag.3.Key"`
	OwnerId                int64  `position:"Query" name:"OwnerId"`
	Tag5Value              string `position:"Query" name:"Tag.5.Value"`
	Tag1Key                string `position:"Query" name:"Tag.1.Key"`
	Tag1Value              string `position:"Query" name:"Tag.1.Value"`
	Encrypted              string `position:"Query" name:"Encrypted"`
	Tag2Value              string `position:"Query" name:"Tag.2.Value"`
	Tag4Key                string `position:"Query" name:"Tag.4.Key"`
	DestinationDescription string `position:"Query" name:"DestinationDescription"`
	Tag3Value              string `position:"Query" name:"Tag.3.Value"`
	RegionId               string `position:"Query" name:"RegionId"`
}

func (req *CopyImageRequest) Invoke(client goaliyun.Client) (*CopyImageResponse, error) {
	resp := &CopyImageResponse{}
	err := client.Request("ecs", "CopyImage", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CopyImageResponse struct {
	RequestId goaliyun.String
	ImageId   goaliyun.String
}
