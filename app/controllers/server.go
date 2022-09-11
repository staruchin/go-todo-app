package controllers

import (
	"fmt"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"text/template"
	"todo_app/app/models"
	"todo_app/config"
)

func generateHTML(w http.ResponseWriter, data interface{}, filenames ...string) {
	var files []string
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("app/views/templates/%s.html", file))
	}

	templates := template.Must(template.ParseFiles(files...))
	// テンプレートのキャッシュを作成
	// template.Must()は独自にエラーチェックを行うため、errorを返り値には持たず、ハンドリングする必要がありません。
	// つまりParseFilesがエラーの場合、panicになる。
	// template.ParseFiles()は可変長引数をとり、その引数としてキャッシュさせたいファイルの名前を指定します。

	templates.ExecuteTemplate(w, "layout", data)
	// defineでテンプレートを定義した場合、ExecuteTemplateでlayoutを明示的に指定する必要がある

	// たぶん基本はテンプレートのhtmlに{{template "content" . }}のように書いておけば
	// 自動でこの場合contentが解決されるんだろう。
	// そうではない、{{template ...}}が書かれていないテンプレートを解決する方法として、
	// ParseFilesして得られたtemplate.TemplateをExecuteないしExetuteTemplateすればいいと思われる。
	// ただ上記の通り{{define ...}}で名前を付けて定義したものはその名前を指定してExecuteTemplate
	// する必要がある、といったところか。

	// template.Mustについては、
	// ページを切り替える度にここが呼ばれるので、毎回ParseFilesしなくていいように
	// キャッシュしておくということだろう。

	// layout.htmlがベースで、その中に書いた{{template ...}}がそれぞれのページに置き換わる。
	// 本generateHTML()に"layout", "top"を指定して呼び出せばtop.htmlの{{define content}}
	// "layout", "signup"を指定して呼び出せばsignup.htmlの{{define "content"}}で置き換わる。
	// さらにその中に{{.}}があれば、templates.ExecuteTemplate(w, "layout", data)の処理により
	// {{.}}がdataで置き換わる。
}

func session(w http.ResponseWriter, r *http.Request) (sess models.Session, err error) {
	cookie, err := r.Cookie("_cookie")
	if err == nil {
		sess = models.Session{UUID: cookie.Value}
		if ok, _ := sess.CheckSession(); !ok {
			err = fmt.Errorf("Invalid session")
		}
	}
	return sess, err
}

var validPath = regexp.MustCompile("^/todos/(edit|update|delete)/([0-9]+)$")

func parseURL(fn func(http.ResponseWriter, *http.Request, int)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// /todos/edit/1
		q := validPath.FindStringSubmatch(r.URL.Path)
		if q == nil {
			http.NotFound(w, r)
			return
		}
		qi, err := strconv.Atoi(q[2])
		if err != nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, qi)
	}
}

func StartMainServer() error {

	files := http.FileServer(http.Dir(config.Config.Static))
	http.Handle("/static/", http.StripPrefix("/static/", files))
	// app/viewsにあるcssとjsを指定したURL（ここでは/static/）に配置したい。
	// top.htmlのようにcssが使えるようになる。

	http.HandleFunc("/", top)
	// 指定したURL（ここでは/）にアクセスしたら指定したハンドラが実行される。
	// つまりここではトップページを開いたら、top()が実行される。
	// その中でgenerateHTML()を呼んでるので、ExecuteTemplateが実行され、
	// htmlが生成されて表示内容が切り替わる。
	http.HandleFunc("/signup", signup)
	// 同様にここでは/signupにアクセスしたら指定したハンドラsignup()が呼ばれる。
	// /signupにアクセスする（ページを開く）＝GETリクエストを投げる場合だけでなく、
	// POSTリクエストを投げる場合もハンドラが実行される。
	// /signupにはPOSTリクエストを投げることもあるので、
	// signup()の中でリクエストのメソッドの種類を見て条件分岐させる。
	http.HandleFunc("/login", login)
	http.HandleFunc("/authenticate", authenticate)
	http.HandleFunc("/logout", logout)
	http.HandleFunc("/todos", index)
	http.HandleFunc("/todos/new", todoNew)
	http.HandleFunc("/todos/save", todoSave)
	http.HandleFunc("/todos/edit/", parseURL(todoEdit))
	// 末尾をスラッシュ/とする。末尾が/でない場合はURLの完全一致が求められる。
	// 末尾が/である場合は、要求されたURLの先頭が登録されたURLと一致するか調べる。
	// つまりURLの後に数値や文字列などが付いていた場合でも、ハンドラ関数に処理を渡すことができる。
	http.HandleFunc("/todos/update/", parseURL(todoUpdate))
	http.HandleFunc("/todos/delete/", parseURL(todoDelete))

	port := os.Getenv("PORT")
	return http.ListenAndServe(":"+port, nil)
	// nilはデフォルトのマルチプレクサ
	// 登録されていないURLにアクセスしたら404 page not foundを返す処理が行われる。
}
