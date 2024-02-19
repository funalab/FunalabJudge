package submission

import "github.com/gin-gonic/gin"

func AddSubmissionHandler(c *gin.Context) {
	// client, exists := c.Get("mongoClient")
	// if !exists {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"Error": "DB client is not available."})
	// }

	// dbName := os.Getenv("DB_NAME")
	// submitCol := os.Getenv("SUBMISSION_COLLECTION")
	// collection := (client.(*mongo.Client)).Database(dbName).Collection(submitCol)

	/*Bind Submission and push into db*/
}
