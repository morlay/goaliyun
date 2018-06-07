package emr

import (
	"github.com/morlay/goaliyun"
)

type ReleaseClusterForInternalRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	ClusterId       string `position:"Query" name:"ClusterId"`
	UserId          string `position:"Query" name:"UserId"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *ReleaseClusterForInternalRequest) Invoke(client goaliyun.Client) (*ReleaseClusterForInternalResponse, error) {
	resp := &ReleaseClusterForInternalResponse{}
	err := client.Request("emr", "ReleaseClusterForInternal", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ReleaseClusterForInternalResponse struct {
	RequestId goaliyun.String
}
