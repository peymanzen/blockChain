package networking

import (
	"bytes"
	"net/http"
)

func SendMessage(nodeAddress string,message []byte) error{
	_, err := http.Post("http://"+nodeAddress+"/message", "application/json", bytes.NewBuffer(message))
    return err
}