package drds

import (
	"github.com/morlay/goaliyun"
)

type CreateReadOnlyAccountRequest struct {
	Password       string `position:"Query" name:"Password"`
	DbName         string `position:"Query" name:"DbName"`
	DrdsInstanceId string `position:"Query" name:"DrdsInstanceId"`
	RegionId       string `position:"Query" name:"RegionId"`
}

func (req *CreateReadOnlyAccountRequest) Invoke(client goaliyun.Client) (*CreateReadOnlyAccountResponse, error) {
	resp := &CreateReadOnlyAccountResponse{}
	err := client.Request("drds", "CreateReadOnlyAccount", "2017-10-16", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CreateReadOnlyAccountResponse struct {
	RequestId goaliyun.String
	Success   bool
	Data      CreateReadOnlyAccountData
}

type CreateReadOnlyAccountData struct {
	DbName         goaliyun.String
	DrdsInstanceId goaliyun.String
	AccountName    goaliyun.String
}
