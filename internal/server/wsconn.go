package server

import (
	"encoding/json"
	"fmt"
	"net"
	"sort"
	"sync"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"github.com/paldraken/sqldebugwatch/internal/types"
)

type dMsg struct {
	Type    string             `json:"type"`
	Payload *types.SqlDebugRow `json:"payload"`
}

type user struct {
	conn   net.Conn
	uid    int
	filter types.Filter
}

type connections struct {
	m    sync.Mutex
	us   []*user
	ucnt int
}

func (c *connections) add(cn net.Conn) *user {
	u := &user{
		conn: cn,
	}

	c.m.Lock()
	{
		u.uid = c.ucnt
		c.us = append(c.us, u)
		c.ucnt++
	}
	c.m.Unlock()
	return u
}

func (c *connections) del(u *user) {
	c.m.Lock()
	defer c.m.Unlock()
	i := sort.Search(len(c.us), func(i int) bool {
		return c.us[i].uid >= u.uid
	})

	if i >= len(c.us) {
		panic("inconsistent state")
	}
	without := make([]*user, len(c.us)-1)
	copy(without[:i], c.us[:i])
	copy(without[i:], c.us[i+1:])
	c.us = without
}

func (c *connections) broadcast(row *types.SqlDebugRow) {

	oMsg := &dMsg{
		Type:    "debug_row",
		Payload: row,
	}

	msg, err := json.Marshal(oMsg)
	if err != nil {
		fmt.Println(err)
	}

	for _, u := range c.us {
		if !u.filter.IsPassed(oMsg.Payload) {
			continue
		}
		err := wsutil.WriteServerMessage(u.conn, ws.OpText, msg)
		if err != nil {
			fmt.Println(err)
		}
	}
}

var cns = &connections{}
