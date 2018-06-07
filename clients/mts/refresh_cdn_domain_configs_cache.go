package mts

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type RefreshCdnDomainConfigsCacheRequest struct {
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Domains              string `position:"Query" name:"Domains"`
	OwnerId              string `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *RefreshCdnDomainConfigsCacheRequest) Invoke(client goaliyun.Client) (*RefreshCdnDomainConfigsCacheResponse, error) {
	resp := &RefreshCdnDomainConfigsCacheResponse{}
	err := client.Request("mts", "RefreshCdnDomainConfigsCache", "2014-06-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type RefreshCdnDomainConfigsCacheResponse struct {
	RequestId     goaliyun.String
	SucessDomains RefreshCdnDomainConfigsCacheSucessDomainList
	FailedDomains RefreshCdnDomainConfigsCacheFailedDomainList
}

type RefreshCdnDomainConfigsCacheSucessDomainList []goaliyun.String

func (list *RefreshCdnDomainConfigsCacheSucessDomainList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]goaliyun.String)
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

type RefreshCdnDomainConfigsCacheFailedDomainList []goaliyun.String

func (list *RefreshCdnDomainConfigsCacheFailedDomainList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]goaliyun.String)
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
