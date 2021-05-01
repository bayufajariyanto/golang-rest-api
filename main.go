package main

func main() {
	app := App{}
	app.Initialize("root", "", "db_go")
	app.Run(":8080")
}
