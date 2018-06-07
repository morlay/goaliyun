package rds

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeDiagnosticReportListRequest struct {
	DBInstanceId string `position:"Query" name:"DBInstanceId"`
	RegionId     string `position:"Query" name:"RegionId"`
}

func (req *DescribeDiagnosticReportListRequest) Invoke(client goaliyun.Client) (*DescribeDiagnosticReportListResponse, error) {
	resp := &DescribeDiagnosticReportListResponse{}
	err := client.Request("rds", "DescribeDiagnosticReportList", "2014-08-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeDiagnosticReportListResponse struct {
	RequestId  goaliyun.String
	ReportList DescribeDiagnosticReportListReportList
}

type DescribeDiagnosticReportListReport struct {
	DiagnosticTime goaliyun.String
	Score          goaliyun.Integer
	StartTime      goaliyun.String
	EndTime        goaliyun.String
	DownloadURL    goaliyun.String
}

type DescribeDiagnosticReportListReportList []DescribeDiagnosticReportListReport

func (list *DescribeDiagnosticReportListReportList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDiagnosticReportListReport)
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
