package main

func main() {
	opearations := []string{
		"1 -> a",
		"a LSHIFT 4 -> b",
		"b OR 7 -> b",
		"ECHO a -> a",
		"ECHO b -> b",
	}

	rt := make(Runtime)
	rt.Run(opearations)
}
