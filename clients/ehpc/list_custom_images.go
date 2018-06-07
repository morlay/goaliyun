package ehpc

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ListCustomImagesRequest struct {
	BaseOsTag       string `position:"Query" name:"BaseOsTag"`
	ImageOwnerAlias string `position:"Query" name:"ImageOwnerAlias"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *ListCustomImagesRequest) Invoke(client goaliyun.Client) (*ListCustomImagesResponse, error) {
	resp := &ListCustomImagesResponse{}
	err := client.Request("ehpc", "ListCustomImages", "2018-04-12", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListCustomImagesResponse struct {
	RequestId goaliyun.String
	Images    ListCustomImagesImageInfoList
}

type ListCustomImagesImageInfo struct {
	Uid               goaliyun.String
	ImageId           goaliyun.String
	ImageName         goaliyun.String
	ImageOwnerAlias   goaliyun.String
	Description       goaliyun.String
	Status            goaliyun.String
	ProductCode       goaliyun.String
	SkuCode           goaliyun.String
	PricingCycle      goaliyun.String
	PostInstallScript goaliyun.String
	BaseOsTag         ListCustomImagesBaseOsTag
}

type ListCustomImagesBaseOsTag struct {
	OsTag        goaliyun.String
	Platform     goaliyun.String
	Version      goaliyun.String
	Architecture goaliyun.String
}

type ListCustomImagesImageInfoList []ListCustomImagesImageInfo

func (list *ListCustomImagesImageInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListCustomImagesImageInfo)
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
