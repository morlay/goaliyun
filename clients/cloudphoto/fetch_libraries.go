package cloudphoto

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type FetchLibrariesRequest struct {
	Size      int64  `position:"Query" name:"Size"`
	StoreName string `position:"Query" name:"StoreName"`
	Page      int64  `position:"Query" name:"Page"`
	RegionId  string `position:"Query" name:"RegionId"`
}

func (req *FetchLibrariesRequest) Invoke(client goaliyun.Client) (*FetchLibrariesResponse, error) {
	resp := &FetchLibrariesResponse{}
	err := client.Request("cloudphoto", "FetchLibraries", "2017-07-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type FetchLibrariesResponse struct {
	Code       goaliyun.String
	Message    goaliyun.String
	TotalCount goaliyun.Integer
	RequestId  goaliyun.String
	Action     goaliyun.String
	Libraries  FetchLibrariesLibraryList
}

type FetchLibrariesLibrary struct {
	LibraryId goaliyun.String
	Ctime     goaliyun.Integer
}

type FetchLibrariesLibraryList []FetchLibrariesLibrary

func (list *FetchLibrariesLibraryList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]FetchLibrariesLibrary)
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
