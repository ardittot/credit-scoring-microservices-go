package main

import (
	"github.com/gin-gonic/gin"
	"fmt"
)

var router *gin.Engine

func main() {

  // Set the router as the default one provided by Gin
  router = gin.Default()

  // Initialize data
  las_status = InitLasStatus()

  // Initialize kafka
  InitKafkaProducer()
  InitKafkaConsumer()
  go func() { 
	_, out_byte := consumeKafka()
	fmt.Printf("Message:\n%s\n", string(out_byte))
  }()

  // Initialize the routes
  initializeRoutes()

  // Start serving the application
  router.Run("0.0.0.0:8000")

  // Terminate kafka
  producer.Close()
}
