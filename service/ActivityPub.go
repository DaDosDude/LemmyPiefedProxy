package service

import (
	"LemmyBeProxy/dto/model/ap"
	"encoding/json"
	"io"
	"net/http"
)

type ActivityPub struct {
}

func NewActivityPub() *ActivityPub {
	return &ActivityPub{}
}

func (receiver *ActivityPub) FetchActor(actorId string) (result ap.Actor, err error) {
	req, err := http.NewRequest(http.MethodGet, actorId, nil)
	if err != nil {
		return
	}
	req.Header.Set("Accept", "application/activity+json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &result)
	return
}
