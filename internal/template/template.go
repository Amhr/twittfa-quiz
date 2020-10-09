package template

import (
	"WhoKnowsMeapp/config"
	"github.com/foolin/echo-template"
	"github.com/labstack/echo"
)

func NewTemplate() *echotemplate.TemplateEngine {
	return echotemplate.Default()
}

func TemplateData(p echo.Map) echo.Map {
	if p == nil {
		p = echo.Map{}
	}
	p["baseurl"] = config.FULL_URL
	p["base_url"] = config.FULL_URL
	p["site_title"] = "TwittFa"
	return p
}
