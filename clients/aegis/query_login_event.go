package aegis

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type QueryLoginEventRequest struct {
	EndTime     string `position:"Query" name:"EndTime"`
	CurrentPage int64  `position:"Query" name:"CurrentPage"`
	StartTime   string `position:"Query" name:"StartTime"`
	Uuid        string `position:"Query" name:"Uuid"`
	Status      int64  `position:"Query" name:"Status"`
	RegionId    string `position:"Query" name:"RegionId"`
}

func (req *QueryLoginEventRequest) Invoke(client goaliyun.Client) (*QueryLoginEventResponse, error) {
	resp := &QueryLoginEventResponse{}
	err := client.Request("aegis", "QueryLoginEvent", "2016-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type QueryLoginEventResponse struct {
	RequestId goaliyun.String
	Code      goaliyun.String
	Success   bool
	Message   goaliyun.String
	Data      QueryLoginEventData
}

type QueryLoginEventData struct {
	List     QueryLoginEventEntityList
	PageInfo QueryLoginEventPageInfo
}

type QueryLoginEventEntity struct {
	Uuid          goaliyun.String
	LoginTime     goaliyun.String
	LoginType     goaliyun.Integer
	LoginTypeName goaliyun.String
	BuyVersion    goaliyun.String
	LoginSourceIp goaliyun.String
	GroupId       goaliyun.Integer
	InstanceName  goaliyun.String
	InstanceId    goaliyun.String
	Ip            goaliyun.String
	Region        goaliyun.String
	Status        goaliyun.Integer
	StatusName    goaliyun.String
	Location      goaliyun.String
	UserName      goaliyun.String
}

type QueryLoginEventPageInfo struct {
	CurrentPage goaliyun.Integer
	PageSize    goaliyun.Integer
	TotalCount  goaliyun.Integer
	Count       goaliyun.Integer
}

type QueryLoginEventEntityList []QueryLoginEventEntity

func (list *QueryLoginEventEntityList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryLoginEventEntity)
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
