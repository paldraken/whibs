package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"github.com/paldraken/whibs/configs"
	"github.com/paldraken/whibs/internal/types"
)

func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	conn, _, _, err := ws.UpgradeHTTP(r, w)
	if err != nil {
		log.Println(err)
	}

	u := cns.add(conn)

	go func() error {
		defer func() {
			conn.Close()
			cns.del(u)
		}()

		for {
			msgStr, op, err := wsutil.ReadClientData(conn)
			if err != nil {
				return err
			}
			if op == ws.OpText {
				command(msgStr, u)
			}
		}
	}()
}

func Start() error {

	if !configs.ServerConfig.Enable {
		return nil
	}

	http.HandleFunc("/ws", wsEndpoint)

	port := fmt.Sprintf("%d", configs.ServerConfig.Port)
	fmt.Println("Server will start on 127.0.0.1:" + port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		return err
	}

	return nil
}

func NotifyWsUsers(row *types.SqlDebugRow) {
	cns.broadcast(row)
}
