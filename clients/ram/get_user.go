package ram

import (
	"github.com/morlay/goaliyun"
)

type GetUserRequest struct {
	UserName string `position:"Query" name:"UserName"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *GetUserRequest) Invoke(client goaliyun.Client) (*GetUserResponse, error) {
	resp := &GetUserResponse{}
	err := client.Request("ram", "GetUser", "2015-05-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type GetUserResponse struct {
	RequestId goaliyun.String
	User      GetUserUser
}

type GetUserUser struct {
	UserId        goaliyun.String
	UserName      goaliyun.String
	DisplayName   goaliyun.String
	MobilePhone   goaliyun.String
	Email         goaliyun.String
	Comments      goaliyun.String
	CreateDate    goaliyun.String
	UpdateDate    goaliyun.String
	LastLoginDate goaliyun.String
}
