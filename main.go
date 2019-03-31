package main

import (
	"fmt"
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine/mail"
)

func main() {
	http.HandleFunc("/", indexHandler)
	appengine.Main() // Starts the server to receive requests
}

// [END gae_go_env_main]
// [START gae_go_env_index]
func indexHandler(w http.ResponseWriter, r *http.Request) {
	// if statement redirects all invalid URLs to the root homepage.
	// Ex: if URL is http://[YOUR_PROJECT_ID].appspot.com/FOO, it will be
	// redirected to http://[YOUR_PROJECT_ID].appspot.com.
	if r.URL.Path == "/daily-mail" {
		dailyMail(w, r)
		return
	}

	fmt.Fprintln(w, "Wrong request!")
}


// const fromEmail = "Me <luongvandu@gmail.com	>"
const fromEmail = "automation-jobs@devsamurai-211600.appspotmail.com"
const mailTitle = "Daily Batch Job Notification"
const confirmMessage = `
This is automation Notification from app engine.
This job is run daily to test.
`

func dailyMail(w http.ResponseWriter, r *http.Request) {

	ctx 	:= appengine.NewContext(r)
	toEmail := []string{"Me <luongvandu@gmail.com>"}
	// addr := r.FormValue("email")
	// url := createConfirmationURL(r)
	msg := &mail.Message{
		Sender 	: fromEmail,
		To 		: toEmail,
		Subject : mailTitle,
		Body 	: fmt.Sprintf(confirmMessage),
	}
	if err := mail.Send(ctx, msg); err != nil {
		log.Errorf(ctx, "Couldn't send email: %v", err)
		fmt.Fprintln(w, "Couldn't send email: %v", err)
	} else {
		fmt.Fprintln(w, "The mail was sent.")
	}
}
