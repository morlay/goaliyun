package drds

import (
	"github.com/morlay/goaliyun"
)

type DescribeCreateDrdsInstanceStatusRequest struct {
	DrdsInstanceId string `position:"Query" name:"DrdsInstanceId"`
	RegionId       string `position:"Query" name:"RegionId"`
}

func (req *DescribeCreateDrdsInstanceStatusRequest) Invoke(client goaliyun.Client) (*DescribeCreateDrdsInstanceStatusResponse, error) {
	resp := &DescribeCreateDrdsInstanceStatusResponse{}
	err := client.Request("drds", "DescribeCreateDrdsInstanceStatus", "2017-10-16", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeCreateDrdsInstanceStatusResponse struct {
	RequestId goaliyun.String
	Success   bool
	Data      DescribeCreateDrdsInstanceStatusData
}

type DescribeCreateDrdsInstanceStatusData struct {
	Status goaliyun.String
}
