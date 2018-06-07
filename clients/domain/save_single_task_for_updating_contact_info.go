package domain

import (
	"github.com/morlay/goaliyun"
)

type SaveSingleTaskForUpdatingContactInfoRequest struct {
	InstanceId          string `position:"Query" name:"InstanceId"`
	ContactType         string `position:"Query" name:"ContactType"`
	UserClientIp        string `position:"Query" name:"UserClientIp"`
	DomainName          string `position:"Query" name:"DomainName"`
	RegistrantProfileId int64  `position:"Query" name:"RegistrantProfileId"`
	AddTransferLock     string `position:"Query" name:"AddTransferLock"`
	Lang                string `position:"Query" name:"Lang"`
	RegionId            string `position:"Query" name:"RegionId"`
}

func (req *SaveSingleTaskForUpdatingContactInfoRequest) Invoke(client goaliyun.Client) (*SaveSingleTaskForUpdatingContactInfoResponse, error) {
	resp := &SaveSingleTaskForUpdatingContactInfoResponse{}
	err := client.Request("domain", "SaveSingleTaskForUpdatingContactInfo", "2018-01-29", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SaveSingleTaskForUpdatingContactInfoResponse struct {
	RequestId goaliyun.String
	TaskNo    goaliyun.String
}
