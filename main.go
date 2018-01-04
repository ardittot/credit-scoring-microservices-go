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
	out := consumeKafka()
	fmt.Printf("%% Message on %s:\n%s\n%s\n", e.TopicPartition, out.ID_Scoring, out.Score)
  }()

  // Initialize the routes
  initializeRoutes()

  // Start serving the application
  router.Run("0.0.0.0:8000")

  // Terminate kafka
  producer.Close()
}
