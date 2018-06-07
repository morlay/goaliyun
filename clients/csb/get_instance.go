package csb

import (
	"github.com/morlay/goaliyun"
)

type GetInstanceRequest struct {
	CsbId    int64  `position:"Query" name:"CsbId"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *GetInstanceRequest) Invoke(client goaliyun.Client) (*GetInstanceResponse, error) {
	resp := &GetInstanceResponse{}
	err := client.Request("csb", "GetInstance", "2017-11-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type GetInstanceResponse struct {
	Code      goaliyun.Integer
	Message   goaliyun.String
	RequestId goaliyun.String
	Data      GetInstanceData
}

type GetInstanceData struct {
	Instance GetInstanceInstance
}

type GetInstanceInstance struct {
	ApprLevel            goaliyun.Integer
	ApprUser1            goaliyun.String
	ApprUser2            goaliyun.String
	BrokerVpcId          goaliyun.String
	BrokerVpcName        goaliyun.String
	ClientVpcId          goaliyun.String
	ClientVpcName        goaliyun.String
	ClusterMembers       goaliyun.Integer
	CredentialGroup      goaliyun.Integer
	CsbAccountId         goaliyun.String
	CsbId                goaliyun.Integer
	DbStatus             goaliyun.Integer
	Description          goaliyun.String
	FrontStatus          goaliyun.String
	GmtCreate            goaliyun.Integer
	GmtModified          goaliyun.Integer
	Id                   goaliyun.Integer
	InstanceCategory     goaliyun.Integer
	InstanceType         goaliyun.Integer
	IpList               goaliyun.String
	IsImported           bool
	IsPublic             bool
	Name                 goaliyun.String
	OwnerId              goaliyun.String
	SentinelCtlStr       goaliyun.String
	SentinelCtrl         goaliyun.Integer
	SentinelGridInterval goaliyun.Integer
	SentinelQps          goaliyun.Integer
	Status               goaliyun.String
	StatusCode           goaliyun.Integer
	TenantId             goaliyun.String
	Testable             bool
	UserId               goaliyun.String
	Visible              bool
	VpcName              goaliyun.String
}
