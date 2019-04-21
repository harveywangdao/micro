package main

import (
	"log"

	"github.com/gin-gonic/gin"

	hello "github.com/harveywangdao/micro/proto/hello"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-web"

	"context"
)

type Say struct{}

var (
	cli1 hello.SayService
	cli2 hello.SayService
	cli3 hello.SayService
)

func (s *Say) Anything(c *gin.Context) {
	log.Print("Received Say.Anything API request")
	c.JSON(200, map[string]string{
		"message": "Hi, this is the Greeter API",
	})
}

func (s *Say) Hello(c *gin.Context) {
	log.Print("Received Say.Hello API request")

	name := c.Param("name")

	response1, err := cli1.Hello(context.TODO(), &hello.Request{
		Name: name,
	})

	response2, err := cli2.Hello(context.TODO(), &hello.Request{
		Name: name,
	})

	response3, err := cli3.Hello(context.TODO(), &hello.Request{
		Name: name,
	})

	if err != nil {
		c.JSON(500, err)
	}

	var response []*hello.Response
	response = append(response, response1, response2, response3)

	c.JSON(200, response)
}

func main() {
	service := web.NewService(
		web.Name("go.micro.api.harvey"),
	)

	service.Init()

	cli1 = hello.NewSayService("go.micro.srv.ms1", client.DefaultClient)
	cli2 = hello.NewSayService("go.micro.srv.ms2", client.DefaultClient)
	cli3 = hello.NewSayService("go.micro.srv.ms3", client.DefaultClient)

	say := new(Say)
	router := gin.Default()
	router.GET("/greeter", say.Anything)
	router.GET("/greeter/:name", say.Hello)

	service.Handle("/", router)

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
