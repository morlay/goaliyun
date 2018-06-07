package iot

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type QueryProductListRequest struct {
	PageSize            int64  `position:"Query" name:"PageSize"`
	CurrentPage         int64  `position:"Query" name:"CurrentPage"`
	AliyunCommodityCode string `position:"Query" name:"AliyunCommodityCode"`
	RegionId            string `position:"Query" name:"RegionId"`
}

func (req *QueryProductListRequest) Invoke(client goaliyun.Client) (*QueryProductListResponse, error) {
	resp := &QueryProductListResponse{}
	err := client.Request("iot", "QueryProductList", "2018-01-20", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type QueryProductListResponse struct {
	RequestId    goaliyun.String
	Success      bool
	ErrorMessage goaliyun.String
	Data         QueryProductListData
}

type QueryProductListData struct {
	CurrentPage goaliyun.Integer
	PageCount   goaliyun.Integer
	PageSize    goaliyun.Integer
	Total       goaliyun.Integer
	List        QueryProductListProductInfoList
}

type QueryProductListProductInfo struct {
	GmtCreate   goaliyun.String
	DataFormat  goaliyun.Integer
	Description goaliyun.String
	DeviceCount goaliyun.Integer
	NodeType    goaliyun.Integer
	ProductKey  goaliyun.String
	ProductName goaliyun.String
}

type QueryProductListProductInfoList []QueryProductListProductInfo

func (list *QueryProductListProductInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryProductListProductInfo)
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
