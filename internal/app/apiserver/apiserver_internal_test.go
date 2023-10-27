package apiserver

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func testAPIserver_HandleHello(t *testing.T) {
	s := New(NewConfiпg())                                   // сервер, который принимает NewConfig()
	rec := httptest.NewRecorder()                            // рекордер
	req, _ := http.NewRequest(http.MethodGet, "/hello", nil) //объекту request передаем метод GET, путь "/hello", nill

	s.handleHello().ServeHTTP(rec, req) //передаем рекордер и запрос

	// с помощью библиотеки assert мы проверим, что наш rec.Body.String() равен "Hello"
	assert.Equal(t, rec.Body.String(), "Hello")
}
