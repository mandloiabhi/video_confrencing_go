package server

import (
	"flag"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofibre/template/html"
	"github.com/gofibre/websocket/v2"
	"golang.org/x/net/websocket"
	"nhooyr.io/websocket"
)

var(
	addr=flag.String("addr,":"",os.Getenv("PORT"),"")
	cert=flag.String("cert","","") // certificate
	key=flag.String("key","","")
)

func Run() error{
	flag.Parse()
	if *addr==":"{
		*addr=":8080"
	}

    engine :=html.New("./views",".html")  // giving access to engine variable all the files of html in view folder
     
	app :=fiber.New(fiber.Config{
		Views: engine,
	})
     
    app.Use(logger.New())
	app.Use(cors.New())



	app.Get("/",handlers.Welcome)
	app.Get("/room/create",handlers.RoomCreate)
	app.Get("/room/uuid",handlers.Room)
	app.Get("/room/uuid/websocket",websocket.New(handlers.RoomWebsocket,websocket.Config{
		HandshakeTimeout: 10*time.Second,
	}))
	app.Get("/room/uuid/chat",handlers.RoomChat)
	app.Get("/room/uuid/chat/websocket",websocket.New(handlers.RoomChatWebsocket))
	app.Get("/room/uuid/viewer/websocket",websocket.New(handlers.RoomViewerWebsocket))
	app.Get("/stream/:ssuid",handlers.Stream)
	app.Get("/stream/:ssuid/websocket",)
	app.Get("/stream/:ssuid/chat/websocket",)
    app.Get("/stream/:ssuid/viewer/websocket",)
}