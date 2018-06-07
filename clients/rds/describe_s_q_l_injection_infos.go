package rds

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeSQLInjectionInfosRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	PageSize             int64  `position:"Query" name:"PageSize"`
	EndTime              string `position:"Query" name:"EndTime"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	StartTime            string `position:"Query" name:"StartTime"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PageNumber           int64  `position:"Query" name:"PageNumber"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeSQLInjectionInfosRequest) Invoke(client goaliyun.Client) (*DescribeSQLInjectionInfosResponse, error) {
	resp := &DescribeSQLInjectionInfosResponse{}
	err := client.Request("rds", "DescribeSQLInjectionInfos", "2014-08-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeSQLInjectionInfosResponse struct {
	RequestId        goaliyun.String
	Engine           goaliyun.String
	TotalRecordCount goaliyun.Integer
	PageNumber       goaliyun.Integer
	PageRecordCount  goaliyun.Integer
	Items            DescribeSQLInjectionInfosSQLInjectionInfoList
}

type DescribeSQLInjectionInfosSQLInjectionInfo struct {
	DBName         goaliyun.String
	SQLText        goaliyun.String
	LatencyTime    goaliyun.String
	HostAddress    goaliyun.String
	ExecuteTime    goaliyun.String
	AccountName    goaliyun.String
	EffectRowCount goaliyun.String
}

type DescribeSQLInjectionInfosSQLInjectionInfoList []DescribeSQLInjectionInfosSQLInjectionInfo

func (list *DescribeSQLInjectionInfosSQLInjectionInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeSQLInjectionInfosSQLInjectionInfo)
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
