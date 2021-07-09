package postgres

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/nu7hatch/gouuid"
)

type QuestionList struct {
    Uuid uuid.UUID
    Results []Question
}

type Question struct {
    gorm.Model
    Uuid string
    Type string
    Title string
}
var (
	questionList = []Question{
		{Uuid: returnUUID(), Type: "onetofour", Title: "How overloaded with pressure at work are you?"},

		{Uuid: returnUUID(), Type: "onetofour", Title: "How overwhelmed do you feel?"},

		{Uuid: returnUUID(), Type: "onetofour", Title: "I am clear what is expected of me at work?"},

		{Uuid: returnUUID(), Type: "onetofour", Title: "I can decide when to take a break?"},

		{Uuid: returnUUID(), Type: "onetofour", Title: "I know how to go about getting my job done?"},

		{Uuid: returnUUID(), Type: "onetofour", Title: "If work gets difficult , my colleagues will help me?"},

		{Uuid: returnUUID(), Type: "onetofour", Title: "I am given supportive feedback on the work I do?"},

		{Uuid: returnUUID(), Type: "onetofour", Title: "I have a say in my own work speed?"},

		{Uuid: returnUUID(), Type: "onetofour", Title: "I am clear what my duties and responsibilities are?"},

		{Uuid: returnUUID(), Type: "onetofour", Title: "I am clear about the goals and objectives for my department?"},

		{Uuid: returnUUID(), Type: "onetofour", Title: "I have a choice in deciding how i do my work?"},

		{Uuid: returnUUID(), Type: "onetofour", Title: "I undertand how my work fits into the overall aim of the organisation"},

		{Uuid: returnUUID(), Type: "onetofour", Title: "I have a choice in deciding what I do at work?"},

		{Uuid: returnUUID(), Type: "onetofour", Title: "I can rely on my line manager to help me out with a work problem?"},

		{Uuid: returnUUID(), Type: "onetofour", Title: "I have sufficient opportunities to question managers about change at work?"},

		{Uuid: returnUUID(), Type: "onetofour", Title: "I receive the respect at work I deserve from my colleagues?"},

		{Uuid: returnUUID(), Type: "onetofour", Title: "Staff are always consulted about changes at work?"},

		{Uuid: returnUUID(), Type: "onetofour", Title: "I can talk to my line manager about hot something that has upset or annoyed me about work?"},

		{Uuid: returnUUID(), Type: "onetofour", Title: "My working time can be flexible?"},

		{Uuid: returnUUID(), Type: "onetofour", Title: "My collegues are willing to listen to my work-related problems?"},

		{Uuid: returnUUID(), Type: "onetofour", Title: "When changes are made at work, I am clear how they will work out in practice?"},

		{Uuid: returnUUID(), Type: "onetofour", Title: "I am supported through emotionally demanding work?"},

		{Uuid: returnUUID(), Type: "onetofour", Title: "My line manager ncourages me at work?"},

		{Uuid: returnUUID(), Type: "yes/no", Title: "Are you able to switch off from your demands at work?"},

		{Uuid: returnUUID(), Type: "yes/no", Title: "Do you feel in control to make changes at work?"},

		{Uuid: returnUUID(), Type: "yes/no", Title: "Do you feel a strong sense of purpose at work?"},

		{Uuid: returnUUID(), Type: "yes/no", Title: "Do you feel supported at work?"},

		{Uuid: returnUUID(), Type: "yes/no", Title: "Do you feel positive about work?"},

		{Uuid: returnUUID(), Type: "yes/no", Title: "Can you speak openly about Mental Health and well-being?"},
	}
)

func returnUUID() string {
	uuid, _ := uuid.NewV4()
	return uuid.String()
}

func createConnection() (*gorm.DB, error)  {
    db, err := gorm.Open( "postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", os.Getenv("PostGres_Host"), os.Getenv("PostGres_Port"), os.Getenv("PostGres_User"), os.Getenv("PostGres_Db"), os.Getenv("PostGres_Password")))

    if err != nil {
        log.Fatal(err)
        return nil, err
    }

    return db, nil
}

func initialSetUp() {
    _, exists := os.LookupEnv("RunSetUp")

    if exists {

        questions := questionList

        db, err := createConnection()

        if err == nil {
            defer db.Close()

            db.AutoMigrate(&Question{})

            for index := range questions {
                db.Create(&questions[index])
            }
        }

    }
    
}


