package recaptcha

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

const requestURL = "https://www.google.com/recaptcha/api/siteverify"

// Check - Проверяем капчу
func Check(r *http.Request, secret string) (ans Ans) {
	// Формируем запрос
	q := url.Values{}
	q.Add("response", r.FormValue("g-recaptcha-response"))
	q.Add("remoteip", r.Header.Get("X-Real-IP"))
	q.Add("secret", secret)

	resp, err := http.PostForm(requestURL, q)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		log.Println("[error]", err)
		return
	}

	// Читаем овтет
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("[error]", err)
		return
	}

	err = json.Unmarshal(content, &ans)
	if err != nil {
		log.Println("[error]", err)
		return
	}

	return
}

// CheckGQL - Проверяем капчу fql
func CheckGQL(response, ip, secret string) (ans Ans) {
	// Формируем запрос
	q := url.Values{}
	q.Add("response", response)
	q.Add("remoteip", ip)
	q.Add("secret", secret)

	resp, err := http.PostForm(requestURL, q)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		log.Println("[error]", err)
		return
	}

	// Читаем овтет
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("[error]", err)
		return
	}

	err = json.Unmarshal(content, &ans)
	if err != nil {
		log.Println("[error]", err)
		return
	}

	return
}
