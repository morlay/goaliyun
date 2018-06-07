package ons

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type OnsEmpowerListRequest struct {
	PreventCache int64  `position:"Query" name:"PreventCache"`
	OnsRegionId  string `position:"Query" name:"OnsRegionId"`
	OnsPlatform  string `position:"Query" name:"OnsPlatform"`
	EmpowerUser  string `position:"Query" name:"EmpowerUser"`
	Topic        string `position:"Query" name:"Topic"`
	RegionId     string `position:"Query" name:"RegionId"`
}

func (req *OnsEmpowerListRequest) Invoke(client goaliyun.Client) (*OnsEmpowerListResponse, error) {
	resp := &OnsEmpowerListResponse{}
	err := client.Request("ons", "OnsEmpowerList", "2017-09-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type OnsEmpowerListResponse struct {
	RequestId goaliyun.String
	HelpUrl   goaliyun.String
	Data      OnsEmpowerListAuthOwnerInfoDoList
}

type OnsEmpowerListAuthOwnerInfoDo struct {
	Topic    goaliyun.String
	Owner    goaliyun.Integer
	Relation goaliyun.Integer
}

type OnsEmpowerListAuthOwnerInfoDoList []OnsEmpowerListAuthOwnerInfoDo

func (list *OnsEmpowerListAuthOwnerInfoDoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]OnsEmpowerListAuthOwnerInfoDo)
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
