package ehpc

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ListVolumesRequest struct {
	PageSize   int64  `position:"Query" name:"PageSize"`
	PageNumber int64  `position:"Query" name:"PageNumber"`
	RegionId   string `position:"Query" name:"RegionId"`
}

func (req *ListVolumesRequest) Invoke(client goaliyun.Client) (*ListVolumesResponse, error) {
	resp := &ListVolumesResponse{}
	err := client.Request("ehpc", "ListVolumes", "2018-04-12", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListVolumesResponse struct {
	RequestId  goaliyun.String
	TotalCount goaliyun.Integer
	PageNumber goaliyun.Integer
	PageSize   goaliyun.Integer
	Volumes    ListVolumesVolumeInfoList
}

type ListVolumesVolumeInfo struct {
	RegionId         goaliyun.String
	ClusterId        goaliyun.String
	ClusterName      goaliyun.String
	VolumeId         goaliyun.String
	VolumeType       goaliyun.String
	VolumeProtocol   goaliyun.String
	VolumeMountpoint goaliyun.String
	RemoteDirectory  goaliyun.String
}

type ListVolumesVolumeInfoList []ListVolumesVolumeInfo

func (list *ListVolumesVolumeInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListVolumesVolumeInfo)
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
