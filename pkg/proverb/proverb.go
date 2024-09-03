package proverb

var List [10]string

func init() {
	List = [10]string{
		"Concurrency is not parallelism.",
		"Channels orchestrate; mutexes serialize.",
		"Make the zero value useful.",
		"A little copying is better than a little dependency.",
		"Syscall must always be guarded with build tags.",
		"Cgo is not Go.",
		"Clear is better than clever.",
		"Reflection is never clear.",
		"Errors are values.",
		"Don't just check errors, handle them gracefully.",
	}
}
