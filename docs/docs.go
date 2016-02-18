package docs

import (
	"encoding/json"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/swagger"
)

const (
    Rootinfo string = `{"apiVersion":"1.0.0","swaggerVersion":"1.2","apis":[{"path":"/users","description":"Operations about user\n"}],"info":{"title":"tantan test API","description":"beego has a very cool tools to autogenerate documents for your API","contact":"1679550318@qq.com","termsOfServiceUrl":"https://github.com/leonardyp","license":"Url http://www.apache.org/licenses/LICENSE-2.0.html"}}`
    Subapi string = `{"/users":{"apiVersion":"1.0.0","swaggerVersion":"1.2","basePath":"","resourcePath":"/users","produces":["application/json","application/xml","text/plain","text/html"],"apis":[{"path":"/","description":"","operations":[{"httpMethod":"POST","nickname":"create","type":"","summary":"使用用户名创建用户,并返回新创建的用户","parameters":[{"paramType":"body","name":"body","description":"\"The user content\"","dataType":"User","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0}],"responseMessages":[{"code":0,"message":"200","responseModel":""},{"code":500,"message":"422","responseModel":""}]}]},{"path":"/","description":"","operations":[{"httpMethod":"GET","nickname":"AllUserAction","type":"","summary":"获取所有用户","responseMessages":[{"code":200,"message":"用户列表","responseModel":""},{"code":500,"message":"","responseModel":""}]}]},{"path":"/:user_id/relationships/:other_user_id","description":"","operations":[{"httpMethod":"PUT","nickname":"create or update","type":"","summary":"create or update","parameters":[{"paramType":"path","name":"user_id","description":"\"用户id\"","dataType":"int","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0},{"paramType":"path","name":"other_user_id","description":"\"其他用户id\"","dataType":"int","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0},{"paramType":"body","name":"body","description":"\"state\"","dataType":"body","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0}],"responseMessages":[{"code":200,"message":"success","responseModel":""},{"code":500,"message":"422","responseModel":""}]}]},{"path":"/:user_id/relationships","description":"","operations":[{"httpMethod":"GET","nickname":"user all relationships","type":"","summary":"user all relationships","parameters":[{"paramType":"path","name":"user_id","description":"\"用户id\"","dataType":"int","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0}],"responseMessages":[{"code":200,"message":"success","responseModel":""},{"code":500,"message":"422","responseModel":""}]}]}]}}`
    BasePath string= "/v1"
)

var rootapi swagger.ResourceListing
var apilist map[string]*swagger.APIDeclaration

func init() {
	if beego.BConfig.WebConfig.EnableDocs {
		err := json.Unmarshal([]byte(Rootinfo), &rootapi)
		if err != nil {
			beego.Error(err)
		}
		err = json.Unmarshal([]byte(Subapi), &apilist)
		if err != nil {
			beego.Error(err)
		}
		beego.GlobalDocAPI["Root"] = rootapi
		for k, v := range apilist {
			for i, a := range v.APIs {
				a.Path = urlReplace(k + a.Path)
				v.APIs[i] = a
			}
			v.BasePath = BasePath
			beego.GlobalDocAPI[strings.Trim(k, "/")] = v
		}
	}
}


func urlReplace(src string) string {
	pt := strings.Split(src, "/")
	for i, p := range pt {
		if len(p) > 0 {
			if p[0] == ':' {
				pt[i] = "{" + p[1:] + "}"
			} else if p[0] == '?' && p[1] == ':' {
				pt[i] = "{" + p[2:] + "}"
			}
		}
	}
	return strings.Join(pt, "/")
}
