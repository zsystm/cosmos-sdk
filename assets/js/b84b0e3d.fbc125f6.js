"use strict";(self.webpackChunkcosmos_sdk_docs=self.webpackChunkcosmos_sdk_docs||[]).push([[6377],{3905:(e,t,n)=>{n.d(t,{Zo:()=>c,kt:()=>m});var i=n(7294);function a(e,t,n){return t in e?Object.defineProperty(e,t,{value:n,enumerable:!0,configurable:!0,writable:!0}):e[t]=n,e}function o(e,t){var n=Object.keys(e);if(Object.getOwnPropertySymbols){var i=Object.getOwnPropertySymbols(e);t&&(i=i.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),n.push.apply(n,i)}return n}function r(e){for(var t=1;t<arguments.length;t++){var n=null!=arguments[t]?arguments[t]:{};t%2?o(Object(n),!0).forEach((function(t){a(e,t,n[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(n)):o(Object(n)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(n,t))}))}return e}function s(e,t){if(null==e)return{};var n,i,a=function(e,t){if(null==e)return{};var n,i,a={},o=Object.keys(e);for(i=0;i<o.length;i++)n=o[i],t.indexOf(n)>=0||(a[n]=e[n]);return a}(e,t);if(Object.getOwnPropertySymbols){var o=Object.getOwnPropertySymbols(e);for(i=0;i<o.length;i++)n=o[i],t.indexOf(n)>=0||Object.prototype.propertyIsEnumerable.call(e,n)&&(a[n]=e[n])}return a}var l=i.createContext({}),p=function(e){var t=i.useContext(l),n=t;return e&&(n="function"==typeof e?e(t):r(r({},t),e)),n},c=function(e){var t=p(e.components);return i.createElement(l.Provider,{value:t},e.children)},d={inlineCode:"code",wrapper:function(e){var t=e.children;return i.createElement(i.Fragment,{},t)}},u=i.forwardRef((function(e,t){var n=e.components,a=e.mdxType,o=e.originalType,l=e.parentName,c=s(e,["components","mdxType","originalType","parentName"]),u=p(n),m=a,f=u["".concat(l,".").concat(m)]||u[m]||d[m]||o;return n?i.createElement(f,r(r({ref:t},c),{},{components:n})):i.createElement(f,r({ref:t},c))}));function m(e,t){var n=arguments,a=t&&t.mdxType;if("string"==typeof e||a){var o=n.length,r=new Array(o);r[0]=u;var s={};for(var l in t)hasOwnProperty.call(t,l)&&(s[l]=t[l]);s.originalType=e,s.mdxType="string"==typeof e?e:a,r[1]=s;for(var p=2;p<o;p++)r[p]=n[p];return i.createElement.apply(null,r)}return i.createElement.apply(null,n)}u.displayName="MDXCreateElement"},8933:(e,t,n)=>{n.r(t),n.d(t,{assets:()=>l,contentTitle:()=>r,default:()=>d,frontMatter:()=>o,metadata:()=>s,toc:()=>p});var i=n(7462),a=(n(7294),n(3905));const o={},r="Specification of Specifications",s={unversionedId:"spec/SPEC-SPEC",id:"version-v0.47/spec/SPEC-SPEC",title:"Specification of Specifications",description:"This file intends to outline the common structure for specifications within",source:"@site/versioned_docs/version-v0.47/spec/SPEC-SPEC.md",sourceDirName:"spec",slug:"/spec/SPEC-SPEC",permalink:"/v0.47/spec/SPEC-SPEC",draft:!1,tags:[],version:"v0.47",frontMatter:{},sidebar:"tutorialSidebar",previous:{title:"Specifications",permalink:"/v0.47/spec/"},next:{title:"Addresses spec",permalink:"/v0.47/spec/addresses/"}},l={},p=[{value:"Tense",id:"tense",level:2},{value:"Pseudo-Code",id:"pseudo-code",level:2},{value:"Common Layout",id:"common-layout",level:2},{value:"Notation for key-value mapping",id:"notation-for-key-value-mapping",level:3}],c={toc:p};function d(e){let{components:t,...n}=e;return(0,a.kt)("wrapper",(0,i.Z)({},c,n,{components:t,mdxType:"MDXLayout"}),(0,a.kt)("h1",{id:"specification-of-specifications"},"Specification of Specifications"),(0,a.kt)("p",null,"This file intends to outline the common structure for specifications within\nthis directory."),(0,a.kt)("h2",{id:"tense"},"Tense"),(0,a.kt)("p",null,"For consistency, specs should be written in passive present tense."),(0,a.kt)("h2",{id:"pseudo-code"},"Pseudo-Code"),(0,a.kt)("p",null,"Generally, pseudo-code should be minimized throughout the spec. Often, simple\nbulleted-lists which describe a function's operations are sufficient and should\nbe considered preferable. In certain instances, due to the complex nature of\nthe functionality being described pseudo-code may the most suitable form of\nspecification. In these cases use of pseudo-code is permissible, but should be\npresented in a concise manner, ideally restricted to only the complex\nelement as a part of a larger description."),(0,a.kt)("h2",{id:"common-layout"},"Common Layout"),(0,a.kt)("p",null,"The following generalized ",(0,a.kt)("inlineCode",{parentName:"p"},"README")," structure should be used to breakdown\nspecifications for modules. The following list is nonbinding and all sections are optional."),(0,a.kt)("ul",null,(0,a.kt)("li",{parentName:"ul"},(0,a.kt)("inlineCode",{parentName:"li"},"# {Module Name}")," - overview of the module"),(0,a.kt)("li",{parentName:"ul"},(0,a.kt)("inlineCode",{parentName:"li"},"## Concepts")," - describe specialized concepts and definitions used throughout the spec"),(0,a.kt)("li",{parentName:"ul"},(0,a.kt)("inlineCode",{parentName:"li"},"## State")," - specify and describe structures expected to marshalled into the store, and their keys"),(0,a.kt)("li",{parentName:"ul"},(0,a.kt)("inlineCode",{parentName:"li"},"## State Transitions")," - standard state transition operations triggered by hooks, messages, etc."),(0,a.kt)("li",{parentName:"ul"},(0,a.kt)("inlineCode",{parentName:"li"},"## Messages")," - specify message structure(s) and expected state machine behaviour(s)"),(0,a.kt)("li",{parentName:"ul"},(0,a.kt)("inlineCode",{parentName:"li"},"## Begin Block")," - specify any begin-block operations"),(0,a.kt)("li",{parentName:"ul"},(0,a.kt)("inlineCode",{parentName:"li"},"## End Block")," - specify any end-block operations"),(0,a.kt)("li",{parentName:"ul"},(0,a.kt)("inlineCode",{parentName:"li"},"## Hooks")," - describe available hooks to be called by/from this module"),(0,a.kt)("li",{parentName:"ul"},(0,a.kt)("inlineCode",{parentName:"li"},"## Events")," - list and describe event tags used"),(0,a.kt)("li",{parentName:"ul"},(0,a.kt)("inlineCode",{parentName:"li"},"## Client")," - list and describe CLI commands and gRPC and REST endpoints"),(0,a.kt)("li",{parentName:"ul"},(0,a.kt)("inlineCode",{parentName:"li"},"## Params")," - list all module parameters, their types (in JSON) and examples"),(0,a.kt)("li",{parentName:"ul"},(0,a.kt)("inlineCode",{parentName:"li"},"## Future Improvements")," - describe future improvements of this module"),(0,a.kt)("li",{parentName:"ul"},(0,a.kt)("inlineCode",{parentName:"li"},"## Tests")," - acceptance tests"),(0,a.kt)("li",{parentName:"ul"},(0,a.kt)("inlineCode",{parentName:"li"},"## Appendix")," - supplementary details referenced elsewhere within the spec")),(0,a.kt)("h3",{id:"notation-for-key-value-mapping"},"Notation for key-value mapping"),(0,a.kt)("p",null,"Within ",(0,a.kt)("inlineCode",{parentName:"p"},"## State")," the following notation ",(0,a.kt)("inlineCode",{parentName:"p"},"->")," should be used to describe key to\nvalue mapping:"),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-text"},"key -> value\n")),(0,a.kt)("p",null,"to represent byte concatenation the ",(0,a.kt)("inlineCode",{parentName:"p"},"|")," may be used. In addition, encoding\ntype may be specified, for example:"),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-text"},"0x00 | addressBytes | address2Bytes -> amino(value_object)\n")),(0,a.kt)("p",null,"Additionally, index mappings may be specified by mapping to the ",(0,a.kt)("inlineCode",{parentName:"p"},"nil")," value, for example:"),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-text"},"0x01 | address2Bytes | addressBytes -> nil\n")))}d.isMDXComponent=!0}}]);