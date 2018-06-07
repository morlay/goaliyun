package rds

import (
	"github.com/morlay/goaliyun"
)

type DescribeCloudDBAServiceRequest struct {
	ServiceRequestParam string `position:"Query" name:"ServiceRequestParam"`
	DBInstanceId        string `position:"Query" name:"DBInstanceId"`
	ServiceRequestType  string `position:"Query" name:"ServiceRequestType"`
	RegionId            string `position:"Query" name:"RegionId"`
}

func (req *DescribeCloudDBAServiceRequest) Invoke(client goaliyun.Client) (*DescribeCloudDBAServiceResponse, error) {
	resp := &DescribeCloudDBAServiceResponse{}
	err := client.Request("rds", "DescribeCloudDBAService", "2014-08-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeCloudDBAServiceResponse struct {
	RequestId goaliyun.String
	ListData  goaliyun.String
	AttrData  goaliyun.String
}
