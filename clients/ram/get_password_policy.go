package ram

import (
	"github.com/morlay/goaliyun"
)

type GetPasswordPolicyRequest struct {
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *GetPasswordPolicyRequest) Invoke(client goaliyun.Client) (*GetPasswordPolicyResponse, error) {
	resp := &GetPasswordPolicyResponse{}
	err := client.Request("ram", "GetPasswordPolicy", "2015-05-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type GetPasswordPolicyResponse struct {
	RequestId      goaliyun.String
	PasswordPolicy GetPasswordPolicyPasswordPolicy
}

type GetPasswordPolicyPasswordPolicy struct {
	MinimumPasswordLength      goaliyun.Integer
	RequireLowercaseCharacters bool
	RequireUppercaseCharacters bool
	RequireNumbers             bool
	RequireSymbols             bool
	HardExpiry                 bool
	MaxPasswordAge             goaliyun.Integer
	PasswordReusePrevention    goaliyun.Integer
	MaxLoginAttemps            goaliyun.Integer
}
