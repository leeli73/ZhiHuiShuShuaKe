package main
import (
	"strings"
	"fmt"
	"bytes"
	"net/http"
	"io/ioutil"
	"github.com/elazarl/goproxy"
)
var Code string = `<script>function query() {
    if ($("#popbox_title").length > 0) {
        $(".popboxes_close")[0].click();
        console.log('关闭窗口');
    }

    if ($("#chapterList .time_ico.fl").nextAll()[2].children[0].style.width === "100%" || $("video").get(0).ended) {
        var num = -1;
        var text = $("#chapterList .time_ico.fl").parent().nextAll()[++num].id;
        while (text === "" ||
               text.substr(0, 5) != "video" ||
               text.substr(0, 7) === "video-0") {
            text = $("#chapterList .time_ico.fl").parent().nextAll()[++num].id;
        }
        $("#chapterList .time_ico.fl").parent().nextAll()[num].click();
    }

    if ($("video").length > 0 && $("video").get(0).playbackRate != 1.5) {
        console.log('切换到1.5倍');
        $(".speedTab15")[0].click();
    }

    if ($("video").get(0).volume > 0) {
        $(".volumeIcon").click();
	}
}
var divObj=document.createElement("div"); 
divObj.innerHTML='<h1 style="font-size:50px;">技术支持来自：Lee QQ:925776327<br>项目整体开源,欢迎Fork和Star <a href="https://github.com/leeli73/ZhiHuiShuShuaKe">github.com/leeli73/ZhiHuiShuShuaKe</a></h1>'; 
var first=document.body.firstChild;
document.body.insertBefore(divObj,first);
window.setInterval(query, 1000);</script>`

func main() {
	proxy := goproxy.NewProxyHttpServer()
	fmt.Println("Proxy已经成功启动...智慧树课程页面监控中...")
	fmt.Println("请将你的浏览器代理设定为127.0.0.1:8080  设置方法请百度，很简单！")
	fmt.Println("支持在内网和外网通过IP访问代理服务")
	fmt.Println("技术支持来自：Lee QQ:925776327")
	fmt.Println("项目整体开源,欢迎Fork和Star 项目地址:https://github.com/leeli73/ZhiHuiShuShuaKe")
	fmt.Println("你也可以通过Tampermonkey、Fiddler等软件，向网页中注入上面Git项目中的Code.js")
	proxy.Verbose = false
	
	proxy.OnRequest().HandleConnect(goproxy.AlwaysMitm)
	proxy.OnRequest().DoFunc(
		func(r *http.Request,ctx *goproxy.ProxyCtx)(*http.Request,*http.Response) {
			r.Header.Set("X-GoProxy","yxorPoG-X")
			return r,nil
	})
	proxy.OnRequest(goproxy.DstHostIs("eol.qhu.edu.cn")).DoFunc(
    func(r *http.Request,ctx *goproxy.ProxyCtx)(*http.Request,*http.Response) {
        return r,goproxy.NewResponse(r,
			goproxy.ContentTypeText,http.StatusForbidden,
			"Don't waste your time!")
	})
	proxy.OnResponse().DoFunc(
		func(r *http.Response, ctx *goproxy.ProxyCtx)*http.Response{
			if strings.Contains(r.Request.URL.Path,"learning/videoList"){
				bs, _ := ioutil.ReadAll(r.Body)
				sbs := string(bs)
				sbs = sbs + Code
				fmt.Println("发现课程页面，刷课代码已经成功注入！")
				bs = []byte(sbs)
				r.Body = ioutil.NopCloser(bytes.NewReader(bs))
			}
			return r
	})
	http.ListenAndServe(":8080", proxy)
}
