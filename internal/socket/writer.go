package socket

//
//import (
//	"github.com/gorilla/websocket"
//	"time"
//)
//
//func Writer(ws *websocket.Conn, function func()) {
//	pingTicker := time.NewTicker(60 * time.Second * 9 / 10)
//	fileTicker := time.NewTicker(10 * time.Second)
//	defer func() {
//		pingTicker.Stop()
//		fileTicker.Stop()
//		err := ws.Close()
//		if err != nil {
//			return
//		}
//	}()
//}
