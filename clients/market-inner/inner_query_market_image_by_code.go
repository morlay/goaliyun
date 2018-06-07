package market_inner

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type InnerQueryMarketImageByCodeRequest struct {
	ImagePc  string `position:"Query" name:"ImagePc"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *InnerQueryMarketImageByCodeRequest) Invoke(client goaliyun.Client) (*InnerQueryMarketImageByCodeResponse, error) {
	resp := &InnerQueryMarketImageByCodeResponse{}
	err := client.Request("market-inner", "InnerQueryMarketImageByCode", "2016-08-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type InnerQueryMarketImageByCodeResponse struct {
	RequestId          goaliyun.String
	Success            bool
	MarketImageProduct InnerQueryMarketImageByCodeMarketImageProduct
}

type InnerQueryMarketImageByCodeMarketImageProduct struct {
	ImagePc             goaliyun.String
	ExtendInfo          goaliyun.String
	ImageRegionInfoList InnerQueryMarketImageByCodeImageRegionInfoList
	SupportChargeTypes  InnerQueryMarketImageByCodeSupportChargeTypeList
}

type InnerQueryMarketImageByCodeImageRegionInfo struct {
	RegionNo     goaliyun.String
	ImageId      goaliyun.String
	ImageVersion goaliyun.String
}

type InnerQueryMarketImageByCodeImageRegionInfoList []InnerQueryMarketImageByCodeImageRegionInfo

func (list *InnerQueryMarketImageByCodeImageRegionInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]InnerQueryMarketImageByCodeImageRegionInfo)
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

type InnerQueryMarketImageByCodeSupportChargeTypeList []goaliyun.String

func (list *InnerQueryMarketImageByCodeSupportChargeTypeList) UnmarshalJSON(data []byte) error {
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
