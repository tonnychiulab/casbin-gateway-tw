<h1 align="center" style="border-bottom: none;">📦⚡️ Casbin Gateway (正體中文版)</h1>
<h3 align="center">看清您電腦上每個 AI 程式設計 Agent 在做什麼，以及它背後的 API 是否真是您付費購買的那個版本。</h3>
<p align="center">
  <a href="https://github.com/tonnychiulab/casbin-gateway-tw/actions/workflows/golangci-lint.yml"><img alt="Lint" src="https://img.shields.io/github/actions/workflow/status/tonnychiulab/casbin-gateway-tw/golangci-lint.yml?branch=master&style=flat-square&logo=github&logoColor=white&label=lint"></a>
  <a href="https://github.com/tonnychiulab/casbin-gateway-tw/actions/workflows/build.yml"><img alt="Build" src="https://img.shields.io/github/actions/workflow/status/tonnychiulab/casbin-gateway-tw/build.yml?branch=master&style=flat-square&logo=github&logoColor=white&label=build"></a>
  <a href="https://pkg.go.dev/github.com/apache/casbin-gateway"><img alt="Go Reference" src="https://img.shields.io/badge/go.dev-reference-00ADD8?style=flat-square&logo=go&logoColor=white"></a>
  <a href="https://github.com/tonnychiulab/casbin-gateway-tw/releases/latest"><img alt="Release" src="https://img.shields.io/github/v/release/tonnychiulab/casbin-gateway-tw?style=flat-square&logo=github&logoColor=white&color=3b82f6"></a>
  <a href="https://github.com/tonnychiulab/casbin-gateway-tw/blob/master/LICENSE"><img alt="License" src="https://img.shields.io/github/license/tonnychiulab/casbin-gateway-tw?style=flat-square&color=3b82f6"></a>
  <a href="https://github.com/tonnychiulab/casbin-gateway-tw/issues"><img alt="Issues" src="https://img.shields.io/github/issues/tonnychiulab/casbin-gateway-tw?style=flat-square&logo=github&logoColor=white&color=3b82f6"></a>
  <a href="https://github.com/tonnychiulab/casbin-gateway-tw/stargazers"><img alt="Stars" src="https://img.shields.io/github/stars/tonnychiulab/casbin-gateway-tw?style=flat-square&logo=github&logoColor=white&color=3b82f6"></a>
  <a href="https://github.com/tonnychiulab/casbin-gateway-tw/network/members"><img alt="Forks" src="https://img.shields.io/github/forks/tonnychiulab/casbin-gateway-tw?style=flat-square&logo=github&logoColor=white&color=3b82f6"></a>
  <a href="https://discord.gg/S5UjpzGZjN"><img alt="Discord" src="https://img.shields.io/discord/1022748306096537660?style=flat-square&logo=discord&logoColor=white&label=discord&color=5865F2"></a>
</p>

<p align="center">
  <a href="./README.md">English</a> | <a href="./README_zh.md">简体中文</a> | <b>繁體中文</b>
</p>

