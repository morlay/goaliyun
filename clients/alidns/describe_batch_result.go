package alidns

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeBatchResultRequest struct {
	Lang         string `position:"Query" name:"Lang"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	TraceId      string `position:"Query" name:"TraceId"`
	RegionId     string `position:"Query" name:"RegionId"`
}

func (req *DescribeBatchResultRequest) Invoke(client goaliyun.Client) (*DescribeBatchResultResponse, error) {
	resp := &DescribeBatchResultResponse{}
	err := client.Request("alidns", "DescribeBatchResult", "2015-01-09", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeBatchResultResponse struct {
	RequestId     goaliyun.String
	TraceId       goaliyun.String
	Status        goaliyun.Integer
	BatchCount    goaliyun.Integer
	SuccessNumber goaliyun.Integer
	FailResults   DescribeBatchResultFailResultList
}

type DescribeBatchResultFailResult struct {
	BatchIndex goaliyun.String
	ErrorCode  goaliyun.String
}

type DescribeBatchResultFailResultList []DescribeBatchResultFailResult

func (list *DescribeBatchResultFailResultList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeBatchResultFailResult)
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
