package crm

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type FindAllContacterRequest struct {
	KpId     int64  `position:"Query" name:"KpId"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *FindAllContacterRequest) Invoke(client goaliyun.Client) (*FindAllContacterResponse, error) {
	resp := &FindAllContacterResponse{}
	err := client.Request("crm", "FindAllContacter", "2015-03-24", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type FindAllContacterResponse struct {
	Success       bool
	ResultCode    goaliyun.String
	ResultMessage goaliyun.String
	Data          FindAllContacterContacterInfoList
}

type FindAllContacterContacterInfo struct {
	ContacterId       goaliyun.Integer
	KpId              goaliyun.Integer
	CustomerId        goaliyun.Integer
	ContacterName     goaliyun.String
	ContacterEmail    goaliyun.String
	ContacterMobile   goaliyun.String
	ContacterPosition goaliyun.String
}

type FindAllContacterContacterInfoList []FindAllContacterContacterInfo

func (list *FindAllContacterContacterInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]FindAllContacterContacterInfo)
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
