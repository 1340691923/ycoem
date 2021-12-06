package middleware

import (
	"fmt"
	"runtime"
	"time"

	"ycoem/lib/jwt"
	"ycoem/lib/logs"
	model "ycoem/lib/models"
	"ycoem/lib/my_error"
	"ycoem/lib/response"
	"ycoem/lib/service/gm_user"
	"ycoem/lib/util"

	"github.com/kataras/iris"
	"go.uber.org/zap"
)

var res response.Response
var err error

func Cors(ctx iris.Context) {
	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.Header("Access-Control-Allow-Headers", "Content-Type") //header的类型
	ctx.Header("Access-Control-Allow-Headers", "x-token")      //header的类型
	ctx.Header("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,PATCH,OPTIONS")
	ctx.Header("Access-Control-Allow-Headers", "Content-Type, Accept, Authorization")
	ctx.Header("Access-Control-Allow-Credentials", "true")

	if ctx.Request().Method == "OPTIONS" {
		return
	}
	ctx.Next()
}

func JwtMiddleware(ctx iris.Context) {

	defer func() {
		if r := recover(); r != nil {
			//打印调用栈信息
			buf := make([]byte, 2048)
			logger := logs.Logger.Sugar()
			n := runtime.Stack(buf, false)
			stackInfo := fmt.Sprintf("%s", buf[:n])
			logger.Errorf("panic stack info %s", stackInfo)
			logger.Errorf("--->JwtMiddleware Error:", r)
		}
	}()

	var claims *jwt.Claims
	if util.GetToken(ctx) == "" {
		err = my_error.NewBusiness(TOKEN_ERROR, INVALID_PARAMS)
		res.Error(ctx, err)
		return
	} else {
		token := util.GetToken(ctx)
		b, err := model.IsBanToken(token)
		if util.FilterRedisNilErr(err) {
			logs.Logger.Error("model.IsBanToken", zap.String("err.Error()", err.Error()))
		}
		if b == true {
			err = my_error.NewBusiness(TOKEN_ERROR, ERROR_AUTH_CHECK_TOKEN_FAIL)
			res.Error(ctx, err)
			return
		}
		var service gm_user.GmUserService
		claims, err = jwt.ParseToken(token)
		if err != nil {
			err = my_error.NewBusiness(TOKEN_ERROR, ERROR_AUTH_CHECK_TOKEN_FAIL)
			res.Error(ctx, err)
			return
		} else if time.Now().Unix() > claims.ExpiresAt {
			err = my_error.NewBusiness(TOKEN_ERROR, ERROR_AUTH_CHECK_TOKEN_TIMEOUT)
			res.Error(ctx, err)
			return
		} else if !service.IsExitUser(claims) {
			err = my_error.NewBusiness(TOKEN_ERROR, ERROR_AUTH_CHECK_TOKEN_TIMEOUT)
			res.Error(ctx, err)
			return
		}
	}

	logs.Logger.Info("GmOperater:", zap.String("uri", ctx.Path()))
	gmOperaterModel := model.GmOperaterModel{
		Operater:       claims.Username,
		OperaterID:     int(claims.ID),
		OperaterAction: ctx.Path(),
	}
	gmOperaterModel.Insert()
	ctx.Next()
}
