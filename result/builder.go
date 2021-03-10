package result

import (
	"git.knowlesea.top/HaoTian/account/status"
	"sync"
)

type builder struct {
	once   *sync.Once
	result *Result
}

func Builder() *builder {
	return &builder{
		once:   &sync.Once{},
		result: &Result{},
	}
}

func (b *builder) Success() *builder {
	b.result.Code = status.OK
	b.result.Message = "Success"
	return b
}

func (b *builder) ErrorCode(code int) *builder {
	b.result.Code = code
	return b
}

func (b *builder) ErrorMessage(message string) *builder {
	b.result.Message = message
	return b
}

func (b *builder) Data(data interface{}) *builder {
	b.result.Data = data
	return b
}

func (b *builder) Build() *Result {
	return b.result
}
