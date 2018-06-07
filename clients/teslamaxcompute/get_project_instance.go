package teslamaxcompute

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type GetProjectInstanceRequest struct {
	PageSize int64  `position:"Query" name:"PageSize"`
	Project  string `position:"Query" name:"Project"`
	PageNum  int64  `position:"Query" name:"PageNum"`
	Region   string `position:"Query" name:"Region"`
	Status   string `position:"Query" name:"Status"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *GetProjectInstanceRequest) Invoke(client goaliyun.Client) (*GetProjectInstanceResponse, error) {
	resp := &GetProjectInstanceResponse{}
	err := client.Request("teslamaxcompute", "GetProjectInstance", "2018-01-04", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type GetProjectInstanceResponse struct {
	Code      goaliyun.Integer
	Message   goaliyun.String
	RequestId goaliyun.String
	Data      GetProjectInstanceData
}

type GetProjectInstanceData struct {
	Total  goaliyun.Integer
	Detail GetProjectInstanceInstanceList
}

type GetProjectInstanceInstance struct {
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

type GetProjectInstanceInstanceList []GetProjectInstanceInstance

func (list *GetProjectInstanceInstanceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetProjectInstanceInstance)
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
