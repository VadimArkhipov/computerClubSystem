package main

func main() {
	path := "test.txt"

	data := readFile(path)

	validate(data)

	club := Club{}
	club.init(data)

	club.processing()
}
