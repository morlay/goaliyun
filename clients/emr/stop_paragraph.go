package emr

import (
	"github.com/morlay/goaliyun"
)

type StopParagraphRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	NoteId          string `position:"Query" name:"NoteId"`
	Id              string `position:"Query" name:"Id"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *StopParagraphRequest) Invoke(client goaliyun.Client) (*StopParagraphResponse, error) {
	resp := &StopParagraphResponse{}
	err := client.Request("emr", "StopParagraph", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type StopParagraphResponse struct {
	RequestId goaliyun.String
}
