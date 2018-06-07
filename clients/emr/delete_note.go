package emr

import (
	"github.com/morlay/goaliyun"
)

type DeleteNoteRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	Id              string `position:"Query" name:"Id"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *DeleteNoteRequest) Invoke(client goaliyun.Client) (*DeleteNoteResponse, error) {
	resp := &DeleteNoteResponse{}
	err := client.Request("emr", "DeleteNote", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteNoteResponse struct {
	RequestId goaliyun.String
}
