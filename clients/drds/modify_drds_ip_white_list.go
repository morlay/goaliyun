package drds

import (
	"github.com/morlay/goaliyun"
)

type ModifyDrdsIpWhiteListRequest struct {
	Mode           string `position:"Query" name:"Mode"`
	DbName         string `position:"Query" name:"DbName"`
	GroupAttribute string `position:"Query" name:"GroupAttribute"`
	IpWhiteList    string `position:"Query" name:"IpWhiteList"`
	DrdsInstanceId string `position:"Query" name:"DrdsInstanceId"`
	GroupName      string `position:"Query" name:"GroupName"`
	RegionId       string `position:"Query" name:"RegionId"`
}

func (req *ModifyDrdsIpWhiteListRequest) Invoke(client goaliyun.Client) (*ModifyDrdsIpWhiteListResponse, error) {
	resp := &ModifyDrdsIpWhiteListResponse{}
	err := client.Request("drds", "ModifyDrdsIpWhiteList", "2017-10-16", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifyDrdsIpWhiteListResponse struct {
	RequestId goaliyun.String
	Success   bool
}
