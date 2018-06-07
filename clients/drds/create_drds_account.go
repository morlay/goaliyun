package drds

import (
	"github.com/morlay/goaliyun"
)

type CreateDrdsAccountRequest struct {
	Password       string `position:"Query" name:"Password"`
	DbName         string `position:"Query" name:"DbName"`
	DrdsInstanceId string `position:"Query" name:"DrdsInstanceId"`
	UserName       string `position:"Query" name:"UserName"`
	RegionId       string `position:"Query" name:"RegionId"`
}

func (req *CreateDrdsAccountRequest) Invoke(client goaliyun.Client) (*CreateDrdsAccountResponse, error) {
	resp := &CreateDrdsAccountResponse{}
	err := client.Request("drds", "CreateDrdsAccount", "2017-10-16", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CreateDrdsAccountResponse struct {
	RequestId goaliyun.String
	Success   bool
}
