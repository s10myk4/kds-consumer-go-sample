package interfaces

type RecordConsumer[In any, Out any] interface {
	Consume(in *In) (*[]Out, error)
}
