package drds

import (
	"github.com/morlay/goaliyun"
)

type ModifyReadOnlyAccountPasswordRequest struct {
	NewPasswd      string `position:"Query" name:"NewPasswd"`
	DbName         string `position:"Query" name:"DbName"`
	AccountName    string `position:"Query" name:"AccountName"`
	OriginPassword string `position:"Query" name:"OriginPassword"`
	DrdsInstanceId string `position:"Query" name:"DrdsInstanceId"`
	RegionId       string `position:"Query" name:"RegionId"`
}

func (req *ModifyReadOnlyAccountPasswordRequest) Invoke(client goaliyun.Client) (*ModifyReadOnlyAccountPasswordResponse, error) {
	resp := &ModifyReadOnlyAccountPasswordResponse{}
	err := client.Request("drds", "ModifyReadOnlyAccountPassword", "2017-10-16", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifyReadOnlyAccountPasswordResponse struct {
	RequestId goaliyun.String
	Success   bool
}
