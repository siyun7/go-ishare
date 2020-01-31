package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

type Languager interface {
	Can(string) string
}

type Someone struct {
	Language string
}

type Nobody struct {
	Age string
}


func (s Someone) Can(name string) string {
	return fmt.Sprintf("%s can program with %s", name, s.Language)
}

func (n Nobody) Can(name string) string {
	return fmt.Sprintf("The %s age is %s", name, n.Age)
}

func Program(L Languager, name string)  {
	log.Println(L.Can(name))
}

func RunDemo1(content *gin.Context)  {
	b := Someone{Language: "go"}
	Program(b, "Thanks")

	c := Nobody{Age: "12"}
	Program(c, "Lee")
}