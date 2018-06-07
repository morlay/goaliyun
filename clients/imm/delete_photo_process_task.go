package imm

import (
	"github.com/morlay/goaliyun"
)

type DeletePhotoProcessTaskRequest struct {
	Project  string `position:"Query" name:"Project"`
	TaskId   string `position:"Query" name:"TaskId"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *DeletePhotoProcessTaskRequest) Invoke(client goaliyun.Client) (*DeletePhotoProcessTaskResponse, error) {
	resp := &DeletePhotoProcessTaskResponse{}
	err := client.Request("imm", "DeletePhotoProcessTask", "2017-09-06", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeletePhotoProcessTaskResponse struct {
	RequestId goaliyun.String
}
