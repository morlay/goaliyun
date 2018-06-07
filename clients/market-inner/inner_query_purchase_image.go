package market_inner

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type InnerQueryPurchaseImageRequest struct {
	ImageId           string `position:"Query" name:"ImageId"`
	ImageProductCode  string `position:"Query" name:"ImageProductCode"`
	AliUid            int64  `position:"Query" name:"AliUid"`
	ChargeType        string `position:"Query" name:"ChargeType"`
	ImagePurchaseType string `position:"Query" name:"ImagePurchaseType"`
	ImageCategory     string `position:"Query" name:"ImageCategory"`
	RegionNo          string `position:"Query" name:"RegionNo"`
	RegionId          string `position:"Query" name:"RegionId"`
}

func (req *InnerQueryPurchaseImageRequest) Invoke(client goaliyun.Client) (*InnerQueryPurchaseImageResponse, error) {
	resp := &InnerQueryPurchaseImageResponse{}
	err := client.Request("market-inner", "InnerQueryPurchaseImage", "2016-08-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type InnerQueryPurchaseImageResponse struct {
	Image InnerQueryPurchaseImageImage
}

type InnerQueryPurchaseImageImage struct {
	ProductCode        goaliyun.String
	ProductSKUCode     goaliyun.String
	ImageId            goaliyun.String
	ImageVersion       goaliyun.String
	RegionId           goaliyun.String
	SupplierId         goaliyun.Integer
	SupplierName       goaliyun.String
	DiskDeviceMappings InnerQueryPurchaseImageDiskDeviceMappingList
}

type InnerQueryPurchaseImageDiskDeviceMapping struct {
	DiskType        goaliyun.String
	Format          goaliyun.String
	SnapshotId      goaliyun.String
	Size            goaliyun.Integer
	Device          goaliyun.String
	ImportOSSBucket goaliyun.String
	ImportOSSObject goaliyun.String
}

type InnerQueryPurchaseImageDiskDeviceMappingList []InnerQueryPurchaseImageDiskDeviceMapping

func (list *InnerQueryPurchaseImageDiskDeviceMappingList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]InnerQueryPurchaseImageDiskDeviceMapping)
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
