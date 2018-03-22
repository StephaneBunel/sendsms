package ovh

import (
	"strings"

	"github.com/romana/rlog"
	"github.com/spf13/viper"
)

func (p *ovhProvider) Config(userConfig *viper.Viper) {
	p.apiConfig.location = userConfig.GetString("api.location")
	p.apiConfig.appKey = userConfig.GetString("api.appKey")
	p.apiConfig.appSecret = userConfig.GetString("api.appSecret")
	p.apiConfig.consumerKey = userConfig.GetString("api.consumerKey")
	p.apiConfig.serviceName = userConfig.GetString("api.serviceName")

	p.smsConfig = map[string]interface{}{
		"charset":           "UTF-8",
		"class":             "phoneDisplay",
		"coding":            "8bit",
		"nostopclause":      true,
		"priority":          "high",
		"senderforresponse": false,
	}

	// Set defaults
	for opt, val := range userConfig.GetStringMap("smsOptions") {
		rlog.Debug("Set user defaut sms configuration: ", opt, " = ", val)
		p.smsConfig[opt] = val
	}

	// Fix options case
	for key, casedopt := range userConfig.GetStringMap("smsOptionsCaps") {
		lcaseopt := strings.ToLower(key)
		if content, exists := p.smsConfig[lcaseopt]; exists {
			p.smsConfig[casedopt.(string)] = content
			delete(p.smsConfig, lcaseopt)
		}
	}
}
