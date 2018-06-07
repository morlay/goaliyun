package ram

import (
	"github.com/morlay/goaliyun"
)

type CreateUserRequest struct {
	Comments    string `position:"Query" name:"Comments"`
	DisplayName string `position:"Query" name:"DisplayName"`
	MobilePhone string `position:"Query" name:"MobilePhone"`
	Email       string `position:"Query" name:"Email"`
	UserName    string `position:"Query" name:"UserName"`
	RegionId    string `position:"Query" name:"RegionId"`
}

func (req *CreateUserRequest) Invoke(client goaliyun.Client) (*CreateUserResponse, error) {
	resp := &CreateUserResponse{}
	err := client.Request("ram", "CreateUser", "2015-05-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CreateUserResponse struct {
	RequestId goaliyun.String
	User      CreateUserUser
}

type CreateUserUser struct {
	UserId      goaliyun.String
	UserName    goaliyun.String
	DisplayName goaliyun.String
	MobilePhone goaliyun.String
	Email       goaliyun.String
	Comments    goaliyun.String
	CreateDate  goaliyun.String
}
