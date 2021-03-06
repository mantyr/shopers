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
    {[]string{"обувь", "Детские баскетбольные кроссовки"},               []string{"Обувь", "Детские", "Баскетбол"}},
    {[]string{"обувь2", "Детские баскетбольные кроссовки"},              []string{"Обувь", "Вторая группа обуви", "Детские", "Баскетбол"}},
    {[]string{"Инвентарь", "Баскетбольные мячи"},                        []string{"Аксессуары", "Мячи"}},  // replace fist category
    {[]string{"Инвентарь", "Баскетбольные мячи", "С фотографией"},       []string{"Аксессуары", "Мячи", "С фотографией"}},  // replace fist category
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
    fields.LoadConfGroup("category_replace_group")
    fields.IsUcFirst(true)

    for _, test := range shoperTests {
        val := fields.Get(test.a)
        val  = fields.GetGroup(val)
        if test_categories(val, test.b) == false {
            t.Errorf("Error replace category name, %q, %q, %q", val, test.a, test.b)
        }
    }
}
