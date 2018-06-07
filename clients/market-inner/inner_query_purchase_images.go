package market_inner

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type InnerQueryPurchaseImagesRequest struct {
	Param    string `position:"Query" name:"Param"`
	AliUid   int64  `position:"Query" name:"AliUid"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *InnerQueryPurchaseImagesRequest) Invoke(client goaliyun.Client) (*InnerQueryPurchaseImagesResponse, error) {
	resp := &InnerQueryPurchaseImagesResponse{}
	err := client.Request("market-inner", "InnerQueryPurchaseImages", "2016-08-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type InnerQueryPurchaseImagesResponse struct {
	TotalCount goaliyun.Integer
	RequestId  goaliyun.String
	ImageList  InnerQueryPurchaseImagesImageList
}

type InnerQueryPurchaseImagesImage struct {
	RequestImageId goaliyun.String
	ChargeType     goaliyun.String
	ProductCode    goaliyun.String
	ProductSKUCode goaliyun.String
	ImageId        goaliyun.String
	ImageVersion   goaliyun.String
	ImageStatus    goaliyun.String
	RegionId       goaliyun.String
	SupplierId     goaliyun.Integer
	SupplierName   goaliyun.String
}

type InnerQueryPurchaseImagesImageList []InnerQueryPurchaseImagesImage

func (list *InnerQueryPurchaseImagesImageList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]InnerQueryPurchaseImagesImage)
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
