package ehpc

import (
	"github.com/morlay/goaliyun"
)

type SetJobUserRequest struct {
	RunasUserPassword string `position:"Query" name:"RunasUserPassword"`
	RunasUser         string `position:"Query" name:"RunasUser"`
	ClusterId         string `position:"Query" name:"ClusterId"`
	RegionId          string `position:"Query" name:"RegionId"`
}

func (req *SetJobUserRequest) Invoke(client goaliyun.Client) (*SetJobUserResponse, error) {
	resp := &SetJobUserResponse{}
	err := client.Request("ehpc", "SetJobUser", "2018-04-12", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SetJobUserResponse struct {
	RequestId goaliyun.String
}
