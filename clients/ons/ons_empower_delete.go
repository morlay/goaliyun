package ons

import (
	"github.com/morlay/goaliyun"
)

type OnsEmpowerDeleteRequest struct {
	PreventCache int64  `position:"Query" name:"PreventCache"`
	OnsRegionId  string `position:"Query" name:"OnsRegionId"`
	OnsPlatform  string `position:"Query" name:"OnsPlatform"`
	EmpowerUser  string `position:"Query" name:"EmpowerUser"`
	Topic        string `position:"Query" name:"Topic"`
	RegionId     string `position:"Query" name:"RegionId"`
}

func (req *OnsEmpowerDeleteRequest) Invoke(client goaliyun.Client) (*OnsEmpowerDeleteResponse, error) {
	resp := &OnsEmpowerDeleteResponse{}
	err := client.Request("ons", "OnsEmpowerDelete", "2017-09-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type OnsEmpowerDeleteResponse struct {
	RequestId goaliyun.String
	HelpUrl   goaliyun.String
}
