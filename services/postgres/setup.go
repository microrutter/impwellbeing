package postgres

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	// "github.com/nu7hatch/gouuid"
)

// type QuestionList struct {
//     Uuid uuid.UUID
//     Results []Question
// }

// type Question struct {
//     gorm.Model
//     Uuid string
//     Type string
//     Title string
// 	Area string
// }

// func returnUUID() string {
// 	uuid, _ := uuid.NewV4()
// 	return uuid.String()
// }

func createConnection() (*gorm.DB, error)  {
    db, err := gorm.Open( "postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", os.Getenv("PostGres_Host"), os.Getenv("PostGres_Port"), os.Getenv("PostGres_User"), os.Getenv("PostGres_Db"), os.Getenv("PostGres_Password")))

    if err != nil {
        log.Fatal(err)
        return nil, err
    }

    return db, nil
}

// func initialSetUp() {
//     _, exists := os.LookupEnv("RunSetUp")

//     if exists {

//         db, err := createConnection()

//         if err == nil {
//             defer db.Close()

//             db.AutoMigrate(&Question{})
//         }

//     }
    
// }


