package constant

// ClashRules 订阅返回的 Clash 规则（自包含、无外部 rule-provider 依赖，Clash Verge 可直接导入）
const ClashRules = `rules:
  - GEOIP,CN,DIRECT
  - MATCH,PROXY
`
