package aegis

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type GetEntityListRequest struct {
	GroupId     int64  `position:"Query" name:"GroupId"`
	PageSize    int64  `position:"Query" name:"PageSize"`
	Remark      string `position:"Query" name:"Remark"`
	EventType   string `position:"Query" name:"EventType"`
	CurrentPage int64  `position:"Query" name:"CurrentPage"`
	RegionNo    string `position:"Query" name:"RegionNo"`
	RegionId    string `position:"Query" name:"RegionId"`
}

func (req *GetEntityListRequest) Invoke(client goaliyun.Client) (*GetEntityListResponse, error) {
	resp := &GetEntityListResponse{}
	err := client.Request("aegis", "GetEntityList", "2016-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type GetEntityListResponse struct {
	RequestId goaliyun.String
	Code      goaliyun.String
	Success   bool
	Message   goaliyun.String
	Data      GetEntityListData
}

type GetEntityListData struct {
	List     GetEntityListEntityList
	PageInfo GetEntityListPageInfo
}

type GetEntityListEntity struct {
	Uuid         goaliyun.String
	GroupId      goaliyun.Integer
	Ip           goaliyun.String
	InstanceName goaliyun.String
	InstanceId   goaliyun.String
	Region       goaliyun.String
	Os           goaliyun.String
	Flag         goaliyun.String
	BuyVersion   goaliyun.String
	AegisOnline  bool
	AegisVersion goaliyun.String
}

type GetEntityListPageInfo struct {
	CurrentPage goaliyun.Integer
	PageSize    goaliyun.Integer
	TotalCount  goaliyun.Integer
	Count       goaliyun.Integer
}

type GetEntityListEntityList []GetEntityListEntity

func (list *GetEntityListEntityList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetEntityListEntity)
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
