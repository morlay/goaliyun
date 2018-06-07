package market_inner

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type InnerQueryAvailableImagesRequest struct {
	SupportIoOptimized string                                `position:"Query" name:"SupportIoOptimized"`
	SnapshotId         string                                `position:"Query" name:"SnapshotId"`
	ImageIds           *InnerQueryAvailableImagesImageIdList `position:"Query" type:"Repeated" name:"ImageId"`
	ImageName          string                                `position:"Query" name:"ImageName"`
	PageSize           int64                                 `position:"Query" name:"PageSize"`
	EcsUnits           *InnerQueryAvailableImagesEcsUnitList `position:"Query" type:"Repeated" name:"EcsUnit"`
	PageIndex          int64                                 `position:"Query" name:"PageIndex"`
	AliUid             int64                                 `position:"Query" name:"AliUid"`
	OsTypes            *InnerQueryAvailableImagesOsTypeList  `position:"Query" type:"Repeated" name:"OsType"`
	RegionNo           string                                `position:"Query" name:"RegionNo"`
	RegionId           string                                `position:"Query" name:"RegionId"`
}

func (req *InnerQueryAvailableImagesRequest) Invoke(client goaliyun.Client) (*InnerQueryAvailableImagesResponse, error) {
	resp := &InnerQueryAvailableImagesResponse{}
	err := client.Request("market-inner", "InnerQueryAvailableImages", "2016-08-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type InnerQueryAvailableImagesEcsUnit struct {
	Core   int64 `name:"Core"`
	Memory int64 `name:"Memory"`
}

type InnerQueryAvailableImagesOsType struct {
	OsKind string `name:"OsKind"`
	OsBit  int64  `name:"OsBit"`
}

type InnerQueryAvailableImagesResponse struct {
	PageNumber goaliyun.Integer
	PageSize   goaliyun.Integer
	TotalCount goaliyun.Integer
	RequestId  goaliyun.String
	ImageList  InnerQueryAvailableImagesImageList
}

type InnerQueryAvailableImagesImage struct {
	ImageId            goaliyun.String
	ImageSize          goaliyun.Integer
	SupplierId         goaliyun.Integer
	SnapshotId         goaliyun.String
	Device             goaliyun.String
	ProductCode        goaliyun.String
	ProductSkuCode     goaliyun.String
	ImageVersion       goaliyun.String
	RegionId           goaliyun.String
	ImageName          goaliyun.String
	SupplierName       goaliyun.String
	OsName             goaliyun.String
	Architecture       goaliyun.String
	Description        goaliyun.String
	ImageOwnerAlias    goaliyun.String
	IsSubscribed       bool
	GmtCreated         goaliyun.String
	SupportIoOptimized bool
	VmType             goaliyun.String
	DiskDeviceMappings InnerQueryAvailableImagesDiskDeviceMappingList
}

type InnerQueryAvailableImagesDiskDeviceMapping struct {
	DiskType        goaliyun.String
	Format          goaliyun.String
	SnapshotId      goaliyun.String
	Size            goaliyun.Integer
	Device          goaliyun.String
	ImportOSSBucket goaliyun.String
	ImportOSSObject goaliyun.String
}

type InnerQueryAvailableImagesImageIdList []string

func (list *InnerQueryAvailableImagesImageIdList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]string)
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

type InnerQueryAvailableImagesEcsUnitList []InnerQueryAvailableImagesEcsUnit

func (list *InnerQueryAvailableImagesEcsUnitList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]InnerQueryAvailableImagesEcsUnit)
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

type InnerQueryAvailableImagesOsTypeList []InnerQueryAvailableImagesOsType

func (list *InnerQueryAvailableImagesOsTypeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]InnerQueryAvailableImagesOsType)
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

type InnerQueryAvailableImagesImageList []InnerQueryAvailableImagesImage

func (list *InnerQueryAvailableImagesImageList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]InnerQueryAvailableImagesImage)
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

type InnerQueryAvailableImagesDiskDeviceMappingList []InnerQueryAvailableImagesDiskDeviceMapping

func (list *InnerQueryAvailableImagesDiskDeviceMappingList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]InnerQueryAvailableImagesDiskDeviceMapping)
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
