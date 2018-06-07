package rds

import (
	"github.com/morlay/goaliyun"
)

type CreateSQLDiagnosisRequest struct {
	EndTime      string `position:"Query" name:"EndTime"`
	DBInstanceId string `position:"Query" name:"DBInstanceId"`
	StartTime    string `position:"Query" name:"StartTime"`
	RegionId     string `position:"Query" name:"RegionId"`
}

func (req *CreateSQLDiagnosisRequest) Invoke(client goaliyun.Client) (*CreateSQLDiagnosisResponse, error) {
	resp := &CreateSQLDiagnosisResponse{}
	err := client.Request("rds", "CreateSQLDiagnosis", "2014-08-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CreateSQLDiagnosisResponse struct {
	RequestId goaliyun.String
	SQLDiagId goaliyun.String
}
