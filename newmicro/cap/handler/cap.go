package handler

import (
	"context"

	log "github.com/micro/micro/v3/service/logger"

	pb "cap/proto"
)

type Cap struct{}

// Return a new handler
func New() *Cap {
	return &Cap{}
}

// Call is a single request handler called via client.Call or the generated client code
func (e *Cap) Call(ctx context.Context, req *pb.Request, rsp *pb.Response) error {
	log.Info("Received Cap.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Cap) Stream(ctx context.Context, req *pb.StreamingRequest, stream pb.Cap_StreamStream) error {
	log.Infof("Received Cap.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Infof("Responding: %d", i)
		if err := stream.Send(&pb.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Cap) PingPong(ctx context.Context, stream pb.Cap_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&pb.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
