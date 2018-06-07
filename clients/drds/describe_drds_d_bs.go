package drds

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeDrdsDBsRequest struct {
	DrdsInstanceId string `position:"Query" name:"DrdsInstanceId"`
	RegionId       string `position:"Query" name:"RegionId"`
}

func (req *DescribeDrdsDBsRequest) Invoke(client goaliyun.Client) (*DescribeDrdsDBsResponse, error) {
	resp := &DescribeDrdsDBsResponse{}
	err := client.Request("drds", "DescribeDrdsDBs", "2017-10-16", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeDrdsDBsResponse struct {
	RequestId goaliyun.String
	Success   bool
	Data      DescribeDrdsDBsDbList
}

type DescribeDrdsDBsDb struct {
	DbName     goaliyun.String
	Status     goaliyun.Integer
	CreateTime goaliyun.String
	Msg        goaliyun.String
	Mode       goaliyun.String
}

type DescribeDrdsDBsDbList []DescribeDrdsDBsDb

func (list *DescribeDrdsDBsDbList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDrdsDBsDb)
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
