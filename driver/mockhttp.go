package driver

import (
	"context"
	"encoding/json"
	"fmt"
)

type Key string

const (
	BodyKey = Key("body")
)

type MockHttpFramework struct{}

func NewMockHttpFramework() *MockHttpFramework {
	return &MockHttpFramework{}
}

func (d *MockHttpFramework) UriParam(ctx context.Context, key string) string {
	s, _ := ctx.Value(Key(key)).(string)
	return s
}

func (d *MockHttpFramework) BindJsonBody(ctx context.Context, modelPointer any) error {
	body, _ := ctx.Value(BodyKey).([]byte)
	return json.Unmarshal(body, modelPointer)
}

func (d *MockHttpFramework) JSON(ctx context.Context, httpCode int, jsonData any) error {
	data, err := json.Marshal(jsonData)
	if err != nil {
		return err
	}
	fmt.Println(string(data))
	return nil
}

// To set mock json body
func (d *MockHttpFramework) ContextWithBody(ctx context.Context, body []byte) context.Context {
	return context.WithValue(ctx, BodyKey, body)
}

func (d *MockHttpFramework) ContextWithUriParam(ctx context.Context, key string, param string) context.Context {
	return context.WithValue(ctx, Key(key), param)
}
