package teslamaxcompute

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type GetClusterInstanceRequest struct {
	Cluster  string `position:"Query" name:"Cluster"`
	PageSize int64  `position:"Query" name:"PageSize"`
	PageNum  int64  `position:"Query" name:"PageNum"`
	Region   string `position:"Query" name:"Region"`
	Status   string `position:"Query" name:"Status"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *GetClusterInstanceRequest) Invoke(client goaliyun.Client) (*GetClusterInstanceResponse, error) {
	resp := &GetClusterInstanceResponse{}
	err := client.Request("teslamaxcompute", "GetClusterInstance", "2018-01-04", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type GetClusterInstanceResponse struct {
	Code      goaliyun.Integer
	Message   goaliyun.String
	RequestId goaliyun.String
	Data      GetClusterInstanceData
}

type GetClusterInstanceData struct {
	Total  goaliyun.Integer
	Detail GetClusterInstanceInstanceList
}

type GetClusterInstanceInstance struct {
	Project         goaliyun.String
	InstanceId      goaliyun.String
	Status          goaliyun.String
	UserAccount     goaliyun.String
	NickName        goaliyun.String
	Cluster         goaliyun.String
	RunTime         goaliyun.String
	CpuUsed         goaliyun.Integer
	CpuRequest      goaliyun.Integer
	CpuUsedTotal    goaliyun.Integer
	CpuUsedRatioMax goaliyun.Float
	CpuUsedRatioMin goaliyun.Float
	MemUsed         goaliyun.Integer
	MemRequest      goaliyun.Integer
	MemUsedTotal    goaliyun.Integer
	MemUsedRatioMax goaliyun.Float
	MemUsedRatioMin goaliyun.Float
	TaskType        goaliyun.String
	SkynetId        goaliyun.String
	QuotaName       goaliyun.String
	QuotaId         goaliyun.Integer
}

type GetClusterInstanceInstanceList []GetClusterInstanceInstance

func (list *GetClusterInstanceInstanceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetClusterInstanceInstance)
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
