func main() {
	log.Println("=== INICIANDO APP ===")
	
	database.Connect()
	
	log.Println("=== DB CONECTADA ===")

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	routes.Routes(router)

	port := os.Getenv("PORT")
	log.Printf("=== PORT value: '%s' ===", port)

	if port == "" {
		port = "10000"
	}

	log.Printf("=== Iniciando en 0.0.0.0:%s ===", port)

	server := &http.Server{
		Addr:    "0.0.0.0:" + port,
		Handler: router,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatal("Error:", err)
	}
}