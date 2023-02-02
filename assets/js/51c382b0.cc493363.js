"use strict";(self.webpackChunkcosmos_sdk_docs=self.webpackChunkcosmos_sdk_docs||[]).push([[3755],{3905:(e,t,n)=>{n.d(t,{Zo:()=>p,kt:()=>m});var r=n(7294);function i(e,t,n){return t in e?Object.defineProperty(e,t,{value:n,enumerable:!0,configurable:!0,writable:!0}):e[t]=n,e}function o(e,t){var n=Object.keys(e);if(Object.getOwnPropertySymbols){var r=Object.getOwnPropertySymbols(e);t&&(r=r.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),n.push.apply(n,r)}return n}function a(e){for(var t=1;t<arguments.length;t++){var n=null!=arguments[t]?arguments[t]:{};t%2?o(Object(n),!0).forEach((function(t){i(e,t,n[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(n)):o(Object(n)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(n,t))}))}return e}function s(e,t){if(null==e)return{};var n,r,i=function(e,t){if(null==e)return{};var n,r,i={},o=Object.keys(e);for(r=0;r<o.length;r++)n=o[r],t.indexOf(n)>=0||(i[n]=e[n]);return i}(e,t);if(Object.getOwnPropertySymbols){var o=Object.getOwnPropertySymbols(e);for(r=0;r<o.length;r++)n=o[r],t.indexOf(n)>=0||Object.prototype.propertyIsEnumerable.call(e,n)&&(i[n]=e[n])}return i}var l=r.createContext({}),u=function(e){var t=r.useContext(l),n=t;return e&&(n="function"==typeof e?e(t):a(a({},t),e)),n},p=function(e){var t=u(e.components);return r.createElement(l.Provider,{value:t},e.children)},c={inlineCode:"code",wrapper:function(e){var t=e.children;return r.createElement(r.Fragment,{},t)}},d=r.forwardRef((function(e,t){var n=e.components,i=e.mdxType,o=e.originalType,l=e.parentName,p=s(e,["components","mdxType","originalType","parentName"]),d=u(n),m=i,g=d["".concat(l,".").concat(m)]||d[m]||c[m]||o;return n?r.createElement(g,a(a({ref:t},p),{},{components:n})):r.createElement(g,a({ref:t},p))}));function m(e,t){var n=arguments,i=t&&t.mdxType;if("string"==typeof e||i){var o=n.length,a=new Array(o);a[0]=d;var s={};for(var l in t)hasOwnProperty.call(t,l)&&(s[l]=t[l]);s.originalType=e,s.mdxType="string"==typeof e?e:i,a[1]=s;for(var u=2;u<o;u++)a[u]=n[u];return r.createElement.apply(null,a)}return r.createElement.apply(null,n)}d.displayName="MDXCreateElement"},9592:(e,t,n)=>{n.r(t),n.d(t,{assets:()=>l,contentTitle:()=>a,default:()=>c,frontMatter:()=>o,metadata:()=>s,toc:()=>u});var r=n(7462),i=(n(7294),n(3905));const o={sidebar_position:1},a="Upgrading Modules",s={unversionedId:"building-modules/upgrade",id:"version-v0.47/building-modules/upgrade",title:"Upgrading Modules",description:"In-Place Store Migrations allow your modules to upgrade to new versions that include breaking changes. This document outlines how to build modules to take advantage of this functionality.",source:"@site/versioned_docs/version-v0.47/building-modules/12-upgrade.md",sourceDirName:"building-modules",slug:"/building-modules/upgrade",permalink:"/v0.47/building-modules/upgrade",draft:!1,tags:[],version:"v0.47",sidebarPosition:1,frontMatter:{sidebar_position:1},sidebar:"tutorialSidebar",previous:{title:"Errors",permalink:"/v0.47/building-modules/errors"},next:{title:"Module Simulation",permalink:"/v0.47/building-modules/simulator"}},l={},u=[{value:"Consensus Version",id:"consensus-version",level:2},{value:"Registering Migrations",id:"registering-migrations",level:2},{value:"Writing Migration Scripts",id:"writing-migration-scripts",level:2}],p={toc:u};function c(e){let{components:t,...n}=e;return(0,i.kt)("wrapper",(0,r.Z)({},p,n,{components:t,mdxType:"MDXLayout"}),(0,i.kt)("h1",{id:"upgrading-modules"},"Upgrading Modules"),(0,i.kt)("admonition",{title:"Synopsis",type:"note"},(0,i.kt)("p",{parentName:"admonition"},(0,i.kt)("a",{parentName:"p",href:"/v0.47/core/upgrade"},"In-Place Store Migrations")," allow your modules to upgrade to new versions that include breaking changes. This document outlines how to build modules to take advantage of this functionality.")),(0,i.kt)("admonition",{type:"note"},(0,i.kt)("h3",{parentName:"admonition",id:"pre-requisite-readings"},"Pre-requisite Readings"),(0,i.kt)("ul",{parentName:"admonition"},(0,i.kt)("li",{parentName:"ul"},(0,i.kt)("a",{parentName:"li",href:"/v0.47/core/upgrade"},"In-Place Store Migration")))),(0,i.kt)("h2",{id:"consensus-version"},"Consensus Version"),(0,i.kt)("p",null,"Successful upgrades of existing modules require each ",(0,i.kt)("inlineCode",{parentName:"p"},"AppModule")," to implement the function ",(0,i.kt)("inlineCode",{parentName:"p"},"ConsensusVersion() uint64"),"."),(0,i.kt)("ul",null,(0,i.kt)("li",{parentName:"ul"},"The versions must be hard-coded by the module developer."),(0,i.kt)("li",{parentName:"ul"},"The initial version ",(0,i.kt)("strong",{parentName:"li"},"must")," be set to 1.")),(0,i.kt)("p",null,"Consensus versions serve as state-breaking versions of app modules and must be incremented when the module introduces breaking changes."),(0,i.kt)("h2",{id:"registering-migrations"},"Registering Migrations"),(0,i.kt)("p",null,"To register the functionality that takes place during a module upgrade, you must register which migrations you want to take place."),(0,i.kt)("p",null,"Migration registration takes place in the ",(0,i.kt)("inlineCode",{parentName:"p"},"Configurator")," using the ",(0,i.kt)("inlineCode",{parentName:"p"},"RegisterMigration")," method. The ",(0,i.kt)("inlineCode",{parentName:"p"},"AppModule")," reference to the configurator is in the ",(0,i.kt)("inlineCode",{parentName:"p"},"RegisterServices")," method."),(0,i.kt)("p",null,"You can register one or more migrations. If you register more than one migration script, list the migrations in increasing order and ensure there are enough migrations that lead to the desired consensus version. For example, to migrate to version 3 of a module, register separate migrations for version 1 and version 2 as shown in the following example:"),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-go"},"func (am AppModule) RegisterServices(cfg module.Configurator) {\n    // --snip--\n    cfg.RegisterMigration(types.ModuleName, 1, func(ctx sdk.Context) error {\n        // Perform in-place store migrations from ConsensusVersion 1 to 2.\n    })\n     cfg.RegisterMigration(types.ModuleName, 2, func(ctx sdk.Context) error {\n        // Perform in-place store migrations from ConsensusVersion 2 to 3.\n    })\n}\n")),(0,i.kt)("p",null,"Since these migrations are functions that need access to a Keeper's store, use a wrapper around the keepers called ",(0,i.kt)("inlineCode",{parentName:"p"},"Migrator")," as shown in this example:"),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-go",metastring:"reference",reference:!0},"https://github.com/cosmos/cosmos-sdk/blob/v0.47.0-rc1/x/bank/keeper/migrations.go#L11-L35\n")),(0,i.kt)("h2",{id:"writing-migration-scripts"},"Writing Migration Scripts"),(0,i.kt)("p",null,"To define the functionality that takes place during an upgrade, write a migration script and place the functions in a ",(0,i.kt)("inlineCode",{parentName:"p"},"migrations/")," directory. For example, to write migration scripts for the bank module, place the functions in ",(0,i.kt)("inlineCode",{parentName:"p"},"x/bank/migrations/"),". Use the recommended naming convention for these functions. For example, ",(0,i.kt)("inlineCode",{parentName:"p"},"v2bank")," is the script that migrates the package ",(0,i.kt)("inlineCode",{parentName:"p"},"x/bank/migrations/v2"),":"),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-go"},"// Migrating bank module from version 1 to 2\nfunc (m Migrator) Migrate1to2(ctx sdk.Context) error {\n    return v2bank.MigrateStore(ctx, m.keeper.storeKey) // v2bank is package `x/bank/migrations/v2`.\n}\n")),(0,i.kt)("p",null,"To see example code of changes that were implemented in a migration of balance keys, check out ",(0,i.kt)("a",{parentName:"p",href:"https://github.com/cosmos/cosmos-sdk/blob/v0.47.0-rc1/x/bank/migrations/v2/store.go#L52-L73"},"migrateBalanceKeys"),". For context, this code introduced migrations of the bank store that updated addresses to be prefixed by their length in bytes as outlined in ",(0,i.kt)("a",{parentName:"p",href:"/v0.47/architecture/adr-028-public-key-addresses"},"ADR-028"),"."))}c.isMDXComponent=!0}}]);