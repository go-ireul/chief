package main

import (
	"context"
	"flag"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"
	"ireul.com/bolt"

	"ireul.com/chief/types"
)

var db *bolt.DB
var dbFile string
var bind string

type chiefServer struct{}

// NewID implements ChiefServer
func (chiefServer) NewID(ctx context.Context, req *types.NewIDRequest) (resp *types.NewIDResponse, err error) {
	var id uint64
	if id, err = newID(db, req.GetName()); err != nil {
		return
	}
	resp = &types.NewIDResponse{
		Code: types.NewIDResponse_OK,
		Name: req.GetName(),
		ID:   id,
	}
	return
}

func main() {
	flag.StringVar(&dbFile, "db", "chief.db", "database file")
	flag.StringVar(&bind, "bind", ":9000", "bind address with format [IP:PORT]")
	flag.Parse()

	var err error
	if db, err = bolt.Open(dbFile, 0600, nil); err != nil {
		log.Println("failed to open db:", dbFile)
		return
	}
	defer db.Close()

	var lis net.Listener
	if lis, err = net.Listen("tcp", bind); err != nil {
		log.Println("failed to listen:", bind)
		return
	}
	s := grpc.NewServer()
	types.RegisterChiefServer(s, chiefServer{})

	// catch signal
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		for range signalChan {
			s.Stop()
			break
		}
	}()

	// start server
	s.Serve(lis)
}
