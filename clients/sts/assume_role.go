package sts

import (
	"github.com/morlay/goaliyun"
)

type AssumeRoleRequest struct {
	RoleArn         string `position:"Query" name:"RoleArn"`
	RoleSessionName string `position:"Query" name:"RoleSessionName"`
	DurationSeconds int64  `position:"Query" name:"DurationSeconds"`
	Policy          string `position:"Query" name:"Policy"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *AssumeRoleRequest) Invoke(client goaliyun.Client) (*AssumeRoleResponse, error) {
	resp := &AssumeRoleResponse{}
	err := client.Request("sts", "AssumeRole", "2015-04-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type AssumeRoleResponse struct {
	RequestId       goaliyun.String
	Credentials     AssumeRoleCredentials
	AssumedRoleUser AssumeRoleAssumedRoleUser
}

type AssumeRoleCredentials struct {
	SecurityToken   goaliyun.String
	AccessKeySecret goaliyun.String
	AccessKeyId     goaliyun.String
	Expiration      goaliyun.String
}

type AssumeRoleAssumedRoleUser struct {
	Arn           goaliyun.String
	AssumedRoleId goaliyun.String
}
