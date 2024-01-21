package apiserver

import (
	"log"
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

// API server
type APIserver struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
	//store  *store.Store
}

// Create new server
func New(config *Config) *APIserver {
	return &APIserver{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

// Start new server
func (s *APIserver) Start() error {
	if err := s.configurelogger(); err != nil {
		return err
	}
	s.configureRouter()
	/* if err := s.configureStore(); err != nil {
		return err
	} */
	s.logger.Info("strating API server")
	return http.ListenAndServe(s.config.BindAddr, s.router)
}

func (s *APIserver) configurelogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}
	s.logger.SetLevel(level)
	return nil
}

func (s *APIserver) configureRouter() {

	router := s.router.StrictSlash(true)
	router.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/404", http.StatusFound)
	})
	router.HandleFunc("/404", s.misshandle())
	router.HandleFunc("/messenger", s.messengerhandle())
	router.HandleFunc("/wallet", s.wallethandle())
	router.HandleFunc("/exchange", s.exchangehandle())
	router.HandleFunc("/blog", s.blogshandle())
	router.HandleFunc(`/blog/{name}`, s.bloghandle())
	router.HandleFunc("/about", s.abouthandle())
	router.HandleFunc("/contacts", s.contacthandle())
	router.HandleFunc("/pivacy", s.privacyhandle())
	router.HandleFunc("/terms", s.termshandle())
	router.HandleFunc("/security", s.securityhandle())
	router.HandleFunc("/", s.mainhandle())
	router.PathPrefix("/assets").Handler(http.FileServer(http.Dir("static")))

}

func (s *APIserver) handlehello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//create html template
		tmpl, err := template.ParseFiles("static/account-address.html")
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		//view html template
		err = tmpl.Execute(w, nil)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
	}
}

func (s *APIserver) mainhandle() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		basetpls := []string{"static/header.html", "static/footer.html", "static/head.html", "static/meta.html"}
		mass := []string{"static/index.html", basetpls[0], basetpls[1], basetpls[2], basetpls[3]}

		//create html template
		tmpl, err := template.ParseFiles(mass...)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}

		err = tmpl.Execute(w, nil)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
	}
}

func (s *APIserver) messengerhandle() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		basetpls := []string{"static/header.html", "static/footer.html", "static/head.html"}
		mass := []string{"static/wallet.html", basetpls[0], basetpls[1], basetpls[2]}
		//create html template
		tmpl, err := template.ParseFiles(mass...)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		//view html template
		err = tmpl.Execute(w, nil)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
	}
}

func (s *APIserver) wallethandle() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		basetpls := []string{"static/header.html", "static/footer.html", "static/head.html"}
		mass := []string{"static/wallet.html", basetpls[0], basetpls[1], basetpls[2]}
		//create html template
		tmpl, err := template.ParseFiles(mass...)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		//view html template
		err = tmpl.Execute(w, nil)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}

	}

}

func (s *APIserver) exchangehandle() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		basetpls := []string{"static/header.html", "static/footer.html", "static/head.html"}
		mass := []string{"static/index.html", basetpls[0], basetpls[1], basetpls[2]}
		//create html template
		tmpl, err := template.ParseFiles(mass...)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		//view html template
		err = tmpl.Execute(w, nil)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}

	}

}

func (s *APIserver) abouthandle() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		basetpls := []string{"static/header.html", "static/footer.html", "static/head.html", "static/meta.html"}
		mass := []string{"static/about.html", basetpls[0], basetpls[1], basetpls[2], basetpls[3]}
		//create html template
		tmpl, err := template.ParseFiles(mass...)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		//view html template
		err = tmpl.Execute(w, nil)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
	}
}

func (s *APIserver) contacthandle() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		mass := []string{"static/contacts.html", "static/header.html", "static/footer.html"}
		//create html template
		tmpl, err := template.ParseFiles(mass...)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		//view html template
		err = tmpl.Execute(w, nil)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
	}
}

func (s *APIserver) blogshandle() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		mass := []string{"static/blog.html", "static/header.html", "static/footer.html"}
		//create html template
		tmpl, err := template.ParseFiles(mass...)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		//view html template
		err = tmpl.Execute(w, nil)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
	}
}

func (s *APIserver) bloghandle() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		mass := []string{"static/blog-article.html", "static/header.html", "static/footer.html"}
		//create html template
		tmpl, err := template.ParseFiles(mass...)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		//view html template
		err = tmpl.Execute(w, nil)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}

	}

}

func (s *APIserver) termshandle() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		basetpls := []string{"static/header.html", "static/footer.html", "static/head.html"}
		mass := []string{"static/terms.html", basetpls[0], basetpls[1], basetpls[2]}
		//create html template
		tmpl, err := template.ParseFiles(mass...)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		//view html template
		err = tmpl.Execute(w, nil)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
	}
}

func (s *APIserver) securityhandle() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		basetpls := []string{"static/header.html", "static/footer.html", "static/head.html"}
		mass := []string{"static/security.html", basetpls[0], basetpls[1], basetpls[2]}
		tmpl, err := template.ParseFiles(mass...)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		//view html template
		err = tmpl.Execute(w, nil)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
	}
}

func (s *APIserver) privacyhandle() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		basetpls := []string{"static/header.html", "static/footer.html", "static/head.html"}
		mass := []string{"static/terms.html", basetpls[0], basetpls[1], basetpls[2]}
		//create html template
		tmpl, err := template.ParseFiles(mass...)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		//view html template
		err = tmpl.Execute(w, nil)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
	}
}

/* func (s *APIserver) configureStore() error {

	st := store.New(s.config.Store)
	if err := st.Open(); err != nil {
		return err
	}
	s.store = st
	return nil

} */

func (s *APIserver) singuphandle() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		tpl, err := template.ParseFiles("static/page-signup.html")
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		tpl.Execute(w, err)

	}
}

/*
func (s *APIserver) registerhandle() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		data := &m.User{
			Email:    r.FormValue("email"),
			Password: r.FormValue("password"),
		}

		s.store.User().Create(data)
		tplpath := "static/page-access.html"
		tpl, err := template.ParseFiles(tplpath)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		tpl.Execute(w, data)

	}
} */
/*
func (s *APIserver) userhandle() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		users, err := s.store.User().GetUsers()
		if err != nil {
			http.Error(w, err.Error(), 400)
		}
		tplpath := "static/users.html"
		tpl, err := template.ParseFiles(tplpath)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		tpl.Execute(w, &users)

	}
}
*/
/*
// test data from Database
func (s *APIserver) testhandle() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		users, err := s.store.User().GetUsers()
		if err != nil {
			http.Error(w, err.Error(), 400)
		}
		fmt.Fprintf(w, "%s\n", users)

	}
}
*/
/* func (s *APIserver) contenthandle() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		url := store.Base
		s := &store.Generated{}
		err := store.GetJson(url, s)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(s.Section)
		tpl, err := template.ParseFiles("static/content.html")
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		err = tpl.Execute(w, &s)
		if err != nil {l
			log.Fatal(err)
		}
	}

} */

func (s *APIserver) misshandle() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		tpl, err := template.ParseFiles("static/404.html")
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		err = tpl.Execute(w, nil)
		if err != nil {
			log.Fatal(err)
		}
	}

}
