package emr

import (
	"github.com/morlay/goaliyun"
)

type RunNoteParagraphsRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	NoteId          string `position:"Query" name:"NoteId"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *RunNoteParagraphsRequest) Invoke(client goaliyun.Client) (*RunNoteParagraphsResponse, error) {
	resp := &RunNoteParagraphsResponse{}
	err := client.Request("emr", "RunNoteParagraphs", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type RunNoteParagraphsResponse struct {
	RequestId goaliyun.String
}
