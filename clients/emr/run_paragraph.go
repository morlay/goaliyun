package emr

import (
	"github.com/morlay/goaliyun"
)

type RunParagraphRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	NoteId          string `position:"Query" name:"NoteId"`
	Id              string `position:"Query" name:"Id"`
	Text            string `position:"Query" name:"Text"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *RunParagraphRequest) Invoke(client goaliyun.Client) (*RunParagraphResponse, error) {
	resp := &RunParagraphResponse{}
	err := client.Request("emr", "RunParagraph", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type RunParagraphResponse struct {
	RequestId goaliyun.String
}
