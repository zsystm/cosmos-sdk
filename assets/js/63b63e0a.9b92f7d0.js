"use strict";(self.webpackChunkcosmos_sdk_docs=self.webpackChunkcosmos_sdk_docs||[]).push([[2383],{3905:(e,t,n)=>{n.d(t,{Zo:()=>m,kt:()=>d});var r=n(7294);function a(e,t,n){return t in e?Object.defineProperty(e,t,{value:n,enumerable:!0,configurable:!0,writable:!0}):e[t]=n,e}function i(e,t){var n=Object.keys(e);if(Object.getOwnPropertySymbols){var r=Object.getOwnPropertySymbols(e);t&&(r=r.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),n.push.apply(n,r)}return n}function o(e){for(var t=1;t<arguments.length;t++){var n=null!=arguments[t]?arguments[t]:{};t%2?i(Object(n),!0).forEach((function(t){a(e,t,n[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(n)):i(Object(n)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(n,t))}))}return e}function s(e,t){if(null==e)return{};var n,r,a=function(e,t){if(null==e)return{};var n,r,a={},i=Object.keys(e);for(r=0;r<i.length;r++)n=i[r],t.indexOf(n)>=0||(a[n]=e[n]);return a}(e,t);if(Object.getOwnPropertySymbols){var i=Object.getOwnPropertySymbols(e);for(r=0;r<i.length;r++)n=i[r],t.indexOf(n)>=0||Object.prototype.propertyIsEnumerable.call(e,n)&&(a[n]=e[n])}return a}var l=r.createContext({}),c=function(e){var t=r.useContext(l),n=t;return e&&(n="function"==typeof e?e(t):o(o({},t),e)),n},m=function(e){var t=c(e.components);return r.createElement(l.Provider,{value:t},e.children)},p={inlineCode:"code",wrapper:function(e){var t=e.children;return r.createElement(r.Fragment,{},t)}},u=r.forwardRef((function(e,t){var n=e.components,a=e.mdxType,i=e.originalType,l=e.parentName,m=s(e,["components","mdxType","originalType","parentName"]),u=c(n),d=a,h=u["".concat(l,".").concat(d)]||u[d]||p[d]||i;return n?r.createElement(h,o(o({ref:t},m),{},{components:n})):r.createElement(h,o({ref:t},m))}));function d(e,t){var n=arguments,a=t&&t.mdxType;if("string"==typeof e||a){var i=n.length,o=new Array(i);o[0]=u;var s={};for(var l in t)hasOwnProperty.call(t,l)&&(s[l]=t[l]);s.originalType=e,s.mdxType="string"==typeof e?e:a,o[1]=s;for(var c=2;c<i;c++)o[c]=n[c];return r.createElement.apply(null,o)}return r.createElement.apply(null,n)}u.displayName="MDXCreateElement"},982:(e,t,n)=>{n.r(t),n.d(t,{assets:()=>l,contentTitle:()=>o,default:()=>p,frontMatter:()=>i,metadata:()=>s,toc:()=>c});var r=n(7462),a=(n(7294),n(3905));const i={},o="ADR 013: Observability",s={unversionedId:"architecture/adr-013-metrics",id:"version-v0.47/architecture/adr-013-metrics",title:"ADR 013: Observability",description:"Changelog",source:"@site/versioned_docs/version-v0.47/architecture/adr-013-metrics.md",sourceDirName:"architecture",slug:"/architecture/adr-013-metrics",permalink:"/v0.47/architecture/adr-013-metrics",draft:!1,tags:[],version:"v0.47",frontMatter:{},sidebar:"tutorialSidebar",previous:{title:"ADR 012: State Accessors",permalink:"/v0.47/architecture/adr-012-state-accessors"},next:{title:"ADR 14: Proportional Slashing",permalink:"/v0.47/architecture/adr-014-proportional-slashing"}},l={},c=[{value:"Changelog",id:"changelog",level:2},{value:"Status",id:"status",level:2},{value:"Context",id:"context",level:2},{value:"Decision",id:"decision",level:2},{value:"Consequences",id:"consequences",level:2},{value:"Positive",id:"positive",level:3},{value:"Negative",id:"negative",level:3},{value:"Neutral",id:"neutral",level:3},{value:"References",id:"references",level:2}],m={toc:c};function p(e){let{components:t,...n}=e;return(0,a.kt)("wrapper",(0,r.Z)({},m,n,{components:t,mdxType:"MDXLayout"}),(0,a.kt)("h1",{id:"adr-013-observability"},"ADR 013: Observability"),(0,a.kt)("h2",{id:"changelog"},"Changelog"),(0,a.kt)("ul",null,(0,a.kt)("li",{parentName:"ul"},"20-01-2020: Initial Draft")),(0,a.kt)("h2",{id:"status"},"Status"),(0,a.kt)("p",null,"Proposed"),(0,a.kt)("h2",{id:"context"},"Context"),(0,a.kt)("p",null,"Telemetry is paramount into debugging and understanding what the application is doing and how it is\nperforming. We aim to expose metrics from modules and other core parts of the Cosmos SDK."),(0,a.kt)("p",null,"In addition, we should aim to support multiple configurable sinks that an operator may choose from.\nBy default, when telemetry is enabled, the application should track and expose metrics that are\nstored in-memory. The operator may choose to enable additional sinks, where we support only\n",(0,a.kt)("a",{parentName:"p",href:"https://prometheus.io/"},"Prometheus")," for now, as it's battle-tested, simple to setup, open source,\nand is rich with ecosystem tooling."),(0,a.kt)("p",null,"We must also aim to integrate metrics into the Cosmos SDK in the most seamless way possible such that\nmetrics may be added or removed at will and without much friction. To do this, we will use the\n",(0,a.kt)("a",{parentName:"p",href:"https://github.com/armon/go-metrics"},"go-metrics")," library."),(0,a.kt)("p",null,"Finally, operators may enable telemetry along with specific configuration options. If enabled, metrics\nwill be exposed via ",(0,a.kt)("inlineCode",{parentName:"p"},"/metrics?format={text|prometheus}")," via the API server."),(0,a.kt)("h2",{id:"decision"},"Decision"),(0,a.kt)("p",null,"We will add an additional configuration block to ",(0,a.kt)("inlineCode",{parentName:"p"},"app.toml")," that defines telemetry settings:"),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-toml"},"###############################################################################\n###                         Telemetry Configuration                         ###\n###############################################################################\n\n[telemetry]\n\n# Prefixed with keys to separate services\nservice-name = {{ .Telemetry.ServiceName }}\n\n# Enabled enables the application telemetry functionality. When enabled,\n# an in-memory sink is also enabled by default. Operators may also enabled\n# other sinks such as Prometheus.\nenabled = {{ .Telemetry.Enabled }}\n\n# Enable prefixing gauge values with hostname\nenable-hostname = {{ .Telemetry.EnableHostname }}\n\n# Enable adding hostname to labels\nenable-hostname-label = {{ .Telemetry.EnableHostnameLabel }}\n\n# Enable adding service to labels\nenable-service-label = {{ .Telemetry.EnableServiceLabel }}\n\n# PrometheusRetentionTime, when positive, enables a Prometheus metrics sink.\nprometheus-retention-time = {{ .Telemetry.PrometheusRetentionTime }}\n")),(0,a.kt)("p",null,"The given configuration allows for two sinks -- in-memory and Prometheus. We create a ",(0,a.kt)("inlineCode",{parentName:"p"},"Metrics"),"\ntype that performs all the bootstrapping for the operator, so capturing metrics becomes seamless."),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-go"},'// Metrics defines a wrapper around application telemetry functionality. It allows\n// metrics to be gathered at any point in time. When creating a Metrics object,\n// internally, a global metrics is registered with a set of sinks as configured\n// by the operator. In addition to the sinks, when a process gets a SIGUSR1, a\n// dump of formatted recent metrics will be sent to STDERR.\ntype Metrics struct {\n  memSink           *metrics.InmemSink\n  prometheusEnabled bool\n}\n\n// Gather collects all registered metrics and returns a GatherResponse where the\n// metrics are encoded depending on the type. Metrics are either encoded via\n// Prometheus or JSON if in-memory.\nfunc (m *Metrics) Gather(format string) (GatherResponse, error) {\n  switch format {\n  case FormatPrometheus:\n    return m.gatherPrometheus()\n\n  case FormatText:\n    return m.gatherGeneric()\n\n  case FormatDefault:\n    return m.gatherGeneric()\n\n  default:\n    return GatherResponse{}, fmt.Errorf("unsupported metrics format: %s", format)\n  }\n}\n')),(0,a.kt)("p",null,"In addition, ",(0,a.kt)("inlineCode",{parentName:"p"},"Metrics")," allows us to gather the current set of metrics at any given point in time. An\noperator may also choose to send a signal, SIGUSR1, to dump and print formatted metrics to STDERR."),(0,a.kt)("p",null,"During an application's bootstrapping and construction phase, if ",(0,a.kt)("inlineCode",{parentName:"p"},"Telemetry.Enabled")," is ",(0,a.kt)("inlineCode",{parentName:"p"},"true"),", the\nAPI server will create an instance of a reference to ",(0,a.kt)("inlineCode",{parentName:"p"},"Metrics")," object and will register a metrics\nhandler accordingly."),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-go"},'func (s *Server) Start(cfg config.Config) error {\n  // ...\n\n  if cfg.Telemetry.Enabled {\n    m, err := telemetry.New(cfg.Telemetry)\n    if err != nil {\n      return err\n    }\n\n    s.metrics = m\n    s.registerMetrics()\n  }\n\n  // ...\n}\n\nfunc (s *Server) registerMetrics() {\n  metricsHandler := func(w http.ResponseWriter, r *http.Request) {\n    format := strings.TrimSpace(r.FormValue("format"))\n\n    gr, err := s.metrics.Gather(format)\n    if err != nil {\n      rest.WriteErrorResponse(w, http.StatusBadRequest, fmt.Sprintf("failed to gather metrics: %s", err))\n      return\n    }\n\n    w.Header().Set("Content-Type", gr.ContentType)\n    _, _ = w.Write(gr.Metrics)\n  }\n\n  s.Router.HandleFunc("/metrics", metricsHandler).Methods("GET")\n}\n')),(0,a.kt)("p",null,"Application developers may track counters, gauges, summaries, and key/value metrics. There is no\nadditional lifting required by modules to leverage profiling metrics. To do so, it's as simple as:"),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-go"},'func (k BaseKeeper) MintCoins(ctx sdk.Context, moduleName string, amt sdk.Coins) error {\n  defer metrics.MeasureSince(time.Now(), "MintCoins")\n  // ...\n}\n')),(0,a.kt)("h2",{id:"consequences"},"Consequences"),(0,a.kt)("h3",{id:"positive"},"Positive"),(0,a.kt)("ul",null,(0,a.kt)("li",{parentName:"ul"},"Exposure into the performance and behavior of an application")),(0,a.kt)("h3",{id:"negative"},"Negative"),(0,a.kt)("h3",{id:"neutral"},"Neutral"),(0,a.kt)("h2",{id:"references"},"References"))}p.isMDXComponent=!0}}]);