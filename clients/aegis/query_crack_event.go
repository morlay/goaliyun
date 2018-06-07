package aegis

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type QueryCrackEventRequest struct {
	EndTime     string `position:"Query" name:"EndTime"`
	CurrentPage int64  `position:"Query" name:"CurrentPage"`
	StartTime   string `position:"Query" name:"StartTime"`
	Uuid        string `position:"Query" name:"Uuid"`
	Status      int64  `position:"Query" name:"Status"`
	RegionId    string `position:"Query" name:"RegionId"`
}

func (req *QueryCrackEventRequest) Invoke(client goaliyun.Client) (*QueryCrackEventResponse, error) {
	resp := &QueryCrackEventResponse{}
	err := client.Request("aegis", "QueryCrackEvent", "2016-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type QueryCrackEventResponse struct {
	RequestId goaliyun.String
	Code      goaliyun.String
	Success   bool
	Message   goaliyun.String
	Data      QueryCrackEventData
}

type QueryCrackEventData struct {
	List     QueryCrackEventEntityList
	PageInfo QueryCrackEventPageInfo
}

type QueryCrackEventEntity struct {
	Uuid           goaliyun.String
	AttackTime     goaliyun.String
	AttackType     goaliyun.Integer
	AttackTypeName goaliyun.String
	BuyVersion     goaliyun.String
	CrackSourceIp  goaliyun.String
	CrackTimes     goaliyun.Integer
	GroupId        goaliyun.Integer
	InstanceName   goaliyun.String
	InstanceId     goaliyun.String
	Ip             goaliyun.String
	Region         goaliyun.String
	Status         goaliyun.Integer
	StatusName     goaliyun.String
	Location       goaliyun.String
	InWhite        goaliyun.Integer
	UserName       goaliyun.String
}

type QueryCrackEventPageInfo struct {
	CurrentPage goaliyun.Integer
	PageSize    goaliyun.Integer
	TotalCount  goaliyun.Integer
	Count       goaliyun.Integer
}

type QueryCrackEventEntityList []QueryCrackEventEntity

func (list *QueryCrackEventEntityList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryCrackEventEntity)
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
