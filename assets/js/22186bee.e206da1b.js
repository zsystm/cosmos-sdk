"use strict";(self.webpackChunkcosmos_sdk_docs=self.webpackChunkcosmos_sdk_docs||[]).push([[3874],{3905:(e,n,t)=>{t.d(n,{Zo:()=>s,kt:()=>f});var i=t(7294);function a(e,n,t){return n in e?Object.defineProperty(e,n,{value:t,enumerable:!0,configurable:!0,writable:!0}):e[n]=t,e}function o(e,n){var t=Object.keys(e);if(Object.getOwnPropertySymbols){var i=Object.getOwnPropertySymbols(e);n&&(i=i.filter((function(n){return Object.getOwnPropertyDescriptor(e,n).enumerable}))),t.push.apply(t,i)}return t}function l(e){for(var n=1;n<arguments.length;n++){var t=null!=arguments[n]?arguments[n]:{};n%2?o(Object(t),!0).forEach((function(n){a(e,n,t[n])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(t)):o(Object(t)).forEach((function(n){Object.defineProperty(e,n,Object.getOwnPropertyDescriptor(t,n))}))}return e}function r(e,n){if(null==e)return{};var t,i,a=function(e,n){if(null==e)return{};var t,i,a={},o=Object.keys(e);for(i=0;i<o.length;i++)t=o[i],n.indexOf(t)>=0||(a[t]=e[t]);return a}(e,n);if(Object.getOwnPropertySymbols){var o=Object.getOwnPropertySymbols(e);for(i=0;i<o.length;i++)t=o[i],n.indexOf(t)>=0||Object.prototype.propertyIsEnumerable.call(e,t)&&(a[t]=e[t])}return a}var c=i.createContext({}),p=function(e){var n=i.useContext(c),t=n;return e&&(t="function"==typeof e?e(n):l(l({},n),e)),t},s=function(e){var n=p(e.components);return i.createElement(c.Provider,{value:n},e.children)},d={inlineCode:"code",wrapper:function(e){var n=e.children;return i.createElement(i.Fragment,{},n)}},m=i.forwardRef((function(e,n){var t=e.components,a=e.mdxType,o=e.originalType,c=e.parentName,s=r(e,["components","mdxType","originalType","parentName"]),m=p(t),f=a,u=m["".concat(c,".").concat(f)]||m[f]||d[f]||o;return t?i.createElement(u,l(l({ref:n},s),{},{components:t})):i.createElement(u,l({ref:n},s))}));function f(e,n){var t=arguments,a=n&&n.mdxType;if("string"==typeof e||a){var o=t.length,l=new Array(o);l[0]=m;var r={};for(var c in n)hasOwnProperty.call(n,c)&&(r[c]=n[c]);r.originalType=e,r.mdxType="string"==typeof e?e:a,l[1]=r;for(var p=2;p<o;p++)l[p]=t[p];return i.createElement.apply(null,l)}return i.createElement.apply(null,t)}m.displayName="MDXCreateElement"},3650:(e,n,t)=>{t.r(n),t.d(n,{assets:()=>c,contentTitle:()=>l,default:()=>d,frontMatter:()=>o,metadata:()=>r,toc:()=>p});var i=t(7462),a=(t(7294),t(3905));const o={sidebar_position:1},l="Confix",r={unversionedId:"tooling/confix",id:"tooling/confix",title:"Confix",description:"Confix is a configuration management tool that allows you to manage your configuration via CLI.",source:"@site/docs/tooling/03-confix.md",sourceDirName:"tooling",slug:"/tooling/confix",permalink:"/main/tooling/confix",draft:!1,tags:[],version:"current",sidebarPosition:1,frontMatter:{sidebar_position:1},sidebar:"tutorialSidebar",previous:{title:"Depinject",permalink:"/main/tooling/depinject"},next:{title:"Hubl",permalink:"/main/tooling/hubl"}},c={},p=[{value:"Installation",id:"installation",level:2},{value:"Add Config Command",id:"add-config-command",level:3},{value:"Using Confix Standalone",id:"using-confix-standalone",level:3},{value:"Usage",id:"usage",level:2},{value:"Get",id:"get",level:3},{value:"Set",id:"set",level:3},{value:"Migrate",id:"migrate",level:3},{value:"Diff",id:"diff",level:3},{value:"Maintainer",id:"maintainer",level:3},{value:"Credits",id:"credits",level:2}],s={toc:p};function d(e){let{components:n,...t}=e;return(0,a.kt)("wrapper",(0,i.Z)({},s,t,{components:n,mdxType:"MDXLayout"}),(0,a.kt)("h1",{id:"confix"},"Confix"),(0,a.kt)("p",null,(0,a.kt)("inlineCode",{parentName:"p"},"Confix")," is a configuration management tool that allows you to manage your configuration via CLI."),(0,a.kt)("p",null,"It is based on the ",(0,a.kt)("a",{parentName:"p",href:"https://github.com/tendermint/tendermint/blob/5013bc3f4a6d64dcc2bf02ccc002ebc9881c62e4/docs/rfc/rfc-019-config-version.md"},"Tendermint RFC 019"),"."),(0,a.kt)("h2",{id:"installation"},"Installation"),(0,a.kt)("h3",{id:"add-config-command"},"Add Config Command"),(0,a.kt)("p",null,"To add the confix tool, it's required to add the ",(0,a.kt)("inlineCode",{parentName:"p"},"ConfigCommand")," to your application's root command file (e.g. ",(0,a.kt)("inlineCode",{parentName:"p"},"simd/cmd/root.go"),")."),(0,a.kt)("p",null,"Import the ",(0,a.kt)("inlineCode",{parentName:"p"},"confixCmd")," package:"),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-go"},'import "cosmossdk.io/tools/confix/cmd"\n')),(0,a.kt)("p",null,"Find the following line:"),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-go"},"initRootCmd(rootCmd, encodingConfig)\n")),(0,a.kt)("p",null,"After that line, add the following:"),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-go"},"rootCmd.AddCommand(\n    confixcmd.ConfigCommand(),\n)\n")),(0,a.kt)("p",null,"The ",(0,a.kt)("inlineCode",{parentName:"p"},"ConfixCommand")," function builds the ",(0,a.kt)("inlineCode",{parentName:"p"},"config")," root command and is defined in the ",(0,a.kt)("inlineCode",{parentName:"p"},"confixCmd")," package (",(0,a.kt)("inlineCode",{parentName:"p"},"cosmossdk.io/tools/confix/cmd"),").\nAn implementation example can be found in ",(0,a.kt)("inlineCode",{parentName:"p"},"simapp"),"."),(0,a.kt)("p",null,"The command will be available as ",(0,a.kt)("inlineCode",{parentName:"p"},"simd config"),"."),(0,a.kt)("h3",{id:"using-confix-standalone"},"Using Confix Standalone"),(0,a.kt)("p",null,"To use Confix standalone, without having to add it in your application, install it with the following command:"),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-bash"},"go install cosmossdk.io/tools/confix/cmd/confix@latest\n")),(0,a.kt)("admonition",{type:"warning"},(0,a.kt)("p",{parentName:"admonition"},"Currently, due to the replace directive in the Confix go.mod, it is not possible to use ",(0,a.kt)("inlineCode",{parentName:"p"},"go install"),".\nBuilding from source or importing in an application is required until that replace directive is removed.")),(0,a.kt)("p",null,"Alternatively, for building from source, simply run ",(0,a.kt)("inlineCode",{parentName:"p"},"make confix"),". The binary will be located in ",(0,a.kt)("inlineCode",{parentName:"p"},"tools/confix"),"."),(0,a.kt)("h2",{id:"usage"},"Usage"),(0,a.kt)("p",null,"Use standalone:"),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-shell"},"confix --help\n")),(0,a.kt)("p",null,"Use in simd:"),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-shell"},"simd config fix --help\n")),(0,a.kt)("h3",{id:"get"},"Get"),(0,a.kt)("p",null,"Get a configuration value, e.g.:"),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-shell"},"simd config get app pruning # gets the value pruning from app.toml\nsimd config get client chain-id # gets the value chain-id from client.toml\n")),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-shell"},"confix get ~/.simapp/config/app.toml pruning # gets the value pruning from app.toml\nconfix get ~/.simapp/config/client.toml chain-id # gets the value chain-id from client.toml\n")),(0,a.kt)("h3",{id:"set"},"Set"),(0,a.kt)("p",null,"Set a configuration value, e.g.:"),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-shell"},'simd config set app pruning "enabled" # sets the value pruning from app.toml\nsimd config set client chain-id "foo-1" # sets the value chain-id from client.toml\n')),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-shell"},'confix set ~/.simapp/config/app.toml pruning "enabled" # sets the value pruning from app.toml\nconfix set ~/.simapp/config/client.toml chain-id "foo-1" # sets the value chain-id from client.toml\n')),(0,a.kt)("h3",{id:"migrate"},"Migrate"),(0,a.kt)("p",null,"Migrate a configuration file to a new version, e.g.:"),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-shell"},"simd config migrate v0.47 # migrates defaultHome/config/app.toml to the latest v0.47 config\n")),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-shell"},"confix migrate v0.47 ~/.simapp/config/app.toml # migrate ~/.simapp/config/app.toml to the latest v0.47 config\n")),(0,a.kt)("h3",{id:"diff"},"Diff"),(0,a.kt)("p",null,"Get the diff between a given configuration file and the default configuration file, e.g.:"),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-shell"},"simd config diff v0.47 # gets the diff between defaultHome/config/app.toml and the latest v0.47 config\n")),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-shell"},"confix diff v0.47 ~/.simapp/config/app.toml # gets the diff between ~/.simapp/config/app.toml and the latest v0.47 config\n")),(0,a.kt)("h3",{id:"maintainer"},"Maintainer"),(0,a.kt)("p",null,"At each SDK modification of the default configuration, add the default SDK config under ",(0,a.kt)("inlineCode",{parentName:"p"},"data/v0.XX-app.toml"),".\nThis allows users to use the tool standalone."),(0,a.kt)("h2",{id:"credits"},"Credits"),(0,a.kt)("p",null,"This project is based on the ",(0,a.kt)("a",{parentName:"p",href:"https://github.com/tendermint/tendermint/blob/5013bc3f4a6d64dcc2bf02ccc002ebc9881c62e4/docs/rfc/rfc-019-config-version.md"},"Tendermint RFC 019")," and their own implementation of ",(0,a.kt)("a",{parentName:"p",href:"https://github.com/tendermint/tendermint/blob/v0.36.x/scripts/confix/confix.go"},"confix"),"."))}d.isMDXComponent=!0}}]);