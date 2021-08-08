package main

import (
	"context"
	"fmt"
	pb "grpc_test_task/get_interface_dumps/ecommerce"
	"log"
	"strings"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type server struct {
	productMap map[string]*pb.GetInterfaceDumpsRequest
}

func (s *server) GetInterfaceDumps(ctx context.Context, in *pb.GetInterfaceDumpsRequest) (*pb.Dumps, error) {

	type DumpData struct {
		Id      string
		Mac     string
		Name    string
		Ip      string
		Dns     string
		Gateway string
	}

	db, err := sqlx.Connect("postgres", "host=db port=5432 user=testuser password='1212' dbname=dump sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}
	limit := fmt.Sprintf("%v", in.LastCount)
	log.Println(limit)
	request_text := "select intdumps.id, dumpdata.name, dumpdata.ip, dumpdata.mac, dumpdata.dns, dumpdata.gateway from intdumps inner join dumpdata on intdumps.id=dumpdata.id_dump order by intdumps.create_date desc limit " + limit + ";"
	place := DumpData{}
	rows, err := db.Queryx(request_text)
	if err != nil {
		log.Fatalln(err)
	}
	test := make([]*pb.Dump, 0)
	for rows.Next() {
		err := rows.StructScan(&place)
		if err != nil {
			log.Println(err)
		}
		id := strings.TrimSpace(place.Id)
		name := strings.TrimSpace(place.Name)
		mac := strings.TrimSpace(place.Mac)
		ip := strings.TrimSpace(place.Ip)
		dns := strings.TrimSpace(place.Dns)
		gateway := strings.TrimSpace(place.Gateway)
		test_dump := &pb.Dump{Id: id, Name: name, Mac: mac, Ip: ip, Dns: dns, Gateway: gateway}
		test = append(test, test_dump)
	}

	dumps := &pb.Dumps{Dumps: test}
	return dumps, nil

}
