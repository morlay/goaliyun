package teslamaxcompute

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type QueryCustomerSaleInfoRequest struct {
	RegionName string `position:"Query" name:"RegionName"`
	RegionId   string `position:"Query" name:"RegionId"`
}

func (req *QueryCustomerSaleInfoRequest) Invoke(client goaliyun.Client) (*QueryCustomerSaleInfoResponse, error) {
	resp := &QueryCustomerSaleInfoResponse{}
	err := client.Request("teslamaxcompute", "QueryCustomerSaleInfo", "2018-01-04", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type QueryCustomerSaleInfoResponse struct {
	Code      goaliyun.Integer
	Message   goaliyun.String
	RequestId goaliyun.String
	Data      QueryCustomerSaleInfoData
}

type QueryCustomerSaleInfoData struct {
	LastUpdate goaliyun.String
	Clusters   QueryCustomerSaleInfoClusterList
}

type QueryCustomerSaleInfoCluster struct {
	Cluster     goaliyun.String
	Region      goaliyun.String
	MachineRoom goaliyun.String
	SaleInfos   QueryCustomerSaleInfoSaleInfoList
}

type QueryCustomerSaleInfoSaleInfo struct {
	SaleMode    goaliyun.String
	Uid         goaliyun.String
	Mem         goaliyun.Integer
	Cpu         goaliyun.Integer
	BizCategory goaliyun.String
	Owner       goaliyun.String
	QueryDate   goaliyun.String
}

type QueryCustomerSaleInfoClusterList []QueryCustomerSaleInfoCluster

func (list *QueryCustomerSaleInfoClusterList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryCustomerSaleInfoCluster)
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

type QueryCustomerSaleInfoSaleInfoList []QueryCustomerSaleInfoSaleInfo

func (list *QueryCustomerSaleInfoSaleInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryCustomerSaleInfoSaleInfo)
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
