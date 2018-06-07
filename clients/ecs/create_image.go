package ecs

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type CreateImageRequest struct {
	DiskDeviceMappings   *CreateImageDiskDeviceMappingList `position:"Query" type:"Repeated" name:"DiskDeviceMapping"`
	Tag4Value            string                            `position:"Query" name:"Tag.4.Value"`
	ResourceOwnerId      int64                             `position:"Query" name:"ResourceOwnerId"`
	SnapshotId           string                            `position:"Query" name:"SnapshotId"`
	Tag2Key              string                            `position:"Query" name:"Tag.2.Key"`
	ClientToken          string                            `position:"Query" name:"ClientToken"`
	Description          string                            `position:"Query" name:"Description"`
	Tag3Key              string                            `position:"Query" name:"Tag.3.Key"`
	Platform             string                            `position:"Query" name:"Platform"`
	Tag1Value            string                            `position:"Query" name:"Tag.1.Value"`
	ImageName            string                            `position:"Query" name:"ImageName"`
	Tag3Value            string                            `position:"Query" name:"Tag.3.Value"`
	Architecture         string                            `position:"Query" name:"Architecture"`
	Tag5Key              string                            `position:"Query" name:"Tag.5.Key"`
	ResourceOwnerAccount string                            `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string                            `position:"Query" name:"OwnerAccount"`
	OwnerId              int64                             `position:"Query" name:"OwnerId"`
	Tag5Value            string                            `position:"Query" name:"Tag.5.Value"`
	Tag1Key              string                            `position:"Query" name:"Tag.1.Key"`
	InstanceId           string                            `position:"Query" name:"InstanceId"`
	Tag2Value            string                            `position:"Query" name:"Tag.2.Value"`
	ImageVersion         string                            `position:"Query" name:"ImageVersion"`
	Tag4Key              string                            `position:"Query" name:"Tag.4.Key"`
	RegionId             string                            `position:"Query" name:"RegionId"`
}

func (req *CreateImageRequest) Invoke(client goaliyun.Client) (*CreateImageResponse, error) {
	resp := &CreateImageResponse{}
	err := client.Request("ecs", "CreateImage", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CreateImageDiskDeviceMapping struct {
	Size       int64  `name:"Size"`
	SnapshotId string `name:"SnapshotId"`
	Device     string `name:"Device"`
	DiskType   string `name:"DiskType"`
}

type CreateImageResponse struct {
	RequestId goaliyun.String
	ImageId   goaliyun.String
}

type CreateImageDiskDeviceMappingList []CreateImageDiskDeviceMapping

func (list *CreateImageDiskDeviceMappingList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]CreateImageDiskDeviceMapping)
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
