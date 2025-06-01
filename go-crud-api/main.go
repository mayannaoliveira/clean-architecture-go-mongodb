// Importação dos pacotes
package main

import (
	"context"
	"go-crud-api/models"
	"net/http"
	"time"

	_"go-crud-api/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection

func connectDB() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}

	collection = client.Database("taskdb").Collection("tasks")
}

// @Summary      Criando nova task
// @Description  Cria uma nova task no banco de dados.
// @Tags         tasks
// @Accept       json
// @Produce      json
// @Param        task  body      models.Task  true  "Dados da Tarefa"
// @Success      201   {object}  map[string]interface{}
// @Failure      400   {object}  map[string]string
// @Failure      500   {object}  map[string]string
// @Router       /task [post]
func createTask(c *gin.Context) {
	var task models.Task // chamada do models > task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task.CreatedAt = time.Now()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := collection.InsertOne(ctx, task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": result.InsertedID})
}

// @Summary      Listar tasks
// @Description  Lista todas as tasks no banco de dados.
// @Tags         tasks
// @Accept       json
// @Produce      json
// @Success      200   {array}   models.Task
// @Failure      500   {object}  map[string]string
// @Router       /tasks [get]
func listTask(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}
	defer cursor.Close(ctx)

	var tasks []models.Task
	if err = cursor.All(ctx, &tasks); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tasks)

}

// @Summary      Atualizando a task
// @Description  Atualiza uma task no banco de dados.
// @Tags         tasks
// @Accept       json
// @Produce      json
// @Param        id    path      string      true  "ID da tarefa"
// @Param        task  body      models.Task true  "Dados da Tarefa"
// @Success      200   {object}  map[string]string
// @Failure      400   {object}  map[string]string
// @Failure      404   {object}  map[string]string
// @Failure      500   {object}  map[string]string
// @Router       /tasks/{id} [put]
func updateTask(c *gin.Context) {
	idParam := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(idParam)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	result, err := collection.UpdateOne(ctx, bson.M{"_id": objID}, bson.M{"$set": task})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if result.MatchedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task não encontrada"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task atualizada com sucesso"})
}

// @Summary      Apagar a task
// @Description  Apaga uma task no banco de dados.
// @Tags         tasks
// @Accept       json
// @Produce      json
// @Param        id    path      string      true  "ID da tarefa"
// @Success      200   {object}  map[string]string
// @Failure      400   {object}  map[string]string
// @Failure      404   {object}  map[string]string
// @Failure      500   {object}  map[string]string
// @Router       /tasks/{id} [delete]		{delete}
func deleteTask(c *gin.Context) {
	idParam := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(idParam)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID Inválido"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	result, err := collection.DeleteOne(ctx, bson.M{"_id": objID})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if result.DeletedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task não encontrada"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Task deletada com sucesso"})
}

// @Title			Task API
// @Version			1.0
// @Description		API para gerenciamento de tarefas com MongoDB.
// @Host			localhost:8000
// @Basepath		/
func main() {
	connectDB()
	router := gin.Default()
	// Endpoint do swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// Mensagem de retorno da api
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Olá, Mundo!"})
	})
	// Routers
	router.GET("/tasks", listTask)         // Router para listar tasks
	router.POST("/task", createTask)       // Router para criar tasks
	router.PUT("/tasks/:id", updateTask)   // Router para atualizar tasks
	router.DELETE("tasks/:id", deleteTask) // Router para apagar tasks

	router.Run(":8000") //Rodando na localhost:8000
}
