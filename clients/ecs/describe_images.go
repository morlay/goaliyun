package ecs

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeImagesRequest struct {
	Tag4Value            string                    `position:"Query" name:"Tag.4.Value"`
	ActionType           string                    `position:"Query" name:"ActionType"`
	ResourceOwnerId      int64                     `position:"Query" name:"ResourceOwnerId"`
	ImageId              string                    `position:"Query" name:"ImageId"`
	SnapshotId           string                    `position:"Query" name:"SnapshotId"`
	Tag2Key              string                    `position:"Query" name:"Tag.2.Key"`
	Usage                string                    `position:"Query" name:"Usage"`
	Tag3Key              string                    `position:"Query" name:"Tag.3.Key"`
	PageNumber           int64                     `position:"Query" name:"PageNumber"`
	ImageOwnerAlias      string                    `position:"Query" name:"ImageOwnerAlias"`
	Tag1Value            string                    `position:"Query" name:"Tag.1.Value"`
	IsSupportIoOptimized string                    `position:"Query" name:"IsSupportIoOptimized"`
	ImageName            string                    `position:"Query" name:"ImageName"`
	IsSupportCloudinit   string                    `position:"Query" name:"IsSupportCloudinit"`
	PageSize             int64                     `position:"Query" name:"PageSize"`
	InstanceType         string                    `position:"Query" name:"InstanceType"`
	Tag3Value            string                    `position:"Query" name:"Tag.3.Value"`
	Architecture         string                    `position:"Query" name:"Architecture"`
	DryRun               string                    `position:"Query" name:"DryRun"`
	Tag5Key              string                    `position:"Query" name:"Tag.5.Key"`
	ResourceOwnerAccount string                    `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string                    `position:"Query" name:"OwnerAccount"`
	ShowExpired          string                    `position:"Query" name:"ShowExpired"`
	OSType               string                    `position:"Query" name:"OSType"`
	OwnerId              int64                     `position:"Query" name:"OwnerId"`
	Tag5Value            string                    `position:"Query" name:"Tag.5.Value"`
	Tag1Key              string                    `position:"Query" name:"Tag.1.Key"`
	Filters              *DescribeImagesFilterList `position:"Query" type:"Repeated" name:"Filter"`
	Tag2Value            string                    `position:"Query" name:"Tag.2.Value"`
	Tag4Key              string                    `position:"Query" name:"Tag.4.Key"`
	Status               string                    `position:"Query" name:"Status"`
	RegionId             string                    `position:"Query" name:"RegionId"`
}

func (req *DescribeImagesRequest) Invoke(client goaliyun.Client) (*DescribeImagesResponse, error) {
	resp := &DescribeImagesResponse{}
	err := client.Request("ecs", "DescribeImages", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeImagesFilter struct {
	Key   string `name:"Key"`
	Value string `name:"Value"`
}

type DescribeImagesResponse struct {
	RequestId  goaliyun.String
	RegionId   goaliyun.String
	TotalCount goaliyun.Integer
	PageNumber goaliyun.Integer
	PageSize   goaliyun.Integer
	Images     DescribeImagesImageList
}

type DescribeImagesImage struct {
	Progress             goaliyun.String
	ImageId              goaliyun.String
	ImageName            goaliyun.String
	ImageVersion         goaliyun.String
	Description          goaliyun.String
	Size                 goaliyun.Integer
	ImageOwnerAlias      goaliyun.String
	IsSupportIoOptimized bool
	IsSupportCloudinit   bool
	OSName               goaliyun.String
	Architecture         goaliyun.String
	Status               goaliyun.String
	ProductCode          goaliyun.String
	IsSubscribed         bool
	CreationTime         goaliyun.String
	IsSelfShared         goaliyun.String
	OSType               goaliyun.String
	Platform             goaliyun.String
	Usage                goaliyun.String
	IsCopied             bool
	DiskDeviceMappings   DescribeImagesDiskDeviceMappingList
	Tags                 DescribeImagesTagList
}

type DescribeImagesDiskDeviceMapping struct {
	SnapshotId      goaliyun.String
	Size            goaliyun.String
	Device          goaliyun.String
	Type            goaliyun.String
	Format          goaliyun.String
	ImportOSSBucket goaliyun.String
	ImportOSSObject goaliyun.String
}

type DescribeImagesTag struct {
	TagKey   goaliyun.String
	TagValue goaliyun.String
}

type DescribeImagesFilterList []DescribeImagesFilter

func (list *DescribeImagesFilterList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeImagesFilter)
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

type DescribeImagesImageList []DescribeImagesImage

func (list *DescribeImagesImageList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeImagesImage)
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

type DescribeImagesDiskDeviceMappingList []DescribeImagesDiskDeviceMapping

func (list *DescribeImagesDiskDeviceMappingList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeImagesDiskDeviceMapping)
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

type DescribeImagesTagList []DescribeImagesTag

func (list *DescribeImagesTagList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeImagesTag)
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
