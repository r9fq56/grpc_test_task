package main

import (
	"context"
	pb "grpc_test_task/make_interface_dump/ecommerce"
	"log"
	"time"

	"github.com/gofrs/uuid"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct {
	productMap map[string]*pb.MakeInterfaceDumpRequest
}

func (s *server) MakeInterfaceDump(ctx context.Context, in *pb.MakeInterfaceDumpRequest) (*pb.Dump, error) {
	out, err := uuid.NewV4()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error while generating Product ID", err)
	}
	in.Id = out.String()
	if s.productMap == nil {
		s.productMap = make(map[string]*pb.MakeInterfaceDumpRequest)
	}
	s.productMap[in.Id] = in

	WriteToDb(in.Id, in.Name, in.Mac, in.Ip, in.Dns, in.Gateway)

	return &pb.Dump{Id: in.Id}, status.New(codes.OK, "").Err()
}

func WriteToDb(id, name, mac, ip, dns, gateway string) {

	// var schema = `
	// CREATE TABLE intdumps (
	// 	id char(36) NOT NULL PRIMARY KEY,
	// 	create_date timestamp
	// );

	// CREATE TABLE dumpdata (
	// 	id_dump char(36),
	// 	mac char(17),
	// 	name char(50),
	// 	ip char(15),
	// 	dns char(15),
	// 	gateway char(30)
	// )`

	// db, err := sqlx.Connect("postgres", "host=db port=5432 user=testuser password='1212' dbname=dump sslmode=disable")
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// // exec the schema or fail; multi-statement Exec behavior varies between
	// // database drivers;  pq will exec them all, sqlite3 won't, ymmv
	// db.MustExec(schema)

	// tx := db.MustBegin()
	// tx.MustExec("INSERT INTO intdumps (id, create_date) VALUES ($1, $2)", "test_id", "2021-08-05 21:00:00")
	// tx.Commit()

	db, err := sqlx.Connect("postgres", "host=db port=5432 user=testuser password='1212' dbname=dump sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}

	dt := time.Now()

	tx := db.MustBegin()
	tx.MustExec("INSERT INTO intdumps (id, create_date) VALUES ($1, $2)", id, dt.Format("01-02-2006 15:04:05"))
	tx.MustExec("INSERT INTO dumpdata (id_dump, name, mac, ip, dns, gateway) VALUES ($1, $2, $3, $4, $5, $6)", id, name, mac, ip, dns, gateway)
	tx.Commit()
}
