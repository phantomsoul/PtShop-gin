package routes

import (
	"github.com/gin-gonic/gin"
	"pt-gin/config"
	"pt-gin/middleware/verify"
)

var (
	vfy *verify.Verify
)

func RegisterApiRouter(c *config.Config, router *gin.Engine) {
	vfy = verify.New(c.Verify.Keys) // 如果是GET方法，就解析querystring的字段，加上ts和appkey去计算sign，如果是POST方法，就取form字段，加上ts和appkey去计算sign
	//checkJwtAcl := auth.Middleware(auth.JwtAuthDriverKey, true)
	//checkJwtNoAcl := auth.Middleware(auth.JwtAuthDriverKey, false)
	//
	//api := router.Group("/api")
	//api.POST("/login", vfy.Verify, controllers.Login) // 带安全签名校验的登录
	//api.GET("/index", controllers.IndexApi)          // 服务可用性验证
	//api.POST("/register", controllers.Register)      // 用户名密码注册
	//api.POST("/login", controllers.Login)            // 用户名密码登录
	//api.POST("/sms/push", controllers.SendSmsCode)   // 发送短信验证码
	//api.POST("/sms/check", controllers.CheckSmsCode) // 校验短信验证码

	//user := api.Group("/user")
	//{
	//	user.POST("/exist/check", controllers.CheckExist)             // 检查用户名是否存在
	//	user.POST("/enterprise/list", controllers.GetUsersEnterprise) // 获取用户管理的企业列表
	//}

	//ctm := api.Group("/ctm")
	//{
	//	ctm.POST("/business/get", checkJwtAcl, controllers.GetBaseBusinessInfo) // 获取工商信息
	//	ctm.POST("/info/add", checkJwtNoAcl, controllers.AddCtmInfo)            // 添加企业信息
	//	ctm.POST("/menu/get", checkJwtNoAcl, controllers.GetCtmMenuAndToken)    // 获取指定企业菜单和Token
	//}
	//
	//rpa := api.Group("/rpa")
	//{
	//	rpa.POST("/bank-receipt/list", controllers.BankReceiptList) // 获取银行回单列表
	//	rpa.POST("/bank-receipt/add", controllers.AddBankReceipt) // 添加银行回单
	//	rpa.POST("/bank-receipt/del", controllers.DelBankReceipt) // 删除银行回单
	//	rpa.POST("/invoice/add", controllers.AddInvoice) // 添加发票
	//	rpa.POST("/invoice/list", controllers.GetInvoice) // 获取发票列表
	//	rpa.POST("/invoice/del", controllers.DelInvoice) // 删除发票
	//	rpa.POST("/scan-data/lists", controllers.ScanDataList) // 获取批量扫描的数据
	//	rpa.POST("/scan-data/del", controllers.DelScanData) // 删除扫描的数据
	//	rpa.POST("/query-match-sys", controllers.QueryMatchSys) // 查询是否匹配系统
	//	rpa.POST("/scan-invoice/info", controllers.ScanInvoiceInfo) // 获取扫描的发票数据
	//	rpa.POST("/scan-invoice-detail/del", controllers.DelScanInvoiceDetail) // 删除扫描的发票详情
	//	rpa.POST("/scan-entry-invoice", controllers.ScanEntryInvoice) // 录入原始凭证发票
	//}

	//voucher := api.Group("/voucher")
	//{
	//	//voucher.POST("/invoice/list", controllers.GetInvoices) // 获取发票列表
	//	//voucher.POST("/invoice/add", controllers.AddBankReceipt)   // 添加发票列表
	//}
	//salary := api.Group("/salary")
	//{
	//	salary.POST("/salary/list", controllers.GetSalary) // 获取某个账期工资
	//	salary.POST("/salary/details", controllers.GetSalaryDetails)   // 获取工资详情列表
	//	salary.POST("/salary/del", controllers.DelSalary)   // 删除工资
	//	salary.POST("/salary/update", controllers.UpdateSalary)   // 删除工资
	//}

	// jwt auth middleware
	//api.Use(checkJwtAcl)
	//{
	//	api.GET("/store", controllers.StoreExample)
	//	api.GET("/db", controllers.DBExample)
	//}
}
