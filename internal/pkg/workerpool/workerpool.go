package workerpool

import (
	"sync"

	"github.com/rs/zerolog/log"
)

type jobCh[T any] chan T
type jobHandler[T any] func(T) error

type pool[T any] struct {
	numOfWorkers int            // количество worker'ов в pool'е
	jobCh        jobCh[T]       // канал с job'ами
	jobHandler   jobHandler[T]  // обработчик job'ы
	wg           sync.WaitGroup // синхронизатор работы worker'ов в pool'e
}

func New[T any](numOfWorkers, jobChSize int, handler jobHandler[T]) *pool[T] {
	jobCh := make(chan T, jobChSize)
	p := pool[T]{
		jobCh:        jobCh,
		numOfWorkers: numOfWorkers,
		jobHandler:   handler,
	}

	for range p.numOfWorkers {
		go p.startWorker()
	}

	return &p
}

// Добавление новой job'ы в worker pool
func (p *pool[T]) Submit(job T) {
	p.wg.Add(1)
	p.jobCh <- job
}

// Остановка и ожидание окончания работы worker pool'а
func (p *pool[T]) StopAndWait() {
	close(p.jobCh)
	p.wg.Wait()
}

// Запуск worker'a
func (p *pool[T]) startWorker() {
	for {
		j := <-p.jobCh
		err := p.jobHandler(j)
		if err != nil {
			log.Error().Msg(err.Error())
		}
		p.wg.Done()
	}
}
