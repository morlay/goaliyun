package emr

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type QueryPriceRequest struct {
	ResourceOwnerId        int64                    `position:"Query" name:"ResourceOwnerId"`
	Period                 int64                    `position:"Query" name:"Period"`
	TaskInstanceType       string                   `position:"Query" name:"TaskInstanceType"`
	TaskDiskType           string                   `position:"Query" name:"TaskDiskType"`
	IoOptimized            string                   `position:"Query" name:"IoOptimized"`
	MasterDiskType         string                   `position:"Query" name:"MasterDiskType"`
	TaskInstanceQuantity   int64                    `position:"Query" name:"TaskInstanceQuantity"`
	MasterInstanceType     string                   `position:"Query" name:"MasterInstanceType"`
	CoreInstanceQuantity   int64                    `position:"Query" name:"CoreInstanceQuantity"`
	Duration               int64                    `position:"Query" name:"Duration"`
	MasterDiskQuantity     int64                    `position:"Query" name:"MasterDiskQuantity"`
	CoreDiskQuantity       int64                    `position:"Query" name:"CoreDiskQuantity"`
	CoreInstanceType       string                   `position:"Query" name:"CoreInstanceType"`
	NetType                string                   `position:"Query" name:"NetType"`
	HostGroups             *QueryPriceHostGroupList `position:"Query" type:"Repeated" name:"HostGroup"`
	ZoneId                 string                   `position:"Query" name:"ZoneId"`
	CoreDiskType           string                   `position:"Query" name:"CoreDiskType"`
	ChargeType             string                   `position:"Query" name:"ChargeType"`
	MasterInstanceQuantity int64                    `position:"Query" name:"MasterInstanceQuantity"`
	TaskDiskQuantity       int64                    `position:"Query" name:"TaskDiskQuantity"`
	RegionId               string                   `position:"Query" name:"RegionId"`
}

func (req *QueryPriceRequest) Invoke(client goaliyun.Client) (*QueryPriceResponse, error) {
	resp := &QueryPriceResponse{}
	err := client.Request("emr", "QueryPrice", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type QueryPriceHostGroup struct {
	HostGroupId     string `name:"HostGroupId"`
	HostGroupName   string `name:"HostGroupName"`
	HostGroupType   string `name:"HostGroupType"`
	ClusterId       string `name:"ClusterId"`
	Comment         string `name:"Comment"`
	CreateType      string `name:"CreateType"`
	ChargeType      string `name:"ChargeType"`
	Period          int64  `name:"Period"`
	NodeCount       int64  `name:"NodeCount"`
	InstanceType    string `name:"InstanceType"`
	DiskType        string `name:"DiskType"`
	DiskCapacity    int64  `name:"DiskCapacity"`
	DiskCount       int64  `name:"DiskCount"`
	SysDiskType     string `name:"SysDiskType"`
	SysDiskCapacity int64  `name:"SysDiskCapacity"`
	AutoRenew       string `name:"AutoRenew"`
	VSwitchId       string `name:"VSwitchId"`
}

type QueryPriceResponse struct {
	RequestId  goaliyun.String
	EmrPrice   goaliyun.String
	EcsPrice   goaliyun.String
	EmrPriceDO QueryPriceEmrPriceDO
	EcsPriceDO QueryPriceEcsPriceDO
}

type QueryPriceEmrPriceDO struct {
	OriginalPrice goaliyun.String
	DiscountPrice goaliyun.String
	TradePrice    goaliyun.String
	TaxPrice      goaliyun.String
	Currency      goaliyun.String
}

type QueryPriceEcsPriceDO struct {
	OriginalPrice goaliyun.String
	DiscountPrice goaliyun.String
	TradePrice    goaliyun.String
	TaxPrice      goaliyun.String
	Currency      goaliyun.String
}

type QueryPriceHostGroupList []QueryPriceHostGroup

func (list *QueryPriceHostGroupList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryPriceHostGroup)
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
