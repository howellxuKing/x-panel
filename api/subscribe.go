package api

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"trojan-panel/model/constant"
	"trojan-panel/model/dto"
	"trojan-panel/model/vo"
	"trojan-panel/service"
)

// ClashSubscribe 获取Clash订阅地址
func ClashSubscribe(c *gin.Context) {
	accountVo := service.GetCurrentAccount(c)
	password, err := service.SelectConnectPassword(&accountVo.Id, &accountVo.Username)
	if err != nil {
		vo.Fail(err.Error(), c)
		return
	}
	vo.Success(fmt.Sprintf("/api/auth/subscribe/%s", base64.StdEncoding.EncodeToString([]byte(password))), c)
}

// ClashSubscribeForSb 获取指定人的Clash订阅地址
func ClashSubscribeForSb(c *gin.Context) {
	var accountRequiredIdDto dto.RequiredIdDto
	_ = c.ShouldBindQuery(&accountRequiredIdDto)
	if err := validate.Struct(&accountRequiredIdDto); err != nil {
		vo.Fail(constant.ValidateFailed, c)
		return
	}
	password, err := service.SelectConnectPassword(accountRequiredIdDto.Id, nil)
	if err != nil {
		vo.Fail(err.Error(), c)
		return
	}
	vo.Success(fmt.Sprintf("/api/auth/subscribe/%s", base64.StdEncoding.EncodeToString([]byte(password))), c)
}

// Subscribe 订阅
func Subscribe(c *gin.Context) {
	token := c.Param("token")
	//userAgent := c.Request.Header.Get("User-Agent")
	tokenDecode, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		vo.Fail(constant.SysError, c)
		return
	}
	pass := string(tokenDecode)

	//if strings.HasPrefix(userAgent, constant.ClashforWindows) {
	account, userInfo, clashConfigYaml, _, err := service.SubscribeClash(pass)
	if err != nil {
		vo.Fail(err.Error(), c)
		return
	}
	result := fmt.Sprintf(`%s
%s`, string(clashConfigYaml), getClashRules())

	c.Header("content-disposition", fmt.Sprintf("attachment; filename=%s.yaml", *account.Username))
	c.Header("profile-update-interval", "12")
	c.Header("subscription-userinfo", userInfo)
	c.String(200, result)
	return
	//}
	//vo.Fail("This client is not supported", c)
}

// getClashRules 返回订阅规则
// 优先读取面板可编辑的规则模板文件（config/template/template-clash-rule.yaml，系统设置→Clash规则 保存后即写入）；
// 文件缺失、为空、为旧版默认内容或依赖 jsDelivr 时，回退到当前内置规则。
// 管理员修改过的模板保持不变。
func getClashRules() string {
	content, err := os.ReadFile(constant.ClashRuleFilePath)
	if err != nil {
		return constant.ClashRules
	}
	rules := strings.TrimSpace(string(content))
	normalizedRules := strings.ReplaceAll(rules, "\r\n", "\n")
	rulesDigest := fmt.Sprintf("%x", sha256.Sum256([]byte(normalizedRules)))
	if rules == "" ||
		strings.Contains(rules, "rule-provider") ||
		strings.Contains(rules, "jsdelivr") ||
		rulesDigest == constant.ClashRulesV233SHA256 ||
		rules == "rules:\n  - GEOIP,CN,DIRECT\n  - MATCH,PROXY" {
		return constant.ClashRules
	}
	return rules
}
