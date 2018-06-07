package ram

import (
	"github.com/morlay/goaliyun"
)

type DeleteUserRequest struct {
	UserName string `position:"Query" name:"UserName"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *DeleteUserRequest) Invoke(client goaliyun.Client) (*DeleteUserResponse, error) {
	resp := &DeleteUserResponse{}
	err := client.Request("ram", "DeleteUser", "2015-05-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteUserResponse struct {
	RequestId goaliyun.String
}
