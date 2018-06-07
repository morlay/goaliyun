package emr

import (
	"github.com/morlay/goaliyun"
)

type SaveParagraphRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	NoteId          string `position:"Query" name:"NoteId"`
	Id              string `position:"Query" name:"Id"`
	Text            string `position:"Query" name:"Text"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *SaveParagraphRequest) Invoke(client goaliyun.Client) (*SaveParagraphResponse, error) {
	resp := &SaveParagraphResponse{}
	err := client.Request("emr", "SaveParagraph", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SaveParagraphResponse struct {
	RequestId goaliyun.String
}
