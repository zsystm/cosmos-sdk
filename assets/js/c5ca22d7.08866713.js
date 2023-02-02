"use strict";(self.webpackChunkcosmos_sdk_docs=self.webpackChunkcosmos_sdk_docs||[]).push([[1917],{3905:(e,t,n)=>{n.d(t,{Zo:()=>d,kt:()=>m});var a=n(7294);function o(e,t,n){return t in e?Object.defineProperty(e,t,{value:n,enumerable:!0,configurable:!0,writable:!0}):e[t]=n,e}function i(e,t){var n=Object.keys(e);if(Object.getOwnPropertySymbols){var a=Object.getOwnPropertySymbols(e);t&&(a=a.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),n.push.apply(n,a)}return n}function r(e){for(var t=1;t<arguments.length;t++){var n=null!=arguments[t]?arguments[t]:{};t%2?i(Object(n),!0).forEach((function(t){o(e,t,n[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(n)):i(Object(n)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(n,t))}))}return e}function l(e,t){if(null==e)return{};var n,a,o=function(e,t){if(null==e)return{};var n,a,o={},i=Object.keys(e);for(a=0;a<i.length;a++)n=i[a],t.indexOf(n)>=0||(o[n]=e[n]);return o}(e,t);if(Object.getOwnPropertySymbols){var i=Object.getOwnPropertySymbols(e);for(a=0;a<i.length;a++)n=i[a],t.indexOf(n)>=0||Object.prototype.propertyIsEnumerable.call(e,n)&&(o[n]=e[n])}return o}var s=a.createContext({}),u=function(e){var t=a.useContext(s),n=t;return e&&(n="function"==typeof e?e(t):r(r({},t),e)),n},d=function(e){var t=u(e.components);return a.createElement(s.Provider,{value:t},e.children)},p={inlineCode:"code",wrapper:function(e){var t=e.children;return a.createElement(a.Fragment,{},t)}},c=a.forwardRef((function(e,t){var n=e.components,o=e.mdxType,i=e.originalType,s=e.parentName,d=l(e,["components","mdxType","originalType","parentName"]),c=u(n),m=o,k=c["".concat(s,".").concat(m)]||c[m]||p[m]||i;return n?a.createElement(k,r(r({ref:t},d),{},{components:n})):a.createElement(k,r({ref:t},d))}));function m(e,t){var n=arguments,o=t&&t.mdxType;if("string"==typeof e||o){var i=n.length,r=new Array(i);r[0]=c;var l={};for(var s in t)hasOwnProperty.call(t,s)&&(l[s]=t[s]);l.originalType=e,l.mdxType="string"==typeof e?e:o,r[1]=l;for(var u=2;u<i;u++)r[u]=n[u];return a.createElement.apply(null,r)}return a.createElement.apply(null,n)}c.displayName="MDXCreateElement"},1007:(e,t,n)=>{n.r(t),n.d(t,{assets:()=>s,contentTitle:()=>r,default:()=>p,frontMatter:()=>i,metadata:()=>l,toc:()=>u});var a=n(7462),o=(n(7294),n(3905));const i={},r="ADR 016: Validator Consensus Key Rotation",l={unversionedId:"architecture/adr-016-validator-consensus-key-rotation",id:"version-v0.47/architecture/adr-016-validator-consensus-key-rotation",title:"ADR 016: Validator Consensus Key Rotation",description:"Changelog",source:"@site/versioned_docs/version-v0.47/architecture/adr-016-validator-consensus-key-rotation.md",sourceDirName:"architecture",slug:"/architecture/adr-016-validator-consensus-key-rotation",permalink:"/v0.47/architecture/adr-016-validator-consensus-key-rotation",draft:!1,tags:[],version:"v0.47",frontMatter:{},sidebar:"tutorialSidebar",previous:{title:"ADR 14: Proportional Slashing",permalink:"/v0.47/architecture/adr-014-proportional-slashing"},next:{title:"ADR 17: Historical Header Module",permalink:"/v0.47/architecture/adr-017-historical-header-module"}},s={},u=[{value:"Changelog",id:"changelog",level:2},{value:"Context",id:"context",level:2},{value:"Decision",id:"decision",level:2},{value:"Pseudo procedure for consensus key rotation",id:"pseudo-procedure-for-consensus-key-rotation",level:3},{value:"Considerations",id:"considerations",level:3},{value:"Workflow",id:"workflow",level:3},{value:"Status",id:"status",level:2},{value:"Consequences",id:"consequences",level:2},{value:"Positive",id:"positive",level:3},{value:"Negative",id:"negative",level:3},{value:"Neutral",id:"neutral",level:3},{value:"References",id:"references",level:2}],d={toc:u};function p(e){let{components:t,...n}=e;return(0,o.kt)("wrapper",(0,a.Z)({},d,n,{components:t,mdxType:"MDXLayout"}),(0,o.kt)("h1",{id:"adr-016-validator-consensus-key-rotation"},"ADR 016: Validator Consensus Key Rotation"),(0,o.kt)("h2",{id:"changelog"},"Changelog"),(0,o.kt)("ul",null,(0,o.kt)("li",{parentName:"ul"},"2019 Oct 23: Initial draft"),(0,o.kt)("li",{parentName:"ul"},"2019 Nov 28: Add key rotation fee")),(0,o.kt)("h2",{id:"context"},"Context"),(0,o.kt)("p",null,"Validator consensus key rotation feature has been discussed and requested for a long time, for the sake of safer validator key management policy (e.g. ",(0,o.kt)("a",{parentName:"p",href:"https://github.com/tendermint/tendermint/issues/1136"},"https://github.com/tendermint/tendermint/issues/1136"),"). So, we suggest one of the simplest form of validator consensus key rotation implementation mostly onto Cosmos SDK."),(0,o.kt)("p",null,"We don't need to make any update on consensus logic in Tendermint because Tendermint does not have any mapping information of consensus key and validator operator key, meaning that from Tendermint point of view, a consensus key rotation of a validator is simply a replacement of a consensus key to another."),(0,o.kt)("p",null,"Also, it should be noted that this ADR includes only the simplest form of consensus key rotation without considering multiple consensus keys concept. Such multiple consensus keys concept shall remain a long term goal of Tendermint and Cosmos SDK."),(0,o.kt)("h2",{id:"decision"},"Decision"),(0,o.kt)("h3",{id:"pseudo-procedure-for-consensus-key-rotation"},"Pseudo procedure for consensus key rotation"),(0,o.kt)("ul",null,(0,o.kt)("li",{parentName:"ul"},"create new random consensus key."),(0,o.kt)("li",{parentName:"ul"},"create and broadcast a transaction with a ",(0,o.kt)("inlineCode",{parentName:"li"},"MsgRotateConsPubKey")," that states the new consensus key is now coupled with the validator operator with signature from the validator's operator key."),(0,o.kt)("li",{parentName:"ul"},"old consensus key becomes unable to participate on consensus immediately after the update of key mapping state on-chain."),(0,o.kt)("li",{parentName:"ul"},"start validating with new consensus key."),(0,o.kt)("li",{parentName:"ul"},"validators using HSM and KMS should update the consensus key in HSM to use the new rotated key after the height ",(0,o.kt)("inlineCode",{parentName:"li"},"h")," when ",(0,o.kt)("inlineCode",{parentName:"li"},"MsgRotateConsPubKey")," committed to the blockchain.")),(0,o.kt)("h3",{id:"considerations"},"Considerations"),(0,o.kt)("ul",null,(0,o.kt)("li",{parentName:"ul"},"consensus key mapping information management strategy",(0,o.kt)("ul",{parentName:"li"},(0,o.kt)("li",{parentName:"ul"},"store history of each key mapping changes in the kvstore."),(0,o.kt)("li",{parentName:"ul"},"the state machine can search corresponding consensus key paired with given validator operator for any arbitrary height in a recent unbonding period."),(0,o.kt)("li",{parentName:"ul"},"the state machine does not need any historical mapping information which is past more than unbonding period."))),(0,o.kt)("li",{parentName:"ul"},"key rotation costs related to LCD and IBC",(0,o.kt)("ul",{parentName:"li"},(0,o.kt)("li",{parentName:"ul"},"LCD and IBC will have traffic/computation burden when there exists frequent power changes"),(0,o.kt)("li",{parentName:"ul"},"In current Tendermint design, consensus key rotations are seen as power changes from LCD or IBC perspective"),(0,o.kt)("li",{parentName:"ul"},"Therefore, to minimize unnecessary frequent key rotation behavior, we limited maximum number of rotation in recent unbonding period and also applied exponentially increasing rotation fee"))),(0,o.kt)("li",{parentName:"ul"},"limits",(0,o.kt)("ul",{parentName:"li"},(0,o.kt)("li",{parentName:"ul"},"a validator cannot rotate its consensus key more than ",(0,o.kt)("inlineCode",{parentName:"li"},"MaxConsPubKeyRotations")," time for any unbonding period, to prevent spam."),(0,o.kt)("li",{parentName:"ul"},"parameters can be decided by governance and stored in genesis file."))),(0,o.kt)("li",{parentName:"ul"},"key rotation fee",(0,o.kt)("ul",{parentName:"li"},(0,o.kt)("li",{parentName:"ul"},"a validator should pay ",(0,o.kt)("inlineCode",{parentName:"li"},"KeyRotationFee")," to rotate the consensus key which is calculated as below"),(0,o.kt)("li",{parentName:"ul"},(0,o.kt)("inlineCode",{parentName:"li"},"KeyRotationFee")," = (max(",(0,o.kt)("inlineCode",{parentName:"li"},"VotingPowerPercentage")," ",(0,o.kt)("em",{parentName:"li"},"100, 1)")," ",(0,o.kt)("inlineCode",{parentName:"li"},"InitialKeyRotationFee"),") * 2^(number of rotations in ",(0,o.kt)("inlineCode",{parentName:"li"},"ConsPubKeyRotationHistory")," in recent unbonding period)"))),(0,o.kt)("li",{parentName:"ul"},"evidence module",(0,o.kt)("ul",{parentName:"li"},(0,o.kt)("li",{parentName:"ul"},"evidence module can search corresponding consensus key for any height from slashing keeper so that it can decide which consensus key is supposed to be used for given height."))),(0,o.kt)("li",{parentName:"ul"},"abci.ValidatorUpdate",(0,o.kt)("ul",{parentName:"li"},(0,o.kt)("li",{parentName:"ul"},"tendermint already has ability to change a consensus key by ABCI communication(",(0,o.kt)("inlineCode",{parentName:"li"},"ValidatorUpdate"),")."),(0,o.kt)("li",{parentName:"ul"},"validator consensus key update can be done via creating new + delete old by change the power to zero."),(0,o.kt)("li",{parentName:"ul"},"therefore, we expect we even do not need to change tendermint codebase at all to implement this feature."))),(0,o.kt)("li",{parentName:"ul"},"new genesis parameters in ",(0,o.kt)("inlineCode",{parentName:"li"},"staking")," module",(0,o.kt)("ul",{parentName:"li"},(0,o.kt)("li",{parentName:"ul"},(0,o.kt)("inlineCode",{parentName:"li"},"MaxConsPubKeyRotations")," : maximum number of rotation can be executed by a validator in recent unbonding period. default value 10 is suggested(11th key rotation will be rejected)"),(0,o.kt)("li",{parentName:"ul"},(0,o.kt)("inlineCode",{parentName:"li"},"InitialKeyRotationFee")," : the initial key rotation fee when no key rotation has happened in recent unbonding period. default value 1atom is suggested(1atom fee for the first key rotation in recent unbonding period)")))),(0,o.kt)("h3",{id:"workflow"},"Workflow"),(0,o.kt)("ol",null,(0,o.kt)("li",{parentName:"ol"},(0,o.kt)("p",{parentName:"li"},"The validator generates a new consensus keypair.")),(0,o.kt)("li",{parentName:"ol"},(0,o.kt)("p",{parentName:"li"},"The validator generates and signs a ",(0,o.kt)("inlineCode",{parentName:"p"},"MsgRotateConsPubKey")," tx with their operator key and new ConsPubKey"),(0,o.kt)("pre",{parentName:"li"},(0,o.kt)("code",{parentName:"pre",className:"language-go"},"type MsgRotateConsPubKey struct {\n    ValidatorAddress  sdk.ValAddress\n    NewPubKey         crypto.PubKey\n}\n"))),(0,o.kt)("li",{parentName:"ol"},(0,o.kt)("p",{parentName:"li"},(0,o.kt)("inlineCode",{parentName:"p"},"handleMsgRotateConsPubKey")," gets ",(0,o.kt)("inlineCode",{parentName:"p"},"MsgRotateConsPubKey"),", calls ",(0,o.kt)("inlineCode",{parentName:"p"},"RotateConsPubKey")," with emits event")),(0,o.kt)("li",{parentName:"ol"},(0,o.kt)("p",{parentName:"li"},(0,o.kt)("inlineCode",{parentName:"p"},"RotateConsPubKey")),(0,o.kt)("ul",{parentName:"li"},(0,o.kt)("li",{parentName:"ul"},"checks if ",(0,o.kt)("inlineCode",{parentName:"li"},"NewPubKey")," is not duplicated on ",(0,o.kt)("inlineCode",{parentName:"li"},"ValidatorsByConsAddr")),(0,o.kt)("li",{parentName:"ul"},"checks if the validator is does not exceed parameter ",(0,o.kt)("inlineCode",{parentName:"li"},"MaxConsPubKeyRotations")," by iterating ",(0,o.kt)("inlineCode",{parentName:"li"},"ConsPubKeyRotationHistory")),(0,o.kt)("li",{parentName:"ul"},"checks if the signing account has enough balance to pay ",(0,o.kt)("inlineCode",{parentName:"li"},"KeyRotationFee")),(0,o.kt)("li",{parentName:"ul"},"pays ",(0,o.kt)("inlineCode",{parentName:"li"},"KeyRotationFee")," to community fund"),(0,o.kt)("li",{parentName:"ul"},"overwrites ",(0,o.kt)("inlineCode",{parentName:"li"},"NewPubKey")," in ",(0,o.kt)("inlineCode",{parentName:"li"},"validator.ConsPubKey")),(0,o.kt)("li",{parentName:"ul"},"deletes old ",(0,o.kt)("inlineCode",{parentName:"li"},"ValidatorByConsAddr")),(0,o.kt)("li",{parentName:"ul"},(0,o.kt)("inlineCode",{parentName:"li"},"SetValidatorByConsAddr")," for ",(0,o.kt)("inlineCode",{parentName:"li"},"NewPubKey")),(0,o.kt)("li",{parentName:"ul"},"Add ",(0,o.kt)("inlineCode",{parentName:"li"},"ConsPubKeyRotationHistory")," for tracking rotation")),(0,o.kt)("pre",{parentName:"li"},(0,o.kt)("code",{parentName:"pre",className:"language-go"},"type ConsPubKeyRotationHistory struct {\n    OperatorAddress         sdk.ValAddress\n    OldConsPubKey           crypto.PubKey\n    NewConsPubKey           crypto.PubKey\n    RotatedHeight           int64\n}\n"))),(0,o.kt)("li",{parentName:"ol"},(0,o.kt)("p",{parentName:"li"},(0,o.kt)("inlineCode",{parentName:"p"},"ApplyAndReturnValidatorSetUpdates")," checks if there is ",(0,o.kt)("inlineCode",{parentName:"p"},"ConsPubKeyRotationHistory")," with ",(0,o.kt)("inlineCode",{parentName:"p"},"ConsPubKeyRotationHistory.RotatedHeight == ctx.BlockHeight()")," and if so, generates 2 ",(0,o.kt)("inlineCode",{parentName:"p"},"ValidatorUpdate")," , one for a remove validator and one for create new validator"),(0,o.kt)("pre",{parentName:"li"},(0,o.kt)("code",{parentName:"pre",className:"language-go"},"abci.ValidatorUpdate{\n    PubKey: tmtypes.TM2PB.PubKey(OldConsPubKey),\n    Power:  0,\n}\n\nabci.ValidatorUpdate{\n    PubKey: tmtypes.TM2PB.PubKey(NewConsPubKey),\n    Power:  v.ConsensusPower(),\n}\n"))),(0,o.kt)("li",{parentName:"ol"},(0,o.kt)("p",{parentName:"li"},"at ",(0,o.kt)("inlineCode",{parentName:"p"},"previousVotes")," Iteration logic of ",(0,o.kt)("inlineCode",{parentName:"p"},"AllocateTokens"),",  ",(0,o.kt)("inlineCode",{parentName:"p"},"previousVote")," using ",(0,o.kt)("inlineCode",{parentName:"p"},"OldConsPubKey")," match up with ",(0,o.kt)("inlineCode",{parentName:"p"},"ConsPubKeyRotationHistory"),", and replace validator for token allocation")),(0,o.kt)("li",{parentName:"ol"},(0,o.kt)("p",{parentName:"li"},"Migrate ",(0,o.kt)("inlineCode",{parentName:"p"},"ValidatorSigningInfo")," and ",(0,o.kt)("inlineCode",{parentName:"p"},"ValidatorMissedBlockBitArray")," from ",(0,o.kt)("inlineCode",{parentName:"p"},"OldConsPubKey")," to ",(0,o.kt)("inlineCode",{parentName:"p"},"NewConsPubKey")))),(0,o.kt)("ul",null,(0,o.kt)("li",{parentName:"ul"},"Note : All above features shall be implemented in ",(0,o.kt)("inlineCode",{parentName:"li"},"staking")," module.")),(0,o.kt)("h2",{id:"status"},"Status"),(0,o.kt)("p",null,"Proposed"),(0,o.kt)("h2",{id:"consequences"},"Consequences"),(0,o.kt)("h3",{id:"positive"},"Positive"),(0,o.kt)("ul",null,(0,o.kt)("li",{parentName:"ul"},"Validators can immediately or periodically rotate their consensus key to have better security policy"),(0,o.kt)("li",{parentName:"ul"},"improved security against Long-Range attacks (",(0,o.kt)("a",{parentName:"li",href:"https://nearprotocol.com/blog/long-range-attacks-and-a-new-fork-choice-rule"},"https://nearprotocol.com/blog/long-range-attacks-and-a-new-fork-choice-rule"),") given a validator throws away the old consensus key(s)")),(0,o.kt)("h3",{id:"negative"},"Negative"),(0,o.kt)("ul",null,(0,o.kt)("li",{parentName:"ul"},"Slash module needs more computation because it needs to lookup corresponding consensus key of validators for each height"),(0,o.kt)("li",{parentName:"ul"},"frequent key rotations will make light client bisection less efficient")),(0,o.kt)("h3",{id:"neutral"},"Neutral"),(0,o.kt)("h2",{id:"references"},"References"),(0,o.kt)("ul",null,(0,o.kt)("li",{parentName:"ul"},"on tendermint repo : ",(0,o.kt)("a",{parentName:"li",href:"https://github.com/tendermint/tendermint/issues/1136"},"https://github.com/tendermint/tendermint/issues/1136")),(0,o.kt)("li",{parentName:"ul"},"on cosmos-sdk repo : ",(0,o.kt)("a",{parentName:"li",href:"https://github.com/cosmos/cosmos-sdk/issues/5231"},"https://github.com/cosmos/cosmos-sdk/issues/5231")),(0,o.kt)("li",{parentName:"ul"},"about multiple consensus keys : ",(0,o.kt)("a",{parentName:"li",href:"https://github.com/tendermint/tendermint/issues/1758#issuecomment-545291698"},"https://github.com/tendermint/tendermint/issues/1758#issuecomment-545291698"))))}p.isMDXComponent=!0}}]);