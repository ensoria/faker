package core_test

import (
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/ensoria/faker/pkg/faker/common/util"
	"github.com/ensoria/faker/pkg/faker/core"
	"github.com/ensoria/faker/pkg/faker/testutil"
)

var _ = Describe("Tests for random time functions", func() {
	randTime := core.NewRandTime(util.RandSeed())

	past30Years := time.Now().Add(-30 * 365 * 24 * time.Hour)
	future30Years := time.Now().Add(30 * 365 * 24 * time.Hour)

	Describe("PastFuture", func() {
		It("should return any random past/future time", func() {
			r := randTime.PastFuture()

			Expect(r).To(BeTemporally(">=", past30Years))
			Expect(r).To(BeTemporally("<=", future30Years))
		})
	})

	Describe("PastFrom", func() {
		It("should return any random past time from the given time", func() {
			r := randTime.PastFrom(past30Years)

			Expect(r).To(BeTemporally(">=", past30Years))
			Expect(r).To(BeTemporally("<=", time.Now()))
		})
	})

	Describe("Past", func() {
		It("should return any random past time", func() {
			r := randTime.Past()

			Expect(r).To(BeTemporally(">=", past30Years))
		})
	})

	Describe("FutureTo", func() {
		It("should return any random future time to the given time", func() {
			r := randTime.FutureTo(future30Years)

			Expect(r).To(BeTemporally(">=", time.Now()))
			Expect(r).To(BeTemporally("<=", future30Years))
		})
	})

	Describe("Future", func() {
		It("should return any random future time", func() {
			r := randTime.Future()

			Expect(r).To(BeTemporally("<=", future30Years))
		})
	})

	Describe("TimeRange", func() {
		It("should return any random time in the given range", func() {
			r := randTime.TimeRange(past30Years, future30Years)

			Expect(r).To(BeTemporally(">=", past30Years))
			Expect(r).To(BeTemporally("<=", future30Years))
		})
	})

	Describe("Duration", func() {
		It("should return any random duration", func() {
			r := randTime.Duration()

			testutil.Output("RandTime.Duration", r)
		})
	})

	Describe("DurationMilliSec", func() {
		It("should return any random duration in milliseconds", func() {
			r := randTime.DurationMilliSec()

			testutil.Output("RandTime.DurationMilliSec", r)
		})
	})

	Describe("DurationSec", func() {
		It("should return any random duration in seconds", func() {
			r := randTime.DurationSec()

			testutil.Output("RandTime.DurationSec", r)
		})
	})

	Describe("DurationMin", func() {
		It("should return any random duration in minutes", func() {
			r := randTime.DurationMin()

			testutil.Output("RandTime.DurationMin", r)
		})
	})

	Describe("DurationHour", func() {
		It("should return any random duration in hours", func() {
			r := randTime.DurationHour()

			testutil.Output("RandTime.DurationHour", r)
		})
	})

	Describe("DurationTo", func() {
		It("should return any random duration to the given time", func() {
			r := randTime.DurationTo(1 * time.Second)

			Expect(r).To(BeNumerically("<=", 1*time.Second))
		})
	})

	Describe("DurationRange", func() {
		It("should return any random duration in the given range", func() {
			r := randTime.DurationRange(1*time.Second, 2*time.Second)

			Expect(r).To(BeNumerically(">=", 1*time.Second))
			Expect(r).To(BeNumerically("<=", 2*time.Second))
		})
	})
})
