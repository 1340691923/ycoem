package concat

import (
	"ycoem/lib/db"
	model "ycoem/lib/models"
	"ycoem/lib/response"
	"ycoem/lib/util"

	jsoniter "github.com/json-iterator/go"
)

type ServiceDepService struct {
	Area string
	Manager string
	Name string
	Support string
}

func(this ServiceDepService)TableName() string{
	return "service_dep"
}

func(this ServiceDepService)ProcessSqlWhere(sqlA db.SelectBuilder)db.SelectBuilder{
	if this.Name !="" {
		sqlA = sqlA.Where(db.Eq{"name":this.Name})
	}
	if this.Manager !="" {
		sqlA = sqlA.Where(db.Eq{"manager":this.Manager})
	}
	if this.Area !="" {
		sqlA = sqlA.Where("area like ? ",db.CreateLike(this.Area))
	}

	return sqlA
}

func(this ServiceDepService)ProcessSqlInsert(sqlA db.InsertBuilder)db.InsertBuilder{
	sqlA = sqlA.Columns("area", "manager","support","name").
		Values(this.Area, this.Manager,this.Support,this.Name)
	return sqlA
}

func(this ServiceDepService)ProcessSqlUpdate(id int,sqlA db.UpdateBuilder)db.UpdateBuilder{
	sqlA = sqlA.
		Where(db.Eq{"id":id}).
		SetMap(map[string]interface{}{
			"area":this.Area,
			"manager":this.Manager,
			"support":this.Support,
			"name":this.Name,
		})
	return sqlA
}

func(this *ServiceDepService)All()(tmp map[string][]response.ServiceDep,err error){
	str := db.Get(model.ServiceDepRedisModel)

	if util.FilterRedisNilErr(err){
		return
	}
	if str != ""{
		var json = jsoniter.ConfigCompatibleWithStandardLibrary
		err = json.Unmarshal([]byte(str),&tmp)
		return
	}
	sqlA := db.SqlBuilder.
		Select("*").
		From(this.TableName())

	sql,args,err := sqlA.ToSql()

	var list []model.ServiceDepModel
	err = db.Sqlx.Select(&list,sql,args...)

	tmpsd := response.ServiceDep{}

	tmp = make(map[string][]response.ServiceDep)

	for _,v := range list{
		tmpsd.Support = v.Support
		tmpsd.Area = v.Area
		tmpsd.Manager = v.Manager
		tmp[v.Name] = append(tmp[v.Name], tmpsd)
	}

	go func() {
		tmp2,err:= jsoniter.Marshal(tmp)
		if err!=nil{
			return
		}
		db.Set(model.ServiceDepRedisModel,string(tmp2))
	}()
	return
}
