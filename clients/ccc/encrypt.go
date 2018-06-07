package ccc

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type EncryptRequest struct {
	PublicKey  string                `position:"Query" name:"PublicKey"`
	PlainTexts *EncryptPlainTextList `position:"Query" type:"Repeated" name:"PlainText"`
	RegionId   string                `position:"Query" name:"RegionId"`
}

func (req *EncryptRequest) Invoke(client goaliyun.Client) (*EncryptResponse, error) {
	resp := &EncryptResponse{}
	err := client.Request("ccc", "Encrypt", "2017-07-05", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type EncryptResponse struct {
	RequestId      goaliyun.String
	Success        bool
	Code           goaliyun.String
	Message        goaliyun.String
	HttpStatusCode goaliyun.Integer
	PublicKey      goaliyun.String
	CypherContents EncryptCypherContentList
}

type EncryptCypherContent struct {
	PlainText  goaliyun.String
	CypherText goaliyun.String
}

type EncryptPlainTextList []string

func (list *EncryptPlainTextList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]string)
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

type EncryptCypherContentList []EncryptCypherContent

func (list *EncryptCypherContentList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]EncryptCypherContent)
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
