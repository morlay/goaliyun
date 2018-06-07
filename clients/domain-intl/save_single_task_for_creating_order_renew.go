package domain_intl

import (
	"github.com/morlay/goaliyun"
)

type SaveSingleTaskForCreatingOrderRenewRequest struct {
	SubscriptionDuration  int64  `position:"Query" name:"SubscriptionDuration"`
	CurrentExpirationDate int64  `position:"Query" name:"CurrentExpirationDate"`
	UserClientIp          string `position:"Query" name:"UserClientIp"`
	DomainName            string `position:"Query" name:"DomainName"`
	Lang                  string `position:"Query" name:"Lang"`
	RegionId              string `position:"Query" name:"RegionId"`
}

func (req *SaveSingleTaskForCreatingOrderRenewRequest) Invoke(client goaliyun.Client) (*SaveSingleTaskForCreatingOrderRenewResponse, error) {
	resp := &SaveSingleTaskForCreatingOrderRenewResponse{}
	err := client.Request("domain-intl", "SaveSingleTaskForCreatingOrderRenew", "2017-12-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SaveSingleTaskForCreatingOrderRenewResponse struct {
	RequestId goaliyun.String
	TaskNo    goaliyun.String
}
