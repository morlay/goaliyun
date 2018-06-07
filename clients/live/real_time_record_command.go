package live

import (
	"github.com/morlay/goaliyun"
)

type RealTimeRecordCommandRequest struct {
	AppName    string `position:"Query" name:"AppName"`
	DomainName string `position:"Query" name:"DomainName"`
	OwnerId    int64  `position:"Query" name:"OwnerId"`
	Command    string `position:"Query" name:"Command"`
	StreamName string `position:"Query" name:"StreamName"`
	RegionId   string `position:"Query" name:"RegionId"`
}

func (req *RealTimeRecordCommandRequest) Invoke(client goaliyun.Client) (*RealTimeRecordCommandResponse, error) {
	resp := &RealTimeRecordCommandResponse{}
	err := client.Request("live", "RealTimeRecordCommand", "2016-11-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type RealTimeRecordCommandResponse struct {
	RequestId goaliyun.String
}
