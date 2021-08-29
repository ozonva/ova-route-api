package flusher_test

import (
	"errors"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"ova_route_api/internal/flusher"
	"ova_route_api/internal/models"
	"ova_route_api/internal/repository/mocks"
)

var _ = Describe("Flusher", func() {
	var (
		ctrl        *gomock.Controller
		mockRepo    *mocks.MockRepo
		testFlusher flusher.Flusher

		routes = []models.Route{
			{ID: 1, UserID: 1, RouteName: "name1", Length: 1},
			{ID: 2, UserID: 1, RouteName: "name2", Length: 2},
			{ID: 3, UserID: 1, RouteName: "name3", Length: 3},
		}
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		mockRepo = mocks.NewMockRepo(ctrl)
		testFlusher = flusher.NewFlusher(3, mockRepo)

	})

	AfterEach(func() {
		ctrl.Finish()
	})

	Describe("Positive set", func() {

		It("Single route", func() {
			oneItem := routes[:1]
			mockRepo.EXPECT().AddEntities(oneItem).Return(nil)
			Expect(testFlusher.Flush(oneItem)).To(BeNil())
		})

		It("Multiple routes", func() {
			mockRepo.EXPECT().AddEntities(routes).Return(nil).AnyTimes()
			Expect(testFlusher.Flush(routes)).To(BeNil())
		})
	})

	Describe("Negative set", func() {

		It("Should return last route", func() {
			gomock.InOrder(
				mockRepo.EXPECT().AddEntities(routes[:2]).Return(nil).Times(1),
				mockRepo.EXPECT().AddEntities(routes[2:]).Return(errors.New("some error")).Times(1),
			)
			testFlusher = flusher.NewFlusher(2, mockRepo)
			result := testFlusher.Flush(routes)
			Expect(len(result)).Should(BeNumerically("==", 1))
		})
	})
})
