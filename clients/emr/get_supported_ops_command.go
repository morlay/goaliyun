package emr

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type GetSupportedOpsCommandRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *GetSupportedOpsCommandRequest) Invoke(client goaliyun.Client) (*GetSupportedOpsCommandResponse, error) {
	resp := &GetSupportedOpsCommandResponse{}
	err := client.Request("emr", "GetSupportedOpsCommand", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type GetSupportedOpsCommandResponse struct {
	RequestId goaliyun.String
	List      GetSupportedOpsCommandOpsCommandCategoryList
}

type GetSupportedOpsCommandOpsCommandCategory struct {
	Category    goaliyun.String
	CommandList GetSupportedOpsCommandOpsCommandInfoList
}

type GetSupportedOpsCommandOpsCommandInfo struct {
	Id          goaliyun.String
	Name        goaliyun.String
	Discription goaliyun.String
	TargetType  goaliyun.String
	Params      goaliyun.String
}

type GetSupportedOpsCommandOpsCommandCategoryList []GetSupportedOpsCommandOpsCommandCategory

func (list *GetSupportedOpsCommandOpsCommandCategoryList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetSupportedOpsCommandOpsCommandCategory)
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

type GetSupportedOpsCommandOpsCommandInfoList []GetSupportedOpsCommandOpsCommandInfo

func (list *GetSupportedOpsCommandOpsCommandInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetSupportedOpsCommandOpsCommandInfo)
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