> [!NOTE]
> **本專案源起與開源授權聲明**：
> 本專案為學習、實驗與資安防護加固用途所建立的 Fork 專案，源自官方優秀的上游專案 **[Apache Casbin-Gateway (Incubating)](https://github.com/apache/casbin-gateway)**。
> 本專案完全遵循原專案的 **[Apache License 2.0](https://github.com/tonnychiulab/casbin-gateway-tw/blob/master/LICENSE)** 開源授權條款，完整保留原專案著作權與作者宣告（Copyright 2022-2026 The Casbin Authors & The Apache Software Foundation）。
> 本 Fork 在原作者深厚的基礎上，額外針對生產環境反向代理架構加固了安全性修補（包含防範反向代理環境下自動登入提權、管理者預設密碼強制輪替、本機檔案瀏覽防目錄遍歷、SSRF 多 IP 完整過濾與 DNS Rebinding 防護等），歡迎社群共同交流與檢驗！

<p align="center"><b>電腦上的每一個程式設計 Agent</b></p>
<p align="center">
  <a href="https://claude.com" title="Claude Code"><img src="https://www.google.com/s2/favicons?domain=claude.com&sz=64" width="30" height="30" alt="Claude Code"></a>
  <a href="https://claude.ai" title="Claude Desktop"><img src="https://www.google.com/s2/favicons?domain=claude.ai&sz=64" width="30" height="30" alt="Claude Desktop"></a>
  <a href="https://openai.com" title="Codex / Codex CLI"><img src="https://www.google.com/s2/favicons?domain=openai.com&sz=64" width="30" height="30" alt="Codex"></a>
  <a href="https://gemini.google.com" title="Gemini CLI"><img src="https://www.google.com/s2/favicons?domain=gemini.google.com&sz=64" width="30" height="30" alt="Gemini CLI"></a>
  <a href="https://cursor.com" title="Cursor / Cursor Agent"><img src="https://www.google.com/s2/favicons?domain=cursor.com&sz=64" width="30" height="30" alt="Cursor"></a>
  <a href="https://windsurf.com" title="Windsurf"><img src="https://www.google.com/s2/favicons?domain=windsurf.com&sz=64" width="30" height="30" alt="Windsurf"></a>
  <a href="https://cline.bot" title="Cline"><img src="https://www.google.com/s2/favicons?domain=cline.bot&sz=64" width="30" height="30" alt="Cline"></a>
  <a href="https://github.com/QwenLM/qwen-code" title="Qwen Code"><img src="https://www.google.com/s2/favicons?domain=qwen.ai&sz=64" width="30" height="30" alt="Qwen Code"></a>
  <a href="https://platform.iflow.cn/en/cli" title="iFlow CLI"><img src="https://www.google.com/s2/favicons?domain=iflow.cn&sz=64" width="30" height="30" alt="iFlow CLI"></a>
  <a href="https://moonshotai.github.io/kimi-code" title="Kimi Code CLI"><img src="https://www.google.com/s2/favicons?domain=kimi.com&sz=64" width="30" height="30" alt="Kimi Code CLI"></a>
  <a href="https://www.codebuddy.ai" title="CodeBuddy Code"><img src="https://www.google.com/s2/favicons?domain=codebuddy.ai&sz=64" width="30" height="30" alt="CodeBuddy Code"></a>
  <a href="https://x.ai/cli" title="Grok Build"><img src="https://www.google.com/s2/favicons?domain=x.ai&sz=64" width="30" height="30" alt="Grok Build"></a>
  <a href="https://roocode.com" title="Roo Code"><img src="https://www.google.com/s2/favicons?domain=roocode.com&sz=64" width="30" height="30" alt="Roo Code"></a>
  <a href="https://github.com/github/copilot-cli" title="Copilot CLI"><img src="https://www.google.com/s2/favicons?domain=github.com&sz=64" width="30" height="30" alt="Copilot CLI"></a>
  <a href="https://continue.dev" title="Continue"><img src="https://www.google.com/s2/favicons?domain=continue.dev&sz=64" width="30" height="30" alt="Continue"></a>
  <a href="https://zed.dev" title="Zed"><img src="https://www.google.com/s2/favicons?domain=zed.dev&sz=64" width="30" height="30" alt="Zed"></a>
  <a href="https://aider.chat" title="Aider"><img src="https://www.google.com/s2/favicons?domain=aider.chat&sz=64" width="30" height="30" alt="Aider"></a>
  <a href="https://github.com/aaif-goose/goose" title="goose"><img src="https://www.google.com/s2/favicons?domain=block.github.io/goose&sz=64" width="30" height="30" alt="goose"></a>
  <a href="https://github.com/charmbracelet/crush" title="Crush"><img src="https://www.google.com/s2/favicons?domain=charm.sh&sz=64" width="30" height="30" alt="Crush"></a>
  <a href="https://factory.ai" title="Droid"><img src="https://www.google.com/s2/favicons?domain=factory.ai&sz=64" width="30" height="30" alt="Droid"></a>
  <a href="https://trae.ai" title="Trae"><img src="https://www.google.com/s2/favicons?domain=trae.ai&sz=64" width="30" height="30" alt="Trae"></a>
  <a href="https://opencode.ai" title="opencode / opencode Desktop"><img src="https://www.google.com/s2/favicons?domain=opencode.ai&sz=64" width="30" height="30" alt="opencode"></a>
  <a href="https://openagentai.org" title="OpenAgent"><img src="https://www.google.com/s2/favicons?domain=openagentai.org&sz=64" width="30" height="30" alt="OpenAgent"></a>
  <a href="https://openclaw.ai" title="OpenClaw"><img src="https://www.google.com/s2/favicons?domain=openclaw.ai&sz=64" width="30" height="30" alt="OpenClaw"></a>
  <a href="https://nousresearch.com" title="Hermes Agent"><img src="https://www.google.com/s2/favicons?domain=nousresearch.com&sz=64" width="30" height="30" alt="Hermes Agent"></a>
  <a href="https://deepseek.com" title="DeepSeek Harness"><img src="https://www.google.com/s2/favicons?domain=deepseek.com&sz=64" width="30" height="30" alt="DeepSeek Harness"></a>
  <a href="https://pi.dev" title="Pi"><img src="https://www.google.com/s2/favicons?domain=pi.dev&sz=64" width="30" height="30" alt="Pi"></a>
</p>

<p align="center"><b>單一入口對接任意模型廠商</b></p>
<p align="center">
  <a href="https://openai.com" title="OpenAI"><img src="https://www.google.com/s2/favicons?domain=openai.com&sz=64" width="28" height="28" alt="OpenAI"></a>
  <a href="https://anthropic.com" title="Anthropic"><img src="https://www.google.com/s2/favicons?domain=anthropic.com&sz=64" width="28" height="28" alt="Anthropic"></a>
  <a href="https://gemini.google.com" title="Google Gemini"><img src="https://www.google.com/s2/favicons?domain=gemini.google.com&sz=64" width="28" height="28" alt="Google Gemini"></a>
  <a href="https://x.ai" title="xAI"><img src="https://www.google.com/s2/favicons?domain=x.ai&sz=64" width="28" height="28" alt="xAI"></a>
  <a href="https://mistral.ai" title="Mistral"><img src="https://www.google.com/s2/favicons?domain=mistral.ai&sz=64" width="28" height="28" alt="Mistral"></a>
  <a href="https://cohere.com" title="Cohere"><img src="https://www.google.com/s2/favicons?domain=cohere.com&sz=64" width="28" height="28" alt="Cohere"></a>
  <a href="https://perplexity.ai" title="Perplexity"><img src="https://www.google.com/s2/favicons?domain=perplexity.ai&sz=64" width="28" height="28" alt="Perplexity"></a>
  <a href="https://deepseek.com" title="DeepSeek"><img src="https://www.google.com/s2/favicons?domain=deepseek.com&sz=64" width="28" height="28" alt="DeepSeek"></a>
  <a href="https://moonshot.cn" title="Moonshot"><img src="https://www.google.com/s2/favicons?domain=moonshot.cn&sz=64" width="28" height="28" alt="Moonshot"></a>
  <a href="https://z.ai" title="Zhipu GLM"><img src="https://www.google.com/s2/favicons?domain=z.ai&sz=64" width="28" height="28" alt="Zhipu GLM"></a>
  <a href="https://dashscope.aliyuncs.com" title="Qwen"><img src="https://www.google.com/s2/favicons?domain=aliyun.com&sz=64" width="28" height="28" alt="Qwen"></a>
  <a href="https://minimaxi.com" title="MiniMax"><img src="https://www.google.com/s2/favicons?domain=minimaxi.com&sz=64" width="28" height="28" alt="MiniMax"></a>
  <a href="https://baichuan-ai.com" title="Baichuan"><img src="https://www.google.com/s2/favicons?domain=baichuan-ai.com&sz=64" width="28" height="28" alt="Baichuan"></a>
  <a href="https://stepfun.com" title="StepFun"><img src="https://www.google.com/s2/favicons?domain=stepfun.com&sz=64" width="28" height="28" alt="StepFun"></a>
  <a href="https://volcengine.com" title="Volcengine Ark"><img src="https://www.google.com/s2/favicons?domain=volcengine.com&sz=64" width="28" height="28" alt="Volcengine Ark"></a>
  <a href="https://cloud.tencent.com" title="Tencent Hunyuan"><img src="https://www.google.com/s2/favicons?domain=cloud.tencent.com&sz=64" width="28" height="28" alt="Tencent Hunyuan"></a>
  <a href="https://baidu.com" title="Baidu ERNIE"><img src="https://www.google.com/s2/favicons?domain=baidu.com&sz=64" width="28" height="28" alt="Baidu ERNIE"></a>
  <a href="https://lingyiwanwu.com" title="01.AI"><img src="https://www.google.com/s2/favicons?domain=lingyiwanwu.com&sz=64" width="28" height="28" alt="01.AI"></a>
  <a href="https://ai21.com" title="AI21 Labs"><img src="https://www.google.com/s2/favicons?domain=ai21.com&sz=64" width="28" height="28" alt="AI21 Labs"></a>
  <a href="https://reka.ai" title="Reka"><img src="https://www.google.com/s2/favicons?domain=reka.ai&sz=64" width="28" height="28" alt="Reka"></a>
  <a href="https://siliconflow.cn" title="SiliconFlow"><img src="https://www.google.com/s2/favicons?domain=siliconflow.cn&sz=64" width="28" height="28" alt="SiliconFlow"></a>
  <a href="https://groq.com" title="Groq"><img src="https://www.google.com/s2/favicons?domain=groq.com&sz=64" width="28" height="28" alt="Groq"></a>
  <a href="https://together.ai" title="Together AI"><img src="https://www.google.com/s2/favicons?domain=together.ai&sz=64" width="28" height="28" alt="Together AI"></a>
  <a href="https://fireworks.ai" title="Fireworks AI"><img src="https://www.google.com/s2/favicons?domain=fireworks.ai&sz=64" width="28" height="28" alt="Fireworks AI"></a>
  <a href="https://novita.ai" title="Novita AI"><img src="https://www.google.com/s2/favicons?domain=novita.ai&sz=64" width="28" height="28" alt="Novita AI"></a>
  <a href="https://deepinfra.com" title="DeepInfra"><img src="https://www.google.com/s2/favicons?domain=deepinfra.com&sz=64" width="28" height="28" alt="DeepInfra"></a>
  <a href="https://modelscope.cn" title="ModelScope"><img src="https://www.google.com/s2/favicons?domain=modelscope.cn&sz=64" width="28" height="28" alt="ModelScope"></a>
  <a href="https://ppinfra.com" title="PPIO"><img src="https://www.google.com/s2/favicons?domain=ppinfra.com&sz=64" width="28" height="28" alt="PPIO"></a>
  <a href="https://cerebras.ai" title="Cerebras"><img src="https://www.google.com/s2/favicons?domain=cerebras.ai&sz=64" width="28" height="28" alt="Cerebras"></a>
  <a href="https://sambanova.ai" title="SambaNova"><img src="https://www.google.com/s2/favicons?domain=sambanova.ai&sz=64" width="28" height="28" alt="SambaNova"></a>
  <a href="https://hyperbolic.xyz" title="Hyperbolic"><img src="https://www.google.com/s2/favicons?domain=hyperbolic.xyz&sz=64" width="28" height="28" alt="Hyperbolic"></a>
  <a href="https://nebius.ai" title="Nebius AI Studio"><img src="https://www.google.com/s2/favicons?domain=nebius.ai&sz=64" width="28" height="28" alt="Nebius AI Studio"></a>
  <a href="https://lambda.ai" title="Lambda"><img src="https://www.google.com/s2/favicons?domain=lambda.ai&sz=64" width="28" height="28" alt="Lambda"></a>
  <a href="https://baseten.co" title="Baseten"><img src="https://www.google.com/s2/favicons?domain=baseten.co&sz=64" width="28" height="28" alt="Baseten"></a>
  <a href="https://build.nvidia.com" title="NVIDIA NIM"><img src="https://www.google.com/s2/favicons?domain=nvidia.com&sz=64" width="28" height="28" alt="NVIDIA NIM"></a>
  <a href="https://ollama.com" title="Ollama"><img src="https://www.google.com/s2/favicons?domain=ollama.com&sz=64" width="28" height="28" alt="Ollama"></a>
  <a href="https://lmstudio.ai" title="LM Studio"><img src="https://www.google.com/s2/favicons?domain=lmstudio.ai&sz=64" width="28" height="28" alt="LM Studio"></a>
  <a href="https://docs.vllm.ai" title="vLLM"><img src="https://www.google.com/s2/favicons?domain=vllm.ai&sz=64" width="28" height="28" alt="vLLM"></a>
  <a href="https://github.com/ggml-org/llama.cpp" title="llama.cpp"><img src="https://www.google.com/s2/favicons?domain=github.com&sz=64" width="28" height="28" alt="llama.cpp"></a>
  <a href="https://openrouter.ai" title="OpenRouter"><img src="https://www.google.com/s2/favicons?domain=openrouter.ai&sz=64" width="28" height="28" alt="OpenRouter"></a>
  <a href="https://aihubmix.com" title="AiHubMix"><img src="https://www.google.com/s2/favicons?domain=aihubmix.com&sz=64" width="28" height="28" alt="AiHubMix"></a>
  <a href="https://302.ai" title="302.AI"><img src="https://www.google.com/s2/favicons?domain=302.ai&sz=64" width="28" height="28" alt="302.AI"></a>
  <a href="https://vercel.com/docs/ai-gateway" title="Vercel AI Gateway"><img src="https://www.google.com/s2/favicons?domain=vercel.com&sz=64" width="28" height="28" alt="Vercel AI Gateway"></a>
  <a href="https://litellm.ai" title="LiteLLM"><img src="https://www.google.com/s2/favicons?domain=litellm.ai&sz=64" width="28" height="28" alt="LiteLLM"></a>
</p>

<p align="center">
  <a href="https://cdn.casbin.org/img/casbin-gateway.gif"><img alt="Casbin Gateway" src="https://cdn.casbin.org/img/casbin-gateway.gif" width="900"></a>
</p>

- **這個 API 金鑰，真的是賣給您的那一個嗎？** 第三方轉發商可能拿便宜模型冒充、偽造快取命中，或者暗中丟棄宣稱支援的參數。[真偽檢測](#殺手鐧那個-key-背後的-api真是賣給你的那個嗎) 會主動探測每個 Provider，給予 A–F 評級；[不需安裝也能單獨測試特定端點](#不需安裝單測一個介面)。
- **每個 Agent 做了什麼、花了多少、傳送到何處。** 包含每條提示詞與工具呼叫，從 Agent 自身的對話歷史紀錄中讀取花費；在 Windows 上更能追蹤每個 Agent 程序將資料上傳至何處，涵蓋任何代理軟體都無法監測的隱蔽連線。
- **所有 Agent，統一集中管控。** 將 Claude Code、Codex、Cursor、Gemini CLI 等無縫連接至 44 個模型廠商中的任意一家，單一廠商額度耗盡或服務中斷時[自動容錯移轉至備用供應商](#provider-無法應答時備用-provider-與模型路由)，規範每個 Agent 的操作權限，並支援直接安裝、升級或版本降級 Agent 本身。

---

## 一鍵安裝執行 (Run it)

僅需一行指令。無需預裝資料庫、Go、Node.js，亦無需任何繁瑣設定。

### Linux 與 macOS：

```bash
curl -fsSL https://raw.githubusercontent.com/tonnychiulab/casbin-gateway-tw/master/scripts/install.sh | bash
```

### Windows (在 PowerShell 中執行)：

```powershell
irm https://raw.githubusercontent.com/tonnychiulab/casbin-gateway-tw/master/scripts/install.ps1 | iex
```

兩者均會自動下載適合本機作業系統與硬體架構的二進位檔案，解壓縮至 `~/.local/share/casbin-gateway`（Windows 下為 `%LOCALAPPDATA%\casbin-gateway`），並將 `casbin-gateway` 指令加入系統 PATH、自動啟動並設定登入時開機自動啟動。執行安裝的終端機視窗將立即釋放還給您。

Gateway 啟動後會在獨立視窗中開啟。關閉視窗後不會結束執行，而是縮小至系統匣（System Tray）於背景繼續運作；點擊系統匣圖示可隨時重新開啟視窗、切換**開機自動啟動**以及徹底結束程式。若您習慣使用瀏覽器操作，或主機未安裝圖形桌面環境，亦可直接連線至 **http://localhost:17000**；使用 `casbin-gateway start` 則僅於背景啟動後端服務。

資料庫採用單一 SQLite 檔案儲存於自身目錄下，開箱即用。

---

## 介面預覽

| 本機上的每一個 Agent | 這些 Agent 身上所具備的每項擴充 |
| :---: | :---: |
| [![Agents](https://cdn.casbin.org/img/casbin-gateway-home.png)](https://cdn.casbin.org/img/casbin-gateway-home.png) | [![Skills, MCP & Prompts](https://cdn.casbin.org/img/casbin-gateway-skills.png)](https://cdn.casbin.org/img/casbin-gateway-skills.png) |
| 顯示每個 Agent 連接的端點、登入帳號、已花費金額及當前執行狀態 | 並列檢視所有 Agent 的技能、MCP 伺服器與提示詞設定，並支援跨 Agent 複製 |

| 每個 Agent 允許執行的權限 | 每個 Agent 當前安裝的版本 |
| :---: | :---: |
| [![權限](https://cdn.casbin.org/img/casbin-gateway-permissions.png)](https://cdn.casbin.org/img/casbin-gateway-permissions.png) | [![Agent 版本](https://cdn.casbin.org/img/casbin-gateway-versions.png)](https://cdn.casbin.org/img/casbin-gateway-versions.png) |
| 提供 40+ 項細部控制開關，編譯為 Casbin 策略，嚴格審核每筆轉發請求 | 對比本機安裝版本與套件管理器最新發布版本，一鍵完成安裝、升級與版本降級 |

| 每個 Agent 的實際花費統計 | 每個模型廠商的集中端點 |
| :---: | :---: |
| [![用量](https://cdn.casbin.org/img/casbin-gateway-usage.png)](https://cdn.casbin.org/img/casbin-gateway-usage.png) | [![新建 Provider](https://cdn.casbin.org/img/casbin-gateway-new-provider.png)](https://cdn.casbin.org/img/casbin-gateway-new-provider.png) |
| 直接讀取 Agent 本機對話歷史紀錄，即使未經 Gateway 轉發的費用亦能精準統計 | 提供 44 種廠商預設範本，或任何 OpenAI / Anthropic 相容的自訂 Base URL |

| Agent 轉發過的每一筆通訊紀錄 | 完整封包檢視，而非僅僅是數字統計 |
| :---: | :---: |
| [![LLM 記錄](https://cdn.casbin.org/img/casbin-gateway-llm-records.png)](https://cdn.casbin.org/img/casbin-gateway-llm-records.png) | [![單條記錄](https://cdn.casbin.org/img/casbin-gateway-llm-record.png)](https://cdn.casbin.org/img/casbin-gateway-llm-record.png) |
| 按模型統計請求次數、Token 用量、快取命中率與成本消耗 | 完整檢視 System Prompt、每條對話訊息以及提供給模型的工具 Schema |

---

## 本 Fork 專案之資安強化項目 (Security Hardening)

為使 Casbin Gateway 能安全部署於現代生產環境（例如位於 Nginx、Caddy、Cloudflare Tunnel 等反向代理背後），本 Fork 版本額外實作並驗證了以下安全性防護：

1. **阻斷反向代理下的管理員自動登入提權 (P0)**：
   修補原設計在單機開發環境依賴 `127.0.0.1` 自動簽發最高權限 Admin Session 的機制。在反向代理架構下強制要求顯式帳號密碼登入，防範外部訪客直接取得系統控制權。
2. **強制管理員預設密碼輪替機制 (P1)**：
   初次部署若系統偵測到仍為初始預設密碼（`admin:123`），系統將限制 API 存取權限並強制重導向至變更密碼頁面，密碼變更完成前全面阻斷受保護的管理端點。
3. **強化 LLM Relay 中繼授權防護 (P1)**：
   修訂中繼轉發驗證邏輯，不再單純信任連線來源為本機 IP，全面落實 Relay Token 簽章驗證，防止付費 LLM 金鑰被未經授權盜刷。
4. **本機目錄瀏覽防路徑遍歷攻擊 (P2)**：
   強化路徑輸入標準化（去除多餘引號、解析絕對路徑 Abs、解析符號連結 EvalSymlinks），嚴格阻斷敏感系統核心目錄（如 `/etc`、`/sys`、`/proc`、`/root` 等）的列舉探測與目錄遍歷攻擊。
5. **SSRF 多 IP 完整解析過濾與 DNS Rebinding 防護 (P2)**：
   出站代理測試全面檢查 DNS 解析的所有 A/AAAA 紀錄及未指定位址（`0.0.0.0`/`::`），並採用固定 IP Dial 連線，杜絕內部私網探測與 DNS 重新綁定攻擊。

---

## 常用指令

- **停止服務**：`casbin-gateway stop`
- **啟動服務**：`casbin-gateway start`
- **查看狀態**：`casbin-gateway status`
- **前景除錯執行**：`casbin-gateway`（使用 `Ctrl+C` 結束）
- **完整解除安裝**：`casbin-gateway uninstall`（自動還原所有 Agent 攔截掛鉤 Hook、開機自啟動設定與桌面捷徑）

---

## 專案架構

Casbin Gateway 主要由兩部分組成：

| 名稱 | 說明 | 開發語言 | 原始碼連結 |
|---|---|---|---|
| **前端** | Web 前端使用者介面 | TypeScript + React + Vite + shadcn/ui | [前端原始碼目錄](https://github.com/tonnychiulab/casbin-gateway-tw/tree/master/web) |
| **後端** | RESTful API 後端服務 | Golang + Beego + XORM + SQLite | [後端原始碼目錄](https://github.com/tonnychiulab/casbin-gateway-tw) |

---

## 開源授權條款 (License)

本專案依循 **[Apache License 2.0](https://github.com/tonnychiulab/casbin-gateway-tw/blob/master/LICENSE)** 授權條款釋出。
原始版權由 **The Casbin Authors** 及 **The Apache Software Foundation** 所有。
本 Fork 專案遵循開源規範，旨在向社群分享安全性加固實踐與學習經驗。
