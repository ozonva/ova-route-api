package saver_test

import (
	"context"
	"ova-route-api/internal/flusher"
	"ova-route-api/internal/models"
	"ova-route-api/internal/repository/mocks"
	"ova-route-api/internal/saver"
	"time"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Flusher", func() {
	var (
		ctrl        *gomock.Controller
		mockRepo    *mocks.MockRepo
		testFlusher flusher.Flusher
		testSaver   saver.Saver

		routes = []models.Route{
			{ID: 1, UserID: 1, RouteName: "name1", Length: 1},
			{ID: 2, UserID: 1, RouteName: "name2", Length: 2},
			{ID: 3, UserID: 1, RouteName: "name3", Length: 3},
		}
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		mockRepo = mocks.NewMockRepo(ctrl)
		testFlusher = flusher.NewFlusher(1, mockRepo)
		testSaver = saver.NewSaver(2, testFlusher)
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	Describe("Positive set", func() {
		It("Flush by timeout", func() {
			gomock.InOrder(
				mockRepo.EXPECT().AddRoutes(routes[:1]).Return(nil).Times(1),
				mockRepo.EXPECT().AddRoutes(routes[1:2]).Return(nil).Times(1),
				mockRepo.EXPECT().AddRoutes(routes[2:3]).Return(nil).Times(1),
			)
			for _, route := range routes {
				testSaver.Save(context.Background(), route)
			}

			Expect(testSaver.BuffSize()).Should(BeNumerically("==", 1))
			time.Sleep(3 * time.Second)
			Expect(testSaver.BuffSize()).Should(BeNumerically("==", 0))
		})

		It("Flush by call Close", func() {
			gomock.InOrder(
				mockRepo.EXPECT().AddRoutes(routes[:1]).Return(nil).Times(1),
				mockRepo.EXPECT().AddRoutes(routes[1:2]).Return(nil).Times(1),
				mockRepo.EXPECT().AddRoutes(routes[2:3]).Return(nil).Times(1),
			)
			for _, route := range routes {
				testSaver.Save(context.Background(), route)
			}

			Expect(testSaver.BuffSize()).Should(BeNumerically("==", 1))
			testSaver.Close(context.Background())
			Expect(testSaver.BuffSize()).Should(BeNumerically("==", 0))
		})
	})
})
