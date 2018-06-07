package teslamaxcompute

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type GetQuotaHistoryInfoRequest struct {
	Cluster   string `position:"Query" name:"Cluster"`
	EndTime   int64  `position:"Query" name:"EndTime"`
	StartTime int64  `position:"Query" name:"StartTime"`
	Region    string `position:"Query" name:"Region"`
	QuotaName string `position:"Query" name:"QuotaName"`
	RegionId  string `position:"Query" name:"RegionId"`
}

func (req *GetQuotaHistoryInfoRequest) Invoke(client goaliyun.Client) (*GetQuotaHistoryInfoResponse, error) {
	resp := &GetQuotaHistoryInfoResponse{}
	err := client.Request("teslamaxcompute", "GetQuotaHistoryInfo", "2018-01-04", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type GetQuotaHistoryInfoResponse struct {
	Code      goaliyun.Integer
	Message   goaliyun.String
	RequestId goaliyun.String
	Data      GetQuotaHistoryInfoDataItemList
}

type GetQuotaHistoryInfoDataItem struct {
	Times goaliyun.Integer
	Point GetQuotaHistoryInfoPoint
}

type GetQuotaHistoryInfoPoint struct {
	CpuMaxQuota GetQuotaHistoryInfoCpuMaxQuota
	CpuMinQuota GetQuotaHistoryInfoCpuMinQuota
	MemUsed     GetQuotaHistoryInfoMemUsed
	CpuUsed     GetQuotaHistoryInfoCpuUsed
	MemMaxQuota GetQuotaHistoryInfoMemMaxQuota
	MemMinQuota GetQuotaHistoryInfoMemMinQuota
}

type GetQuotaHistoryInfoCpuMaxQuota struct {
	Min goaliyun.Float
	Max goaliyun.Float
	Avg goaliyun.Float
}

type GetQuotaHistoryInfoCpuMinQuota struct {
	Min goaliyun.Float
	Max goaliyun.Float
	Avg goaliyun.Float
}

type GetQuotaHistoryInfoMemUsed struct {
	Min goaliyun.Float
	Max goaliyun.Float
	Avg goaliyun.Float
}

type GetQuotaHistoryInfoCpuUsed struct {
	Min goaliyun.Float
	Max goaliyun.Float
	Avg goaliyun.Float
}

type GetQuotaHistoryInfoMemMaxQuota struct {
	Min goaliyun.Float
	Max goaliyun.Float
	Avg goaliyun.Float
}

type GetQuotaHistoryInfoMemMinQuota struct {
	Min goaliyun.Float
	Max goaliyun.Float
	Avg goaliyun.Float
}

type GetQuotaHistoryInfoDataItemList []GetQuotaHistoryInfoDataItem

func (list *GetQuotaHistoryInfoDataItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetQuotaHistoryInfoDataItem)
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
