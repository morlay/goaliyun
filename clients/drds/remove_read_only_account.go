package drds

import (
	"github.com/morlay/goaliyun"
)

type RemoveReadOnlyAccountRequest struct {
	DbName         string `position:"Query" name:"DbName"`
	AccountName    string `position:"Query" name:"AccountName"`
	DrdsInstanceId string `position:"Query" name:"DrdsInstanceId"`
	RegionId       string `position:"Query" name:"RegionId"`
}

func (req *RemoveReadOnlyAccountRequest) Invoke(client goaliyun.Client) (*RemoveReadOnlyAccountResponse, error) {
	resp := &RemoveReadOnlyAccountResponse{}
	err := client.Request("drds", "RemoveReadOnlyAccount", "2017-10-16", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type RemoveReadOnlyAccountResponse struct {
	RequestId goaliyun.String
	Success   bool
}
