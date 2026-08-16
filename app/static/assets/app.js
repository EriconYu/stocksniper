const state = {
  meta: null,
  cache: new Map(),
  route: "dashboard",
};

const content = document.querySelector("#content");
const navigation = document.querySelector("#navigation");
const pageTitle = document.querySelector("#pageTitle");
const sidebar = document.querySelector("#sidebar");
const contactModal = document.querySelector("#contactModal");
const toast = document.querySelector("#toast");

const pageNames = {
  dashboard: "研究总览",
  market: "市场环境",
  advice: "今日观察",
  strategies: "策略中心",
  ranking: "综合评分",
  backtest: "回测验证",
  boards: "板块分析",
  signals: "信号追踪",
  data: "数据管理",
  logs: "执行日志",
};

function escapeHTML(value) {
  return String(value ?? "")
    .replaceAll("&", "&amp;")
    .replaceAll("<", "&lt;")
    .replaceAll(">", "&gt;")
    .replaceAll('"', "&quot;")
    .replaceAll("'", "&#039;");
}

function numberClass(value) {
  if (Number(value) > 0) return "text-up";
  if (Number(value) < 0) return "text-down";
  return "";
}

function signed(value, suffix = "%") {
  const numeric = Number(value);
  return `${numeric > 0 ? "+" : ""}${numeric.toFixed(2)}${suffix}`;
}

async function api(path, options = {}) {
  const response = await fetch(`/api/v1/${path}`, {
    headers: { "Content-Type": "application/json" },
    ...options,
  });
  if (!response.ok) throw new Error(`请求失败 (${response.status})`);
  return response.json();
}

async function cached(path) {
  if (!state.cache.has(path)) state.cache.set(path, api(path));
  return state.cache.get(path);
}

function notice() {
  return `
    <div class="notice-band">
      <span>当前为开源展示版，页面数据仅用于产品演示，不构成投资建议。</span>
      <a href="https://facai.stackbang.com" target="_blank" rel="noreferrer">进入正式版 →</a>
    </div>`;
}

function sectionHead(title, description = "", action = "") {
  return `
    <div class="section-head">
      <div><h2>${escapeHTML(title)}</h2>${description ? `<p>${escapeHTML(description)}</p>` : ""}</div>
      ${action}
    </div>`;
}

async function renderDashboard() {
  const data = await cached("overview");
  const metrics = data.metrics.map((item) => `
    <article class="metric-card">
      <span>${escapeHTML(item.label)}</span>
      <strong>${escapeHTML(item.value)}</strong> <em>${escapeHTML(item.unit)}</em>
      <small>${escapeHTML(item.trend)}</small>
    </article>`).join("");
  const rows = data.watchlist.map((item) => `
    <tr>
      <td><strong>${escapeHTML(item.symbol)}</strong></td><td>${escapeHTML(item.name)}</td>
      <td><span class="score">${item.score}</span></td><td><span class="tag">${escapeHTML(item.signal)}</span></td>
      <td class="${numberClass(item.change)}">${signed(item.change)}</td>
    </tr>`).join("");

  content.innerHTML = `
    ${notice()}
    ${sectionHead("核心快照", "用于展示信息密度与工作流，非实时行情")}
    <div class="metric-grid">${metrics}</div>
    ${sectionHead("市场脉搏", "近 20 个演示观察点")}
    <div class="dashboard-grid">
      <section class="panel">
        <div class="panel-head"><h3>市场温度轨迹</h3><span>${escapeHTML(data.as_of)}</span></div>
        <div class="chart-wrap"><canvas id="marketChart" aria-label="市场温度演示曲线"></canvas></div>
      </section>
      <section class="panel">
        <div class="panel-head"><h3>全市场宽度</h3><span>演示统计</span></div>
        <div class="breadth">
          <div><span>上涨</span><strong class="text-up">${data.breadth.up}</strong></div>
          <div><span>平盘</span><strong>${data.breadth.flat}</strong></div>
          <div><span>下跌</span><strong class="text-down">${data.breadth.down}</strong></div>
        </div>
      </section>
    </div>
    ${sectionHead("重点观察", "综合评分与信号均为演示内容", `<button class="button button-secondary" data-action="export">导出列表</button>`)}
    <div class="data-scroll"><table class="data-table">
      <thead><tr><th>代码</th><th>名称</th><th>评分</th><th>状态</th><th>涨跌幅</th></tr></thead>
      <tbody>${rows}</tbody>
    </table></div>`;
  requestAnimationFrame(() => drawLineChart(document.querySelector("#marketChart"), data.market_curve));
}

