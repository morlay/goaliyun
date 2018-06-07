package imm

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type GetProjectRequest struct {
	Project  string `position:"Query" name:"Project"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *GetProjectRequest) Invoke(client goaliyun.Client) (*GetProjectResponse, error) {
	resp := &GetProjectResponse{}
	err := client.Request("imm", "GetProject", "2017-09-06", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type GetProjectResponse struct {
	RequestId   goaliyun.String
	Project     goaliyun.String
	ServiceRole goaliyun.String
	Endpoint    goaliyun.String
	CreateTime  goaliyun.String
	ModifyTime  goaliyun.String
	Indexers    GetProjectIndexersItemList
	Engines     GetProjectEnginesItemList
}

type GetProjectIndexersItem struct {
	Name   goaliyun.String
	Status goaliyun.String
}

type GetProjectEnginesItem struct {
	Name   goaliyun.String
	JobTtl goaliyun.Integer
}

type GetProjectIndexersItemList []GetProjectIndexersItem

func (list *GetProjectIndexersItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetProjectIndexersItem)
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

type GetProjectEnginesItemList []GetProjectEnginesItem

func (list *GetProjectEnginesItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetProjectEnginesItem)
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
