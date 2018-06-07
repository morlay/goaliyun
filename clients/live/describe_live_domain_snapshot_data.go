package live

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeLiveDomainSnapshotDataRequest struct {
	DomainName string `position:"Query" name:"DomainName"`
	EndTime    string `position:"Query" name:"EndTime"`
	StartTime  string `position:"Query" name:"StartTime"`
	OwnerId    int64  `position:"Query" name:"OwnerId"`
	RegionId   string `position:"Query" name:"RegionId"`
}

func (req *DescribeLiveDomainSnapshotDataRequest) Invoke(client goaliyun.Client) (*DescribeLiveDomainSnapshotDataResponse, error) {
	resp := &DescribeLiveDomainSnapshotDataResponse{}
	err := client.Request("live", "DescribeLiveDomainSnapshotData", "2016-11-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeLiveDomainSnapshotDataResponse struct {
	RequestId         goaliyun.String
	SnapshotDataInfos DescribeLiveDomainSnapshotDataSnapshotDataInfoList
}

type DescribeLiveDomainSnapshotDataSnapshotDataInfo struct {
	Date  goaliyun.String
	Total goaliyun.Integer
}

type DescribeLiveDomainSnapshotDataSnapshotDataInfoList []DescribeLiveDomainSnapshotDataSnapshotDataInfo

func (list *DescribeLiveDomainSnapshotDataSnapshotDataInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeLiveDomainSnapshotDataSnapshotDataInfo)
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
