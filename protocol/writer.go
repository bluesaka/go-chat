package protocol

import (
	"fmt"
	"io"
)

type CommandWriter struct {
	writer io.Writer
}

func NewCommandWriter(writer io.Writer) *CommandWriter {
	return &CommandWriter{
		writer: writer,
	}
}

func (w *CommandWriter) writeString(msg string) error {
	_, err := w.writer.Write([]byte(msg))

	return err
}

func (w *CommandWriter) Write(c interface{}) error {
	var err error

	switch v:= c.(type) {  // 类型断言 https://studygolang.com/articles/12355?fr=sidebar
		case SendCommand:
			err = w.writeString(fmt.Sprintf("SEND %v\n", v.Message))
		case MessageCommand:
			err = w.writeString(fmt.Sprintf("MESSAGE %v %v\n", v.Name, v.Message))
		case NameCommand:
			err = w.writeString(fmt.Sprintf("NAME %v\n", v.Name))
		default:
			err = UnknownCommand
	}

	return err
}


