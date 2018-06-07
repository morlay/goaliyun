package jarvis

import (
	"github.com/morlay/goaliyun"
)

type ModifyAccessWhiteListAutoShareRequest struct {
	SrcIP         string `position:"Query" name:"SrcIP"`
	SourceIp      string `position:"Query" name:"SourceIp"`
	AutoConfig    int64  `position:"Query" name:"AutoConfig"`
	ProductName   string `position:"Query" name:"ProductName"`
	WhiteListType int64  `position:"Query" name:"WhiteListType"`
	Lang          string `position:"Query" name:"Lang"`
	SourceCode    string `position:"Query" name:"SourceCode"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *ModifyAccessWhiteListAutoShareRequest) Invoke(client goaliyun.Client) (*ModifyAccessWhiteListAutoShareResponse, error) {
	resp := &ModifyAccessWhiteListAutoShareResponse{}
	err := client.Request("jarvis", "ModifyAccessWhiteListAutoShare", "2018-02-06", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifyAccessWhiteListAutoShareResponse struct {
	RequestId goaliyun.String
	Module    goaliyun.String
}
