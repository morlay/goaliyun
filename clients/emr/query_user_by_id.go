package emr

import (
	"github.com/morlay/goaliyun"
)

type QueryUserByIdRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *QueryUserByIdRequest) Invoke(client goaliyun.Client) (*QueryUserByIdResponse, error) {
	resp := &QueryUserByIdResponse{}
	err := client.Request("emr", "QueryUserById", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type QueryUserByIdResponse struct {
	RequestId goaliyun.String
	User      QueryUserByIdUser
}

type QueryUserByIdUser struct {
	Id              goaliyun.String
	AliyunId        goaliyun.String
	AliyunOmtId     goaliyun.String
	UserId          goaliyun.String
	Email           goaliyun.String
	Status          goaliyun.String
	DefaultSecGroup goaliyun.String
	RegionId        goaliyun.String
	ChannelId       goaliyun.String
}
