package rds

import (
	"github.com/morlay/goaliyun"
)

type CreateDiagnosticReportRequest struct {
	EndTime      string `position:"Query" name:"EndTime"`
	DBInstanceId string `position:"Query" name:"DBInstanceId"`
	StartTime    string `position:"Query" name:"StartTime"`
	RegionId     string `position:"Query" name:"RegionId"`
}

func (req *CreateDiagnosticReportRequest) Invoke(client goaliyun.Client) (*CreateDiagnosticReportResponse, error) {
	resp := &CreateDiagnosticReportResponse{}
	err := client.Request("rds", "CreateDiagnosticReport", "2014-08-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CreateDiagnosticReportResponse struct {
	RequestId goaliyun.String
	ReportId  goaliyun.String
}
