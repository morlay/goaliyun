package pvtz

import (
	"github.com/morlay/goaliyun"
)

type DescribeUserServiceStatusRequest struct {
	UserClientIp string `position:"Query" name:"UserClientIp"`
	Lang         string `position:"Query" name:"Lang"`
	RegionId     string `position:"Query" name:"RegionId"`
}

func (req *DescribeUserServiceStatusRequest) Invoke(client goaliyun.Client) (*DescribeUserServiceStatusResponse, error) {
	resp := &DescribeUserServiceStatusResponse{}
	err := client.Request("pvtz", "DescribeUserServiceStatus", "2018-01-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeUserServiceStatusResponse struct {
	RequestId goaliyun.String
	Status    goaliyun.String
}
