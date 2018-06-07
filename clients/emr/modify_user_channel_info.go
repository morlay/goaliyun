package emr

import (
	"github.com/morlay/goaliyun"
)

type ModifyUserChannelInfoRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	ChannelId       string `position:"Query" name:"ChannelId"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *ModifyUserChannelInfoRequest) Invoke(client goaliyun.Client) (*ModifyUserChannelInfoResponse, error) {
	resp := &ModifyUserChannelInfoResponse{}
	err := client.Request("emr", "ModifyUserChannelInfo", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifyUserChannelInfoResponse struct {
	RequestId goaliyun.String
	Success   bool
}
