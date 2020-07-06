package initialize

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
)

// customize Logically delete
var LogicDelete = func(scope *gorm.Scope) {
	tn := scope.TableName()
	pk := scope.PrimaryField().Field.Uint()
	cmd := fmt.Sprintf("UPDATE %s SET is_deleted = 1 WHERE id = %d", tn, pk)
	scope.DB().Exec(cmd)
}

// update create_time and update_time before create
var BeforeCreate = func(scope *gorm.Scope) {
	if scope.HasColumn("create_time") {
		_ = scope.SetColumn("create_time", time.Now().Unix())
	}
	if scope.HasColumn("update_time") {
		_ = scope.SetColumn("update_time", time.Now().Unix())
	}
	if scope.HasColumn("is_deleted") {
		_ = scope.SetColumn("is_deleted", 0)
	}
}

// update update_time before update
var BeforeUpdate = func(scope *gorm.Scope) {
	if scope.HasColumn("update_time") {
		_ = scope.SetColumn("update_time", time.Now().Unix())
	}
}

// exclude deleted item
var BeforeQuery = func(scope *gorm.Scope) {
	scope.Search = scope.Search.Where("is_deleted = ?", 0)
}
