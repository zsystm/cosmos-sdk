"use strict";(self.webpackChunkcosmos_sdk_docs=self.webpackChunkcosmos_sdk_docs||[]).push([[1050],{3905:(e,n,t)=>{t.d(n,{Zo:()=>c,kt:()=>m});var i=t(7294);function a(e,n,t){return n in e?Object.defineProperty(e,n,{value:t,enumerable:!0,configurable:!0,writable:!0}):e[n]=t,e}function r(e,n){var t=Object.keys(e);if(Object.getOwnPropertySymbols){var i=Object.getOwnPropertySymbols(e);n&&(i=i.filter((function(n){return Object.getOwnPropertyDescriptor(e,n).enumerable}))),t.push.apply(t,i)}return t}function o(e){for(var n=1;n<arguments.length;n++){var t=null!=arguments[n]?arguments[n]:{};n%2?r(Object(t),!0).forEach((function(n){a(e,n,t[n])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(t)):r(Object(t)).forEach((function(n){Object.defineProperty(e,n,Object.getOwnPropertyDescriptor(t,n))}))}return e}function l(e,n){if(null==e)return{};var t,i,a=function(e,n){if(null==e)return{};var t,i,a={},r=Object.keys(e);for(i=0;i<r.length;i++)t=r[i],n.indexOf(t)>=0||(a[t]=e[t]);return a}(e,n);if(Object.getOwnPropertySymbols){var r=Object.getOwnPropertySymbols(e);for(i=0;i<r.length;i++)t=r[i],n.indexOf(t)>=0||Object.prototype.propertyIsEnumerable.call(e,t)&&(a[t]=e[t])}return a}var s=i.createContext({}),d=function(e){var n=i.useContext(s),t=n;return e&&(t="function"==typeof e?e(n):o(o({},n),e)),t},c=function(e){var n=d(e.components);return i.createElement(s.Provider,{value:n},e.children)},u={inlineCode:"code",wrapper:function(e){var n=e.children;return i.createElement(i.Fragment,{},n)}},p=i.forwardRef((function(e,n){var t=e.components,a=e.mdxType,r=e.originalType,s=e.parentName,c=l(e,["components","mdxType","originalType","parentName"]),p=d(t),m=a,h=p["".concat(s,".").concat(m)]||p[m]||u[m]||r;return t?i.createElement(h,o(o({ref:n},c),{},{components:t})):i.createElement(h,o({ref:n},c))}));function m(e,n){var t=arguments,a=n&&n.mdxType;if("string"==typeof e||a){var r=t.length,o=new Array(r);o[0]=p;var l={};for(var s in n)hasOwnProperty.call(n,s)&&(l[s]=n[s]);l.originalType=e,l.mdxType="string"==typeof e?e:a,o[1]=l;for(var d=2;d<r;d++)o[d]=t[d];return i.createElement.apply(null,o)}return i.createElement.apply(null,t)}p.displayName="MDXCreateElement"},722:(e,n,t)=>{t.r(n),t.d(n,{assets:()=>s,contentTitle:()=>o,default:()=>u,frontMatter:()=>r,metadata:()=>l,toc:()=>d});var i=t(7462),a=(t(7294),t(3905));const r={},o="ADR 009: Evidence Module",l={unversionedId:"architecture/adr-009-evidence-module",id:"architecture/adr-009-evidence-module",title:"ADR 009: Evidence Module",description:"Changelog",source:"@site/docs/architecture/adr-009-evidence-module.md",sourceDirName:"architecture",slug:"/architecture/adr-009-evidence-module",permalink:"/main/architecture/adr-009-evidence-module",draft:!1,tags:[],version:"current",frontMatter:{},sidebar:"tutorialSidebar",previous:{title:"ADR 008: Decentralized Computer Emergency Response Team (dCERT) Group",permalink:"/main/architecture/adr-008-dCERT-group"},next:{title:"ADR 010: Modular AnteHandler",permalink:"/main/architecture/adr-010-modular-antehandler"}},s={},d=[{value:"Changelog",id:"changelog",level:2},{value:"Status",id:"status",level:2},{value:"Context",id:"context",level:2},{value:"Decision",id:"decision",level:2},{value:"Types",id:"types",level:3},{value:"Routing &amp; Handling",id:"routing--handling",level:3},{value:"Submission",id:"submission",level:3},{value:"Genesis",id:"genesis",level:3},{value:"Consequences",id:"consequences",level:2},{value:"Positive",id:"positive",level:3},{value:"Negative",id:"negative",level:3},{value:"Neutral",id:"neutral",level:3},{value:"References",id:"references",level:2}],c={toc:d};function u(e){let{components:n,...t}=e;return(0,a.kt)("wrapper",(0,i.Z)({},c,t,{components:n,mdxType:"MDXLayout"}),(0,a.kt)("h1",{id:"adr-009-evidence-module"},"ADR 009: Evidence Module"),(0,a.kt)("h2",{id:"changelog"},"Changelog"),(0,a.kt)("ul",null,(0,a.kt)("li",{parentName:"ul"},"2019 July 31: Initial draft"),(0,a.kt)("li",{parentName:"ul"},"2019 October 24: Initial implementation")),(0,a.kt)("h2",{id:"status"},"Status"),(0,a.kt)("p",null,"Accepted"),(0,a.kt)("h2",{id:"context"},"Context"),(0,a.kt)("p",null,"In order to support building highly secure, robust and interoperable blockchain\napplications, it is vital for the Cosmos SDK to expose a mechanism in which arbitrary\nevidence can be submitted, evaluated and verified resulting in some agreed upon\npenalty for any misbehavior committed by a validator, such as equivocation (double-voting),\nsigning when unbonded, signing an incorrect state transition (in the future), etc.\nFurthermore, such a mechanism is paramount for any\n",(0,a.kt)("a",{parentName:"p",href:"https://github.com/cosmos/ics/blob/master/ibc/2_IBC_ARCHITECTURE.md"},"IBC")," or\ncross-chain validation protocol implementation in order to support the ability\nfor any misbehavior to be relayed back from a collateralized chain to a primary\nchain so that the equivocating validator(s) can be slashed."),(0,a.kt)("h2",{id:"decision"},"Decision"),(0,a.kt)("p",null,"We will implement an evidence module in the Cosmos SDK supporting the following\nfunctionality:"),(0,a.kt)("ul",null,(0,a.kt)("li",{parentName:"ul"},"Provide developers with the abstractions and interfaces necessary to define\ncustom evidence messages, message handlers, and methods to slash and penalize\naccordingly for misbehavior."),(0,a.kt)("li",{parentName:"ul"},"Support the ability to route evidence messages to handlers in any module to\ndetermine the validity of submitted misbehavior."),(0,a.kt)("li",{parentName:"ul"},"Support the ability, through governance, to modify slashing penalties of any\nevidence type."),(0,a.kt)("li",{parentName:"ul"},"Querier implementation to support querying params, evidence types, params, and\nall submitted valid misbehavior.")),(0,a.kt)("h3",{id:"types"},"Types"),(0,a.kt)("p",null,"First, we define the ",(0,a.kt)("inlineCode",{parentName:"p"},"Evidence")," interface type. The ",(0,a.kt)("inlineCode",{parentName:"p"},"x/evidence")," module may implement\nits own types that can be used by many chains (e.g. ",(0,a.kt)("inlineCode",{parentName:"p"},"CounterFactualEvidence"),").\nIn addition, other modules may implement their own ",(0,a.kt)("inlineCode",{parentName:"p"},"Evidence")," types in a similar\nmanner in which governance is extensible. It is important to note any concrete\ntype implementing the ",(0,a.kt)("inlineCode",{parentName:"p"},"Evidence")," interface may include arbitrary fields such as\nan infraction time. We want the ",(0,a.kt)("inlineCode",{parentName:"p"},"Evidence")," type to remain as flexible as possible."),(0,a.kt)("p",null,"When submitting evidence to the ",(0,a.kt)("inlineCode",{parentName:"p"},"x/evidence")," module, the concrete type must provide\nthe validator's consensus address, which should be known by the ",(0,a.kt)("inlineCode",{parentName:"p"},"x/slashing"),"\nmodule (assuming the infraction is valid), the height at which the infraction\noccurred and the validator's power at same height in which the infraction occurred."),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-go"},"type Evidence interface {\n  Route() string\n  Type() string\n  String() string\n  Hash() HexBytes\n  ValidateBasic() error\n\n  // The consensus address of the malicious validator at time of infraction\n  GetConsensusAddress() ConsAddress\n\n  // Height at which the infraction occurred\n  GetHeight() int64\n\n  // The total power of the malicious validator at time of infraction\n  GetValidatorPower() int64\n\n  // The total validator set power at time of infraction\n  GetTotalPower() int64\n}\n")),(0,a.kt)("h3",{id:"routing--handling"},"Routing & Handling"),(0,a.kt)("p",null,"Each ",(0,a.kt)("inlineCode",{parentName:"p"},"Evidence")," type must map to a specific unique route and be registered with\nthe ",(0,a.kt)("inlineCode",{parentName:"p"},"x/evidence")," module. It accomplishes this through the ",(0,a.kt)("inlineCode",{parentName:"p"},"Router")," implementation."),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-go"},"type Router interface {\n  AddRoute(r string, h Handler) Router\n  HasRoute(r string) bool\n  GetRoute(path string) Handler\n  Seal()\n}\n")),(0,a.kt)("p",null,"Upon successful routing through the ",(0,a.kt)("inlineCode",{parentName:"p"},"x/evidence")," module, the ",(0,a.kt)("inlineCode",{parentName:"p"},"Evidence")," type\nis passed through a ",(0,a.kt)("inlineCode",{parentName:"p"},"Handler"),". This ",(0,a.kt)("inlineCode",{parentName:"p"},"Handler")," is responsible for executing all\ncorresponding business logic necessary for verifying the evidence as valid. In\naddition, the ",(0,a.kt)("inlineCode",{parentName:"p"},"Handler")," may execute any necessary slashing and potential jailing.\nSince slashing fractions will typically result from some form of static functions,\nallow the ",(0,a.kt)("inlineCode",{parentName:"p"},"Handler")," to do this provides the greatest flexibility. An example could\nbe ",(0,a.kt)("inlineCode",{parentName:"p"},"k * evidence.GetValidatorPower()")," where ",(0,a.kt)("inlineCode",{parentName:"p"},"k")," is an on-chain parameter controlled\nby governance. The ",(0,a.kt)("inlineCode",{parentName:"p"},"Evidence")," type should provide all the external information\nnecessary in order for the ",(0,a.kt)("inlineCode",{parentName:"p"},"Handler")," to make the necessary state transitions.\nIf no error is returned, the ",(0,a.kt)("inlineCode",{parentName:"p"},"Evidence")," is considered valid."),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-go"},"type Handler func(Context, Evidence) error\n")),(0,a.kt)("h3",{id:"submission"},"Submission"),(0,a.kt)("p",null,(0,a.kt)("inlineCode",{parentName:"p"},"Evidence")," is submitted through a ",(0,a.kt)("inlineCode",{parentName:"p"},"MsgSubmitEvidence")," message type which is internally\nhandled by the ",(0,a.kt)("inlineCode",{parentName:"p"},"x/evidence")," module's ",(0,a.kt)("inlineCode",{parentName:"p"},"SubmitEvidence"),"."),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-go"},"type MsgSubmitEvidence struct {\n  Evidence\n}\n\nfunc handleMsgSubmitEvidence(ctx Context, keeper Keeper, msg MsgSubmitEvidence) Result {\n  if err := keeper.SubmitEvidence(ctx, msg.Evidence); err != nil {\n    return err.Result()\n  }\n\n  // emit events...\n\n  return Result{\n    // ...\n  }\n}\n")),(0,a.kt)("p",null,"The ",(0,a.kt)("inlineCode",{parentName:"p"},"x/evidence")," module's keeper is responsible for matching the ",(0,a.kt)("inlineCode",{parentName:"p"},"Evidence")," against\nthe module's router and invoking the corresponding ",(0,a.kt)("inlineCode",{parentName:"p"},"Handler")," which may include\nslashing and jailing the validator. Upon success, the submitted evidence is persisted."),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-go"},"func (k Keeper) SubmitEvidence(ctx Context, evidence Evidence) error {\n  handler := keeper.router.GetRoute(evidence.Route())\n  if err := handler(ctx, evidence); err != nil {\n    return ErrInvalidEvidence(keeper.codespace, err)\n  }\n\n  keeper.setEvidence(ctx, evidence)\n  return nil\n}\n")),(0,a.kt)("h3",{id:"genesis"},"Genesis"),(0,a.kt)("p",null,"Finally, we need to represent the genesis state of the ",(0,a.kt)("inlineCode",{parentName:"p"},"x/evidence")," module. The\nmodule only needs a list of all submitted valid infractions and any necessary params\nfor which the module needs in order to handle submitted evidence. The ",(0,a.kt)("inlineCode",{parentName:"p"},"x/evidence"),"\nmodule will naturally define and route native evidence types for which it'll most\nlikely need slashing penalty constants for."),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-go"},"type GenesisState struct {\n  Params       Params\n  Infractions  []Evidence\n}\n")),(0,a.kt)("h2",{id:"consequences"},"Consequences"),(0,a.kt)("h3",{id:"positive"},"Positive"),(0,a.kt)("ul",null,(0,a.kt)("li",{parentName:"ul"},"Allows the state machine to process misbehavior submitted on-chain and penalize\nvalidators based on agreed upon slashing parameters."),(0,a.kt)("li",{parentName:"ul"},"Allows evidence types to be defined and handled by any module. This further allows\nslashing and jailing to be defined by more complex mechanisms."),(0,a.kt)("li",{parentName:"ul"},"Does not solely rely on Tendermint to submit evidence.")),(0,a.kt)("h3",{id:"negative"},"Negative"),(0,a.kt)("ul",null,(0,a.kt)("li",{parentName:"ul"},"No easy way to introduce new evidence types through governance on a live chain\ndue to the inability to introduce the new evidence type's corresponding handler")),(0,a.kt)("h3",{id:"neutral"},"Neutral"),(0,a.kt)("ul",null,(0,a.kt)("li",{parentName:"ul"},"Should we persist infractions indefinitely? Or should we rather rely on events?")),(0,a.kt)("h2",{id:"references"},"References"),(0,a.kt)("ul",null,(0,a.kt)("li",{parentName:"ul"},(0,a.kt)("a",{parentName:"li",href:"https://github.com/cosmos/ics"},"ICS")),(0,a.kt)("li",{parentName:"ul"},(0,a.kt)("a",{parentName:"li",href:"https://github.com/cosmos/ics/blob/master/ibc/1_IBC_ARCHITECTURE.md"},"IBC Architecture")),(0,a.kt)("li",{parentName:"ul"},(0,a.kt)("a",{parentName:"li",href:"https://github.com/tendermint/spec/blob/7b3138e69490f410768d9b1ffc7a17abc23ea397/spec/consensus/fork-accountability.md"},"Tendermint Fork Accountability"))))}u.isMDXComponent=!0}}]);