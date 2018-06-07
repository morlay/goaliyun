package emr

import (
	"github.com/morlay/goaliyun"
)

type DeleteParagraphRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	NoteId          string `position:"Query" name:"NoteId"`
	Id              string `position:"Query" name:"Id"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *DeleteParagraphRequest) Invoke(client goaliyun.Client) (*DeleteParagraphResponse, error) {
	resp := &DeleteParagraphResponse{}
	err := client.Request("emr", "DeleteParagraph", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteParagraphResponse struct {
	RequestId goaliyun.String
}
