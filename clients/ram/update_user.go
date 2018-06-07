package ram

import (
	"github.com/morlay/goaliyun"
)

type UpdateUserRequest struct {
	NewUserName    string `position:"Query" name:"NewUserName"`
	NewDisplayName string `position:"Query" name:"NewDisplayName"`
	NewMobilePhone string `position:"Query" name:"NewMobilePhone"`
	NewComments    string `position:"Query" name:"NewComments"`
	NewEmail       string `position:"Query" name:"NewEmail"`
	UserName       string `position:"Query" name:"UserName"`
	RegionId       string `position:"Query" name:"RegionId"`
}

func (req *UpdateUserRequest) Invoke(client goaliyun.Client) (*UpdateUserResponse, error) {
	resp := &UpdateUserResponse{}
	err := client.Request("ram", "UpdateUser", "2015-05-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type UpdateUserResponse struct {
	RequestId goaliyun.String
	User      UpdateUserUser
}

type UpdateUserUser struct {
	UserId      goaliyun.String
	UserName    goaliyun.String
	DisplayName goaliyun.String
	MobilePhone goaliyun.String
	Email       goaliyun.String
	Comments    goaliyun.String
	CreateDate  goaliyun.String
	UpdateDate  goaliyun.String
}
