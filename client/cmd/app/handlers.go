package main

import (
	"carsharing-system/list/pb"
	"context"
	"encoding/json"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func (app *App) getAPIContent(url string, templateData interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	json.Unmarshal(bodyBytes, templateData)
	return nil
}

func (app *App) viewCars(w http.ResponseWriter, r *http.Request) {
	fmt.Println("client")
	cc, err := grpc.Dial("localhost:4000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer cc.Close()

	c := pb.NewCarServiceClient(cc)

	getCars(c, w)
}

func (app *App) viewUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("client")
	cc, err := grpc.Dial("localhost:4000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer cc.Close()

	c := pb.NewCarServiceClient(cc)

	getUsers(c, w)
}

func getCars(c pb.CarServiceClient, w http.ResponseWriter) {
	fmt.Println("getting cars info")
	req := &pb.ListCarReq{}

	stream, err := c.ListCars(context.Background(), req)
	if err != nil {
		log.Fatalf("error rpc: %v", err)
	}
	var cars []*pb.Car
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("error: %v", err)
		}
		cars = append(cars, res.GetCar())
	}

	json.NewEncoder(w).Encode(cars)
}

func getUsers(c pb.CarServiceClient, w http.ResponseWriter) {
	fmt.Println("getting users info")
	req := &pb.ListUserReq{}

	stream, err := c.ListUsers(context.Background(), req)
	if err != nil {
		log.Fatalf("error rpc: %v", err)
	}
	var users []*pb.User
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("error: %v", err)
		}
		users = append(users, res.GetUser())
	}
	json.NewEncoder(w).Encode(users)
}
