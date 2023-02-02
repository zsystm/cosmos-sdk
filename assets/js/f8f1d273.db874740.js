"use strict";(self.webpackChunkcosmos_sdk_docs=self.webpackChunkcosmos_sdk_docs||[]).push([[9238],{3905:(e,t,r)=>{r.d(t,{Zo:()=>p,kt:()=>m});var a=r(7294);function n(e,t,r){return t in e?Object.defineProperty(e,t,{value:r,enumerable:!0,configurable:!0,writable:!0}):e[t]=r,e}function i(e,t){var r=Object.keys(e);if(Object.getOwnPropertySymbols){var a=Object.getOwnPropertySymbols(e);t&&(a=a.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),r.push.apply(r,a)}return r}function o(e){for(var t=1;t<arguments.length;t++){var r=null!=arguments[t]?arguments[t]:{};t%2?i(Object(r),!0).forEach((function(t){n(e,t,r[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(r)):i(Object(r)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(r,t))}))}return e}function s(e,t){if(null==e)return{};var r,a,n=function(e,t){if(null==e)return{};var r,a,n={},i=Object.keys(e);for(a=0;a<i.length;a++)r=i[a],t.indexOf(r)>=0||(n[r]=e[r]);return n}(e,t);if(Object.getOwnPropertySymbols){var i=Object.getOwnPropertySymbols(e);for(a=0;a<i.length;a++)r=i[a],t.indexOf(r)>=0||Object.prototype.propertyIsEnumerable.call(e,r)&&(n[r]=e[r])}return n}var l=a.createContext({}),c=function(e){var t=a.useContext(l),r=t;return e&&(r="function"==typeof e?e(t):o(o({},t),e)),r},p=function(e){var t=c(e.components);return a.createElement(l.Provider,{value:t},e.children)},u={inlineCode:"code",wrapper:function(e){var t=e.children;return a.createElement(a.Fragment,{},t)}},d=a.forwardRef((function(e,t){var r=e.components,n=e.mdxType,i=e.originalType,l=e.parentName,p=s(e,["components","mdxType","originalType","parentName"]),d=c(r),m=n,f=d["".concat(l,".").concat(m)]||d[m]||u[m]||i;return r?a.createElement(f,o(o({ref:t},p),{},{components:r})):a.createElement(f,o({ref:t},p))}));function m(e,t){var r=arguments,n=t&&t.mdxType;if("string"==typeof e||n){var i=r.length,o=new Array(i);o[0]=d;var s={};for(var l in t)hasOwnProperty.call(t,l)&&(s[l]=t[l]);s.originalType=e,s.mdxType="string"==typeof e?e:n,o[1]=s;for(var c=2;c<i;c++)o[c]=r[c];return a.createElement.apply(null,o)}return a.createElement.apply(null,r)}d.displayName="MDXCreateElement"},4418:(e,t,r)=>{r.r(t),r.d(t,{assets:()=>l,contentTitle:()=>o,default:()=>u,frontMatter:()=>i,metadata:()=>s,toc:()=>c});var a=r(7462),n=(r(7294),r(3905));const i={},o="ADR 006: Secret Store Replacement",s={unversionedId:"architecture/adr-006-secret-store-replacement",id:"architecture/adr-006-secret-store-replacement",title:"ADR 006: Secret Store Replacement",description:"Changelog",source:"@site/docs/architecture/adr-006-secret-store-replacement.md",sourceDirName:"architecture",slug:"/architecture/adr-006-secret-store-replacement",permalink:"/main/architecture/adr-006-secret-store-replacement",draft:!1,tags:[],version:"current",frontMatter:{},sidebar:"tutorialSidebar",previous:{title:"ADR 004: Split Denomination Keys",permalink:"/main/architecture/adr-004-split-denomination-keys"},next:{title:"ADR 007: Specialization Groups",permalink:"/main/architecture/adr-007-specialization-groups"}},l={},c=[{value:"Changelog",id:"changelog",level:2},{value:"Context",id:"context",level:2},{value:"Decision",id:"decision",level:2},{value:"Status",id:"status",level:2},{value:"Consequences",id:"consequences",level:2},{value:"Positive",id:"positive",level:3},{value:"Negative",id:"negative",level:3},{value:"Neutral",id:"neutral",level:3},{value:"References",id:"references",level:2}],p={toc:c};function u(e){let{components:t,...r}=e;return(0,n.kt)("wrapper",(0,a.Z)({},p,r,{components:t,mdxType:"MDXLayout"}),(0,n.kt)("h1",{id:"adr-006-secret-store-replacement"},"ADR 006: Secret Store Replacement"),(0,n.kt)("h2",{id:"changelog"},"Changelog"),(0,n.kt)("ul",null,(0,n.kt)("li",{parentName:"ul"},"July 29th, 2019: Initial draft"),(0,n.kt)("li",{parentName:"ul"},"September 11th, 2019: Work has started"),(0,n.kt)("li",{parentName:"ul"},"November 4th: Cosmos SDK changes merged in"),(0,n.kt)("li",{parentName:"ul"},"November 18th: Gaia changes merged in")),(0,n.kt)("h2",{id:"context"},"Context"),(0,n.kt)("p",null,"Currently, a Cosmos SDK application's CLI directory stores key material and metadata in a plain text database in the user\u2019s home directory.  Key material is encrypted by a passphrase, protected by bcrypt hashing algorithm. Metadata (e.g. addresses, public keys, key storage details) is available in plain text."),(0,n.kt)("p",null,"This is not desirable for a number of reasons. Perhaps the biggest reason is insufficient security protection of key material and metadata. Leaking the plain text allows an attacker to surveil what keys a given computer controls via a number of techniques, like compromised dependencies without any privilege execution. This could be followed by a more targeted attack on a particular user/computer."),(0,n.kt)("p",null,"All modern desktop computers OS (Ubuntu, Debian, MacOS, Windows) provide a built-in secret store that is designed to allow applications to store information that is isolated from all other applications and requires passphrase entry to access the data."),(0,n.kt)("p",null,"We are seeking solution that provides a common abstraction layer to the many different backends and reasonable fallback for minimal platforms that don\u2019t provide a native secret store."),(0,n.kt)("h2",{id:"decision"},"Decision"),(0,n.kt)("p",null,"We recommend replacing the current Keybase backend based on LevelDB with ",(0,n.kt)("a",{parentName:"p",href:"https://github.com/99designs/keyring"},"Keyring")," by 99 designs. This application is designed to provide a common abstraction and uniform interface between many secret stores and is used by AWS Vault application by 99-designs application."),(0,n.kt)("p",null,"This appears to fulfill the requirement of protecting both key material and metadata from rouge software on a user\u2019s machine."),(0,n.kt)("h2",{id:"status"},"Status"),(0,n.kt)("p",null,"Accepted"),(0,n.kt)("h2",{id:"consequences"},"Consequences"),(0,n.kt)("h3",{id:"positive"},"Positive"),(0,n.kt)("p",null,"Increased safety for users."),(0,n.kt)("h3",{id:"negative"},"Negative"),(0,n.kt)("p",null,"Users must manually migrate."),(0,n.kt)("p",null,"Testing against all supported backends is difficult."),(0,n.kt)("p",null,"Running tests locally on a Mac require numerous repetitive password entries."),(0,n.kt)("h3",{id:"neutral"},"Neutral"),(0,n.kt)("p",null,"{neutral consequences}"),(0,n.kt)("h2",{id:"references"},"References"),(0,n.kt)("ul",null,(0,n.kt)("li",{parentName:"ul"},"#4754 Switch secret store to the keyring secret store (original PR by @poldsam) ","[",(0,n.kt)("strong",{parentName:"li"},"CLOSED"),"]"),(0,n.kt)("li",{parentName:"ul"},"#5029 Add support for github.com/99designs/keyring-backed keybases ","[",(0,n.kt)("strong",{parentName:"li"},"MERGED"),"]"),(0,n.kt)("li",{parentName:"ul"},"#5097 Add keys migrate command ","[",(0,n.kt)("strong",{parentName:"li"},"MERGED"),"]"),(0,n.kt)("li",{parentName:"ul"},"#5180 Drop on-disk keybase in favor of keyring ","[",(0,n.kt)("em",{parentName:"li"},"PENDING_REVIEW"),"]"),(0,n.kt)("li",{parentName:"ul"},"cosmos/gaia#164 Drop on-disk keybase in favor of keyring (gaia's changes) ","[",(0,n.kt)("em",{parentName:"li"},"PENDING_REVIEW"),"]")))}u.isMDXComponent=!0}}]);