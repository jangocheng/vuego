package model

import (
	"encoding/json"

	"gopkg.in/workanator/vuego.v1/errors"
)

// UniqueModel represents a map object which data is shared and valid through screen session.
// The model is equivalent to the following code snippet.
//
//    new Vue({
//      data: function() { return {/*MODEL*/} }
//    })
type UniqueModel struct {
	BasicModel
}

func (m *UniqueModel) Field(path ...string) ModelInitialer {
	return ModelInitial{
		Modeler: &FieldModel{
			Owner: m,
			Path:  path,
		},
	}
}

func (m *UniqueModel) Markup() (string, error) {
	m.RLock()
	defer m.RUnlock()

	if data, err := json.Marshal(m.BasicModel.container); err != nil {
		return "", errors.ErrMarkupFailed{
			Tag:    "script",
			Reason: err,
		}
	} else {
		return "<script>new Vue({data:function(){return " + string(data) + "}})</script>", nil
	}
}
