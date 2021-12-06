package options

import "ycoem/lib/db"

func Options(column,tableName string)(list []string,err error){

	sql,args,err := db.SqlBuilder.Select(column).From(tableName).GroupBy(column).ToSql()

	rows,err := db.Sqlx.Queryx(sql,args...)
	if err!=nil{
		return
	}
	defer rows.Close()
	var tmp string
	for rows.Next(){
		rows.Scan(&tmp)
		list = append(list, tmp)
	}
	return
}
