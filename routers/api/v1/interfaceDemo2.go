package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

type Gopher interface {
	Program(string) string
}

type Student struct {
	Name string
}

func (S Student) Program(language string) string {
	return fmt.Sprintf("%s 会写 %s，叫他 Gopher。", S.Name, language)}

func (S Student) Run(language string) string {
	return fmt.Sprintf("%s 也会写 %s", S.Name, language)
}

func Go(body Gopher)  {
	log.Println(body.Program("Go"))
}

type PHPer interface {
	Do(string) string
}

type Teacher struct {
	Name string
}

func (T Teacher) Do(language string) string {
	return fmt.Sprintf("%s 会教 %s，叫他 PHPer。", T.Name, language)
}

func Php(body PHPer)  {
	log.Println(body.Do("PHP"))
}

type Pythoner interface {
	Run(string) string
}

type Roommate struct {
	Name string
}

func (R Roommate) Run(language string) string {
	return fmt.Sprintf("%s 会学 %s，叫她 Pythoner。", R.Name, language)
}

func Python(body Pythoner) {
	log.Println(body.Run("Python"))
}
type AwesomeDeveloper interface {
	Gopher
	Pythoner
}

func Development(a AwesomeDeveloper) {
	log.Println(a.Program("go"))
	log.Println(a.Run("python"))
}

func Demo2(content *gin.Context)  {
	s := Student{Name: "谢小路"}
	t := Teacher{Name: "谢小人"}
	r := Roommate{Name: "谢小甲"}

	Go(s)
	Php(t)
	Python(r)
	Development(s)
	//Development(r)
}
