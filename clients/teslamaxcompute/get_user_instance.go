package teslamaxcompute

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type GetUserInstanceRequest struct {
	PageSize int64  `position:"Query" name:"PageSize"`
	PageNum  int64  `position:"Query" name:"PageNum"`
	Region   string `position:"Query" name:"Region"`
	User     string `position:"Query" name:"User"`
	Status   string `position:"Query" name:"Status"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *GetUserInstanceRequest) Invoke(client goaliyun.Client) (*GetUserInstanceResponse, error) {
	resp := &GetUserInstanceResponse{}
	err := client.Request("teslamaxcompute", "GetUserInstance", "2018-01-04", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type GetUserInstanceResponse struct {
	Code      goaliyun.Integer
	Message   goaliyun.String
	RequestId goaliyun.String
	Data      GetUserInstanceData
}

type GetUserInstanceData struct {
	Total  goaliyun.Integer
	Detail GetUserInstanceInstanceList
}

type GetUserInstanceInstance struct {
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

type GetUserInstanceInstanceList []GetUserInstanceInstance

func (list *GetUserInstanceInstanceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetUserInstanceInstance)
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
