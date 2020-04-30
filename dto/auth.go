package dto

/**
 * Token用户信息结构
 */
type JwtUserInfo struct {
	UserID   int64  `json:"user_id"`   //用户ID
	UserName string `json:"user_name"` //用户名（手机号）
	CtmID    int64  `json:"ctm_id"`    //当前登录企业ID
	RoleID   string `json:"role_id"`   //当前企业角色
	IsAdmin  int8   `json:"is_admin"`  //是否为当前企业管理员
	IsAgent  int8   `json:"is_agent"`  //当前企业是否为代理服务商身份
}