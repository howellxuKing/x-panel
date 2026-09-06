<div align="center">
<img src="https://raw.githubusercontent.com/howellxuKing/x-panel-ui/main/public/logo.png" width="120" alt="X Panel" />
<h1>X Panel</h1>
<p>多用户代理管理面板（Trojan Panel 定制版）——一键安装 · 订阅即用 · 去品牌重做</p>
</div>

## 简介

X Panel 是一个基于 [Trojan Panel](https://github.com/trojanpanel) 深度定制的多用户代理管理面板，支持 Xray（VLESS / VLESS+Reality / VMess / Trojan / Shadowsocks）、Trojan-Go、Hysteria、Hysteria2、NaiveProxy，自带用户管理、流量统计、到期管理、一键伪装站。

**本仓库为 X Panel 后端**（Go 编写的 API + 订阅服务），配套仓库：

| 组件 | 仓库 | 镜像 |
|---|---|---|
| 后端（本仓库） | `howellxuKing/x-panel` | `ghcr.io/howellxuking/x-panel` |
| 前端 UI | `howellxuKing/x-panel-ui` | `ghcr.io/howellxuking/x-panel-ui` |
| 一键安装脚本 | `howellxuKing/install-script` | — |
| 内核 | Trojan Panel Core（未改动） | `jonssonyan/trojan-panel-core` |

## 与上游的主要差异

1. **订阅干净化**：去掉 cdn.jsdelivr.net 的 rule-providers，订阅规则自包含 → **Clash Verge 秒导入**，不再卡下载
2. **内置分流规则**：国内电商/社交直连（DIRECT）+ 国外 AI（OpenAI/Claude/Google 等）走代理（PROXY）+ 兜底
3. **去品牌重做**：更名 X Panel（浏览器标题/种子数据/logo），移除「项目地址/项目文档」入口
4. **镜像走 GHCR**：前端/后端镜像由 GitHub Actions 自动构建推送，可匿名拉取

## 一键安装

```bash
source <(curl -L https://raw.githubusercontent.com/howellxuKing/install-script/main/install_script.sh)
```

菜单说明：

| 选项 | 功能 |
|---|---|
| 1 / 2 / 3 | 安装 前端 UI / 后端 / 内核 |
| 8 / 9 | 更新 前端 / 后端 |
| 其他 | 卸载、重置管理员密码、伪装站等 |

安装完成后访问 `https://你的域名:8888` 登录管理面板。

## 支持的协议

- Xray：VLESS、VLESS+Reality、VMess、Trojan、Shadowsocks
- Trojan-Go、Hysteria、Hysteria2、NaiveProxy

客户端：Clash Verge / Clash Meta、v2rayN、Shadowrocket 等。

## 开发构建

后端（本仓库）：

```bash
go build   # 或参考 compile.bat（Windows 交叉编译）
```

前端见 [x-panel-ui](https://github.com/howellxuKing/x-panel-ui)（`npm run build`，vue-cli 4 + Node 22 需 `NODE_OPTIONS=--openssl-legacy-provider`）。

## 致谢

本项目是 [Trojan Panel](https://github.com/trojanpanel) 的定制分支，感谢原作者 jonssonyan 与 Trojan Panel 社区的开源贡献。

## License

MIT
