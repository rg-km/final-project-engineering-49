package main

import (
	"net/http/httptest"
	"strings"

	"belajar-golang/handler"
	"belajar-golang/repository"
	"belajar-golang/router"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("JWT", func() {
	Describe("/login", func() {
		It("should return non authorized when given wrong credential", func() {
			bodyReader := strings.NewReader(`{"email": "alamin@gmail.com", "password": "alamin"}`)
			repo := repository.NewDB(db)
			h := handler.NewRepo(repo)
			r := router.Setuprouter(h)
			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST"  , "/login", bodyReader)
			r.ServeHTTP(w, req)

			Expect(w.Code).To(Equal(401))
		})

		It("should return authorized when given correct credential", func() {
		
			bodyReader := strings.NewReader(`{"email": "alamin.ibad@gmail.com", "password": "alamin123"}`)
			repo := repository.NewDB(db)
			h := handler.NewRepo(repo)
			r := router.Setuprouter(h)
			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST"  , "/login", bodyReader)
			r.ServeHTTP(w, req)

			Expect(w.Code).To(Equal(200))
		})
	})
})
