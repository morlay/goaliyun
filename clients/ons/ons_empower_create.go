package ons

import (
	"github.com/morlay/goaliyun"
)

type OnsEmpowerCreateRequest struct {
	PreventCache int64  `position:"Query" name:"PreventCache"`
	OnsRegionId  string `position:"Query" name:"OnsRegionId"`
	OnsPlatform  string `position:"Query" name:"OnsPlatform"`
	EmpowerUser  string `position:"Query" name:"EmpowerUser"`
	Topic        string `position:"Query" name:"Topic"`
	Relation     int64  `position:"Query" name:"Relation"`
	RegionId     string `position:"Query" name:"RegionId"`
}

func (req *OnsEmpowerCreateRequest) Invoke(client goaliyun.Client) (*OnsEmpowerCreateResponse, error) {
	resp := &OnsEmpowerCreateResponse{}
	err := client.Request("ons", "OnsEmpowerCreate", "2017-09-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type OnsEmpowerCreateResponse struct {
	RequestId goaliyun.String
	HelpUrl   goaliyun.String
}
