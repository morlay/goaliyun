package vpc

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type RemoveBandwidthPackageIpsRequest struct {
	RemovedIpAddressess  *RemoveBandwidthPackageIpsRemovedIpAddressesList `position:"Query" type:"Repeated" name:"RemovedIpAddresses"`
	ResourceOwnerId      int64                                            `position:"Query" name:"ResourceOwnerId"`
	BandwidthPackageId   string                                           `position:"Query" name:"BandwidthPackageId"`
	ResourceOwnerAccount string                                           `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string                                           `position:"Query" name:"ClientToken"`
	OwnerAccount         string                                           `position:"Query" name:"OwnerAccount"`
	OwnerId              int64                                            `position:"Query" name:"OwnerId"`
	RegionId             string                                           `position:"Query" name:"RegionId"`
}

func (req *RemoveBandwidthPackageIpsRequest) Invoke(client goaliyun.Client) (*RemoveBandwidthPackageIpsResponse, error) {
	resp := &RemoveBandwidthPackageIpsResponse{}
	err := client.Request("vpc", "RemoveBandwidthPackageIps", "2016-04-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type RemoveBandwidthPackageIpsResponse struct {
	RequestId goaliyun.String
}

type RemoveBandwidthPackageIpsRemovedIpAddressesList []string

func (list *RemoveBandwidthPackageIpsRemovedIpAddressesList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]string)
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	for _, v := range m {
		*list = v
		break
	}
	return nil
}
