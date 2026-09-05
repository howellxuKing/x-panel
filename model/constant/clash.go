package constant

// ClashRules 订阅返回的 Clash 规则（自包含、无外部 rule-provider 依赖，Clash Verge 可直接导入）
const ClashRules = `rules:
  # 1. 本地网络与私人网段直连
  - GEOIP,private,DIRECT,no-resolve
  - IP-CIDR,127.0.0.0/8,DIRECT,no-resolve
  - IP-CIDR,172.16.0.0/12,DIRECT,no-resolve
  - IP-CIDR,192.168.0.0/16,DIRECT,no-resolve
  - IP-CIDR,10.0.0.0/8,DIRECT,no-resolve
  # 2. 常用国内公共 DNS / API 直连
  - IP-CIDR,223.5.5.5/32,DIRECT,no-resolve
  - IP-CIDR,119.29.29.29/32,DIRECT,no-resolve
  - IP-CIDR,180.76.76.76/32,DIRECT,no-resolve
  # 3. 国内软件/网站直连
  - DOMAIN-KEYWORD,weixin,DIRECT
  - DOMAIN-KEYWORD,wechat,DIRECT
  - DOMAIN-KEYWORD,qq,DIRECT
  - DOMAIN-SUFFIX,qq.com,DIRECT
  - DOMAIN-SUFFIX,tencent.com,DIRECT
  - DOMAIN-KEYWORD,taobao,DIRECT
  - DOMAIN-KEYWORD,alipay,DIRECT
  - DOMAIN-KEYWORD,jd,DIRECT
  - DOMAIN-SUFFIX,tmall.com,DIRECT
  - DOMAIN-SUFFIX,pinduoduo.com,DIRECT
  - DOMAIN-KEYWORD,baidu,DIRECT
  - DOMAIN-SUFFIX,baidu.com,DIRECT
  - DOMAIN-SUFFIX,163.com,DIRECT
  - DOMAIN-SUFFIX,sohu.com,DIRECT
  - DOMAIN-SUFFIX,sina.com.cn,DIRECT
  - DOMAIN-KEYWORD,bilibili,DIRECT
  - DOMAIN-SUFFIX,bilibili.com,DIRECT
  - DOMAIN-SUFFIX,hdslb.com,DIRECT
  - DOMAIN-KEYWORD,iqiyi,DIRECT
  - DOMAIN-KEYWORD,youku,DIRECT
  - DOMAIN-SUFFIX,music.163.com,DIRECT
  - DOMAIN-KEYWORD,steamcontent,DIRECT
  - DOMAIN-SUFFIX,cm.steampowered.com,DIRECT
  - DOMAIN-SUFFIX,epgames.com,DIRECT
  - GEOIP,CN,DIRECT
  # 4. 国外服务强制走代理
  - DOMAIN-KEYWORD,openai,Proxy
  - DOMAIN-KEYWORD,chatgpt,Proxy
  - DOMAIN-KEYWORD,anthropic,Proxy
  - DOMAIN-KEYWORD,claude,Proxy
  - DOMAIN-KEYWORD,google,Proxy
  - DOMAIN-KEYWORD,github,Proxy
  - DOMAIN-KEYWORD,youtube,Proxy
  - DOMAIN-KEYWORD,telegram,Proxy
  - DOMAIN-SUFFIX,twitter.com,Proxy
  - DOMAIN-SUFFIX,x.com,Proxy
  # 5. 兜底：其余流量走代理
  - MATCH,PROXY
`
