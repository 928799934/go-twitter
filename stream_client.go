package gotwitter

import (
	"bufio"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/928799934/go-twitter/internal/util"
)

type StreamClient[T util.Response] struct {
	response *http.Response
	stream   *bufio.Scanner
}

func newStreamClient[T util.Response](resp *http.Response) (*StreamClient[T], error) {
	if resp == nil {
		return nil, errors.New("HTTP Response is nil")
	}

	if resp.Close {
		return nil, errors.New("HTTP Response body has already closed")
	}

	return &StreamClient[T]{
		response: resp,
		stream:   bufio.NewScanner(resp.Body),
	}, nil
}

func (s *StreamClient[T]) Receive() bool {
	if s == nil {
		return false
	}
	return s.stream.Scan()
}

func (s *StreamClient[T]) Stop() {
	if s == nil {
		return
	}
	s.response.Body.Close()
}

func safeUnmarshal(input []byte, target interface{}) error {
	if len(input) == 0 {
		return nil
	}
	return json.Unmarshal(input, target)
}

func (s *StreamClient[T]) Read() (T, error) {
	var n T
	if s == nil {
		return n, errors.New("StreamClient is nil")
	}

	t := s.stream.Text()
	out := new(T)
	if err := safeUnmarshal([]byte(t), out); err != nil {
		return n, err
	}

	return *out, nil
}
