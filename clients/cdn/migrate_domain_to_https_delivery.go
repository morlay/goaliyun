package cdn

import (
	"github.com/morlay/goaliyun"
)

type MigrateDomainToHttpsDeliveryRequest struct {
	PrivateKey        string `position:"Query" name:"PrivateKey"`
	ServerCertificate string `position:"Query" name:"ServerCertificate"`
	SecurityToken     string `position:"Query" name:"SecurityToken"`
	OwnerAccount      string `position:"Query" name:"OwnerAccount"`
	DomainName        string `position:"Query" name:"DomainName"`
	OwnerId           int64  `position:"Query" name:"OwnerId"`
	RegionId          string `position:"Query" name:"RegionId"`
}

func (req *MigrateDomainToHttpsDeliveryRequest) Invoke(client goaliyun.Client) (*MigrateDomainToHttpsDeliveryResponse, error) {
	resp := &MigrateDomainToHttpsDeliveryResponse{}
	err := client.Request("cdn", "MigrateDomainToHttpsDelivery", "2014-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type MigrateDomainToHttpsDeliveryResponse struct {
	RequestId goaliyun.String
}
