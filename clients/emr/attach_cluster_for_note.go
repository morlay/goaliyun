package emr

import (
	"github.com/morlay/goaliyun"
)

type AttachClusterForNoteRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	Id              string `position:"Query" name:"Id"`
	ClusterId       string `position:"Query" name:"ClusterId"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *AttachClusterForNoteRequest) Invoke(client goaliyun.Client) (*AttachClusterForNoteResponse, error) {
	resp := &AttachClusterForNoteResponse{}
	err := client.Request("emr", "AttachClusterForNote", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type AttachClusterForNoteResponse struct {
	RequestId goaliyun.String
}
