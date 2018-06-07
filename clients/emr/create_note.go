package emr

import (
	"github.com/morlay/goaliyun"
)

type CreateNoteRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	Name            string `position:"Query" name:"Name"`
	ClusterId       string `position:"Query" name:"ClusterId"`
	Type            string `position:"Query" name:"Type"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *CreateNoteRequest) Invoke(client goaliyun.Client) (*CreateNoteResponse, error) {
	resp := &CreateNoteResponse{}
	err := client.Request("emr", "CreateNote", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CreateNoteResponse struct {
	RequestId goaliyun.String
	Id        goaliyun.String
	Paragraph goaliyun.String
}
