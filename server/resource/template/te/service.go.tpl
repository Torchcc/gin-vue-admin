package service

import (
	"fmt"

	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

// @title    Create{{.StructName}}
// @description   create a {{.StructName}}
// @param     {{.Abbreviation}}               model.{{.StructName}}
// @auth                     （2020/04/05  20:22）
// @return    err             error

func Create{{.StructName}}({{.Abbreviation}} model.{{.StructName}}) (err error) {
	err = global.BIZ_DB.Create(&{{.Abbreviation}}).Error
	return err
}

// @title    Delete{{.StructName}}
// @description   delete a {{.StructName}}
// @auth                     （2020/04/05  20:22）
// @param     {{.Abbreviation}}               model.{{.StructName}}
// @return                    error

func Delete{{.StructName}}({{.Abbreviation}} model.{{.StructName}}) (err error) {
	err = global.BIZ_DB.Delete({{.Abbreviation}}).Error
	return err
}

// @title    Delete{{.StructName}}ByIds
// @description   delete {{.StructName}}s
// @auth                     （2020/04/05  20:22）
// @param     {{.Abbreviation}}               model.{{.StructName}}
// @return                    error

func Delete{{.StructName}}ByIds(ids request.IdsReq) (err error) {
    db := global.BIZ_DB.Model(&model.{{.StructName}}{})
	cmd := fmt.Sprintf("UPDATE %s SET is_deleted = 1 WHERE id IN (?)", model.{{.StructName}}{}.TableName())
	err = db.Exec(cmd, ids.Ids).Error
	return err
}

// @title    Update{{.StructName}}
// @description   update a {{.StructName}}
// @param     {{.Abbreviation}}          *model.{{.StructName}}
// @auth                     （2020/04/05  20:22）
// @return                    error

func Update{{.StructName}}({{.Abbreviation}} *model.{{.StructName}}) (err error) {
	err = global.BIZ_DB.Save({{.Abbreviation}}).Error
	return err
}

// @title    Get{{.StructName}}
// @description   get the info of a {{.StructName}}
// @auth                     （2020/04/05  20:22）
// @param     id              uint
// @return                    error
// @return    {{.StructName}}        {{.StructName}}

func Get{{.StructName}}(id uint) (err error, {{.Abbreviation}} model.{{.StructName}}) {
	err = global.BIZ_DB.Where("id = ?", id).First(&{{.Abbreviation}}).Error
	return
}

// @title    Get{{.StructName}}InfoList
// @description   get {{.StructName}} list by pagination, 分页获取用户列表
// @auth                     （2020/04/05  20:22）
// @param     info            PageInfo
// @return                    error

func Get{{.StructName}}InfoList(info request.{{.StructName}}Search) (err error, list interface{}, total int) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.BIZ_DB.Model(&model.{{.StructName}}{}).Where("is_deleted = ?", 0)
    var {{.Abbreviation}}s []model.{{.StructName}}
    // 如果有条件搜索 下方会自动创建搜索语句
        {{- range .Fields}}
            {{- if .FieldSearchType}}
                {{- if eq .FieldType "string" }}
    if info.{{.FieldName}} != "" {
        db = db.Where("{{.ColumnName}} {{.FieldSearchType}} ?",{{if eq .FieldSearchType "LIKE"}}"%"+ {{ end }}info.{{.FieldName}}{{if eq .FieldSearchType "LIKE"}}+"%"{{ end }})
    }
                {{- else if eq .FieldType "bool" }}
    if info.{{.FieldName}} != nil {
        db = db.Where("{{.ColumnName}} {{.FieldSearchType}} ?",{{if eq .FieldSearchType "LIKE"}}"%"+{{ end }}info.{{.FieldName}}{{if eq .FieldSearchType "LIKE"}}+"%"{{ end }})
    }
                {{- else if eq .FieldType "int" }}
    if info.{{.FieldName}} != 0 {
        db = db.Where("{{.ColumnName}} {{.FieldSearchType}} ?",{{if eq .FieldSearchType "LIKE"}}"%"+{{ end }}info.{{.FieldName}}{{if eq .FieldSearchType "LIKE"}}+"%"{{ end }})
    }
                {{- else if eq .FieldType "float64" }}
    if info.{{.FieldName}} != 0 {
        db = db.Where("{{.ColumnName}} {{.FieldSearchType}} ?",{{if eq .FieldSearchType "LIKE"}}"%"+{{ end }}info.{{.FieldName}}{{if eq .FieldSearchType "LIKE"}}+"%"{{ end }})
    }
                {{- else if eq .FieldType "time.Time" }}
    if !info.{{.FieldName}}.IsZero() {
         db = db.Where("{{.ColumnName}} {{.FieldSearchType}} ?",{{if eq .FieldSearchType "LIKE"}}"%"+{{ end }}info.{{.FieldName}}{{if eq .FieldSearchType "LIKE"}}+"%"{{ end }})
    }
                {{- end }}
        {{- end }}
    {{- end }}
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&{{.Abbreviation}}s).Error
	return err, {{.Abbreviation}}s, total
}