package handler

import (
	"fmt"
	"net/url"

	"git-dev.linecorp.com/tadafumi-yoshihara/idpgo/idp"
	"github.com/gin-gonic/gin"
)

const (
	baseURL    = "https://biz-idp-internal-api.line-apps-beta.com/"
	token      = "uMOQO7iJmwqA-zhNJVeRd-HEIOmQPf5AUbSVbBS4AuI"
	rootURL    = "https://console.brain.line-beta.biz/"
	idpURL     = "https://account.line-beta.biz/"
	apiURL     = "https://idp-test.line-apps-beta.com"
)

var (
	idpClient *idp.Client
	env      string
	corsOrigin  map[string]string
)

func init() {
	var err error
	idpClient, err = idp.New(baseURL, token, 1000)
	if err != nil {
		panic(err)
	}
	env = "local"
	
	corsOrigin = map[string]string {
		"local": "http://localhost:9000",
		"beta":  "https://console.brain.line-beta.biz",
	}
}

func SetEnv(e string) {
	env = e
	fmt.Printf("Setting env to %s\n", env)
}

func Root(ctx *gin.Context) {

	fmt.Println("0000")
	params := gin.H{}
	params["idpUrl"] = idpURL
	params["rootUrl"] = rootURL

	ses := getCookie(ctx)

	//ses = "ssss"

	if ses != "" {
		fmt.Println("111111")
		rs, err := idpClient.GetSession(ses)

		//err = nil
		//rs = &idp.SessionResponse{}
		//rs.User = idp.UserResponse{}
		//rs.User.BusinessID = "hoge"

		if err == nil && rs.User.BusinessID != "" {
			fmt.Println("22222")
			user := rs.User
			params["user"] = user
			params["rootUrlE"] = url.QueryEscape(rootURL)
			ctx.HTML(200, "index.html", params)
			return
		}
		fmt.Printf("Failed to get IdP session: %v, %v\n", rs, err)
	}

	ctx.HTML(200, "index.html", params)
}

func GetLoginStatus(ctx *gin.Context) {

	params := gin.H{}
	params["idpUrl"] = idpURL
	params["rootUrl"] = rootURL

	ses := getCookie(ctx)

	//ses = "ssss"
	fmt.Printf(ses)

	if ses != "" {
		rs, err := idpClient.GetSession(ses)

		//err = nil
		//rs = &idp.SessionResponse{}
		//rs.User = idp.UserResponse{}
		//rs.User.BusinessID = "hoge"

		if err == nil && rs.User.BusinessID != "" {
			user := rs.User
			params["user"] = user
			params["rootUrlE"] = url.QueryEscape(rootURL)
			// ctx.HTML(200, "my.html"c, params)
			ctx.Header("Access-Control-Allow-Origin", corsOrigin[env])
			ctx.Header("Access-Control-Allow-Credentials", "true")
			ctx.JSON(200, params)
			return
		}
		fmt.Printf("Failed to get IdP session: %v, %v\n", rs, err)
	}

	// ctx.HTML(200, "index.html", params)
	ctx.Header("Access-Control-Allow-Origin", corsOrigin[env])
	ctx.Header("Access-Control-Allow-Credentials", "true")
	ctx.JSON(200, params)
}

func GetLogoutURL(ctx *gin.Context) {
	ses := getCookie(ctx)
	lr, err := idpClient.GenerateLogoutURI(ses, rootURL)
	if err != nil || lr.Status != "success" {
		fmt.Printf("Failed to get generate logout URL: %+v, %v\n", lr, err)
		ctx.AbortWithStatus(500)
	}
	ctx.Header("Access-Control-Allow-Credentials", "true")
	ctx.Header("Access-Control-Allow-Origin", corsOrigin[env])

	ret := gin.H{}
	ret["logoutUrl"] = lr.LogoutURI
	ctx.JSON(200, ret)
}

func Logout(ctx *gin.Context) {
	ses := getCookie(ctx)
	lr, err := idpClient.GenerateLogoutURI(ses, rootURL)
	if err != nil || lr.Status != "success" {
		fmt.Printf("Failed to get generate logout URL: %+v, %v\n", lr, err)
		ctx.AbortWithStatus(500)
	}
	ctx.Header("Access-Control-Allow-Credentials", "true")
	ctx.Header("Access-Control-Allow-Origin", "http://localhost:9000")
	ctx.Redirect(302, lr.LogoutURI)
}

func Login(ctx *gin.Context) {
	// ses := getCookie(ctx)
	url := idp.GenerateLoginURI()

	ctx.Header("Access-Control-Allow-Credentials", "true")
	ctx.Header("Access-Control-Allow-Origin", "http://localhost:9000")
	ctx.Redirect(302, url)
}

func getCookie(ctx *gin.Context) string {
	// ctx.Cookie() does not work
	// https://github.com/gin-gonic/gin/issues/1717
	cookie, err := ctx.Request.Cookie("ses")
	if err != nil {
		return ""
	}
	return cookie.Value
}
