"use strict";(self.webpackChunkcosmos_sdk_docs=self.webpackChunkcosmos_sdk_docs||[]).push([[1822],{3905:(e,t,n)=>{n.d(t,{Zo:()=>u,kt:()=>m});var i=n(7294);function a(e,t,n){return t in e?Object.defineProperty(e,t,{value:n,enumerable:!0,configurable:!0,writable:!0}):e[t]=n,e}function o(e,t){var n=Object.keys(e);if(Object.getOwnPropertySymbols){var i=Object.getOwnPropertySymbols(e);t&&(i=i.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),n.push.apply(n,i)}return n}function r(e){for(var t=1;t<arguments.length;t++){var n=null!=arguments[t]?arguments[t]:{};t%2?o(Object(n),!0).forEach((function(t){a(e,t,n[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(n)):o(Object(n)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(n,t))}))}return e}function l(e,t){if(null==e)return{};var n,i,a=function(e,t){if(null==e)return{};var n,i,a={},o=Object.keys(e);for(i=0;i<o.length;i++)n=o[i],t.indexOf(n)>=0||(a[n]=e[n]);return a}(e,t);if(Object.getOwnPropertySymbols){var o=Object.getOwnPropertySymbols(e);for(i=0;i<o.length;i++)n=o[i],t.indexOf(n)>=0||Object.prototype.propertyIsEnumerable.call(e,n)&&(a[n]=e[n])}return a}var c=i.createContext({}),s=function(e){var t=i.useContext(c),n=t;return e&&(n="function"==typeof e?e(t):r(r({},t),e)),n},u=function(e){var t=s(e.components);return i.createElement(c.Provider,{value:t},e.children)},p={inlineCode:"code",wrapper:function(e){var t=e.children;return i.createElement(i.Fragment,{},t)}},d=i.forwardRef((function(e,t){var n=e.components,a=e.mdxType,o=e.originalType,c=e.parentName,u=l(e,["components","mdxType","originalType","parentName"]),d=s(n),m=a,h=d["".concat(c,".").concat(m)]||d[m]||p[m]||o;return n?i.createElement(h,r(r({ref:t},u),{},{components:n})):i.createElement(h,r({ref:t},u))}));function m(e,t){var n=arguments,a=t&&t.mdxType;if("string"==typeof e||a){var o=n.length,r=new Array(o);r[0]=d;var l={};for(var c in t)hasOwnProperty.call(t,c)&&(l[c]=t[c]);l.originalType=e,l.mdxType="string"==typeof e?e:a,r[1]=l;for(var s=2;s<o;s++)r[s]=n[s];return i.createElement.apply(null,r)}return i.createElement.apply(null,n)}d.displayName="MDXCreateElement"},6967:(e,t,n)=>{n.r(t),n.d(t,{assets:()=>c,contentTitle:()=>r,default:()=>p,frontMatter:()=>o,metadata:()=>l,toc:()=>s});var i=n(7462),a=(n(7294),n(3905));const o={},r="ADR 007: Specialization Groups",l={unversionedId:"architecture/adr-007-specialization-groups",id:"version-v0.47/architecture/adr-007-specialization-groups",title:"ADR 007: Specialization Groups",description:"Changelog",source:"@site/versioned_docs/version-v0.47/architecture/adr-007-specialization-groups.md",sourceDirName:"architecture",slug:"/architecture/adr-007-specialization-groups",permalink:"/v0.47/architecture/adr-007-specialization-groups",draft:!1,tags:[],version:"v0.47",frontMatter:{},sidebar:"tutorialSidebar",previous:{title:"ADR 006: Secret Store Replacement",permalink:"/v0.47/architecture/adr-006-secret-store-replacement"},next:{title:"ADR 008: Decentralized Computer Emergency Response Team (dCERT) Group",permalink:"/v0.47/architecture/adr-008-dCERT-group"}},c={},s=[{value:"Changelog",id:"changelog",level:2},{value:"Context",id:"context",level:2},{value:"Decision",id:"decision",level:2},{value:"Status",id:"status",level:2},{value:"Consequences",id:"consequences",level:2},{value:"Positive",id:"positive",level:3},{value:"Negative",id:"negative",level:3},{value:"Neutral",id:"neutral",level:3},{value:"References",id:"references",level:2}],u={toc:s};function p(e){let{components:t,...n}=e;return(0,a.kt)("wrapper",(0,i.Z)({},u,n,{components:t,mdxType:"MDXLayout"}),(0,a.kt)("h1",{id:"adr-007-specialization-groups"},"ADR 007: Specialization Groups"),(0,a.kt)("h2",{id:"changelog"},"Changelog"),(0,a.kt)("ul",null,(0,a.kt)("li",{parentName:"ul"},"2019 Jul 31: Initial Draft")),(0,a.kt)("h2",{id:"context"},"Context"),(0,a.kt)("p",null,'This idea was first conceived of in order to fulfill the use case of the\ncreation of a decentralized Computer Emergency Response Team (dCERT), whose\nmembers would be elected by a governing community and would fulfill the role of\ncoordinating the community under emergency situations. This thinking\ncan be further abstracted into the conception of "blockchain specialization\ngroups".'),(0,a.kt)("p",null,"The creation of these groups are the beginning of specialization capabilities\nwithin a wider blockchain community which could be used to enable a certain\nlevel of delegated responsibilities. Examples of specialization which could be\nbeneficial to a blockchain community include: code auditing, emergency response,\ncode development etc. This type of community organization paves the way for\nindividual stakeholders to delegate votes by issue type, if in the future\ngovernance proposals include a field for issue type."),(0,a.kt)("h2",{id:"decision"},"Decision"),(0,a.kt)("p",null,"A specialization group can be broadly broken down into the following functions\n(herein containing examples):"),(0,a.kt)("ul",null,(0,a.kt)("li",{parentName:"ul"},"Membership Admittance"),(0,a.kt)("li",{parentName:"ul"},"Membership Acceptance"),(0,a.kt)("li",{parentName:"ul"},"Membership Revocation",(0,a.kt)("ul",{parentName:"li"},(0,a.kt)("li",{parentName:"ul"},"(probably) Without Penalty",(0,a.kt)("ul",{parentName:"li"},(0,a.kt)("li",{parentName:"ul"},"member steps down (self-Revocation)"),(0,a.kt)("li",{parentName:"ul"},"replaced by new member from governance"))),(0,a.kt)("li",{parentName:"ul"},"(probably) With Penalty",(0,a.kt)("ul",{parentName:"li"},(0,a.kt)("li",{parentName:"ul"},"due to breach of soft-agreement (determined through governance)"),(0,a.kt)("li",{parentName:"ul"},"due to breach of hard-agreement (determined by code)"))))),(0,a.kt)("li",{parentName:"ul"},"Execution of Duties",(0,a.kt)("ul",{parentName:"li"},(0,a.kt)("li",{parentName:"ul"},"Special transactions which only execute for members of a specialization\ngroup (for example, dCERT members voting to turn off transaction routes in\nan emergency scenario)"))),(0,a.kt)("li",{parentName:"ul"},"Compensation",(0,a.kt)("ul",{parentName:"li"},(0,a.kt)("li",{parentName:"ul"},"Group compensation (further distribution decided by the specialization group)"),(0,a.kt)("li",{parentName:"ul"},"Individual compensation for all constituents of a group from the\ngreater community")))),(0,a.kt)("p",null,"Membership admittance to a specialization group could take place over a wide\nvariety of mechanisms. The most obvious example is through a general vote among\nthe entire community, however in certain systems a community may want to allow\nthe members already in a specialization group to internally elect new members,\nor maybe the community may assign a permission to a particular specialization\ngroup to appoint members to other 3rd party groups. The sky is really the limit\nas to how membership admittance can be structured. We attempt to capture\nsome of these possiblities in a common interface dubbed the ",(0,a.kt)("inlineCode",{parentName:"p"},"Electionator"),". For\nits initial implementation as a part of this ADR we recommend that the general\nelection abstraction (",(0,a.kt)("inlineCode",{parentName:"p"},"Electionator"),") is provided as well as a basic\nimplementation of that abstraction which allows for a continuous election of\nmembers of a specialization group."),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-golang"},'// The Electionator abstraction covers the concept space for\n// a wide variety of election kinds.  \ntype Electionator interface {\n\n    // is the election object accepting votes.\n    Active() bool\n\n    // functionality to execute for when a vote is cast in this election, here\n    // the vote field is anticipated to be marshalled into a vote type used\n    // by an election.\n    //\n    // NOTE There are no explicit ids here. Just votes which pertain specifically\n    // to one electionator. Anyone can create and send a vote to the electionator item\n    // which will presumably attempt to marshal those bytes into a particular struct\n    // and apply the vote information in some arbitrary way. There can be multiple\n    // Electionators within the Cosmos-Hub for multiple specialization groups, votes\n    // would need to be routed to the Electionator upstream of here.\n    Vote(addr sdk.AccAddress, vote []byte)\n\n    // here lies all functionality to authenticate and execute changes for\n    // when a member accepts being elected\n    AcceptElection(sdk.AccAddress)\n\n    // Register a revoker object\n    RegisterRevoker(Revoker)\n\n    // No more revokers may be registered after this function is called\n    SealRevokers()\n\n    // register hooks to call when an election actions occur\n    RegisterHooks(ElectionatorHooks)\n\n    // query for the current winner(s) of this election based on arbitrary\n    // election ruleset\n    QueryElected() []sdk.AccAddress\n\n    // query metadata for an address in the election this\n    // could include for example position that an address\n    // is being elected for within a group\n    //\n    // this metadata may be directly related to\n    // voting information and/or privileges enabled\n    // to members within a group.\n    QueryMetadata(sdk.AccAddress) []byte\n}\n\n// ElectionatorHooks, once registered with an Electionator,\n// trigger execution of relevant interface functions when\n// Electionator events occur.\ntype ElectionatorHooks interface {\n    AfterVoteCast(addr sdk.AccAddress, vote []byte)\n    AfterMemberAccepted(addr sdk.AccAddress)\n    AfterMemberRevoked(addr sdk.AccAddress, cause []byte)\n}\n\n// Revoker defines the function required for a membership revocation rule-set\n// used by a specialization group. This could be used to create self revoking,\n// and evidence based revoking, etc. Revokers types may be created and\n// reused for different election types.\n//\n// When revoking the "cause" bytes may be arbitrarily marshalled into evidence,\n// memos, etc.\ntype Revoker interface {\n    RevokeName() string      // identifier for this revoker type\n    RevokeMember(addr sdk.AccAddress, cause []byte) error\n}\n')),(0,a.kt)("p",null,"Certain level of commonality likely exists between the existing code within\n",(0,a.kt)("inlineCode",{parentName:"p"},"x/governance")," and required functionality of elections. This common\nfunctionality should be abstracted during implementation. Similarly for each\nvote implementation client CLI/REST functionality should be abstracted\nto be reused for multiple elections."),(0,a.kt)("p",null,"The specialization group abstraction firstly extends the ",(0,a.kt)("inlineCode",{parentName:"p"},"Electionator"),"\nbut also further defines traits of the group."),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-golang"},"type SpecializationGroup interface {\n    Electionator\n    GetName() string\n    GetDescription() string\n\n    // general soft contract the group is expected\n    // to fulfill with the greater community\n    GetContract() string\n\n    // messages which can be executed by the members of the group\n    Handler(ctx sdk.Context, msg sdk.Msg) sdk.Result\n\n    // logic to be executed at endblock, this may for instance\n    // include payment of a stipend to the group members\n    // for participation in the security group.\n    EndBlocker(ctx sdk.Context)\n}\n")),(0,a.kt)("h2",{id:"status"},"Status"),(0,a.kt)("blockquote",null,(0,a.kt)("p",{parentName:"blockquote"},"Proposed")),(0,a.kt)("h2",{id:"consequences"},"Consequences"),(0,a.kt)("h3",{id:"positive"},"Positive"),(0,a.kt)("ul",null,(0,a.kt)("li",{parentName:"ul"},"increases specialization capabilities of a blockchain"),(0,a.kt)("li",{parentName:"ul"},"improve abstractions in ",(0,a.kt)("inlineCode",{parentName:"li"},"x/gov/")," such that they can be used with specialization groups")),(0,a.kt)("h3",{id:"negative"},"Negative"),(0,a.kt)("ul",null,(0,a.kt)("li",{parentName:"ul"},"could be used to increase centralization within a community")),(0,a.kt)("h3",{id:"neutral"},"Neutral"),(0,a.kt)("h2",{id:"references"},"References"),(0,a.kt)("ul",null,(0,a.kt)("li",{parentName:"ul"},(0,a.kt)("a",{parentName:"li",href:"/v0.47/architecture/adr-008-dCERT-group"},"dCERT ADR"))))}p.isMDXComponent=!0}}]);