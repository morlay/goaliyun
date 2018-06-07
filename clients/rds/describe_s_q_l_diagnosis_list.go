package rds

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeSQLDiagnosisListRequest struct {
	DBInstanceId string `position:"Query" name:"DBInstanceId"`
	RegionId     string `position:"Query" name:"RegionId"`
}

func (req *DescribeSQLDiagnosisListRequest) Invoke(client goaliyun.Client) (*DescribeSQLDiagnosisListResponse, error) {
	resp := &DescribeSQLDiagnosisListResponse{}
	err := client.Request("rds", "DescribeSQLDiagnosisList", "2014-08-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeSQLDiagnosisListResponse struct {
	RequestId   goaliyun.String
	SQLDiagList DescribeSQLDiagnosisListSQLDiagList
}

type DescribeSQLDiagnosisListSQLDiag struct {
	SQLDiagId goaliyun.String
	StartTime goaliyun.String
	EndTime   goaliyun.String
	Status    goaliyun.Integer
	Progress  goaliyun.Integer
}

type DescribeSQLDiagnosisListSQLDiagList []DescribeSQLDiagnosisListSQLDiag

func (list *DescribeSQLDiagnosisListSQLDiagList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeSQLDiagnosisListSQLDiag)
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
