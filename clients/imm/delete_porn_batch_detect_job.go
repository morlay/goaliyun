package imm

import (
	"github.com/morlay/goaliyun"
)

type DeletePornBatchDetectJobRequest struct {
	JobId    string `position:"Query" name:"JobId"`
	Project  string `position:"Query" name:"Project"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *DeletePornBatchDetectJobRequest) Invoke(client goaliyun.Client) (*DeletePornBatchDetectJobResponse, error) {
	resp := &DeletePornBatchDetectJobResponse{}
	err := client.Request("imm", "DeletePornBatchDetectJob", "2017-09-06", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeletePornBatchDetectJobResponse struct {
	RequestId goaliyun.String
}
