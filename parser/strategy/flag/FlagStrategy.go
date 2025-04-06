package flag

type Strategy interface {
	Execute(channel chan string) chan string
}
