package shopers

import (
    "github.com/mantyr/conf"
    "github.com/mantyr/runner"
    "strings"
)

func NewShopers() (s *Shopers) {
    s = new(Shopers)
    return
}

// Заменяет названия категорий на те что указаны в конфиге по следующему принципу:
//  1. Вначале текущая цепочка категорий
//  2. Потом по уменьшению до корневой
//  3. Потом в категории default
func (s *Shopers) Get(categories []string) (result []string) {
    var key_category  string
    var key_value_old string
    var is bool

    for key, category := range categories {
        category = runner.Trim(category)
        key_category = strings.ToLower(category)

        if s.is_uc_first {
            key_value_old = runner.UcFirst(category)
        } else {
            key_value_old = category
        }

        is = false

        if key > 0 {
            for i := len(result)-1; i >= 0; i-- {
                section_replace := strings.ToLower("category: \""+strings.Join(result[0:i+1], "\" \"")+"\"")

                if s.conf_file.Is(key_category, section_replace) {
                    arr := strings.Split(s.conf_file.Get(key_category, section_replace), " | ")
                    for _, arr_item := range arr {
                        arr_item = runner.Trim(arr_item)
                        result = append(result, arr_item)
                    }
                    is = true
                    break;
                }
            }
        }

        if !is {
            if s.conf_file.Is(key_category, "default") {
                arr := strings.Split(s.conf_file.Get(key_category, "default"), " | ")
                for _, arr_item := range arr {
                    arr_item = runner.Trim(arr_item)
                    result = append(result, arr_item)
                }
            } else {
                result = append(result, key_value_old)
            }
        }
    }

    return result
}

// Нужно ли делать первую букву исходного названия категории большой (заменяемые из конфига не модифицируются)
func (s *Shopers) IsUcFirst(status bool) {
    s.is_uc_first = status
}

func (s *Shopers) LoadConf(name string) *Shopers {
    s.conf_file = conf.GetFile(name)
    return s
}