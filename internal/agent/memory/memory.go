package memory

import (
	"fmt"
	"sync"

	"github.com/ulixes-bloom/ya-metrics/internal/pkg/errors"
	"github.com/ulixes-bloom/ya-metrics/internal/pkg/metrics"
)

type storage struct {
	metrics map[string]metrics.Metric
	mutex   sync.RWMutex
}

func NewStorage() *storage {
	m := storage{
		metrics: make(map[string]metrics.Metric, metrics.MetricsCount),
	}
	return &m
}

func (s *storage) Set(metric metrics.Metric) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	switch metric.MType {
	case metrics.Counter:
		cur, ok := s.metrics[metric.ID]
		if ok {
			newDelta := (metric.GetDelta() + cur.GetDelta())
			metric.Delta = &newDelta
			s.metrics[metric.ID] = metric
		} else {
			s.metrics[metric.ID] = metric
		}
	case metrics.Gauge:
		s.metrics[metric.ID] = metric
	default:
		return errors.ErrMetricTypeNotImplemented
	}

	return nil
}

func (s *storage) SetAll(meticsSlice []metrics.Metric) error {
	for _, m := range meticsSlice {
		if err := s.Set(m); err != nil {
			return fmt.Errorf("memory.setAll: %w", err)
		}
	}

	return nil
}

func (s *storage) GetAll() map[string]metrics.Metric {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	return s.metrics
}
