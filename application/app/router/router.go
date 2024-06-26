package router

import (
	"net/http"
	"strconv"
	"todo-app/model"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func getTickets(c *gin.Context, db *gorm.DB) {
	tickets := make([]model.Ticket, 0)
	err := model.GetTickets(&tickets, db)
	if err == nil {
		c.JSON(http.StatusOK, tickets)
	} else {
		c.AbortWithStatus(http.StatusInternalServerError)
	}
}

func getCompletedTickets(c *gin.Context, db *gorm.DB) {
	tickets := make([]model.Ticket, 0)
	err := model.GetTicketsByStatus(&tickets, true, db)
	if err == nil {
		c.JSON(http.StatusOK, tickets)
	} else {
		c.AbortWithStatus(http.StatusNotFound)
	}
}

func getUncompletedTickets(c *gin.Context, db *gorm.DB) {
	tickets := make([]model.Ticket, 0)
	err := model.GetTicketsByStatus(&tickets, false, db)
	if err == nil {
		c.JSON(http.StatusOK, tickets)
	} else {
		c.AbortWithStatus(http.StatusNotFound)
	}
}

func getTicket(c *gin.Context, db *gorm.DB) {
	ticket := model.Ticket{}
	id, err := strconv.ParseUint(c.Params.ByName("id"), 10, 64)
	if err != nil {
		c.AbortWithStatus(http.StatusUnprocessableEntity)
	}
	err = model.GetTicket(&ticket, uint(id), db)
	if err == nil {
		c.JSON(http.StatusOK, ticket)
	} else {
		c.AbortWithStatus(http.StatusNotFound)
	}
}

func _createTicket(c *gin.Context, db *gorm.DB) {
	title := c.Params.ByName("title")
	ticket := model.Ticket{Title: title}
	err := model.CreateTicket(&ticket, db)
	if err == nil {
		c.JSON(http.StatusOK, ticket)
	} else {
		c.AbortWithStatus(http.StatusInternalServerError)
	}
}

func createTicket(c *gin.Context, db *gorm.DB) {
	ticket := model.Ticket{}
	err := c.BindJSON(&ticket)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}
	err = model.CreateTicket(&ticket, db)
	if err == nil {
		c.JSON(http.StatusOK, ticket)
	} else {
		c.AbortWithStatus(http.StatusInternalServerError)
	}
}

func completeTicket(c *gin.Context, db *gorm.DB) {
	id, err := strconv.ParseUint(c.Params.ByName("id"), 10, 64)
	if err != nil {
		c.AbortWithStatus(http.StatusUnprocessableEntity)
	}
	err = model.CompleteTicket(uint(id), db)
	if err == nil {
		c.JSON(http.StatusOK, id)
	} else {
		c.AbortWithStatus(http.StatusNotFound)
	}
}

func deleteTicket(c *gin.Context, db *gorm.DB) {
	id, err := strconv.ParseUint(c.Params.ByName("id"), 10, 64)
	if err != nil {
		c.AbortWithStatus(http.StatusUnprocessableEntity)
	}
	err = model.DeleteTicket(uint(id), db)
	if err == nil {
		c.JSON(http.StatusOK, id)
	} else {
		c.AbortWithStatus(http.StatusNotFound)
	}
}

func SetUpRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	r.Use(cors.Default())

	app := r.Group("/app")
	app.GET("tickets", func(c *gin.Context) { getTickets(c, db) })
	app.GET("completed-tickets", func(c *gin.Context) { getCompletedTickets(c, db) })
	app.GET("uncompleted-tickets", func(c *gin.Context) { getUncompletedTickets(c, db) })
	app.GET("ticket/:id", func(c *gin.Context) { getTicket(c, db) })
	app.POST("ticket", func(c *gin.Context) { createTicket(c, db) })
	app.PUT("complete-ticket/:id", func(c *gin.Context) { completeTicket(c, db) })
	app.DELETE("ticket/:id", func(c *gin.Context) { deleteTicket(c, db) })

	return r
}
