package drds

import (
	"github.com/morlay/goaliyun"
)

type ModifyDrdsDBPasswdRequest struct {
	NewPasswd      string `position:"Query" name:"NewPasswd"`
	DbName         string `position:"Query" name:"DbName"`
	DrdsInstanceId string `position:"Query" name:"DrdsInstanceId"`
	RegionId       string `position:"Query" name:"RegionId"`
}

func (req *ModifyDrdsDBPasswdRequest) Invoke(client goaliyun.Client) (*ModifyDrdsDBPasswdResponse, error) {
	resp := &ModifyDrdsDBPasswdResponse{}
	err := client.Request("drds", "ModifyDrdsDBPasswd", "2017-10-16", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifyDrdsDBPasswdResponse struct {
	RequestId goaliyun.String
	Success   bool
}
