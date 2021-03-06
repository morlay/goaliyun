package domain

import (
	"github.com/morlay/goaliyun"
)

type SaveSingleTaskForSynchronizingDnsHostRequest struct {
	InstanceId string `position:"Query" name:"InstanceId"`
	Lang       string `position:"Query" name:"Lang"`
	RegionId   string `position:"Query" name:"RegionId"`
}

func (req *SaveSingleTaskForSynchronizingDnsHostRequest) Invoke(client goaliyun.Client) (*SaveSingleTaskForSynchronizingDnsHostResponse, error) {
	resp := &SaveSingleTaskForSynchronizingDnsHostResponse{}
	err := client.Request("domain", "SaveSingleTaskForSynchronizingDnsHost", "2018-01-29", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SaveSingleTaskForSynchronizingDnsHostResponse struct {
	RequestId goaliyun.String
	TaskNo    goaliyun.String
}
