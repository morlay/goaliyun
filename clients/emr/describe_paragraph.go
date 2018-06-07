package emr

import (
	"github.com/morlay/goaliyun"
)

type DescribeParagraphRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	NoteId          string `position:"Query" name:"NoteId"`
	Id              string `position:"Query" name:"Id"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *DescribeParagraphRequest) Invoke(client goaliyun.Client) (*DescribeParagraphResponse, error) {
	resp := &DescribeParagraphResponse{}
	err := client.Request("emr", "DescribeParagraph", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeParagraphResponse struct {
	RequestId goaliyun.String
	Paragraph goaliyun.String
}
