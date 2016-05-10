package shopers

import (
    "github.com/mantyr/conf"
    "testing"
//    "fmt"
)

func init() {
    conf.SetDefaultCatalog("./testdata")
}

type shoperTest struct{
    a []string
    b []string
}

var shoperTests = []shoperTest{
    {[]string{"shoes", "Basketball"},   []string{"Обувь", "Кроссовки баскетбольные"}},
    {[]string{"Обувь", "Basketball"},   []string{"Обувь", "Кроссовки баскетбольные"}},
    {[]string{"обувь", "Basketball"},   []string{"Обувь", "Кроссовки баскетбольные"}},
    {[]string{"clothes", "Basketball"}, []string{"Одежда", "Для баскетбола"}},
    {[]string{"Basketball"},            []string{"Баскетбол"}},
    {[]string{"Slides"},                []string{"Сланцы, балетки"}},
    {[]string{"обувь", "сланцы"},       []string{"Обувь", "Сланцы, балетки"}},
    {[]string{"обувь", "сланцы", "летние", "Красивые", "Великолепные"},  []string{"Обувь", "Сланцы, балетки", "Летние", "Красивые", "Великолепные балетки"}},
}

func test_categories(a []string, b[]string) bool {
    if len(a) != len(b) {
        return false
    }
    for key, value := range a {
        if b[key] != value {
            return false
        }
    }
    return true
}

func TestShopers(t *testing.T) {
    fields := NewShopers()
    fields.LoadConf("category_replace")
    fields.IsUcFirst(true)

    for _, test := range shoperTests {
        val := fields.Get(test.a)
        if test_categories(val, test.b) == false {
            t.Errorf("Error replace category name, %q, %q, %q", val, test.a, test.b)
        }
    }
}
