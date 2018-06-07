package jarvis

import (
	"github.com/morlay/goaliyun"
)

type CreateUidWhiteListGroupRequest struct {
	Note            string `position:"Query" name:"Note"`
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	SourceIp        string `position:"Query" name:"SourceIp"`
	DstPort         int64  `position:"Query" name:"DstPort"`
	InstanceIdList  string `position:"Query" name:"InstanceIdList"`
	LiveTime        int64  `position:"Query" name:"LiveTime"`
	ProductName     string `position:"Query" name:"ProductName"`
	WhiteListType   int64  `position:"Query" name:"WhiteListType"`
	Lang            string `position:"Query" name:"Lang"`
	SrcUid          string `position:"Query" name:"SrcUid"`
	SourceCode      string `position:"Query" name:"SourceCode"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *CreateUidWhiteListGroupRequest) Invoke(client goaliyun.Client) (*CreateUidWhiteListGroupResponse, error) {
	resp := &CreateUidWhiteListGroupResponse{}
	err := client.Request("jarvis", "CreateUidWhiteListGroup", "2018-02-06", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CreateUidWhiteListGroupResponse struct {
	RequestId goaliyun.String
	Module    goaliyun.String
}
