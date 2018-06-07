package qualitycheck

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type GetDataSetListRequest struct {
	JsonStr  string `position:"Query" name:"JsonStr"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *GetDataSetListRequest) Invoke(client goaliyun.Client) (*GetDataSetListResponse, error) {
	resp := &GetDataSetListResponse{}
	err := client.Request("qualitycheck", "GetDataSetList", "2016-08-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type GetDataSetListResponse struct {
	RequestId goaliyun.String
	Success   bool
	Code      goaliyun.String
	Message   goaliyun.String
	Count     goaliyun.Integer
	Data      GetDataSetListDataSetList
}

type GetDataSetListDataSet struct {
	SetId         goaliyun.Integer
	SetName       goaliyun.String
	SetDomain     goaliyun.String
	SetRoleArn    goaliyun.String
	SetBucketName goaliyun.String
	SetFolderName goaliyun.String
	ChannelType   goaliyun.Integer
	CreateType    goaliyun.Integer
}

type GetDataSetListDataSetList []GetDataSetListDataSet

func (list *GetDataSetListDataSetList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetDataSetListDataSet)
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
