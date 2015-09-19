package rester_test

import (
	. "github.com/Iceyer/rester"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type User struct {
	Name  string
	Email string
}

var _ = Describe("Rester", func() {
	aoiURL := "https://api.iceyer.org/v0/test"
	restAPI := Rest("Rester", aoiURL)

	BeforeEach(func() {
	})

	Describe("Rester Request Test", func() {
		Context("Rester Create Request", func() {
			req := restAPI.Create(&User{})
			It("Should Be POST http request", func() {
				Expect(req.Url).To(Equal(aoiURL))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Payload).To(Equal(&User{}))
			})
		})

		Context("Rester Retrieve Request", func() {
			username := "Iceyer"
			req := restAPI.Retrieve(username).WithResponeAs(&User{})
			It("Should Be GET http request", func() {
				Expect(req.Url).To(Equal(aoiURL + "/" + username))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Result).To(Equal(&User{}))
			})
		})

		Context("Rester List Request", func() {
			req := restAPI.List("").WithResponeAs(&[]User{})
			It("Should Be GET http request", func() {
				Expect(req.Url).To(Equal(aoiURL))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Result).To(Equal(&[]User{}))
			})
		})

		Context("Rester Update Request", func() {
			username := "Iceyer"
			req := restAPI.Update(username, &User{})
			It("Should Be PUT http request", func() {
				Expect(req.Url).To(Equal(aoiURL + "/" + username))
				Expect(req.Method).To(Equal("PUT"))
				Expect(req.Payload).To(Equal(&User{}))
			})
		})

		Context("Rester Delete Request", func() {
			username := "Iceyer"
			req := restAPI.Delete(username).WithResponeAs(&User{})
			It("Should Be DELETE http request", func() {
				Expect(req.Url).To(Equal(aoiURL + "/" + username))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.Result).To(Equal(&User{}))
			})
		})

		Context("Rester Remove Request", func() {
			req := restAPI.Remove("").WithResponeAs(&[]User{})
			It("Should Be PUT http request", func() {
				Expect(req.Url).To(Equal(aoiURL))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.Result).To(Equal(&[]User{}))
			})
		})
	})
})
