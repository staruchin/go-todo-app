package controllers

import (
	"log"
	"net/http"
	"todo_app/app/models"
)

func top(w http.ResponseWriter, r *http.Request) {
	// t, err := template.ParseFiles("app/views/templates/top.html")
	// // 指定したファイルを解析してtemplate.Templateオブジェクトを取得する。
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// // t.Execute(w, nil)
	// t.Execute(w, "Hello")
	// // テンプレートを実行する。
	// // 第一引数はio.Writer, 第二引数は任意のデータを指定する。
	// // 指定したデータはテンプレートファイル（ここではtop.html）
	// // に記載した{{.}}を置き換えられる。

	_, err := session(w, r)
	if err != nil {
		generateHTML(w, "Hello", "layout", "public_navbar", "top")
	} else {
		http.Redirect(w, r, "/todos", http.StatusFound)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/", http.StatusFound)
	} else {
		user, err := sess.GetUserBySession()
		if err != nil {
			log.Println(err)
		}
		todos, _ := user.GetTodosByUser()
		user.Todos = todos
		generateHTML(w, user, "layout", "private_navbar", "index")
		// dataとしてオブジェクトを渡すこともできる。
		// テンプレート（ここではindex.html）の中では、{{.Name}}のようにプロパティにアクセスできる。
		// さらに{{range .Todos}} {{.Content}} {{end}}のようにスライスの中身にもループしてアクセスできる。
	}
}

func todoNew(w http.ResponseWriter, r *http.Request) {
	_, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusFound)
	} else {
		generateHTML(w, nil, "layout", "private_navbar", "todo_new")
	}
}

func todoSave(w http.ResponseWriter, r *http.Request) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusFound)
	} else {
		err = r.ParseForm()
		if err != nil {
			log.Println(err)
		}
		user, err := sess.GetUserBySession()
		if err != nil {
			log.Println(err)
		}
		content := r.PostFormValue("content")
		if err := user.CreateTodo(content); err != nil {
			log.Println(err)
		}

		http.Redirect(w, r, "/todos", http.StatusFound)
	}
}

func todoEdit(w http.ResponseWriter, r *http.Request, id int) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusFound)
	} else {
		_, err := sess.GetUserBySession()
		if err != nil {
			log.Println(err)
		}
		todo, err := models.GetTodo(id)
		if err != nil {
			log.Println(err)
		}
		generateHTML(w, todo, "layout", "private_navbar", "todo_edit")
	}
}

func todoUpdate(w http.ResponseWriter, r *http.Request, id int) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusFound)
	} else {
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
		}
		user, err := sess.GetUserBySession()
		if err != nil {
			log.Println(err)
		}
		content := r.PostFormValue("content")
		todo := &models.Todo{ID: id, Content: content, UserID: user.ID}
		if err := todo.UpdateTodo(); err != nil {
			log.Println(err)
		}
		http.Redirect(w, r, "/todos", http.StatusFound)
	}
}

func todoDelete(w http.ResponseWriter, r *http.Request, id int) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusFound)
	} else {
		_, err := sess.GetUserBySession()
		if err != nil {
			log.Println(err)
		}
		todo, err := models.GetTodo(id)
		if err != nil {
			log.Println(err)
		}
		if err := todo.DeleteTodo(); err != nil {
			log.Println(err)
		}
		http.Redirect(w, r, "/todos", http.StatusFound)
	}
}
