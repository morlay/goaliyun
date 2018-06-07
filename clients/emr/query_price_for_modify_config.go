package emr

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type QueryPriceForModifyConfigRequest struct {
	ResourceOwnerId   int64                                          `position:"Query" name:"ResourceOwnerId"`
	ModifyConfigSpecs *QueryPriceForModifyConfigModifyConfigSpecList `position:"Query" type:"Repeated" name:"ModifyConfigSpec"`
	ClusterId         string                                         `position:"Query" name:"ClusterId"`
	RegionId          string                                         `position:"Query" name:"RegionId"`
}

func (req *QueryPriceForModifyConfigRequest) Invoke(client goaliyun.Client) (*QueryPriceForModifyConfigResponse, error) {
	resp := &QueryPriceForModifyConfigResponse{}
	err := client.Request("emr", "QueryPriceForModifyConfig", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type QueryPriceForModifyConfigModifyConfigSpec struct {
	HostGroupId     string `name:"HostGroupId"`
	NewInstanceType string `name:"NewInstanceType"`
	NewDiskSize     int64  `name:"NewDiskSize"`
}

type QueryPriceForModifyConfigResponse struct {
	RequestId  goaliyun.String
	EcsId      goaliyun.String
	EmrPriceDO QueryPriceForModifyConfigEmrPriceDO
	EcsPriceDO QueryPriceForModifyConfigEcsPriceDO
}

type QueryPriceForModifyConfigEmrPriceDO struct {
	OriginalPrice goaliyun.String
	DiscountPrice goaliyun.String
	TradePrice    goaliyun.String
	TaxPrice      goaliyun.String
	Currency      goaliyun.String
}

type QueryPriceForModifyConfigEcsPriceDO struct {
	OriginalPrice goaliyun.String
	DiscountPrice goaliyun.String
	TradePrice    goaliyun.String
	TaxPrice      goaliyun.String
	Currency      goaliyun.String
}

type QueryPriceForModifyConfigModifyConfigSpecList []QueryPriceForModifyConfigModifyConfigSpec

func (list *QueryPriceForModifyConfigModifyConfigSpecList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryPriceForModifyConfigModifyConfigSpec)
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
