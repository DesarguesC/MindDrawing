package validation

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"sync"
)

type CustomValidator struct {
	once     sync.Once
	validate *validator.Validate
}

func (c *CustomValidator) Validate(i interface{}) error {
	(*c).ToInit()
	fmt.Println("HERE!")
	fmt.Println(i)
	fmt.Println((*c).validate == nil)
	return (*c).validate.Struct(i)
}

func (c *CustomValidator) ToInit() {
	(*c).once.Do(func() {
		(*c).validate = validator.New()
		//fmt.Println(c.validate == nil)
	})
}
