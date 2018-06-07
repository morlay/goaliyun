package ehpc

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ListImagesRequest struct {
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *ListImagesRequest) Invoke(client goaliyun.Client) (*ListImagesResponse, error) {
	resp := &ListImagesResponse{}
	err := client.Request("ehpc", "ListImages", "2018-04-12", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListImagesResponse struct {
	RequestId goaliyun.String
	OsTags    ListImagesOsInfoList
}

type ListImagesOsInfo struct {
	OsTag        goaliyun.String
	Platform     goaliyun.String
	Version      goaliyun.String
	Architecture goaliyun.String
}

type ListImagesOsInfoList []ListImagesOsInfo

func (list *ListImagesOsInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListImagesOsInfo)
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
