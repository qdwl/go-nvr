// Code generated by goctl. DO NOT EDIT.
package types

type Base struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type UserLoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserLoginToken struct {
	AccessToken  string `json:"accessToken"`
	AccessExpire int64  `json:"accessExpire"`
	RefreshAfter int64  `json:"refreshAfter"`
}

type UserLoginResp struct {
	Base
	Data UserLoginToken `json:"data"`
}

type UserRole struct {
	Id     int64  `json:"id,optional"`
	Name   string `json:"name"`
	Type   string `json:"type"`
	Remark string `json:"remark"`
}

type UserRoleAddReq struct {
	UserRole
}

type UserRoleAddResp struct {
	Base
}

type UserRoleListReq struct {
	Page int    `form:"page,optional"`
	Size int    `form:"size,optional"`
	Type string `form:"type,optional"`
}

type UserRoleList struct {
	List  []*UserRole `json:"list"`
	Count int64       `json:"count"`
}

type UserRoleListResp struct {
	Base
	Data UserRoleList `json:"data"`
}

type UserRoleDeleteReq struct {
	Id int64 `path:"id"`
}

type UserRoleDeleteResp struct {
	Base
}

type UserRoleUpdateReq struct {
	UserRole
}

type UserRoleUpdateResp struct {
	Base
}

type UserInfo struct {
	Id       int64  `json:"id,optional"`
	UserId   string `json:"userId"`
	UserName string `json:"userName"`
	Password string `json:"password"`
	RoleId   int64  `json:"roleId"`
	Remark   string `json:"remark"`
}

type UserInfoAddReq struct {
	UserInfo
}

type UserInfoAddResp struct {
	Base
}

type UserInfoDeleteReq struct {
	Id int64 `path:"id"`
}

type UserInfoDeleteResp struct {
	Base
}

type UserInfoUpdateReq struct {
	UserInfo
}

type UserInfoUpdateResp struct {
	Base
}

type UserInfoListReq struct {
	Page int `form:"page,optional"`
	Size int `form:"size,optional"`
}

type UserInfoList struct {
	List  []*UserInfo `json:"list"`
	Count int64       `json:"count"`
}

type UserInfoListResp struct {
	Base
	Data UserInfoList `json:"data"`
}
