package teslamaxcompute

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type GetQuotaInstanceRequest struct {
	Cluster   string `position:"Query" name:"Cluster"`
	PageSize  int64  `position:"Query" name:"PageSize"`
	QuotaId   string `position:"Query" name:"QuotaId"`
	PageNum   int64  `position:"Query" name:"PageNum"`
	Region    string `position:"Query" name:"Region"`
	QuotaName string `position:"Query" name:"QuotaName"`
	Status    string `position:"Query" name:"Status"`
	RegionId  string `position:"Query" name:"RegionId"`
}

func (req *GetQuotaInstanceRequest) Invoke(client goaliyun.Client) (*GetQuotaInstanceResponse, error) {
	resp := &GetQuotaInstanceResponse{}
	err := client.Request("teslamaxcompute", "GetQuotaInstance", "2018-01-04", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type GetQuotaInstanceResponse struct {
	Code      goaliyun.Integer
	Message   goaliyun.String
	RequestId goaliyun.String
	Data      GetQuotaInstanceData
}

type GetQuotaInstanceData struct {
	Total  goaliyun.Integer
	Detail GetQuotaInstanceInstanceList
}

type GetQuotaInstanceInstance struct {
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
	User            goaliyun.String
	IsRealOwner     goaliyun.String
	ProjectOwner    goaliyun.String
	CollectTime     goaliyun.String
}

type GetQuotaInstanceInstanceList []GetQuotaInstanceInstance

func (list *GetQuotaInstanceInstanceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetQuotaInstanceInstance)
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
