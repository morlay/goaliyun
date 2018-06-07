package imm

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DetectTagRequest struct {
	SrcUris  string `position:"Query" name:"SrcUris"`
	Project  string `position:"Query" name:"Project"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *DetectTagRequest) Invoke(client goaliyun.Client) (*DetectTagResponse, error) {
	resp := &DetectTagResponse{}
	err := client.Request("imm", "DetectTag", "2017-09-06", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DetectTagResponse struct {
	RequestId      goaliyun.String
	SuccessNum     goaliyun.String
	SuccessDetails DetectTagSuccessDetailsItemList
	FailDetails    DetectTagFailDetailsItemList
}

type DetectTagSuccessDetailsItem struct {
	SrcUri goaliyun.String
	Tags   DetectTagTagsItemList
}

type DetectTagTagsItem struct {
	TagId         goaliyun.String
	TagLevel      goaliyun.String
	TagName       goaliyun.String
	ParentTagId   goaliyun.String
	ParentTagName goaliyun.String
	TagScore      goaliyun.String
}

type DetectTagFailDetailsItem struct {
	SrcUri goaliyun.String
	Reason goaliyun.String
}

type DetectTagSuccessDetailsItemList []DetectTagSuccessDetailsItem

func (list *DetectTagSuccessDetailsItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DetectTagSuccessDetailsItem)
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

type DetectTagFailDetailsItemList []DetectTagFailDetailsItem

func (list *DetectTagFailDetailsItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DetectTagFailDetailsItem)
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

type DetectTagTagsItemList []DetectTagTagsItem

func (list *DetectTagTagsItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DetectTagTagsItem)
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
