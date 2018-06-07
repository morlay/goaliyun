package market_inner

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type InnerQueryMarketImageRequest struct {
	ImagePc  string `position:"Query" name:"ImagePc"`
	AliUid   int64  `position:"Query" name:"AliUid"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *InnerQueryMarketImageRequest) Invoke(client goaliyun.Client) (*InnerQueryMarketImageResponse, error) {
	resp := &InnerQueryMarketImageResponse{}
	err := client.Request("market-inner", "InnerQueryMarketImage", "2016-08-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type InnerQueryMarketImageResponse struct {
	RequestId          goaliyun.String
	Success            bool
	MarketImageProduct InnerQueryMarketImageMarketImageProduct
}

type InnerQueryMarketImageMarketImageProduct struct {
	ImagePc             goaliyun.String
	ExtendInfo          goaliyun.String
	ImageRegionInfoList InnerQueryMarketImageImageRegionInfoList
	SupportChargeTypes  InnerQueryMarketImageSupportChargeTypeList
}

type InnerQueryMarketImageImageRegionInfo struct {
	RegionNo     goaliyun.String
	ImageId      goaliyun.String
	ImageVersion goaliyun.String
}

type InnerQueryMarketImageImageRegionInfoList []InnerQueryMarketImageImageRegionInfo

func (list *InnerQueryMarketImageImageRegionInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]InnerQueryMarketImageImageRegionInfo)
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

type InnerQueryMarketImageSupportChargeTypeList []goaliyun.String

func (list *InnerQueryMarketImageSupportChargeTypeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]goaliyun.String)
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
