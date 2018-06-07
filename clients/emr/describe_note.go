package emr

import (
	"github.com/morlay/goaliyun"
)

type DescribeNoteRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	Id              string `position:"Query" name:"Id"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *DescribeNoteRequest) Invoke(client goaliyun.Client) (*DescribeNoteResponse, error) {
	resp := &DescribeNoteResponse{}
	err := client.Request("emr", "DescribeNote", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeNoteResponse struct {
	RequestId  goaliyun.String
	Id         goaliyun.String
	Name       goaliyun.String
	ClusterId  goaliyun.String
	Type       goaliyun.String
	Paragraphs goaliyun.String
}
