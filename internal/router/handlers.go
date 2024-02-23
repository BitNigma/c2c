package router

import (
	"fmt"
	"log"
	"time"

	"p2p/internal/model"
	"p2p/internal/userstore"

	"github.com/gofiber/fiber/v2"
)

type Auth struct {
	Islogin    bool
	Email      string
	KYC        bool
	ErrorPass  error
	ErrorEmail error
}

// Handle main page
func HandleHome(c *fiber.Ctx) error {

	auth := &Auth{}
	auth.Islogin = true

	session := c.Cookies("session_id")
	if len(session) < 1 {
		auth.Islogin = false
	}

	// Render index template
	return c.Render("./static/index.html", fiber.Map{
		"Title":   "C2C processing",
		"Islogin": auth.Islogin,
	})
}

// Handle POST request
func HandleLoginAuth(c *fiber.Ctx) error {

	auth := &Auth{}
	cookie := c.Cookies("session_id")
	if len(cookie) > 1 {
		auth.Islogin = true
		return c.Redirect("/")
	}

	sess, err := store.Get(c)
	if err != nil {
		log.Println("can't get session store", err)
		return err
	}

	//Create User
	user := model.User{
		Email:    c.FormValue("email"),
		Password: c.FormValue("password"),
	}

	store := userstore.NewStore()

	if err = store.CheckAuthUser(user); err != nil {
		switch err {
		case userstore.ErrorByEmail:
			auth.ErrorEmail = err
		case userstore.ErrorPassword:
			auth.ErrorPass = err
		}

		return c.Render("./static/login.html", fiber.Map{
			"Error":   err,
			"Islogin": auth.Islogin,
		})
	}
	sess.Set("session_id", "super-test")
	if err = sess.Save(); err != nil {
		log.Println("can't save user session in store", err)
		return err
	}
	return c.Redirect("/personal/kyc")
}

// Render Sign In page
func HandleLogin(c *fiber.Ctx) error {

	sess, err := store.Get(c)
	if err != nil {
		log.Println("can't get user cookie", err)
	}

	value := sess.Get("session_id")
	if value == nil {
		sess.Destroy()
		return c.Render("./static/login.html", fiber.Map{
			"Jonny":   "Hello, World!",
			"Islogin": false,
		})
	}

	return c.Redirect("/personal/kyc")
}

// Handle Post request to Create New User
func HandleRegister(c *fiber.Ctx) error {

	sess, err := store.Get(c)
	if err != nil {
		log.Println(err)
	}

	cookie := c.Cookies("session_id")
	if len(cookie) > 1 {
		c.Redirect("/personal/kyc", fiber.StatusAccepted)
	}

	//Create User
	user := model.User{
		Email:    c.FormValue("email"),
		Password: c.FormValue("password"),
	}

	store := userstore.NewStore()
	if err = store.CreateUser(c, user); err != nil {
		log.Println("can't create user", err)
		return c.Redirect("/signup", fiber.StatusMethodNotAllowed)
	}

	val := &fiber.Cookie{
		Name: "session_id",
		// Set expiry date to the past
		Expires:  time.Now().Add(-(time.Hour * 2)),
		HTTPOnly: true,
		Value:    "super_test",
		SameSite: "lax",
	}
	c.Cookie(val)
	defer sess.Save()
	sess.Set(val.Name, val.Value)

	return c.Redirect("/")
}

// Hanlde Get request
func HandleSignUp(c *fiber.Ctx) error {

	sess, err := store.Get(c)
	if err != nil {
		log.Println("can't get user cookie", err)
	}

	value := sess.Get("session_id")
	if value == nil {
		sess.Destroy()
		return c.Render("./static/signup.html", fiber.Map{
			"Jonny":   "Hello, World!",
			"Islogin": false,
		})
	}

	return c.Redirect("/personal/kyc")
}

func HandleRecover(c *fiber.Ctx) error {
	return c.Render("./static/recover.html", fiber.Map{
		"Jonny": "Hello, World!",
	})
}

func HandleKyc(c *fiber.Ctx) error {
	return c.Render("./static/account-overview.html", fiber.Map{
		"Jonny": "Hello, World!",
	})
}

func HandlePayments(c *fiber.Ctx) error {
	return c.Render("./static/account-billing.html", fiber.Map{
		"Jonny": "Hello, World!",
	})
}

func HandleKycInfo(c *fiber.Ctx) error {

	file_front, err := c.FormFile("front")
	if err != nil {
		log.Println("can't get front file", err)
	}

	file_back, err := c.FormFile("front")
	if err != nil {
		log.Println("can't get front file", err)
	}

	dist_front := fmt.Sprintf("/uloads/%s/%s", "", file_front.Filename)
	dist_back := fmt.Sprintf("/uloads/%s/%s", "", file_back.Filename)

	c.SaveFile(file_front, dist_front)
	c.SaveFile(file_back, dist_back)

	return c.Render("/personal/kyc", fiber.Map{
		"KycStatusBio":     "inporgress",
		"KycStatusAddress": "f",
	})
}

func CheckAuth(handler fiber.Handler) fiber.Handler {

	return func(c *fiber.Ctx) error {
		session := c.Cookies("session_id")
		if len(session) < 1 {
			c.ClearCookie("session_id")
		}
		return handler(c)
	}
}
