package emr

import (
	"github.com/morlay/goaliyun"
)

type GetOpsCommandResultRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	EndCursor       int64  `position:"Query" name:"EndCursor"`
	StartCursor     int64  `position:"Query" name:"StartCursor"`
	ClusterId       string `position:"Query" name:"ClusterId"`
	TaskId          int64  `position:"Query" name:"TaskId"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *GetOpsCommandResultRequest) Invoke(client goaliyun.Client) (*GetOpsCommandResultResponse, error) {
	resp := &GetOpsCommandResultResponse{}
	err := client.Request("emr", "GetOpsCommandResult", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type GetOpsCommandResultResponse struct {
	RequestId goaliyun.String
	Content   goaliyun.String
	Cursor    goaliyun.Integer
	Finished  bool
}
