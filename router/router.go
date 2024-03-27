package router

import "github.com/gin-gonic/gin"

func Initialize() {
	// Inicializa o Router utilizando as configurações Default do Gin
	r := gin.Default()
	// Definindo uma rota (Rotas são definidas com verbos, nesse caso GET)
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	// Inicializando a API
	// Da para trocar a porta da aplicação trocando dentro do Run
	r.Run(":8080")
}
