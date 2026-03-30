package core

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/ensoria/gofake/pkg/faker/common/log"
)

// RandTime provides methods for generating random time.Time and time.Duration values.
// Note: Unlike the datetime generator, which returns formatted date strings,
// RandTime returns time.Time and time.Duration values directly.
//
// ランダムなtime.Timeおよびtime.Durationの値を生成するメソッドを提供する構造体。
// 注: datetimeジェネレーターはフォーマットされた日時文字列を返すが、
// RandTimeはtime.Timeやtime.Durationの値を直接返す。

type RandTime struct {
	rand *rand.Rand
}

// NewRandTime creates a new RandTime instance with the given random source.
// 指定されたランダムソースで新しいRandTimeインスタンスを作成する。
func NewRandTime(rand *rand.Rand) *RandTime {
	return &RandTime{
		rand,
	}
}

// PastFuture returns a random time between 30 years ago and 30 years from now.
// 30年前から30年後までのランダムな日時を返す。
func (r *RandTime) PastFuture() time.Time {
	past30Years := time.Now().Add(-30 * 365 * 24 * time.Hour)
	future30Years := time.Now().Add(30 * 365 * 24 * time.Hour)
	return r.TimeRange(past30Years, future30Years)
}

// PastFrom returns a random time between the given past time and now.
// 指定した過去の日時から現在までのランダムな日時を返す。
func (r *RandTime) PastFrom(from time.Time) time.Time {
	if from.After(time.Now()) {
		errMsg := fmt.Sprintf("Invalid past time: from=%#v", from)
		log.WrongUsage(errMsg, 1)
		return time.Time{}
	}
	return r.TimeRange(from, time.Now())
}

// Past returns a random time between 30 years ago and now.
// 30年前から現在までのランダムな日時を返す。
func (r *RandTime) Past() time.Time {
	past30Years := time.Now().Add(-30 * 365 * 24 * time.Hour)
	return r.PastFrom(past30Years)
}

// FutureTo returns a random time between now and the given future time.
// 現在から指定した未来の日時までのランダムな日時を返す。
func (r *RandTime) FutureTo(to time.Time) time.Time {
	if to.Before(time.Now()) {
		errMsg := fmt.Sprintf("Invalid future time: to=%#v", to)
		log.WrongUsage(errMsg, 1)
		return time.Time{}
	}
	return r.TimeRange(time.Now(), to)
}

// Future returns a random time between now and 30 years from now.
// 現在から30年後までのランダムな日時を返す。
func (r *RandTime) Future() time.Time {
	future30Years := time.Now().Add(30 * 365 * 24 * time.Hour)
	return r.FutureTo(future30Years)
}

// TimeRange returns a random time within the specified range [from, to].
// 指定された範囲[from, to]の中でランダムな日時を返す。
func (r *RandTime) TimeRange(from time.Time, to time.Time) time.Time {
	if from.After(to) {
		errMsg := fmt.Sprintf("Invalid range: from=%#v, to=%#v", from, to)
		log.WrongUsage(errMsg, 1)
		return time.Time{}
	}
	diff := to.Sub(from)
	randomDiff := time.Duration(r.rand.Int63n(int64(diff)))
	return from.Add(randomDiff)
}

// Duration returns a random duration.
// ランダムなDurationを返す。
func (r *RandTime) Duration() time.Duration {
	return time.Duration(r.rand.Int63())
}

// DurationMilliSec returns a random duration up to 1 second.
// 1秒以下のランダムなDurationを返す。
func (r *RandTime) DurationMilliSec() time.Duration {
	return r.DurationTo(1000 * time.Millisecond)
}

// DurationSec returns a random duration up to 1 minute.
// 1分以下のランダムなDurationを返す。
func (r *RandTime) DurationSec() time.Duration {
	return r.DurationTo(60 * time.Second)
}

// DurationMin returns a random duration up to 1 hour.
// 1時間以下のランダムなDurationを返す。
func (r *RandTime) DurationMin() time.Duration {
	return r.DurationTo(60 * time.Minute)
}

// DurationHour returns a random duration up to 1 day.
// 1日以下のランダムなDurationを返す。
func (r *RandTime) DurationHour() time.Duration {
	return r.DurationTo(24 * time.Hour)
}

// DurationTo returns a random duration up to the specified maximum (inclusive).
// 指定した最大値以下のランダムなDurationを返す。toで指定した値も含まれる。
func (r *RandTime) DurationTo(to time.Duration) time.Duration {
	return time.Duration(r.rand.Int63n(int64(to + 1)))
}

// DurationRange returns a random duration within the specified range [from, to] (inclusive).
// 指定された範囲[from, to]の中でランダムなDurationを返す。fromとtoで指定した値も含まれる。
func (r *RandTime) DurationRange(from time.Duration, to time.Duration) time.Duration {
	if from > to {
		errMsg := fmt.Sprintf("Invalid range: from=%#v, to=%#v", from, to)
		log.WrongUsage(errMsg, 1)
		return 0
	}
	diff := to - from
	randomDiff := time.Duration(r.rand.Int63n(int64(diff + 1)))
	return from + randomDiff
}
