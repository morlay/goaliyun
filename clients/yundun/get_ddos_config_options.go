package yundun

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type GetDdosConfigOptionsRequest struct {
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *GetDdosConfigOptionsRequest) Invoke(client goaliyun.Client) (*GetDdosConfigOptionsResponse, error) {
	resp := &GetDdosConfigOptionsResponse{}
	err := client.Request("yundun", "GetDdosConfigOptions", "2015-04-16", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type GetDdosConfigOptionsResponse struct {
	RequestId                  goaliyun.String
	RequestThresholdOptions1   GetDdosConfigOptionsRequestThresholdOptionList
	RequestThresholdOptions2   GetDdosConfigOptionsRequestThresholdOptionList
	ConnectionThresholdOptions GetDdosConfigOptionsConnectionThresholdOptionList
	QpsOptions1                GetDdosConfigOptionsQpsOptions1List
	QpsOptions2                GetDdosConfigOptionsQpsOptions2List
}

type GetDdosConfigOptionsRequestThresholdOption struct {
	Bps goaliyun.Integer
	Pps goaliyun.Integer
}

type GetDdosConfigOptionsConnectionThresholdOption struct {
	Sipconn goaliyun.Integer
	Sipnew  goaliyun.Integer
}

type GetDdosConfigOptionsRequestThresholdOptionList []GetDdosConfigOptionsRequestThresholdOption

func (list *GetDdosConfigOptionsRequestThresholdOptionList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetDdosConfigOptionsRequestThresholdOption)
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

type GetDdosConfigOptionsConnectionThresholdOptionList []GetDdosConfigOptionsConnectionThresholdOption

func (list *GetDdosConfigOptionsConnectionThresholdOptionList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetDdosConfigOptionsConnectionThresholdOption)
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

type GetDdosConfigOptionsQpsOptions1List []goaliyun.String

func (list *GetDdosConfigOptionsQpsOptions1List) UnmarshalJSON(data []byte) error {
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

type GetDdosConfigOptionsQpsOptions2List []goaliyun.String

func (list *GetDdosConfigOptionsQpsOptions2List) UnmarshalJSON(data []byte) error {
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