async function renderMarket() {
  const data = await cached("market");
  const indices = data.indices.map((item) => `
    <div><span>${escapeHTML(item.name)}</span><strong>${escapeHTML(item.value)}</strong><small class="${numberClass(item.change)}">${signed(item.change)}</small></div>`).join("");
  const dimensions = data.dimensions.map((item) => `
    <div class="dimension-row">
      <strong>${escapeHTML(item.label)}</strong><div class="progress"><i style="width:${item.score}%"></i></div>
      <span>${item.score}</span><span>${escapeHTML(item.status)}</span>
    </div>`).join("");
  content.innerHTML = `
    ${notice()}
    ${sectionHead("指数概览", "固定演示快照")}
    <div class="market-strip">${indices}</div>
    ${sectionHead("环境判断", "多维状态聚合展示")}
    <div class="dashboard-grid">
      <section class="panel">
        <span class="section-kicker">当前状态</span>
        <h2>${escapeHTML(data.regime)}</h2>
        <p style="color:var(--muted);font-size:11px">模型置信度 ${(data.confidence * 100).toFixed(0)}%，仅用于界面展示。</p>
      </section>
      <section class="panel"><div class="dimension-list">${dimensions}</div></section>
    </div>`;
}

async function renderAdvice() {
  const [overview, signals] = await Promise.all([cached("overview"), cached("signals")]);
  const focus = overview.watchlist.slice(0, 3).map((item) => `
    <tr><td>${escapeHTML(item.symbol)}</td><td>${escapeHTML(item.name)}</td><td><span class="score">${item.score}</span></td><td>${escapeHTML(item.signal)}</td></tr>`).join("");
  content.innerHTML = `
    ${notice()}
    ${sectionHead("今日观察清单", "展示版不会生成真实买卖建议")}
    <div class="data-scroll"><table class="data-table"><thead><tr><th>代码</th><th>名称</th><th>评分</th><th>观察原因</th></tr></thead><tbody>${focus}</tbody></table></div>
    ${sectionHead("盘后摘要", "基于演示信号的界面样例")}
    <section class="panel">
      <div class="panel-head"><h3>结构性机会仍需等待确认</h3><span>${signals.length} 条演示信号</span></div>
      <p style="color:var(--muted);font-size:11px;line-height:1.8">市场环境显示为震荡偏强，量价结构有所改善。正式版会结合完整数据质量、策略准入和风险过滤生成研究结论。</p>
      <button class="button button-primary" data-action="trade-assistant">查看完整建议</button>
    </section>`;
}

async function renderStrategies() {
  const strategies = await cached("strategies");
  const cards = strategies.map((item) => `
    <article class="strategy-card">
      <header><h3>${escapeHTML(item.name)}</h3><span class="tag ${escapeHTML(item.status)}">${item.status === "active" ? "运行中" : "观察中"}</span></header>
      <p>${escapeHTML(item.description)}</p>
      <div class="strategy-meta"><span>${escapeHTML(item.last_run)}</span><span>${item.candidates} 个演示候选</span></div>
      <div class="strategy-actions">
        <button class="button button-secondary" data-action="strategy-run">查看规则</button>
        <button class="button button-primary" data-action="strategy-run">运行策略</button>
      </div>
    </article>`).join("");
  content.innerHTML = `${notice()}${sectionHead("策略目录", "公开版展示研究方向，不提供规则阈值和执行代码")}<div class="strategy-grid">${cards}</div>`;
}

async function renderRanking() {
  const rankings = await cached("rankings");
  const rows = rankings.map((item) => `
    <tr><td><strong>${item.rank}</strong></td><td>${escapeHTML(item.symbol)}</td><td>${escapeHTML(item.name)}</td>
    <td><span class="score">${item.score}</span></td><td>${item.trend}</td><td>${item.volume}</td><td>${item.risk}</td></tr>`).join("");
  content.innerHTML = `
    ${notice()}${sectionHead("综合评分榜", "因子权重与计算过程未包含在社区版", `<button class="button button-secondary" data-action="export">导出排名</button>`)}
    <div class="data-scroll"><table class="data-table"><thead><tr><th>排名</th><th>代码</th><th>名称</th><th>综合</th><th>趋势</th><th>量能</th><th>风险</th></tr></thead><tbody>${rows}</tbody></table></div>`;
}

