package service

import (
	"gomoku-server/data/po"
	"gomoku-server/pkg/app"
)

type UserService struct {
}

func (u UserService) Test(aw *app.Wrapper) app.Result {

	type userReq struct {
		ID       int    `form:"id" binding:"required"`
		Username string `form:"username"`
		Credit   string `form:"credit"`
	}

	var tuser userReq
	err := aw.Ctx.ShouldBind(&tuser)

	if err != nil {
		return aw.Error(err.Error())
	}

	user := po.User{tuser.ID, tuser.Username, tuser.Credit}
	
	userDao.SaveUser(&user)

	return aw.Success(user)
}
