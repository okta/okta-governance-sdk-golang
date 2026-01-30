package governance

import (
	"runtime"

	"github.com/okta/okta-sdk-golang/v6/okta"
)

type UserAgent struct {
	goVersion string

	osName string

	osVersion string

	config *okta.Configuration
}

func NewUserAgent(config *okta.Configuration) UserAgent {
	ua := UserAgent{}
	ua.config = config
	ua.goVersion = runtime.Version()
	ua.osName = runtime.GOOS
	ua.osVersion = runtime.GOARCH

	return ua
}

func (ua UserAgent) String() string {
	userAgentString := "okta-governance-sdk-golang/" + VERSION + " "
	userAgentString += "golang/" + ua.goVersion + " "
	userAgentString += ua.osName + "/" + ua.osVersion

	if ua.config.UserAgentExtra != "" {
		userAgentString += " " + ua.config.UserAgentExtra
	}

	return userAgentString
}
