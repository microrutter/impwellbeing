package mongo

import (
    "context"
    "log"
    "os"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "github.com/nu7hatch/gouuid"
)

type QuestionList struct {
    Uuid uuid.UUID
    Results []Question
}

type Question struct {
    Title string
    Result  int
}

func CheckConnection() string {
    
    client := getConnection()
    
    err := client.Ping(context.TODO(), nil)

    if err != nil {
        log.Fatal(err)
        return err.Error()
    }

    return "Connected to MongoDB!"

}

func getConnection() *mongo.Client {
    clientOptions := options.Client().ApplyURI(os.Getenv("MongoDB"))
    client, err := mongo.Connect(context.TODO(), clientOptions)

    if err != nil {
        log.Fatal(err)
    }

    return client
}