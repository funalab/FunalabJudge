package submission

import (
	"context"
	"go-test/judge"
	"go-test/myMath"
	"go-test/myTypes"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type submissionRequest struct {
	ProblemId     int32     `form:"problemId"`
	SubmittedDate time.Time `form:"submittedDate"`
}

func AddSubmissionHandler(c *gin.Context) {
	// add submission document
	var sr submissionRequest
	if err := c.Bind(&sr); err != nil {
		log.Println(err.Error())
		c.JSON(400, gin.H{"err": err.Error()})
	}
	s := makeSubmissionDocument(c, &sr)
	client, exists := c.Get("mongoClient")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "DB client is not available."})
	}
	dbName := os.Getenv("DB_NAME")
	submitCol := os.Getenv("SUBMISSION_COLLECTION")
	col := (client.(*mongo.Client)).Database(dbName).Collection(submitCol)
	col.InsertOne(context.TODO(), *s)

	// save posted files
	form, _ := c.MultipartForm()
	files := form.File["files[]"]
	for _, file := range files {
		c.SaveUploadedFile(file, filepath.Join(os.Getenv("EXEC_DIR"), strconv.Itoa(int(s.Id)), file.Filename))
	}
	// コンパイル&実行プロセスのマルチスレッド予約
	go judge.JudgeProcess(c, *s)
	c.JSON(400, nil)
}

func makeSubmissionDocument(c *gin.Context, sr *submissionRequest) *myTypes.Submission {
	client, exists := c.Get("mongoClient")
	if !exists {
		return nil
	}
	col := (client.(*mongo.Client)).Database(os.Getenv("DB_NAME")).Collection(os.Getenv("USERS_COLLECTION"))
	filter := bson.M{"userName": c.Param("userName")}
	var u myTypes.User
	err := col.FindOne(context.TODO(), filter).Decode(&u)
	if err != nil {
		log.Println("Failed to find such user.")
		return nil
	}
	var s myTypes.Submission
	s.Id = int32(getMaxSubmissionId(c)) + 1
	s.UserId = u.UserId
	s.ProblemId = sr.ProblemId
	s.SubmittedDate = sr.SubmittedDate
	s.Status = "WJ"

	/*Map Results*/
	col = (client.(*mongo.Client)).Database(os.Getenv("DB_NAME")).Collection(os.Getenv("PROBLEMS_COLLECTION"))
	filter = bson.M{"problemId": s.ProblemId}
	var p myTypes.Problem
	err = col.FindOne(context.TODO(), filter).Decode(&p)
	if err != nil {
		log.Println("Failed to find such problem.")
		return nil
	}
	nt := len(p.TestcaseWithPaths)
	s.Results = make([]myTypes.Result, nt)

	for index, ele := range s.Results {
		ele.TestId = index + 1
		ele.Status = "WJ"
		s.Results[index] = ele
	}
	return &s
}

func getMaxSubmissionId(c *gin.Context) int {
	dbName := os.Getenv("DB_NAME")
	subCol := os.Getenv("SUBMISSION_COLLECTION")
	client, exists := c.Get("mongoClient")
	if !exists {
		log.Fatalln("DB client is not available.")
	}
	collection := (client.(*mongo.Client)).Database(dbName).Collection(subCol)
	cur, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Fatalln("Failed to find all submission information.")
	}
	var submissions []myTypes.Submission
	if err = cur.All(context.TODO(), &submissions); err != nil {
		log.Fatalln("Failed to fetch all submission information.")
	}
	maxSubmissionId := -1
	for _, submission := range submissions {
		submissionId := int(submission.Id)
		maxSubmissionId = myMath.MaxInt(submissionId, maxSubmissionId)
	}
	return maxSubmissionId
}
