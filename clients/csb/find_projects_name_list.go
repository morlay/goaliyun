package csb

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type FindProjectsNameListRequest struct {
	OperationFlag string `position:"Query" name:"OperationFlag"`
	CsbId         int64  `position:"Query" name:"CsbId"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *FindProjectsNameListRequest) Invoke(client goaliyun.Client) (*FindProjectsNameListResponse, error) {
	resp := &FindProjectsNameListResponse{}
	err := client.Request("csb", "FindProjectsNameList", "2017-11-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type FindProjectsNameListResponse struct {
	Code      goaliyun.Integer
	Message   goaliyun.String
	RequestId goaliyun.String
	Data      FindProjectsNameListData
}

type FindProjectsNameListData struct {
	ProjectNameList FindProjectsNameListProjectNameListList
}

type FindProjectsNameListProjectNameListList []goaliyun.String

func (list *FindProjectsNameListProjectNameListList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]goaliyun.String)
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
