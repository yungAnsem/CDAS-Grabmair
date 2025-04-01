// main.go

package main

func main() {
	a := App{}
	a.Initialize(
		"postgres",
		"password",
		"postgres")

	a.Run(":8010")
}
