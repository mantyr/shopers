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
    result = make([]string, len(categories))
    var key_category string
    var is bool

    for key, category := range categories {
        key_category = strings.ToLower(runner.Trim(category))

        if s.is_uc_first {
            result[key] = runner.UcFirst(category)
        } else {
            result[key] = category
        }

        if key == 0 {
            if s.conf_file.Is(key_category, "default") {
                result[key] = s.conf_file.Get(key_category, "default")
            }
        } else {
            is = false
            for i := key-1; i >= 0; i-- {
                section_replace := strings.ToLower("category: \""+strings.Join(result[0:i+1], "\" \"")+"\"")

                if s.conf_file.Is(key_category, section_replace) {
                    result[key] = s.conf_file.Get(key_category, section_replace)
                    is = true
                    break;
                }
            }
            if !is && s.conf_file.Is(key_category, "default") {
                result[key] = s.conf_file.Get(key_category, "default")
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