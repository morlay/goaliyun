package emr

import (
	"github.com/morlay/goaliyun"
)

type DetachClusterForNoteRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	Id              string `position:"Query" name:"Id"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *DetachClusterForNoteRequest) Invoke(client goaliyun.Client) (*DetachClusterForNoteResponse, error) {
	resp := &DetachClusterForNoteResponse{}
	err := client.Request("emr", "DetachClusterForNote", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DetachClusterForNoteResponse struct {
	RequestId goaliyun.String
}
