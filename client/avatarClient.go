package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"log"
	"net/http"
	avatar "rpc-hello/avatarRpc"
)

func main() {

	r := gin.Default()

	r.GET("/rpc/avatar", func(c *gin.Context) {
		getMovies(c)
	})

	// Run http server
	if err := r.Run(":8052"); err != nil {
		log.Fatalf("could not run server: %v", err)
	}
}

func getMovies(c *gin.Context) {

	// Set up a connection to the server.
	conn, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()

	client := avatar.NewAvatarMovieClient(conn)
	movieid := c.DefaultQuery("movieid","3")
	input :=avatar.Input{Movieid:movieid}
    inputs :=[]*avatar.Input{&input}
	req := &avatar.AvatarMovieRequest{Input:inputs}
	res, err := client.GetAvatarMovie(c, req)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": fmt.Sprint(res.Output.Data.GetStructValue().Fields),
		"code": fmt.Sprint(res.Output.Code),
	})
	fmt.Sprint(res.Output.Data.GetStructValue().Fields["ui_type"].GetStructValue())
}
