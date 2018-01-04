package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "strconv"
    "fmt"
)

func GetStatus(c *gin.Context) {
    _, out_byte := consumeKafka()
    fmt.Printf("Message:\n%s\n", string(out_byte))
    //fmt.Printf("Message:\n%+v\n", output)
    //c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": output})
    c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": las_status})
}

func GetStatusSingle(c *gin.Context) {
    param_id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        // handle error
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    }
    result, flag := las_status.Get(uint64(param_id))
    if flag {
        c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": result})
    } else {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    }
}

func CreateStatus(c *gin.Context) {
    var las_t_scoring       Las_t_scoring
    var las_t_scoring_clean Las_t_scoring_clean
    var las_status_datum    Las_status
    if err := c.ShouldBindJSON(&las_t_scoring); err == nil {
        las_t_scoring_clean = las_t_scoring.ToClean()
        las_status_datum = las_t_scoring_clean.Score()
        las_status.Add(las_status_datum)
	las_status_datum.ProduceKafka() // Produce data to Kafka topic
        c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": las_status_datum})
    } else {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    }
}

func DeleteStatus(c *gin.Context) {
    param_id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        // handle error
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    }
    las_status.Delete(uint64(param_id))
    c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": las_status})
}