function renderWorkbench(type) {
  const configs = {
    backtest: { title: "回测参数", button: "开始回测", feature: "backtest", heading: "回测引擎未包含", text: "完整版本提供交易日历、成交约束、费用模型和策略绩效分析。" },
    data: { title: "数据任务", button: "开始同步", feature: "data-sync", heading: "数据链路未包含", text: "完整版本提供多源行情、增量更新、质量检查和任务追踪。" },
  };
  const config = configs[type];
  const extra = type === "backtest"
    ? `<label>策略</label><select><option>趋势观察</option><option>量价观察</option></select><label>日期范围</label><input value="最近 12 个月" readonly />`
    : `<label>任务类型</label><select><option>日线增量同步</option><option>股票列表同步</option><option>完整性检查</option></select><label>数据范围</label><input value="全部演示标的" readonly />`;
  content.innerHTML = `
    ${notice()}${sectionHead(config.title, "参数控件保留用于展示完整工作流")}
    <section class="feature-workbench">
      <form class="feature-form" data-feature-form data-feature="${config.feature}">
        <label>股票代码</label><input name="symbol" value="600000.SH" maxlength="32" />${extra}
        <button class="button button-primary" type="submit">${config.button}</button>
      </form>
      <div class="feature-preview"><div class="locked-state">
        <div class="locked-mark">PRO</div><h3>${config.heading}</h3><p>${config.text}</p>
        <div class="locked-actions"><a class="button button-primary" href="https://facai.stackbang.com" target="_blank" rel="noreferrer">访问正式版</a><button class="button button-secondary" data-contact>联系作者</button></div>
      </div></div>
    </section>`;
}

async function renderBoards() {
  const boards = await cached("boards");
  const cards = boards.map((item) => `
    <article class="board-card"><h3>${escapeHTML(item.name)}</h3><strong>${item.heat}</strong><footer><span>热度</span><span class="${numberClass(item.change)}">${signed(item.change)}</span></footer></article>`).join("");
  content.innerHTML = `${notice()}${sectionHead("板块热度", "行业与概念聚合结果为演示数据")}<div class="board-grid">${cards}</div>`;
}

async function renderSignals() {
  const signals = await cached("signals");
  const rows = signals.map((item) => `
    <div class="signal-row"><time>${escapeHTML(item.time)}</time><span class="tag ${escapeHTML(item.level)}">${item.level === "risk" ? "风险" : item.level === "focus" ? "重点" : "观察"}</span><strong>${escapeHTML(item.title)}</strong><span>${escapeHTML(item.target)}</span><span>${escapeHTML(item.source)}</span></div>`).join("");
  content.innerHTML = `${notice()}${sectionHead("信号时间线", "演示信号不会触发通知或交易")}<div class="signal-list">${rows}</div>`;
}

async function renderLogs() {
  const activity = await cached("activity");
  const rows = activity.map((item) => `<div class="log-row"><time>${escapeHTML(item.time)}</time><span class="tag">${escapeHTML(item.module)}</span><span>${escapeHTML(item.message)}</span></div>`).join("");
  content.innerHTML = `${notice()}${sectionHead("最近活动", "固定演示日志，不连接任务调度器")}<div class="log-list">${rows}</div>`;
}

const renderers = {
  dashboard: renderDashboard,
  market: renderMarket,
  advice: renderAdvice,
  strategies: renderStrategies,
  ranking: renderRanking,
  backtest: () => renderWorkbench("backtest"),
  boards: renderBoards,
  signals: renderSignals,
  data: () => renderWorkbench("data"),
  logs: renderLogs,
};

async function renderRoute(route) {
  const normalized = renderers[route] ? route : "dashboard";
  state.route = normalized;
  pageTitle.textContent = pageNames[normalized];
  document.title = `${pageNames[normalized]} · StockSniper`;
  document.querySelectorAll(".nav-item").forEach((item) => item.classList.toggle("active", item.dataset.route === normalized));
  content.innerHTML = `<div class="loading-state">正在加载演示界面...</div>`;
  try {
    await renderers[normalized]();
  } catch (error) {
    content.innerHTML = `<div class="empty-state">${escapeHTML(error.message)}，请稍后重试。</div>`;
  }
  sidebar.classList.remove("open");
}

