package main

func main() {
	f_jpg := MakeAddSuffix("jpg")
	f_png := MakeAddSuffix("png")

	println(f_jpg("file"))
	println(f_png("file"))
}

func MakeAddSuffix(suffix string) func(string) string {
	return func(fileName string) string {
		return fileName + "." + suffix
	}
}
