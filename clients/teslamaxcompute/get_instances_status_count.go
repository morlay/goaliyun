package teslamaxcompute

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type GetInstancesStatusCountRequest struct {
	Cluster   string `position:"Query" name:"Cluster"`
	QuotaId   string `position:"Query" name:"QuotaId"`
	Region    string `position:"Query" name:"Region"`
	QuotaName string `position:"Query" name:"QuotaName"`
	RegionId  string `position:"Query" name:"RegionId"`
}

func (req *GetInstancesStatusCountRequest) Invoke(client goaliyun.Client) (*GetInstancesStatusCountResponse, error) {
	resp := &GetInstancesStatusCountResponse{}
	err := client.Request("teslamaxcompute", "GetInstancesStatusCount", "2018-01-04", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type GetInstancesStatusCountResponse struct {
	Code      goaliyun.Integer
	Message   goaliyun.String
	RequestId goaliyun.String
	Data      GetInstancesStatusCountDataItemList
}

type GetInstancesStatusCountDataItem struct {
	Status goaliyun.String
	Size   goaliyun.Integer
}

type GetInstancesStatusCountDataItemList []GetInstancesStatusCountDataItem

func (list *GetInstancesStatusCountDataItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetInstancesStatusCountDataItem)
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
