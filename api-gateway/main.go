package main

import (
	"hello-kubernetes/proto"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("hello-service:50076", grpc.WithTransportCredentials((insecure.NewCredentials())))
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := proto.NewHelloServiceClient(conn)

	r := gin.Default()
	r.GET("/hello", HelloHandler(client))
	r.Run(":8080")
}
func HelloHandler(client proto.HelloServiceClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		name := c.Query("name")
		response, err := client.SayHello(c, &proto.HelloRequest{
			Name: name,
		})
		if err != nil {
			c.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(200, gin.H{"message": response.Message})
	}

}
