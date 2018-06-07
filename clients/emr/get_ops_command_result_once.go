package emr

import (
	"github.com/morlay/goaliyun"
)

type GetOpsCommandResultOnceRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	ClusterId       string `position:"Query" name:"ClusterId"`
	TaskId          int64  `position:"Query" name:"TaskId"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *GetOpsCommandResultOnceRequest) Invoke(client goaliyun.Client) (*GetOpsCommandResultOnceResponse, error) {
	resp := &GetOpsCommandResultOnceResponse{}
	err := client.Request("emr", "GetOpsCommandResultOnce", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type GetOpsCommandResultOnceResponse struct {
	RequestId goaliyun.String
	Content   goaliyun.String
}
