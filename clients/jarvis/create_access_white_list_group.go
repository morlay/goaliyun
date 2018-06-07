package jarvis

import (
	"github.com/morlay/goaliyun"
)

type CreateAccessWhiteListGroupRequest struct {
	Note            string `position:"Query" name:"Note"`
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	SrcIP           string `position:"Query" name:"SrcIP"`
	SourceIp        string `position:"Query" name:"SourceIp"`
	DstPort         int64  `position:"Query" name:"DstPort"`
	InstanceIdList  string `position:"Query" name:"InstanceIdList"`
	LiveTime        int64  `position:"Query" name:"LiveTime"`
	ProductName     string `position:"Query" name:"ProductName"`
	WhiteListType   int64  `position:"Query" name:"WhiteListType"`
	Lang            string `position:"Query" name:"Lang"`
	SourceCode      string `position:"Query" name:"SourceCode"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *CreateAccessWhiteListGroupRequest) Invoke(client goaliyun.Client) (*CreateAccessWhiteListGroupResponse, error) {
	resp := &CreateAccessWhiteListGroupResponse{}
	err := client.Request("jarvis", "CreateAccessWhiteListGroup", "2018-02-06", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CreateAccessWhiteListGroupResponse struct {
	RequestId goaliyun.String
	Module    goaliyun.String
}
