package main

import (
	_ "little-contacts/routers"
	"log"
	"os"
	"strconv"

	context "github.com/astaxie/beego/context"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
)

func main() {
	if os.Getenv("PORT") != "" {
		log.Println("Env $PORT :", os.Getenv("PORT"))
		port, err := strconv.Atoi(os.Getenv("PORT"))
		if err != nil {
			log.Fatal(err)
			log.Fatal("$PORT must be set")
		}
		log.Println("port : ", port)
		beego.BConfig.Listen.HTTPSPort = port
	}
	if os.Getenv("BEEGO_ENV") != "" {
		log.Println("Env $BEEGO_ENV :", os.Getenv("BEEGO_ENV"))
		beego.BConfig.RunMode = os.Getenv("BEEGO_ENV")
	}

	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins:     []string{"localhost:8080", "localhost:3000", "*"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "OPTIONS", "DELETE"},
		AllowHeaders:     []string{"Origin", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "X-Requested-With", "Content-Type", "Accept", "Connection", "Upgrade", "Token", "Authorization", "Websocket", "Set-Cookie", "withCredentials"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Origin", "Connection", "Upgrade", "Token", "Authorization", "Websocket", "Set-Cookie", "withCredentials"},
		AllowCredentials: true,
	}))

	beego.InsertFilter("*", beego.BeforeRouter, func(ctx *context.Context) {
		ctx.Output.Header("Server", "bded553bd17da4d5409ba131684770bc")
		ctx.Output.Header("Content-Security-Policy", "default-src 'self'; style-src 'self' https://fonts.googleapis.com https://www.gstatic.com 'unsafe-inline' blob: ; font-src 'self' data:  https://fonts.gstatic.com https://cdnjs.cloudflare.com; img-src 'self' data: ; connect-src 'self' https://viacep.com.br https://sentry.io; script-src 'self' 'unsafe-inline' 'unsafe-eval' data: https://www.google.com https://www.gstatic.com https://fonts.gstatic.com https://cdnjs.cloudflare.com *.mathjax.org ; frame-src https://www.google.com ; child-src https://www.google.com ; object-src 'self' blob: ;")
		ctx.Output.Header("X-Frame-Options", "DENY")
		ctx.Output.Header("X-Xss-Protection", "1; mode=block")
		ctx.Output.Header("X-Content-Type-Options", "nosniff")
		ctx.Output.Header("Referrer-Policy", "no-referrer")
		ctx.Output.Header("Strict-Transport-Security", "max-age=31536000")
		ctx.Output.Header("Cache-Control", "no-cache, no-store, must-revalidate")
		ctx.Output.Header("Pragma", "no-cache")
		ctx.Output.Header("Expires", "0")
		//ctx.Output.Header("Public-Key-Pins", `pin-sha256="`+fingerprint+`"; max-age=2592000;`)
	})

	beego.Run()
}
