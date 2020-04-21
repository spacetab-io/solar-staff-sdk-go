package rest

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMakeSignature(t *testing.T) {
	bodyParamsInp := map[string]interface{}{
		"client_id": 6,
		"action":    "workers_list",
	}

	if !assert.Equal(t, "19861f409729a42c2a8c0c636cfa0a4fb845e8fb", makeSignature("salt", bodyParamsInp)) {
		t.FailNow()
	}
}
