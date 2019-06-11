package {{package_name}}

//This file is generate by scripts,don't edit it
//

import (
    "changit.cn/contra/bot/db"
)

func LoadBaseData() {
    {% for table in table_list -%}
    {{table.struct_name}}Cache.LoadAll()
    {% endfor -%}

    {% for table in table_list -%}
    db.BaseDataCaches["{{table.struct_name}}"] = {{table.struct_name}}Cache
    {% endfor -%}

}
