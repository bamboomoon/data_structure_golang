package commentTeacher

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/axgle/mahonia"
	"github.com/pkg/errors"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func BeginComment() {
	http.HandleFunc("/loginPage", loginPage)
	http.HandleFunc("/login", login)
	http.Handle("/", http.StripPrefix("/image/", http.FileServer(http.Dir("image"))))
	http.HandleFunc("/jjj", yzm)
	http.ListenAndServe(":9090", nil)
}

var coo *http.Cookie
var _VIEWSTATEValue string

func yzm(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Header, r.Form)
}

//获取 cookie
func getCookie() bool {
	resp, err := http.Get("http://jwweb.scujcc.cn/default2.aspx")
	if err != nil {
		fmt.Println(err)
		return false
	}
	if err != nil {
		log.Fatal(err)
		return false
	}

	coo = resp.Cookies()[0]

	//使用goquery
	document, _ := goquery.NewDocumentFromResponse(resp)
	_VIEWSTATEValue, _ = document.Find("input").First().Attr("value")
	return true
}

//获取验证码图片
func getImage() bool {
	s := getCookie()
	if s == false {
		return false
	}
	resp, err := sendARequest("GET", "http://jwweb.scujcc.cn/CheckCode.aspx", nil, nil, false, nil)
	defer resp.Body.Close()
	if err != nil {
		return false
	}
	//保存图片
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return false
	}
	ioutil.WriteFile("./image/authCode.gif", data, 0666)

	return true
}

//返回登录页面
func loginPage(w http.ResponseWriter, res *http.Request) {
	if res.Method == "GET" {
		s := getImage()
		if s == false {
			return
		}
		t, err := template.ParseFiles("index.html")
		fmt.Println(res.Form)
		if err != nil {
			log.Fatal(err)
			return
		}

		if coo == nil {
			log.Fatal("没有cookie")
			return
		}
		t.Execute(w, nil)
	}
}

//开始登录
func login(w http.ResponseWriter, res *http.Request) {
	fmt.Println("login")
	if res.Method == "GET" {
		return
	}
	res.ParseForm()
	fmt.Println(res.Form)
	name := res.Form.Get("name")
	form := url.Values{}
	form.Add("__VIEWSTATE", _VIEWSTATEValue)
	form.Add("txtUserName", name)
	form.Add("TextBox2", res.Form.Get("pass"))
	form.Add("txtSecretCode", res.Form.Get("authCode")) //验证码
	form.Add("RadioButtonList1", mahonia.NewEncoder("GBK").ConvertString("学生"))
	form.Add("Button1", "")
	form.Add("lbLanguage", "")
	form.Add("hidPdrs", "")
	form.Add("hidsc", "")

	var headerForm map[string]string = make(map[string]string)
	headerForm["Content-Type"] = "application/x-www-form-urlencoded"
	headerForm["Host"] = "jwweb.scujcc.cn"
	headerForm["Origin"] = "http://jwweb.scujcc.cn"
	headerForm["Referer"] = "http://jwweb.scujcc.cn/default2.aspx"
	sendARequest("POST", "http://jwweb.scujcc.cn/default2.aspx", form, headerForm, true, func(location *url.URL) {
		gradeURL := loginRedirect(location)
		//开启评级
		giveGradeToTeacher(gradeURL, location.String())
	})

}

//登录重定向 返回教务质量评级的URL
func loginRedirect(location *url.URL) (herf string) {
	resp, err := sendARequest("GET", fmt.Sprintf("%s", location), nil, nil, false, nil)
	if err != nil {
		log.Fatalln("loginRedirect() err:", err)
	}
	//goquery 抓取质量评教 URL
	document, err := goquery.NewDocumentFromResponse(resp) //goquert 会关闭 resp.body
	if err != nil {
		log.Fatalln("loginRedirect document Error :", err)
		return
	}
	hrefValue, _ := document.Find("a").Eq(9).Attr("href")
	defer resp.Body.Close()
	return hrefValue
}

//开始评级
func giveGradeToTeacher(gradeURL, location string) {
	var hederMap map[string]string = make(map[string]string)
	hederMap["Referer"] = location
	resp, err := sendARequest("GET", fmt.Sprintf("http://jwweb.scujcc.cn/%s", gradeURL), nil, hederMap, false, nil)
	if err != nil {
		log.Fatalln("giveGradeToTeacher err:", err)
		return
	}
	getCourseFormActionURL(resp)
}

//在方法外必须关闭 body
func sendARequest(method, url string, postForm url.Values, headerForm map[string]string, isStopRedirect bool, redricetFunc func(location *url.URL)) (*http.Response, error) {
	c := http.Client{}
	//请求
	req, err := http.NewRequest(method, url, strings.NewReader(postForm.Encode()))
	if err != nil {
		log.Fatalln("创建请求出错:", err)
		return nil, errors.New("创建请求出错")
	}
	if headerForm != nil {
		for k, v := range headerForm {
			req.Header.Add(k, v)
		}
	}
	//阻止重定向
	if isStopRedirect == true {
		c.CheckRedirect = func(req *http.Request, via []*http.Request) error {
			fmt.Println(req.Header, via)
			return errors.New("err")
		}
	}
	//增加Cookie
	req.Header.Add("Cookie", fmt.Sprintf("%s=%s", coo.Name, coo.Value))
	//发送
	resp, err := c.Do(req)
	if err != nil && isStopRedirect == true {
		if resp.StatusCode == 302 {
			location, _ := resp.Location()
			redricetFunc(location)
		}
	}
	return resp, nil
}

//获取每个课程的评级链接
func getCourseFormActionURL(resp *http.Response) {

	document, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		log.Println("获取课程链接 error:", err)
		return
	}

	sec := document.Find("tbody").Eq(1).Find("a")
	sec.Each(func(i int, se *goquery.Selection) {
		value, _ := se.Attr("onclick")
		//分割字符串 获取到真正的评级地址
		//fmt.Println(value)
		ss := strings.Split(value, ",")
		urlPath := strings.Split(ss[0], "'")[1]
		go beginTeacher(fmt.Sprintf("http://jwweb.scujcc.cn/%s", urlPath))
	})
	return
}

//提交评教内容
var __VIEWSTATE string = ""

func beginTeacher(url_string string) {
	if __VIEWSTATE == "" {
		resp, err := sendARequest("GET", url_string, nil, nil, false, nil)
		document, err := goquery.NewDocumentFromResponse(resp)
		if err != nil {
			log.Fatal("beginTearch Begin err :", err)
			return
		}
		__VIEWSTATE, _ = document.Find("input").Eq(0).Attr("value")
	}

	form := url.Values{}
	form.Add("__VIEWSTATE", __VIEWSTATE)
	for i := 2; i <= 14; i++ {
		form.Add(fmt.Sprintf("DataGrid1:_ctl%d:txt_pf", i), "100")
	}
	form.Add("txt_pjxx", mahonia.NewEncoder("gbk").ConvertString("33老师很好!很认真负责。时不时的开课堂讨论课！很好的带起了我们学习的积极性"))
	form.Add("Button1", mahonia.NewEncoder("gbk").ConvertString("保  存"))
	form.Add("TextBox1", "")

	var headerMap map[string]string = make(map[string]string)
	headerMap["Referer"] = url_string
	headerMap["Content-Type"] = "application/x-www-form-urlencoded"

	resp, err := sendARequest("POST", url_string, form, headerMap, false, nil)
	if err != nil {
		log.Fatalln("teacher post :", err)
		return
	}
	defer resp.Body.Close()
}
