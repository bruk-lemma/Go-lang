// package main

// import (
// 	"log"
// 	"net/smtp"
// )

func init() {
	log.SetPrefix("Trace:")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
	log.Println("init started")
}
func main() {

	//pritnln writes to the standard logger
	//log.Println("main started")

	//Fatalln is pritln followed by a call to os.Exit(1)

	//log.Fatalln("fatal message")

	//Panicln is pritln() followed by a call to panic()
	//log.Panicln("panic message")

	//
	//fmt.Println("Real worls use of logs when tying to conect to smtp server")

	//connect to remote smtp server.

	client, err := smtp.Dial("smtp.smail.com:25")
	if err != nil {
		log.Fatalln(err)
	}
	client.Data()

}

func realworld_use_of_logs() {
	//connect to remote smtp server.

	client, err := smtp.Dial("smtp.smail.com:25")
	if err != nil {
		log.Fatalln(err)
	}
	client.Data()
}
