package handler

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHandler_jsonError(t *testing.T) {
	msg := "Json Message Error"
	result := jsonError(msg)

	require.Equal(t, []byte(`{"message":"Json Message Error"}`), result)
}
