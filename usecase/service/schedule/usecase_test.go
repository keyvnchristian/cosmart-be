package schedule_test

import (
	"cosmart/mocks/mocks_repository"
	"cosmart/model"
	"cosmart/usecase/service/schedule"
	"errors"
	"github.com/golang/mock/gomock"
	"testing"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestSchedule(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Schedule Suite")
}

var _ = Describe("Books Test", func() {
	var (
		mockCtrl   *gomock.Controller
		repository *mocks_repository.MockRepository
		usecase    schedule.ScheduleUsecase

		schedReq model.ScheduleRequest
		storeReq model.Schedule
		resp     []model.Schedule
	)

	AfterEach(func() {
		time.Sleep(300 * time.Millisecond)
		mockCtrl.Finish()
	})

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		mockCtrl.Finish()

		repository = mocks_repository.NewMockRepository(mockCtrl)
		usecase = schedule.NewScheduleUsecase(repository)
		schedReq = model.ScheduleRequest{
			UserID:         1,
			Title:          "Jesus is Lord",
			Author:         "Paul the Apostle",
			Edition:        2,
			PickupSchedule: "2023-11-10",
		}

		date, _ := time.Parse(time.DateOnly, schedReq.PickupSchedule)
		storeReq = model.Schedule{
			UserID:     schedReq.UserID,
			Title:      schedReq.Title,
			Author:     schedReq.Author,
			Edition:    schedReq.Edition,
			PickUpTime: date,
		}

		resp = []model.Schedule{
			{
				UserID:     schedReq.UserID,
				Title:      schedReq.Title,
				Author:     schedReq.Author,
				Edition:    schedReq.Edition,
				PickUpTime: date,
			},
		}
	})

	Describe("Add Schedule", func() {
		Context("error", func() {
			It("will return err on store book", func() {
				repository.EXPECT().StoreBook(storeReq).Return(resp, errors.New("error"))

				_, err := usecase.CreateSchedule(schedReq)
				Expect(err).ToNot(BeNil())
			})
		})

		Context("success", func() {
			It("will add new schedule to the list", func() {
				repository.EXPECT().StoreBook(storeReq).Return(resp, nil)

				resp, err := usecase.CreateSchedule(schedReq)
				Expect(len(resp)).To(Equal(1))
				Expect(err).To(BeNil())
			})
		})
	})
})
