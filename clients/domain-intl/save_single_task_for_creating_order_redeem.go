package domain_intl

import (
	"github.com/morlay/goaliyun"
)

type SaveSingleTaskForCreatingOrderRedeemRequest struct {
	CurrentExpirationDate int64  `position:"Query" name:"CurrentExpirationDate"`
	UserClientIp          string `position:"Query" name:"UserClientIp"`
	DomainName            string `position:"Query" name:"DomainName"`
	Lang                  string `position:"Query" name:"Lang"`
	RegionId              string `position:"Query" name:"RegionId"`
}

func (req *SaveSingleTaskForCreatingOrderRedeemRequest) Invoke(client goaliyun.Client) (*SaveSingleTaskForCreatingOrderRedeemResponse, error) {
	resp := &SaveSingleTaskForCreatingOrderRedeemResponse{}
	err := client.Request("domain-intl", "SaveSingleTaskForCreatingOrderRedeem", "2017-12-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SaveSingleTaskForCreatingOrderRedeemResponse struct {
	RequestId goaliyun.String
	TaskNo    goaliyun.String
}
