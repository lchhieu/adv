package controllers

import (
	//"bee-project/models"
	"github.com/astaxie/beego"
	"net"
	//"strconv"
)

func init() {

}

type AnalyticResult struct {
	IdSupplySource int
	Url            string
	Width, Height  int
	Path, Domain   string
	ResponseType   string
	CacheBuster    string
	Cookie         string
	UserAgent      string
	Referer        string
	Ip             string
}

type AnalyticController struct {
	beego.Controller
}

func (this *AnalyticController) Get() {
	var result AnalyticResult
	//result.IdSupplySource, _ = strconv.Atoi(this.Input().Get("i"))
	this.Ctx.Input.Bind(&result.IdSupplySource, "i")
	//result.Width, _ = strconv.Atoi(this.Input().Get("w"))
	this.Ctx.Input.Bind(&result.Url, "w")
	this.Ctx.Input.Bind(&result.Width, "w")
	this.Ctx.Input.Bind(&result.Height, "h")
	this.Ctx.Input.Bind(&result.Domain, "p")
	this.Ctx.Input.Bind(&result.ResponseType, "vp")
	result.Cookie = this.Ctx.Request.Header.Get("Cookie")
	result.UserAgent = this.Ctx.Request.Header.Get("User-Agent")
	result.Referer = this.Ctx.Request.Header.Get("Referer")
	ip, _, _ := net.SplitHostPort(this.Ctx.Request.RemoteAddr)
	result.Ip = ip
	//userIP := net.ParseIP(ip)
	//this.Redirect(result.Referer, 302)
	this.Data["xml"] = result
	this.ServeXML()
}
