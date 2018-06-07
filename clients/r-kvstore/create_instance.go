package r_kvstore

import (
	"github.com/morlay/goaliyun"
)

type CreateInstanceRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	NodeType             string `position:"Query" name:"NodeType"`
	CouponNo             string `position:"Query" name:"CouponNo"`
	NetworkType          string `position:"Query" name:"NetworkType"`
	EngineVersion        string `position:"Query" name:"EngineVersion"`
	InstanceClass        string `position:"Query" name:"InstanceClass"`
	Capacity             int64  `position:"Query" name:"Capacity"`
	Password             string `position:"Query" name:"Password"`
	SecurityToken        string `position:"Query" name:"SecurityToken"`
	InstanceType         string `position:"Query" name:"InstanceType"`
	BusinessInfo         string `position:"Query" name:"BusinessInfo"`
	Period               string `position:"Query" name:"Period"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	SrcDBInstanceId      string `position:"Query" name:"SrcDBInstanceId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	BackupId             string `position:"Query" name:"BackupId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Token                string `position:"Query" name:"Token"`
	VSwitchId            string `position:"Query" name:"VSwitchId"`
	InstanceName         string `position:"Query" name:"InstanceName"`
	VpcId                string `position:"Query" name:"VpcId"`
	ZoneId               string `position:"Query" name:"ZoneId"`
	ChargeType           string `position:"Query" name:"ChargeType"`
	Config               string `position:"Query" name:"Config"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *CreateInstanceRequest) Invoke(client goaliyun.Client) (*CreateInstanceResponse, error) {
	resp := &CreateInstanceResponse{}
	err := client.Request("r-kvstore", "CreateInstance", "2015-01-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CreateInstanceResponse struct {
	RequestId        goaliyun.String
	InstanceId       goaliyun.String
	InstanceName     goaliyun.String
	ConnectionDomain goaliyun.String
	Port             goaliyun.Integer
	UserName         goaliyun.String
	InstanceStatus   goaliyun.String
	RegionId         goaliyun.String
	Capacity         goaliyun.Integer
	QPS              goaliyun.Integer
	Bandwidth        goaliyun.Integer
	Connections      goaliyun.Integer
	ZoneId           goaliyun.String
	Config           goaliyun.String
	ChargeType       goaliyun.String
	EndTime          goaliyun.String
	NodeType         goaliyun.String
	NetworkType      goaliyun.String
	VpcId            goaliyun.String
	VSwitchId        goaliyun.String
	PrivateIpAddr    goaliyun.String
}
