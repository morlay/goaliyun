package emr

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type QueryClusterOrdersRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	Id              string `position:"Query" name:"Id"`
	Status          string `position:"Query" name:"Status"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *QueryClusterOrdersRequest) Invoke(client goaliyun.Client) (*QueryClusterOrdersResponse, error) {
	resp := &QueryClusterOrdersResponse{}
	err := client.Request("emr", "QueryClusterOrders", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type QueryClusterOrdersResponse struct {
	RequestId goaliyun.String
	OrderList QueryClusterOrdersOrderList
}

type QueryClusterOrdersOrder struct {
	OrderId         goaliyun.String
	CreateTimestamp goaliyun.Integer
}

type QueryClusterOrdersOrderList []QueryClusterOrdersOrder

func (list *QueryClusterOrdersOrderList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryClusterOrdersOrder)
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
