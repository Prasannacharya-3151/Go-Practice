package intermediate
import (
	"log"
	"os"
)

func main(){
	//The log package is used to record application events, errors, and debugging information. It provides timestamped log messages and can write logs to the terminal or files.
	file, err := os.OpenFile( //open file for logging
		"app.log",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0644,
	)

	if err != nil {
		log.Fatal(err) //log+stop program also the print message + time
	}

	defer file.Close()

	log.SetOutput(file) //sends logs to the terminal .dont print to terminal. write logs into the files
    //stroes messages
	log.Println("Applicaion Started")

	log.Println("User Uploded file")
}