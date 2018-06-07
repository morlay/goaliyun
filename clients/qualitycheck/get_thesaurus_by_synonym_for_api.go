package qualitycheck

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type GetThesaurusBySynonymForApiRequest struct {
	JsonStr  string `position:"Query" name:"JsonStr"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *GetThesaurusBySynonymForApiRequest) Invoke(client goaliyun.Client) (*GetThesaurusBySynonymForApiResponse, error) {
	resp := &GetThesaurusBySynonymForApiResponse{}
	err := client.Request("qualitycheck", "GetThesaurusBySynonymForApi", "2016-08-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type GetThesaurusBySynonymForApiResponse struct {
	RequestId goaliyun.String
	Success   bool
	Code      goaliyun.String
	Message   goaliyun.String
	Data      GetThesaurusBySynonymForApiThesaurusPoList
}

type GetThesaurusBySynonymForApiThesaurusPo struct {
	Id          goaliyun.Integer
	Business    goaliyun.String
	SynonymList GetThesaurusBySynonymForApiSynonymListList
}

type GetThesaurusBySynonymForApiThesaurusPoList []GetThesaurusBySynonymForApiThesaurusPo

func (list *GetThesaurusBySynonymForApiThesaurusPoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetThesaurusBySynonymForApiThesaurusPo)
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

type GetThesaurusBySynonymForApiSynonymListList []goaliyun.String

func (list *GetThesaurusBySynonymForApiSynonymListList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]goaliyun.String)
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
