package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var Counters = struct {
	AddSucceeds      prometheus.Counter
	AddMultiSucceeds prometheus.Counter
	UpdateSucceeds   prometheus.Counter
	DeleteSucceeds   prometheus.Counter
}{
	AddSucceeds:      addSucceeds,
	AddMultiSucceeds: addMultiSucceeds,
	UpdateSucceeds:   updateSucceeds,
	DeleteSucceeds:   deleteSucceeds,
}

var (
	addSucceeds = promauto.NewCounter(prometheus.CounterOpts{
		Name: "ova_song_api_create_calls_count_succeed",
		Help: "The total number of successful executions of the ADD method",
	})

	addMultiSucceeds = promauto.NewCounter(prometheus.CounterOpts{
		Name: "ova_song_api_create_multi_calls_count_succeed",
		Help: "The total number of successful executions of the ADDMULTI method",
	})

	updateSucceeds = promauto.NewCounter(prometheus.CounterOpts{
		Name: "ova_song_api_update_calls_count_succeed",
		Help: "The total number of successful executions of the UPDATE method",
	})

	deleteSucceeds = promauto.NewCounter(prometheus.CounterOpts{
		Name: "ova_song_api_delete_calls_count_succeed",
		Help: "The total number of successful executions of the DELETE method",
	})
)
