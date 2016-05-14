package shopers

import (
    "github.com/mantyr/conf"
)

type Shopers struct {
    conf_file  conf.ConfigFile
    conf_group conf.ConfigFile
    is_uc_first bool
}
