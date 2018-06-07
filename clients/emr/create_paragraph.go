package emr

import (
	"github.com/morlay/goaliyun"
)

type CreateParagraphRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	NoteId          string `position:"Query" name:"NoteId"`
	Text            string `position:"Query" name:"Text"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *CreateParagraphRequest) Invoke(client goaliyun.Client) (*CreateParagraphResponse, error) {
	resp := &CreateParagraphResponse{}
	err := client.Request("emr", "CreateParagraph", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CreateParagraphResponse struct {
	RequestId goaliyun.String
	Id        goaliyun.String
}
