package main

func main() {
	v := secret.Memory("my-fake-key")
	err := v.Set("demo_key", "some random value")

}