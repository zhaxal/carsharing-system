package main

import (
	"carsharing-system/cars/pkg/models"
	"carsharing-system/list/pb"
	models2 "carsharing-system/users/pkg/models"
	"database/sql"
	"flag"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
	"strconv"
)

type CarServiceServer struct {
	pb.UnimplementedCarServiceServer
}

var db *sql.DB

func main() {
	fmt.Println("Starting server on port :4000...")
	listener, err := net.Listen("tcp", ":4000")

	if err != nil {
		log.Fatalf("Unable to listen on port :4000: %v", err)
	}

	s := grpc.NewServer()
	srv := &CarServiceServer{}
	pb.RegisterCarServiceServer(s, srv)

	dsn := flag.String("dsn", "root:aserty1234@/car_sharing?parseTime=true", "MySQL DSN")

	flag.Parse()

	db = connect(*dsn)

	go func() {
		if err := s.Serve(listener); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()

	fmt.Println("Server succesfully started on port :4000")

	c := make(chan os.Signal)

	signal.Notify(c, os.Interrupt)
	<-c

	// After receiving CTRL+C Properly stop the server
	fmt.Println("\nStopping the server...")
	s.Stop()
	listener.Close()
	fmt.Println("Closing MySQL connection")
	defer db.Close()
	fmt.Println("Done.")
}

func (s *CarServiceServer) ListCars(req *pb.ListCarReq, stream pb.CarService_ListCarsServer) error {

	stmt := `SELECT id, name, price, available, experience_required FROM car`

	rows, err := db.Query(stmt)
	if err != nil {
		return err
	}

	defer rows.Close()

	for rows.Next() {
		data := &models.Car{}
		err := rows.Scan(&data.ID, &data.Name, &data.Price, &data.Available, &data.ExpReq)

		stream.Send(&pb.ListCarRes{
			Car: &pb.Car{
				Id:        int64(data.ID),
				Name:      data.Name,
				Price:     int64(data.Price),
				Available: strconv.FormatBool(data.Available),
				ExpReq:    int64(data.ExpReq),
			},
		})

		if err != nil {
			return err
		}

	}

	if err = rows.Err(); err != nil {
		return err
	}
	return nil
}

func (s *CarServiceServer) ListUsers(req *pb.ListUserReq, stream pb.CarService_ListUsersServer) error {
	stmt := `SELECT id, name, surname, birthdate, experience FROM user`

	rows, err := db.Query(stmt)
	if err != nil {
		return err
	}

	defer rows.Close()
	for rows.Next() {
		data := &models2.User{}
		err := rows.Scan(&data.ID, &data.Name, &data.Surname, &data.Birthdate, &data.Experience)

		stream.Send(&pb.ListUserRes{
			User: &pb.User{
				Id:         int64(data.ID),
				Name:       data.Name,
				Surname:    data.Surname,
				Birthdate:  data.Birthdate.String(),
				Experience: int64(data.Experience),
			},
		})

		if err != nil {
			return err
		}
	}

	return nil
}

func connect(dsn string) *sql.DB {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	return db
}
