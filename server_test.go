package app

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"testing"
)

type mockStore struct {
	s map[string]string
}

func newMockStore() Store {
	return &mockStore{s: make(map[string]string)}
}

func (m *mockStore) Set(k, v string) error  {
	m.s[k] = v
	return nil
}

func (m *mockStore) Get(k string) (string, error) {
	if v, ok := m.s[k]; ok {
		return v, nil
	}
	return "", errors.New("key not found")
}

func TestHandleSet(t *testing.T) {
	s := newMockStore()
	handler := newHandler(s)

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/set", nil)
	r.URL.RawQuery = "key=test&value=test"

	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = r
	handler.handleSet(ctx)

	if w.Code != http.StatusOK {
		t.Fatalf("expected code to be 200 OK. Got %d", w.Code)
	}
}

func TestRedisStore_SetBadRequest(t *testing.T) {
	s := newMockStore()
	handler := newHandler(s)

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/set", nil)

	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = r
	handler.handleSet(ctx)

	if w.Code != http.StatusBadRequest {
		t.Fatalf("expected code to be 400 bad request. Got %d", w.Code)
	}
}

func TestRedisStore_Get(t *testing.T) {
	s := newMockStore()
	handler := newHandler(s)

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/set", nil)
	r.URL.RawQuery = "key=test&value=test"

	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = r
	handler.handleSet(ctx)

	if w.Code != http.StatusOK {
		t.Fatalf("expected http code to be 200 OK. Got %d", w.Code)
	}

	w_ := httptest.NewRecorder()
	r_ := httptest.NewRequest("GET", "/get", nil)
	r_.URL.RawQuery = "key=test"

	ctx_, _ := gin.CreateTestContext(w_)
	ctx_.Request = r_
	handler.handleGet(ctx_)

	if w_.Code != http.StatusOK {
		t.Fatalf("expected http code to be 200 OK. Got %d", w_.Code)
	}
}

func TestRedisStore_Get_BadRequest(t *testing.T) {
	s := newMockStore()
	handler := newHandler(s)

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/get", nil)

	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = r
	handler.handleGet(ctx)

	if w.Code != http.StatusBadRequest {
		t.Fatalf("expected code to be 400 Bad request. Got %d", w.Code)
	}
}