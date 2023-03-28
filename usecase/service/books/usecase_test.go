package books_test

import (
	"context"
	"cosmart/mocks/mocks_routing/mocks_openlibrary"
	"cosmart/model"
	"cosmart/routing/openlibrary"
	"cosmart/usecase/service/books"
	"errors"
	"github.com/golang/mock/gomock"
	"testing"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestBooks(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Books Suite")
}

var _ = Describe("Books Test", func() {
	var (
		matcherAny         gomock.Matcher
		mockCtrl           *gomock.Controller
		openLibraryService *mocks_openlibrary.MockLibraryService
		booksUsecase       books.BooksUsecase

		req     openlibrary.LibraryRequest
		resp    openlibrary.LibraryResponse
		bookReq model.BookRequest
	)

	AfterEach(func() {
		time.Sleep(300 * time.Millisecond)
		mockCtrl.Finish()
	})

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		matcherAny = gomock.Any()
		mockCtrl.Finish()

		openLibraryService = mocks_openlibrary.NewMockLibraryService(mockCtrl)
		booksUsecase = books.NewBookUsecase(openLibraryService, 3*time.Second)

		req = openlibrary.LibraryRequest{
			Name:     "love",
			IsDetail: false,
		}

		resp = openlibrary.LibraryResponse{
			Name: "love",
			Works: []openlibrary.Works{
				{
					Title:   "Love of God",
					Edition: 1995,
					Authors: []openlibrary.Authors{
						{
							Name: "Paul the Apostle",
						},
					},
				},
			},
		}

		bookReq = model.BookRequest{
			Genre: "love",
		}
	})

	Describe("Fetch Books", func() {
		Context("error", func() {
			It("will return error on fetch from OpenLibrary", func() {
				openLibraryService.EXPECT().FetchBook(matcherAny, req).Return(resp, errors.New("error"))

				_, err := booksUsecase.FetchBooks(context.Background(), bookReq)
				Expect(err).NotTo(BeNil())
			})
		})

		Context("Success", func() {
			It("will return list of books", func() {
				openLibraryService.EXPECT().FetchBook(matcherAny, req).Return(resp, nil)

				resp, err := booksUsecase.FetchBooks(context.Background(), bookReq)
				Expect(len(resp)).To(Equal(1))
				Expect(err).To(BeNil())
			})
		})
	})
})
