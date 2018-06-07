package market_inner

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type InnerDescribeImageInstancesRequest struct {
	EcsInstanceId        string                                              `position:"Query" name:"EcsInstanceId"`
	PageSize             int64                                               `position:"Query" name:"PageSize"`
	ImageNo              string                                              `position:"Query" name:"ImageNo"`
	ProductName          string                                              `position:"Query" name:"ProductName"`
	PageNumber           int64                                               `position:"Query" name:"PageNumber"`
	ImageInstanceIdLists *InnerDescribeImageInstancesImageInstanceIdListList `position:"Query" type:"Repeated" name:"ImageInstanceIdList"`
	RegionNo             string                                              `position:"Query" name:"RegionNo"`
	RegionId             string                                              `position:"Query" name:"RegionId"`
}

func (req *InnerDescribeImageInstancesRequest) Invoke(client goaliyun.Client) (*InnerDescribeImageInstancesResponse, error) {
	resp := &InnerDescribeImageInstancesResponse{}
	err := client.Request("market-inner", "InnerDescribeImageInstances", "2016-08-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type InnerDescribeImageInstancesResponse struct {
	PageNumber goaliyun.Integer
	PageSize   goaliyun.Integer
	TotalCount goaliyun.Integer
	RequestId  goaliyun.String
	ImageList  InnerDescribeImageInstancesImageList
}

type InnerDescribeImageInstancesImage struct {
	InstanceId     goaliyun.Integer
	OrderId        goaliyun.Integer
	SupplierId     goaliyun.Integer
	SupplierName   goaliyun.String
	ProductCode    goaliyun.String
	ProductSkuCode goaliyun.String
	ProductName    goaliyun.String
	Status         goaliyun.String
	ChargeType     goaliyun.String
	BeganOn        goaliyun.Integer
	EndOn          goaliyun.Integer
	CreatedOn      goaliyun.Integer
	RemainTime     goaliyun.Integer
	RegionNo       goaliyun.String
	ImageNo        goaliyun.String
	ImageVersion   goaliyun.String
	EcsInstanceId  goaliyun.String
}

type InnerDescribeImageInstancesImageInstanceIdListList []int64

func (list *InnerDescribeImageInstancesImageInstanceIdListList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]int64)
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

type InnerDescribeImageInstancesImageList []InnerDescribeImageInstancesImage

func (list *InnerDescribeImageInstancesImageList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]InnerDescribeImageInstancesImage)
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
