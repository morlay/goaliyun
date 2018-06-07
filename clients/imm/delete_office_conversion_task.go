package imm

import (
	"github.com/morlay/goaliyun"
)

type DeleteOfficeConversionTaskRequest struct {
	Project  string `position:"Query" name:"Project"`
	TaskId   string `position:"Query" name:"TaskId"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *DeleteOfficeConversionTaskRequest) Invoke(client goaliyun.Client) (*DeleteOfficeConversionTaskResponse, error) {
	resp := &DeleteOfficeConversionTaskResponse{}
	err := client.Request("imm", "DeleteOfficeConversionTask", "2017-09-06", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteOfficeConversionTaskResponse struct {
	RequestId goaliyun.String
}
