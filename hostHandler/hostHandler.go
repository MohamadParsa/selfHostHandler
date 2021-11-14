package hostHandler

import (
	"net/http"
	"io/ioutil"
	"fmt"
	"strings"
	"os"
	"bufio"
)
type handler struct{
	// http.Handler is an interface from net/http, It enables us to overwrite the ServeHTTP func .
	// so we can define a method with a receiver such as ( this *handler ) and overwrite the ServeHTTP func
	http.Handler
	handlerSetting Settings
}
type Settings struct{
	// physical web pages address. based onhandler the Main package.
	RootWebPath 		string
	//////////////////////////////////////////////////////
	//				Custom Pages Address				//
	//////////////////////////////////////////////////////
	NotFindePagePath 	string
	DefaultPageName 	string
}
func (this *handler) ServeHTTP(w http.ResponseWriter,req *http.Request)  {
	
	if (this.handlerSetting.RootWebPath==""){
		w.Write([]byte("Error Code: WPH-P:001-19"))// web page address is't define.
	}
	reqURLPath:=req.URL.Path
	if (reqURLPath=="" || reqURLPath=="/"){
		reqURLPath=this.handlerSetting.DefaultPageName
	}
	fmt.Println(this.handlerSetting.RootWebPath+reqURLPath)
	pageData,errorInReadingPage:=os.Open(this.handlerSetting.RootWebPath+reqURLPath)
	
	if errorInReadingPage==nil	{
		bufferReader := bufio.NewReader(pageData)
		var contentType string=getContentType(reqURLPath)
		
		w.Header().Add("Content Type",contentType)
		bufferReader.WriteTo(w)
	} else {
		//Error handling with not found page or a simple text
		notFindePagePath:=this.handlerSetting.NotFindePagePath
	fmt.Println(this.handlerSetting.RootWebPath+notFindePagePath)

		if(notFindePagePath != ""){
			notFoundPageData,errorInReadingNotFoundPage:=ioutil.ReadFile((string(this.handlerSetting.RootWebPath+notFindePagePath)))
			if errorInReadingNotFoundPage==nil {
				w.Write(notFoundPageData)
			} else {
				notFindePagePath=""
			}
		}
		if(notFindePagePath == ""){
			w.WriteHeader(404)
			w.Write([]byte("404 - "+http.StatusText(404)))
		} 
	}
}
func hostHandler()  {

}
func RunHostHandler(Routing string, port string,setting Settings)  {
	fmt.Println("Start")
	http.Handle(Routing,&handler{handlerSetting:setting})
	http.ListenAndServe(":"+port,nil)
	fmt.Println("End")

}
func getContentType(reqURLPath string) string{
	contentType:="text/plain";
	switch  {
	case strings.HasSuffix(reqURLPath,".css"):
		contentType="text/css"
	case strings.HasSuffix(reqURLPath,".html"):
		contentType="text/html"
	case strings.HasSuffix(reqURLPath,".js"):
		contentType="application/javascript"
	case strings.HasSuffix(reqURLPath,".png"):
		contentType="image/png"
	case strings.HasSuffix(reqURLPath,".mp4"):
		contentType="video/mp4"
	}
	return contentType
}