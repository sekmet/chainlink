package adapters

import (
	"errors"

	simplejson "github.com/bitly/go-simplejson"
	"github.com/smartcontractkit/chainlink/store"
	"github.com/smartcontractkit/chainlink/store/models"
)

type JsonParse struct {
	Path []string `json:"path"`
}

func (jpa *JsonParse) Perform(input models.RunResult, _ *store.Store) models.RunResult {
	js, err := simplejson.NewJson([]byte(input.Value()))
	if err != nil {
		return models.RunResultWithError(err)
	}

	js, err = checkEarlyPath(js, jpa.Path)
	if err != nil {
		return models.RunResultWithError(err)
	}

	rval, ok := js.CheckGet(jpa.Path[len(jpa.Path)-1])
	if !ok {
		return models.RunResult{}
	}

	return models.RunResultWithValue(rval.MustString())
}

func checkEarlyPath(js *simplejson.Json, path []string) (*simplejson.Json, error) {
	var ok bool
	for _, k := range path[:len(path)-1] {
		js, ok = js.CheckGet(k)
		if !ok {
			return js, errors.New("No value could be found for the key '" + k + "'")
		}
	}
	return js, nil
}
