package Structs

import "github.com/golang-jwt/jwt"

//配置发送邮件
type Email struct {
	Host     string `mapstructure:"host"`     // 发件人邮箱
	Smtp     string `mapstructure:"smtp"`     // smtp服务器
	Addr     string `mapstructure:"addr"`     // 发件人邮箱地址
	Name     string `mapstructure:"name"`     // 发件人名称
	Password string `mapstructure:"password"` // 发件人邮箱密码
	Expires  int    `mapstructure:"expires"`  // 过期时间
}

//高级权限设置
type JWT struct {
	Key     string `mapstructure:"key"`     // 签名密钥
	Issuer  string `mapstructure:"issuer"`  // 发行人
	Expires int64  `mapstructure:"expires"` // 过期时间
}

// jwt
type Claims struct {
	UID      uint   // 用户 ID
	Username string // 用户名
	Password string //用户密码
	jwt.StandardClaims
}

//邮箱
type Emails struct {
	Email string `json:"email" binding:"email"` // 邮箱
}

//用户注册
type Regist struct {
	Username   string `json:"username" binding:"min=3,max=20"`   // 登录名
	Password   string `json:"password" binding:"min=8,max=32"`   // 密码
	Email      string `json:"email" binding:"email"`             // 邮箱
	VerifyCode string `json:"verify_code" binding:"min=6,max=6"` // 验证码
}
