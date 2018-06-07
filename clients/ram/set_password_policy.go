package ram

import (
	"github.com/morlay/goaliyun"
)

type SetPasswordPolicyRequest struct {
	RequireNumbers             string `position:"Query" name:"RequireNumbers"`
	PasswordReusePrevention    int64  `position:"Query" name:"PasswordReusePrevention"`
	RequireUppercaseCharacters string `position:"Query" name:"RequireUppercaseCharacters"`
	MaxPasswordAge             int64  `position:"Query" name:"MaxPasswordAge"`
	MaxLoginAttemps            int64  `position:"Query" name:"MaxLoginAttemps"`
	HardExpiry                 string `position:"Query" name:"HardExpiry"`
	MinimumPasswordLength      int64  `position:"Query" name:"MinimumPasswordLength"`
	RequireLowercaseCharacters string `position:"Query" name:"RequireLowercaseCharacters"`
	RequireSymbols             string `position:"Query" name:"RequireSymbols"`
	RegionId                   string `position:"Query" name:"RegionId"`
}

func (req *SetPasswordPolicyRequest) Invoke(client goaliyun.Client) (*SetPasswordPolicyResponse, error) {
	resp := &SetPasswordPolicyResponse{}
	err := client.Request("ram", "SetPasswordPolicy", "2015-05-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SetPasswordPolicyResponse struct {
	RequestId      goaliyun.String
	PasswordPolicy SetPasswordPolicyPasswordPolicy
}

type SetPasswordPolicyPasswordPolicy struct {
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
