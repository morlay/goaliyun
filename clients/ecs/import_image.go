package ecs

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ImportImageRequest struct {
	DiskDeviceMappings   *ImportImageDiskDeviceMappingList `position:"Query" type:"Repeated" name:"DiskDeviceMapping"`
	ResourceOwnerId      int64                             `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string                            `position:"Query" name:"ResourceOwnerAccount"`
	ImageName            string                            `position:"Query" name:"ImageName"`
	RoleName             string                            `position:"Query" name:"RoleName"`
	Description          string                            `position:"Query" name:"Description"`
	OSType               string                            `position:"Query" name:"OSType"`
	OwnerId              int64                             `position:"Query" name:"OwnerId"`
	Platform             string                            `position:"Query" name:"Platform"`
	Architecture         string                            `position:"Query" name:"Architecture"`
	RegionId             string                            `position:"Query" name:"RegionId"`
}

func (req *ImportImageRequest) Invoke(client goaliyun.Client) (*ImportImageResponse, error) {
	resp := &ImportImageResponse{}
	err := client.Request("ecs", "ImportImage", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ImportImageDiskDeviceMapping struct {
	Format        string `name:"Format"`
	OSSBucket     string `name:"OSSBucket"`
	OSSObject     string `name:"OSSObject"`
	DiskImSize    int64  `name:"DiskImSize"`
	DiskImageSize int64  `name:"DiskImageSize"`
	Device        string `name:"Device"`
}

type ImportImageResponse struct {
	RequestId goaliyun.String
	TaskId    goaliyun.String
	RegionId  goaliyun.String
	ImageId   goaliyun.String
}

type ImportImageDiskDeviceMappingList []ImportImageDiskDeviceMapping

func (list *ImportImageDiskDeviceMappingList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ImportImageDiskDeviceMapping)
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
