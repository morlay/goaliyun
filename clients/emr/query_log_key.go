package emr

import (
	"github.com/morlay/goaliyun"
)

type QueryLogKeyRequest struct {
	JobId           string `position:"Query" name:"JobId"`
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	InstanceId      string `position:"Query" name:"InstanceId"`
	LogName         string `position:"Query" name:"LogName"`
	ClusterId       string `position:"Query" name:"ClusterId"`
	KeyBase         string `position:"Query" name:"KeyBase"`
	ContainerId     string `position:"Query" name:"ContainerId"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *QueryLogKeyRequest) Invoke(client goaliyun.Client) (*QueryLogKeyResponse, error) {
	resp := &QueryLogKeyResponse{}
	err := client.Request("emr", "QueryLogKey", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type QueryLogKeyResponse struct {
	RequestId goaliyun.String
	LogKey    goaliyun.String
}
