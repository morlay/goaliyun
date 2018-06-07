package jarvis

import (
	"github.com/morlay/goaliyun"
)

type CreateCpmcPunishFeedBackRequest struct {
	FeedBack     int64  `position:"Query" name:"FeedBack"`
	SrcIP        string `position:"Query" name:"SrcIP"`
	SourceIp     string `position:"Query" name:"SourceIp"`
	DstPort      int64  `position:"Query" name:"DstPort"`
	ProtocolName string `position:"Query" name:"ProtocolName"`
	SrcPort      int64  `position:"Query" name:"SrcPort"`
	PunishType   string `position:"Query" name:"PunishType"`
	GmtCreate    string `position:"Query" name:"GmtCreate"`
	DstIP        string `position:"Query" name:"DstIP"`
	Lang         string `position:"Query" name:"Lang"`
	SourceCode   string `position:"Query" name:"SourceCode"`
	RegionId     string `position:"Query" name:"RegionId"`
}

func (req *CreateCpmcPunishFeedBackRequest) Invoke(client goaliyun.Client) (*CreateCpmcPunishFeedBackResponse, error) {
	resp := &CreateCpmcPunishFeedBackResponse{}
	err := client.Request("jarvis", "CreateCpmcPunishFeedBack", "2018-02-06", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CreateCpmcPunishFeedBackResponse struct {
	RequestId goaliyun.String
	Module    goaliyun.String
}
