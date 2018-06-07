package imm

import (
	"github.com/morlay/goaliyun"
)

type DeleteConvertOfficeFormatTaskRequest struct {
	Project  string `position:"Query" name:"Project"`
	TaskId   string `position:"Query" name:"TaskId"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *DeleteConvertOfficeFormatTaskRequest) Invoke(client goaliyun.Client) (*DeleteConvertOfficeFormatTaskResponse, error) {
	resp := &DeleteConvertOfficeFormatTaskResponse{}
	err := client.Request("imm", "DeleteConvertOfficeFormatTask", "2017-09-06", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteConvertOfficeFormatTaskResponse struct {
	RequestId goaliyun.String
}
