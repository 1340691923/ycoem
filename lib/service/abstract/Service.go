package abstract

import (
	"ycoem/lib/db"
	"ycoem/lib/logs"
)

type Service interface {
	ProcessSqlWhere(sqlA db.SelectBuilder) db.SelectBuilder
	TableName() string
	ProcessSqlInsert(sqlA db.InsertBuilder) db.InsertBuilder
	ProcessSqlUpdate(id int, sqlA db.UpdateBuilder) db.UpdateBuilder
}

func SearchAll(service Service, list interface{}) (err error) {
	sqlA := db.SqlBuilder.Select("*").
		From(service.TableName())
	sqlA = service.ProcessSqlWhere(sqlA)
	sql, args, err := sqlA.ToSql()
	logs.Logger.Sugar().Info(sql, args)
	err = db.Sqlx.Select(list, sql, args...)
	return
}

func FindOne(service Service, obj interface{}) (err error) {
	sqlA := db.SqlBuilder.Select("*").
		From(service.TableName())
	sqlA = service.ProcessSqlWhere(sqlA)
	sql, args, err := sqlA.Limit(1).ToSql()
	logs.Logger.Sugar().Info(sql, args)
	err = db.Sqlx.Get(obj, sql, args...)
	return
}

//为了省时间 只用于单表操作
func SearchList(service Service, page, limit int, list interface{}) (err error) {
	sqlA := db.SqlBuilder.Select("*").
		From(service.TableName())
	sqlA = service.ProcessSqlWhere(sqlA)
	sql, args, err := sqlA.
		Limit(uint64(limit)).
		Offset(db.CreatePage(page, limit)).
		ToSql()
	logs.Logger.Sugar().Info(sql, args)
	err = db.Sqlx.Select(list, sql, args...)
	return
}

func Count(service Service) (count int, err error) {
	sqlA := db.SqlBuilder.
		Select("count(*)").
		From(service.TableName())
	sqlA = service.ProcessSqlWhere(sqlA)
	sql, args, err := sqlA.ToSql()
	logs.Logger.Sugar().Info(sql, args)
	err = db.Sqlx.Get(&count, sql, args...)
	return
}

func Options(column string, service Service) (list []string, err error) {
	sqlA := db.SqlBuilder.Select(column).From(service.TableName())
	sqlA = service.ProcessSqlWhere(sqlA)
	sql, args, err := sqlA.GroupBy(column).ToSql()
	logs.Logger.Sugar().Info(sql, args)
	rows, err := db.Sqlx.Queryx(sql, args...)
	if err != nil {
		return
	}
	defer rows.Close()
	var tmp string
	for rows.Next() {
		rows.Scan(&tmp)
		list = append(list, tmp)
	}
	return
}

func Delete(id int, service Service) (err error) {
	sql, args, err := db.SqlBuilder.Delete(service.TableName()).Where(db.Eq{"id": id}).ToSql()
	if err != nil {
		return
	}
	logs.Logger.Sugar().Info(sql, args)
	_, err = db.Sqlx.Exec(sql, args...)
	return
}

func DeleteJoin(id int, parent, node Service, nodeColumn string) (err error) {
	err = Delete(id, parent)
	if err != nil {
		return
	}
	// 删除 子表数据
	sql, args, err := db.SqlBuilder.Delete(node.TableName()).Where(db.Eq{nodeColumn: id}).ToSql()
	if err != nil {
		return
	}
	logs.Logger.Sugar().Info(sql, args)
	_, err = db.Sqlx.Exec(sql, args...)

	return
}

func DeleteByColumn(column string, parma interface{}, service Service) (err error) {
	sql, args, err := db.SqlBuilder.Delete(service.TableName()).Where(db.Eq{column: parma}).ToSql()
	if err != nil {
		return
	}
	logs.Logger.Sugar().Info(sql, args)
	_, err = db.Sqlx.Exec(sql, args...)
	return
}

func Insert(service Service) (lastInsertId int, err error) {
	sqlA := db.SqlBuilder.
		Insert(service.TableName())

	sqlA = service.ProcessSqlInsert(sqlA)

	sql, args, err := sqlA.ToSql()
	logs.Logger.Sugar().Info(sql, args)
	res, err := db.Sqlx.Exec(sql, args...)
	if err != nil {
		return
	}
	id, err := res.LastInsertId()

	lastInsertId = int(id)
	return
}

func Update(id int, service Service) (err error) {
	sqlA := db.SqlBuilder.
		Update(service.TableName())
	sqlA = service.ProcessSqlUpdate(id, sqlA)
	sql, args, err := sqlA.ToSql()
	logs.Logger.Sugar().Info(sql, args)
	_, err = db.Sqlx.Exec(sql, args...)
	return
}
