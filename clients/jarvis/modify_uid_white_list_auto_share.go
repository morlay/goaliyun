package jarvis

import (
	"github.com/morlay/goaliyun"
)

type ModifyUidWhiteListAutoShareRequest struct {
	SourceIp      string `position:"Query" name:"SourceIp"`
	AutoConfig    int64  `position:"Query" name:"AutoConfig"`
	ProductName   string `position:"Query" name:"ProductName"`
	WhiteListType int64  `position:"Query" name:"WhiteListType"`
	Lang          string `position:"Query" name:"Lang"`
	SrcUid        string `position:"Query" name:"SrcUid"`
	SourceCode    string `position:"Query" name:"SourceCode"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *ModifyUidWhiteListAutoShareRequest) Invoke(client goaliyun.Client) (*ModifyUidWhiteListAutoShareResponse, error) {
	resp := &ModifyUidWhiteListAutoShareResponse{}
	err := client.Request("jarvis", "ModifyUidWhiteListAutoShare", "2018-02-06", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifyUidWhiteListAutoShareResponse struct {
	RequestId goaliyun.String
	Module    goaliyun.String
}
