package ocs

import (
	"github.com/morlay/goaliyun"
)

type CreateInstanceRequest struct {
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OcsInstanceName      string `position:"Query" name:"OcsInstanceName"`
	Password             string `position:"Query" name:"Password"`
	Capacity             int64  `position:"Query" name:"Capacity"`
	Region               string `position:"Query" name:"Region"`
	IzNo                 string `position:"Query" name:"IzNo"`
	NetworkType          string `position:"Query" name:"NetworkType"`
	VpcId                string `position:"Query" name:"VpcId"`
	VSwitchId            string `position:"Query" name:"VSwitchId"`
	PrivateIp            string `position:"Query" name:"PrivateIp"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *CreateInstanceRequest) Invoke(client goaliyun.Client) (*CreateInstanceResponse, error) {
	resp := &CreateInstanceResponse{}
	err := client.Request("ocs", "CreateInstance", "2015-03-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CreateInstanceResponse struct {
	RequestId   goaliyun.String
	OcsInstance CreateInstanceOcsInstance
}

type CreateInstanceOcsInstance struct {
	OcsInstanceId     goaliyun.String
	OcsInstanceName   goaliyun.String
	Capacity          goaliyun.Integer
	Qps               goaliyun.Integer
	Bandwidth         goaliyun.Integer
	Connections       goaliyun.Integer
	ConnectionDomain  goaliyun.String
	Port              goaliyun.Integer
	UserName          goaliyun.String
	RegionId          goaliyun.String
	OcsInstanceStatus goaliyun.String
	GmtCreated        goaliyun.String
	EndTime           goaliyun.String
	ChargeType        goaliyun.String
	IzId              goaliyun.String
	NetworkType       goaliyun.String
	VpcId             goaliyun.String
	VSwitchId         goaliyun.String
	PrivateIp         goaliyun.String
}
