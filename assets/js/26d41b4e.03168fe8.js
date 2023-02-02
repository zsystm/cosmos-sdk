"use strict";(self.webpackChunkcosmos_sdk_docs=self.webpackChunkcosmos_sdk_docs||[]).push([[9730],{3905:(e,r,n)=>{n.d(r,{Zo:()=>s,kt:()=>m});var t=n(7294);function a(e,r,n){return r in e?Object.defineProperty(e,r,{value:n,enumerable:!0,configurable:!0,writable:!0}):e[r]=n,e}function o(e,r){var n=Object.keys(e);if(Object.getOwnPropertySymbols){var t=Object.getOwnPropertySymbols(e);r&&(t=t.filter((function(r){return Object.getOwnPropertyDescriptor(e,r).enumerable}))),n.push.apply(n,t)}return n}function i(e){for(var r=1;r<arguments.length;r++){var n=null!=arguments[r]?arguments[r]:{};r%2?o(Object(n),!0).forEach((function(r){a(e,r,n[r])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(n)):o(Object(n)).forEach((function(r){Object.defineProperty(e,r,Object.getOwnPropertyDescriptor(n,r))}))}return e}function l(e,r){if(null==e)return{};var n,t,a=function(e,r){if(null==e)return{};var n,t,a={},o=Object.keys(e);for(t=0;t<o.length;t++)n=o[t],r.indexOf(n)>=0||(a[n]=e[n]);return a}(e,r);if(Object.getOwnPropertySymbols){var o=Object.getOwnPropertySymbols(e);for(t=0;t<o.length;t++)n=o[t],r.indexOf(n)>=0||Object.prototype.propertyIsEnumerable.call(e,n)&&(a[n]=e[n])}return a}var c=t.createContext({}),p=function(e){var r=t.useContext(c),n=r;return e&&(n="function"==typeof e?e(r):i(i({},r),e)),n},s=function(e){var r=p(e.components);return t.createElement(c.Provider,{value:r},e.children)},d={inlineCode:"code",wrapper:function(e){var r=e.children;return t.createElement(t.Fragment,{},r)}},u=t.forwardRef((function(e,r){var n=e.components,a=e.mdxType,o=e.originalType,c=e.parentName,s=l(e,["components","mdxType","originalType","parentName"]),u=p(n),m=a,f=u["".concat(c,".").concat(m)]||u[m]||d[m]||o;return n?t.createElement(f,i(i({ref:r},s),{},{components:n})):t.createElement(f,i({ref:r},s))}));function m(e,r){var n=arguments,a=r&&r.mdxType;if("string"==typeof e||a){var o=n.length,i=new Array(o);i[0]=u;var l={};for(var c in r)hasOwnProperty.call(r,c)&&(l[c]=r[c]);l.originalType=e,l.mdxType="string"==typeof e?e:a,i[1]=l;for(var p=2;p<o;p++)i[p]=n[p];return t.createElement.apply(null,i)}return t.createElement.apply(null,n)}u.displayName="MDXCreateElement"},6467:(e,r,n)=>{n.r(r),n.d(r,{assets:()=>c,contentTitle:()=>i,default:()=>d,frontMatter:()=>o,metadata:()=>l,toc:()=>p});var t=n(7462),a=(n(7294),n(3905));const o={sidebar_position:1},i="RunTx recovery middleware",l={unversionedId:"core/runtx_middleware",id:"core/runtx_middleware",title:"RunTx recovery middleware",description:"BaseApp.runTx() function handles Go panics that might occur during transactions execution, for example, keeper has faced an invalid state and paniced.",source:"@site/docs/core/11-runtx_middleware.md",sourceDirName:"core",slug:"/core/runtx_middleware",permalink:"/main/core/runtx_middleware",draft:!1,tags:[],version:"current",sidebarPosition:1,frontMatter:{sidebar_position:1},sidebar:"tutorialSidebar",previous:{title:"Object-Capability Model",permalink:"/main/core/ocap"},next:{title:"Cosmos Blockchain Simulator",permalink:"/main/core/simulation"}},c={},p=[{value:"Interface",id:"interface",level:2},{value:"Custom RecoveryHandler register",id:"custom-recoveryhandler-register",level:2},{value:"Example",id:"example",level:2}],s={toc:p};function d(e){let{components:r,...n}=e;return(0,a.kt)("wrapper",(0,t.Z)({},s,n,{components:r,mdxType:"MDXLayout"}),(0,a.kt)("h1",{id:"runtx-recovery-middleware"},"RunTx recovery middleware"),(0,a.kt)("p",null,(0,a.kt)("inlineCode",{parentName:"p"},"BaseApp.runTx()")," function handles Go panics that might occur during transactions execution, for example, keeper has faced an invalid state and paniced.\nDepending on the panic type different handler is used, for instance the default one prints an error log message.\nRecovery middleware is used to add custom panic recovery for Cosmos SDK application developers."),(0,a.kt)("p",null,"More context can found in the corresponding ",(0,a.kt)("a",{parentName:"p",href:"/main/architecture/adr-022-custom-panic-handling"},"ADR-022")," and the implementation in ",(0,a.kt)("a",{parentName:"p",href:"https://github.com/cosmos/cosmos-sdk/blob/v0.47.0-rc1/baseapp/recovery.go"},"recovery.go"),"."),(0,a.kt)("h2",{id:"interface"},"Interface"),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-go",metastring:"reference",reference:!0},"https://github.com/cosmos/cosmos-sdk/blob/v0.47.0-rc1/baseapp/recovery.go#L11-L14\n")),(0,a.kt)("p",null,(0,a.kt)("inlineCode",{parentName:"p"},"recoveryObj")," is a return value for ",(0,a.kt)("inlineCode",{parentName:"p"},"recover()")," function from the ",(0,a.kt)("inlineCode",{parentName:"p"},"buildin")," Go package."),(0,a.kt)("p",null,(0,a.kt)("strong",{parentName:"p"},"Contract:")),(0,a.kt)("ul",null,(0,a.kt)("li",{parentName:"ul"},"RecoveryHandler returns ",(0,a.kt)("inlineCode",{parentName:"li"},"nil")," if ",(0,a.kt)("inlineCode",{parentName:"li"},"recoveryObj")," wasn't handled and should be passed to the next recovery middleware;"),(0,a.kt)("li",{parentName:"ul"},"RecoveryHandler returns a non-nil ",(0,a.kt)("inlineCode",{parentName:"li"},"error")," if ",(0,a.kt)("inlineCode",{parentName:"li"},"recoveryObj")," was handled;")),(0,a.kt)("h2",{id:"custom-recoveryhandler-register"},"Custom RecoveryHandler register"),(0,a.kt)("p",null,(0,a.kt)("inlineCode",{parentName:"p"},"BaseApp.AddRunTxRecoveryHandler(handlers ...RecoveryHandler)")),(0,a.kt)("p",null,"BaseApp method adds recovery middleware to the default recovery chain."),(0,a.kt)("h2",{id:"example"},"Example"),(0,a.kt)("p",null,'Lets assume we want to emit the "Consensus failure" chain state if some particular error occurred.'),(0,a.kt)("p",null,"We have a module keeper that panics:"),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-go"},'func (k FooKeeper) Do(obj interface{}) {\n    if obj == nil {\n        // that shouldn\'t happen, we need to crash the app\n        err := sdkErrors.Wrap(fooTypes.InternalError, "obj is nil")\n        panic(err)\n    }\n}\n')),(0,a.kt)("p",null,"By default that panic would be recovered and an error message will be printed to log. To override that behaviour we should register a custom RecoveryHandler:"),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-go"},'// Cosmos SDK application constructor\ncustomHandler := func(recoveryObj interface{}) error {\n    err, ok := recoveryObj.(error)\n    if !ok {\n        return nil\n    }\n\n    if fooTypes.InternalError.Is(err) {\n        panic(fmt.Errorf("FooKeeper did panic with error: %w", err))\n    }\n\n    return nil\n}\n\nbaseApp := baseapp.NewBaseApp(...)\nbaseApp.AddRunTxRecoveryHandler(customHandler)\n')))}d.isMDXComponent=!0}}]);