function drawLineChart(canvas, values) {
  if (!canvas || !values.length) return;
  const rect = canvas.getBoundingClientRect();
  const ratio = window.devicePixelRatio || 1;
  canvas.width = Math.max(1, Math.round(rect.width * ratio));
  canvas.height = Math.max(1, Math.round(rect.height * ratio));
  const ctx = canvas.getContext("2d");
  ctx.scale(ratio, ratio);
  const width = rect.width;
  const height = rect.height;
  const padding = { top: 14, right: 12, bottom: 24, left: 32 };
  const min = Math.min(...values) - 4;
  const max = Math.max(...values) + 4;
  const x = (index) => padding.left + index * (width - padding.left - padding.right) / (values.length - 1);
  const y = (value) => padding.top + (max - value) * (height - padding.top - padding.bottom) / (max - min);

  ctx.strokeStyle = "#e2e8e5";
  ctx.lineWidth = 1;
  ctx.fillStyle = "#77817d";
  ctx.font = "10px system-ui";
  for (let step = 0; step < 4; step += 1) {
    const value = min + (max - min) * step / 3;
    const lineY = y(value);
    ctx.beginPath(); ctx.moveTo(padding.left, lineY); ctx.lineTo(width - padding.right, lineY); ctx.stroke();
    ctx.fillText(value.toFixed(0), 4, lineY + 3);
  }
  ctx.strokeStyle = "#167a5b";
  ctx.lineWidth = 2;
  ctx.beginPath();
  values.forEach((value, index) => index === 0 ? ctx.moveTo(x(index), y(value)) : ctx.lineTo(x(index), y(value)));
  ctx.stroke();
  values.forEach((value, index) => {
    if (index % 3 !== 0 && index !== values.length - 1) return;
    ctx.fillStyle = "#ffffff"; ctx.strokeStyle = "#167a5b"; ctx.lineWidth = 2;
    ctx.beginPath(); ctx.arc(x(index), y(value), 3, 0, Math.PI * 2); ctx.fill(); ctx.stroke();
  });
}

function openContact() {
  contactModal.hidden = false;
  document.body.style.overflow = "hidden";
}

function closeContact() {
  contactModal.hidden = true;
  document.body.style.overflow = "";
}

function showToast(message) {
  toast.textContent = message;
  toast.hidden = false;
  window.clearTimeout(showToast.timer);
  showToast.timer = window.setTimeout(() => { toast.hidden = true; }, 3500);
}

async function runShowcaseAction(feature, symbol = null) {
  const result = await api("showcase-action", { method: "POST", body: JSON.stringify({ feature, symbol: symbol || null }) });
  showToast(result.message);
  window.setTimeout(openContact, 500);
}

document.addEventListener("click", (event) => {
  const nav = event.target.closest("[data-route]");
  if (nav) {
    window.location.hash = nav.dataset.route;
    return;
  }
  if (event.target.closest("[data-contact]")) openContact();
  const action = event.target.closest("[data-action]");
  if (action) runShowcaseAction(action.dataset.action).catch((error) => showToast(error.message));
});

document.addEventListener("submit", (event) => {
  const form = event.target.closest("[data-feature-form]");
  if (!form) return;
  event.preventDefault();
  const values = new FormData(form);
  runShowcaseAction(form.dataset.feature, values.get("symbol")).catch((error) => showToast(error.message));
});

document.querySelector("#menuButton").addEventListener("click", () => sidebar.classList.toggle("open"));
document.querySelector("#closeContact").addEventListener("click", closeContact);
contactModal.addEventListener("click", (event) => { if (event.target === contactModal) closeContact(); });
document.addEventListener("keydown", (event) => { if (event.key === "Escape") closeContact(); });
window.addEventListener("hashchange", () => renderRoute(window.location.hash.slice(1)));
window.addEventListener("resize", () => { if (state.route === "dashboard") renderRoute("dashboard"); });

async function bootstrap() {
  try {
    state.meta = await api("meta");
    navigation.innerHTML = state.meta.navigation.map((item) => `
      <button class="nav-item" data-route="${escapeHTML(item.id)}">${escapeHTML(item.label)}</button>`).join("");
    await renderRoute(window.location.hash.slice(1) || "dashboard");
  } catch (error) {
    content.innerHTML = `<div class="empty-state">${escapeHTML(error.message)}，请确认服务已启动。</div>`;
  }
}

bootstrap();

