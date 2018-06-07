package rds

import (
	"github.com/morlay/goaliyun"
)

type RequestServiceOfCloudDBARequest struct {
	ServiceRequestParam string `position:"Query" name:"ServiceRequestParam"`
	DBInstanceId        string `position:"Query" name:"DBInstanceId"`
	ServiceRequestType  string `position:"Query" name:"ServiceRequestType"`
	RegionId            string `position:"Query" name:"RegionId"`
}

func (req *RequestServiceOfCloudDBARequest) Invoke(client goaliyun.Client) (*RequestServiceOfCloudDBAResponse, error) {
	resp := &RequestServiceOfCloudDBAResponse{}
	err := client.Request("rds", "RequestServiceOfCloudDBA", "2014-08-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type RequestServiceOfCloudDBAResponse struct {
	RequestId goaliyun.String
	ListData  goaliyun.String
	AttrData  goaliyun.String
}
