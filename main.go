package main

import (
	"context"
	"flag"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/boltdb/bolt"
	"google.golang.org/grpc"

	"github.com/yankeguo/chief/types"
)

var db *bolt.DB
var dbFile string
var bind string

var shard uint64
var shardPrefix uint64
var shardMask = ^uint64(1) >> 4
var shardSize = uint64(1) << 60

var exitCode int

type chiefServer struct{}

// NewID implements ChiefServer
func (chiefServer) NewID(ctx context.Context, req *types.NewIDRequest) (resp *types.NewIDResponse, err error) {
	var id uint64
	if id, err = newID(db, req.GetName()); err != nil {
		return
	}
	resp = &types.NewIDResponse{
		Name: req.GetName(),
		ID:   id,
	}
	return
}

func main() {
	// defered os.Exit with code
	defer os.Exit(exitCode)

	// parse flags
	flag.StringVar(&dbFile, "db", "chief.db", "database file")
	flag.StringVar(&bind, "bind", ":9000", "bind address with format [IP:PORT]")
	flag.Uint64Var(&shard, "shard", 1, "shard configuration")
	flag.Parse()

	// check shard
	if shard < 1 || shard > 16 {
		log.Println("shard should be a integer from 1 to 16")
		exitCode = 1
		return
	}

	// shift shard
	shardPrefix = shard << 60 // 16 = 2 ** 4

	// open DB
	var err error
	if db, err = bolt.Open(dbFile, 0660, nil); err != nil {
		log.Println("failed to open db:", dbFile)
		exitCode = 1
		return
	}
	defer db.Close()

	// listen
	var lis net.Listener
	if lis, err = net.Listen("tcp", bind); err != nil {
		log.Println("failed to listen:", bind)
		exitCode = 1
		return
	}
	s := grpc.NewServer()
	types.RegisterChiefServer(s, chiefServer{})

	// catch signal
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		for sig := range signalChan {
			log.Printf("exiting on signal [%v]\n", sig)
			s.Stop()
			break
		}
	}()

	// start server
	log.Printf("listening at [%v], db file at [%v]\n", bind, dbFile)
	s.Serve(lis)
}
