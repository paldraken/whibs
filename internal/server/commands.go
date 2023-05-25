package server

import (
	"encoding/json"
	"fmt"

	"github.com/paldraken/sqldebugwatch/internal/types"
)

type baseCommand struct {
	Command string `json:"command"`
}

type filterCommand struct {
	baseCommand
	Payload *types.Filter `json:"payload"`
}

func command(msgRaw []byte, u *user) {

	cmd := &baseCommand{}
	err := json.Unmarshal(msgRaw, cmd)
	if err != nil {
		fmt.Println(err)
		return
	}

	switch cmd.Command {
	case "filter":
		apllyFilterForUser(msgRaw, u)
	}
}

func apllyFilterForUser(msgRaw []byte, u *user) error {
	fcmd := &filterCommand{}

	err := json.Unmarshal(msgRaw, fcmd)
	if err != nil {
		return err
	}

	u.filter = *fcmd.Payload

	return nil
}
