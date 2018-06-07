package imm

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type PutProjectRequest struct {
	Indexers    string `position:"Query" name:"Indexers"`
	Engines     string `position:"Query" name:"Engines"`
	ServiceRole string `position:"Query" name:"ServiceRole"`
	Project     string `position:"Query" name:"Project"`
	RegionId    string `position:"Query" name:"RegionId"`
}

func (req *PutProjectRequest) Invoke(client goaliyun.Client) (*PutProjectResponse, error) {
	resp := &PutProjectResponse{}
	err := client.Request("imm", "PutProject", "2017-09-06", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type PutProjectResponse struct {
	RequestId   goaliyun.String
	Project     goaliyun.String
	CreateTime  goaliyun.String
	ModifyTime  goaliyun.String
	ServiceRole goaliyun.String
	Engines     PutProjectEnginesItemList
	Indexers    PutProjectIndexersItemList
}

type PutProjectEnginesItem struct {
	Name   goaliyun.String
	JobTtl goaliyun.Integer
}

type PutProjectIndexersItem struct {
	Name   goaliyun.String
	Status goaliyun.String
}

type PutProjectEnginesItemList []PutProjectEnginesItem

func (list *PutProjectEnginesItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]PutProjectEnginesItem)
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

type PutProjectIndexersItemList []PutProjectIndexersItem

func (list *PutProjectIndexersItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]PutProjectIndexersItem)
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
