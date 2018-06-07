package rds

import (
	"github.com/morlay/goaliyun"
)

type LoginDBInstancefromCloudDBARequest struct {
	ServiceRequestParam string `position:"Query" name:"ServiceRequestParam"`
	DBInstanceId        string `position:"Query" name:"DBInstanceId"`
	ServiceRequestType  string `position:"Query" name:"ServiceRequestType"`
	RegionId            string `position:"Query" name:"RegionId"`
}

func (req *LoginDBInstancefromCloudDBARequest) Invoke(client goaliyun.Client) (*LoginDBInstancefromCloudDBAResponse, error) {
	resp := &LoginDBInstancefromCloudDBAResponse{}
	err := client.Request("rds", "LoginDBInstancefromCloudDBA", "2014-08-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type LoginDBInstancefromCloudDBAResponse struct {
	RequestId goaliyun.String
	ListData  goaliyun.String
	AttrData  goaliyun.String
}
