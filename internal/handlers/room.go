package handlers

import (
	"fmt"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	guuid "github.com/google/uuid"
	"nhooyr.io/websocket"
)

func RoomCreate( c *fiber.Ctx) error {
  return c.Redirect(fmt.Sprintf("/room/%s",gguid.New().String))
}

func Room( c *fiber.Ctx) error {
	uuid:=c.Params("uuid")
	if uuid==""{
		c.Status(400)
		return nil
	}

	uuid,suuid,_:=createOrGetRoom=(uuid)
}

func RoomWebsocket(c * websocket.Conn){   // so c has params that user passes to our backend from frontend
  uuid:= c.params("uuid")
  if uuid ==""{
	return 
  }

  _,_,room:=createOrGetRoom=(uuid)

}

func createOrGetRoom(uuid string) (string,string,Room)
{
	
}