package imm

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type IndexTagRequest struct {
	SrcUris  string `position:"Query" name:"SrcUris"`
	Project  string `position:"Query" name:"Project"`
	SetId    string `position:"Query" name:"SetId"`
	Force    string `position:"Query" name:"Force"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *IndexTagRequest) Invoke(client goaliyun.Client) (*IndexTagResponse, error) {
	resp := &IndexTagResponse{}
	err := client.Request("imm", "IndexTag", "2017-09-06", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type IndexTagResponse struct {
	RequestId       goaliyun.String
	SetId           goaliyun.String
	SuccessIndexNum goaliyun.String
	FailDetails     IndexTagFailDetailsItemList
	SuccessDetails  IndexTagSuccessDetailsItemList
}

type IndexTagFailDetailsItem struct {
	SrcUri goaliyun.String
	Reason goaliyun.String
}

type IndexTagSuccessDetailsItem struct {
	SrcUri goaliyun.String
	Tags   IndexTagTagsItemList
}

type IndexTagTagsItem struct {
	TagId         goaliyun.String
	TagLevel      goaliyun.String
	TagName       goaliyun.String
	ParentTagId   goaliyun.String
	ParentTagName goaliyun.String
	TagScore      goaliyun.String
}

type IndexTagFailDetailsItemList []IndexTagFailDetailsItem

func (list *IndexTagFailDetailsItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]IndexTagFailDetailsItem)
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

type IndexTagSuccessDetailsItemList []IndexTagSuccessDetailsItem

func (list *IndexTagSuccessDetailsItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]IndexTagSuccessDetailsItem)
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

type IndexTagTagsItemList []IndexTagTagsItem

func (list *IndexTagTagsItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]IndexTagTagsItem)
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
