package service

import (
	"context"
	"math/rand"
	"sync/atomic"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/slimsamil/KreisliGo-go-ss2023/src/livetickgo/grpc/livetickgo"
	"google.golang.org/protobuf/types/known/emptypb"
)

const retryTime = 5 * time.Second

type LivetickgoService struct {
	livetickgo.LivetickgoService
	counter    int32
	queue      chan livetickgo.Event
	retryQueue chan livetickgo.Event
	stop       chan struct{}
}

func NewLivetickgoService() *LivetickgoService {
	rand.Seed(time.Now().UnixNano())
	return &LivetickgoService{counter: 0,
		queue:      make(chan livetickgo.Event),
		retryQueue: make(chan livetickgo.Event),
		stop:       make(chan struct{}),
	}
}

func (s *LivetickgoService) SendEvent(_ context.Context, event *livetickgo.Event) (*emptypb.Empty, error) {
	entry := log.WithField("event", event)
	entry.Info(("Received event"))
	s.processEvent(event)
	return &emptypb.Empty{}, nil
}

func (s *LivetickgoService) processEvent(event *livetickgo.Event) {
	entry := log.WithField("event", event)
	go func(event livetickgo.Event) {
		entry.Info("Start processing event")
		event.Id = s.getUniqueId()
		time.Sleep(time.Duration(rand.Intn(9)+1) * time.Second)
		s.queue <- event
		entry.Info("Processing event finished")
	}(*event)
}

func (s *LivetickgoService) getUniqueId() int32 {
	return atomic.AddInt32(&s.counter, 1)
}

func (s *LivetickgoService) ProcessEvents(stream livetickgo.BankTransfer_ProcessEventsServer) error {
	return func() error {
		for {
			select {
			case <-stream.Context().Done():
				log.Info("Watching events cancelled from the client side")
				return nil
			case event := <-s.queue:
				id := event.Id
				entry := log.WithField("event", event)
				entry.Info("Sending event")
				if err := stream.Send(&event); err != nil {
					s.requeueEvent(&event)
					entry.WithError(err).Error("Error sending event")
					return err
				}
				entry.Info("Event sent. Waiting for processing response")
				response, err := stream.Recv()
				if err != nil {
					s.requeueEvent(&event)
					entry.WithError(err).Error("Error receiving processing response")
					return err
				}
				if response.Id != id {
					// NOTE: this is just a guard and not happening as event is local per connection
					entry.Error("Received processing response of a different event")
				} else {
					entry.Info("Processing response received")
				}
			}
		}
	}()
}

func (s *LivetickgoService) requeueEvent(event *livetickgo.Event) {
	entry := log.WithField("event", event)
	go func(event livetickgo.Event) {
		entry.Infof("Requeuing event. Wait for %f seconds", retryTime.Seconds())
		time.Sleep(retryTime)
		s.retryQueue <- event
		entry.Info("Requeued event")
	}(*event)
}

func (s *LivetickgoService) Start() {
	log.Info("Starting livetickgo service")
	go func() {
		for {
			select {
			case <-s.stop:
				break
			case event := <-s.retryQueue:
				s.queue <- event
			}
		}
	}()
}

func (s *LivetickgoService) Stop() {
	log.Info("Stopping livetickgo service")
	close(s.stop)
}
