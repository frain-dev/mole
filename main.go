package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/stripe/smokescreen/cmd"
	"github.com/stripe/smokescreen/pkg/smokescreen"
)

func main() {
	conf, err := cmd.NewConfiguration(nil, nil)
	if err != nil {
		log.Fatal(err)
	}

	if conf == nil {
		os.Exit(1)
	}

	password, ok := os.LookupEnv("PROXY_PASSWORD")
	if !ok {
		fmt.Println("missing environment variable: PROXY_PASSWORD")
		os.Exit(1)
	}

	conf.RoleFromRequest = func(request *http.Request) (string, error) {
		fail := func(err string) (string, error) { return "", fmt.Errorf(err) }

		auth := strings.SplitN(request.Header.Get("Proxy-Authorization"), " ", 2)

		if len(auth) != 2 || auth[0] != "Basic" {
			return fail("authorization failed")
		}

		payload, _ := base64.StdEncoding.DecodeString(auth[1])
		pair := strings.SplitN(string(payload), ":", 2)

		if len(pair) != 2 || pair[1] != password {
			return fail("authorization failed")
		}
		return "authed", nil
	}

	conf.Healthcheck = http.HandlerFunc(OK)
	smokescreen.StartWithConfig(conf, nil)

}

func OK(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("OK"))
}
