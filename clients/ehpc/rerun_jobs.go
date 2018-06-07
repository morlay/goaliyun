package ehpc

import (
	"github.com/morlay/goaliyun"
)

type RerunJobsRequest struct {
	Jobs      string `position:"Query" name:"Jobs"`
	ClusterId string `position:"Query" name:"ClusterId"`
	RegionId  string `position:"Query" name:"RegionId"`
}

func (req *RerunJobsRequest) Invoke(client goaliyun.Client) (*RerunJobsResponse, error) {
	resp := &RerunJobsResponse{}
	err := client.Request("ehpc", "RerunJobs", "2018-04-12", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type RerunJobsResponse struct {
	RequestId goaliyun.String
}
