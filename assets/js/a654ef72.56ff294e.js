"use strict";(self.webpackChunkcosmos_sdk_docs=self.webpackChunkcosmos_sdk_docs||[]).push([[5181],{3905:(e,t,n)=>{n.d(t,{Zo:()=>c,kt:()=>h});var a=n(7294);function i(e,t,n){return t in e?Object.defineProperty(e,t,{value:n,enumerable:!0,configurable:!0,writable:!0}):e[t]=n,e}function o(e,t){var n=Object.keys(e);if(Object.getOwnPropertySymbols){var a=Object.getOwnPropertySymbols(e);t&&(a=a.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),n.push.apply(n,a)}return n}function r(e){for(var t=1;t<arguments.length;t++){var n=null!=arguments[t]?arguments[t]:{};t%2?o(Object(n),!0).forEach((function(t){i(e,t,n[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(n)):o(Object(n)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(n,t))}))}return e}function s(e,t){if(null==e)return{};var n,a,i=function(e,t){if(null==e)return{};var n,a,i={},o=Object.keys(e);for(a=0;a<o.length;a++)n=o[a],t.indexOf(n)>=0||(i[n]=e[n]);return i}(e,t);if(Object.getOwnPropertySymbols){var o=Object.getOwnPropertySymbols(e);for(a=0;a<o.length;a++)n=o[a],t.indexOf(n)>=0||Object.prototype.propertyIsEnumerable.call(e,n)&&(i[n]=e[n])}return i}var l=a.createContext({}),p=function(e){var t=a.useContext(l),n=t;return e&&(n="function"==typeof e?e(t):r(r({},t),e)),n},c=function(e){var t=p(e.components);return a.createElement(l.Provider,{value:t},e.children)},d={inlineCode:"code",wrapper:function(e){var t=e.children;return a.createElement(a.Fragment,{},t)}},m=a.forwardRef((function(e,t){var n=e.components,i=e.mdxType,o=e.originalType,l=e.parentName,c=s(e,["components","mdxType","originalType","parentName"]),m=p(n),h=i,u=m["".concat(l,".").concat(h)]||m[h]||d[h]||o;return n?a.createElement(u,r(r({ref:t},c),{},{components:n})):a.createElement(u,r({ref:t},c))}));function h(e,t){var n=arguments,i=t&&t.mdxType;if("string"==typeof e||i){var o=n.length,r=new Array(o);r[0]=m;var s={};for(var l in t)hasOwnProperty.call(t,l)&&(s[l]=t[l]);s.originalType=e,s.mdxType="string"==typeof e?e:i,r[1]=s;for(var p=2;p<o;p++)r[p]=n[p];return a.createElement.apply(null,r)}return a.createElement.apply(null,n)}m.displayName="MDXCreateElement"},1328:(e,t,n)=>{n.r(t),n.d(t,{assets:()=>l,contentTitle:()=>r,default:()=>d,frontMatter:()=>o,metadata:()=>s,toc:()=>p});var a=n(7462),i=(n(7294),n(3905));const o={sidebar_position:1},r="Transactions",s={unversionedId:"core/transactions",id:"core/transactions",title:"Transactions",description:"Transactions are objects created by end-users to trigger state changes in the application.",source:"@site/docs/core/01-transactions.md",sourceDirName:"core",slug:"/core/transactions",permalink:"/main/core/transactions",draft:!1,tags:[],version:"current",sidebarPosition:1,frontMatter:{sidebar_position:1},sidebar:"tutorialSidebar",previous:{title:"BaseApp",permalink:"/main/core/baseapp"},next:{title:"Context",permalink:"/main/core/context"}},l={},p=[{value:"Transactions",id:"transactions-1",level:2},{value:"Type Definition",id:"type-definition",level:2},{value:"Signing Transactions",id:"signing-transactions",level:3},{value:"<code>SIGN_MODE_DIRECT</code> (preferred)",id:"sign_mode_direct-preferred",level:4},{value:"<code>SIGN_MODE_LEGACY_AMINO_JSON</code>",id:"sign_mode_legacy_amino_json",level:4},{value:"Other Sign Modes",id:"other-sign-modes",level:4},{value:"<code>SIGN_MODE_DIRECT_AUX</code>",id:"sign_mode_direct_aux",level:4},{value:"<code>SIGN_MODE_TEXTUAL</code>",id:"sign_mode_textual",level:4},{value:"Custom Sign modes",id:"custom-sign-modes",level:4},{value:"Transaction Process",id:"transaction-process",level:2},{value:"Messages",id:"messages",level:3},{value:"Transaction Generation",id:"transaction-generation",level:3},{value:"Broadcasting the Transaction",id:"broadcasting-the-transaction",level:3},{value:"CLI",id:"cli",level:4},{value:"gRPC",id:"grpc",level:4},{value:"REST",id:"rest",level:4},{value:"Tendermint RPC",id:"tendermint-rpc",level:4}],c={toc:p};function d(e){let{components:t,...n}=e;return(0,i.kt)("wrapper",(0,a.Z)({},c,n,{components:t,mdxType:"MDXLayout"}),(0,i.kt)("h1",{id:"transactions"},"Transactions"),(0,i.kt)("admonition",{title:"Synopsis",type:"note"},(0,i.kt)("p",{parentName:"admonition"},(0,i.kt)("inlineCode",{parentName:"p"},"Transactions")," are objects created by end-users to trigger state changes in the application.")),(0,i.kt)("admonition",{type:"note"},(0,i.kt)("h3",{parentName:"admonition",id:"pre-requisite-readings"},"Pre-requisite Readings"),(0,i.kt)("ul",{parentName:"admonition"},(0,i.kt)("li",{parentName:"ul"},(0,i.kt)("a",{parentName:"li",href:"/main/basics/app-anatomy"},"Anatomy of a Cosmos SDK Application")))),(0,i.kt)("h2",{id:"transactions-1"},"Transactions"),(0,i.kt)("p",null,"Transactions are comprised of metadata held in ",(0,i.kt)("a",{parentName:"p",href:"/main/core/context"},"contexts")," and ",(0,i.kt)("a",{parentName:"p",href:"/main/building-modules/messages-and-queries"},(0,i.kt)("inlineCode",{parentName:"a"},"sdk.Msg"),"s")," that trigger state changes within a module through the module's Protobuf ",(0,i.kt)("a",{parentName:"p",href:"/main/building-modules/msg-services"},(0,i.kt)("inlineCode",{parentName:"a"},"Msg")," service"),"."),(0,i.kt)("p",null,"When users want to interact with an application and make state changes (e.g. sending coins), they create transactions. Each of a transaction's ",(0,i.kt)("inlineCode",{parentName:"p"},"sdk.Msg")," must be signed using the private key associated with the appropriate account(s), before the transaction is broadcasted to the network. A transaction must then be included in a block, validated, and approved by the network through the consensus process. To read more about the lifecycle of a transaction, click ",(0,i.kt)("a",{parentName:"p",href:"/main/basics/tx-lifecycle"},"here"),"."),(0,i.kt)("h2",{id:"type-definition"},"Type Definition"),(0,i.kt)("p",null,"Transaction objects are Cosmos SDK types that implement the ",(0,i.kt)("inlineCode",{parentName:"p"},"Tx")," interface"),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-go",metastring:"reference",reference:!0},"https://github.com/cosmos/cosmos-sdk/blob/v0.47.0-rc1/types/tx_msg.go#L42-L50\n")),(0,i.kt)("p",null,"It contains the following methods:"),(0,i.kt)("ul",null,(0,i.kt)("li",{parentName:"ul"},(0,i.kt)("strong",{parentName:"li"},"GetMsgs:")," unwraps the transaction and returns a list of contained ",(0,i.kt)("inlineCode",{parentName:"li"},"sdk.Msg"),"s - one transaction may have one or multiple messages, which are defined by module developers."),(0,i.kt)("li",{parentName:"ul"},(0,i.kt)("strong",{parentName:"li"},"ValidateBasic:")," lightweight, ",(0,i.kt)("a",{parentName:"li",href:"/main/basics/tx-lifecycle#types-of-checks"},(0,i.kt)("em",{parentName:"a"},"stateless"))," checks used by ABCI messages ",(0,i.kt)("a",{parentName:"li",href:"/main/core/baseapp#checktx"},(0,i.kt)("inlineCode",{parentName:"a"},"CheckTx"))," and ",(0,i.kt)("a",{parentName:"li",href:"/main/core/baseapp#delivertx"},(0,i.kt)("inlineCode",{parentName:"a"},"DeliverTx"))," to make sure transactions are not invalid. For example, the ",(0,i.kt)("a",{parentName:"li",href:"https://github.com/cosmos/cosmos-sdk/tree/main/x/auth"},(0,i.kt)("inlineCode",{parentName:"a"},"auth"))," module's ",(0,i.kt)("inlineCode",{parentName:"li"},"ValidateBasic")," function checks that its transactions are signed by the correct number of signers and that the fees do not exceed what the user's maximum. Note that this function is to be distinct from ",(0,i.kt)("inlineCode",{parentName:"li"},"sdk.Msg")," ",(0,i.kt)("a",{parentName:"li",href:"/main/basics/tx-lifecycle#ValidateBasic"},(0,i.kt)("inlineCode",{parentName:"a"},"ValidateBasic"))," methods, which perform basic validity checks on messages only. When ",(0,i.kt)("a",{parentName:"li",href:"/main/core/baseapp#runtx"},(0,i.kt)("inlineCode",{parentName:"a"},"runTx"))," is checking a transaction created from the ",(0,i.kt)("a",{parentName:"li",href:"https://github.com/cosmos/cosmos-sdk/tree/main/x/auth/spec"},(0,i.kt)("inlineCode",{parentName:"a"},"auth"))," module, it first runs ",(0,i.kt)("inlineCode",{parentName:"li"},"ValidateBasic")," on each message, then runs the ",(0,i.kt)("inlineCode",{parentName:"li"},"auth")," module AnteHandler which calls ",(0,i.kt)("inlineCode",{parentName:"li"},"ValidateBasic")," for the transaction itself.")),(0,i.kt)("p",null,"As a developer, you should rarely manipulate ",(0,i.kt)("inlineCode",{parentName:"p"},"Tx")," directly, as ",(0,i.kt)("inlineCode",{parentName:"p"},"Tx")," is really an intermediate type used for transaction generation. Instead, developers should prefer the ",(0,i.kt)("inlineCode",{parentName:"p"},"TxBuilder")," interface, which you can learn more about ",(0,i.kt)("a",{parentName:"p",href:"#transaction-generation"},"below"),"."),(0,i.kt)("h3",{id:"signing-transactions"},"Signing Transactions"),(0,i.kt)("p",null,"Every message in a transaction must be signed by the addresses specified by its ",(0,i.kt)("inlineCode",{parentName:"p"},"GetSigners"),". The Cosmos SDK currently allows signing transactions in two different ways."),(0,i.kt)("h4",{id:"sign_mode_direct-preferred"},(0,i.kt)("inlineCode",{parentName:"h4"},"SIGN_MODE_DIRECT")," (preferred)"),(0,i.kt)("p",null,"The most used implementation of the ",(0,i.kt)("inlineCode",{parentName:"p"},"Tx")," interface is the Protobuf ",(0,i.kt)("inlineCode",{parentName:"p"},"Tx")," message, which is used in ",(0,i.kt)("inlineCode",{parentName:"p"},"SIGN_MODE_DIRECT"),":"),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-protobuf",metastring:"reference",reference:!0},"https://github.com/cosmos/cosmos-sdk/blob/v0.47.0-rc1/proto/cosmos/tx/v1beta1/tx.proto#L13-L26\n")),(0,i.kt)("p",null,"Because Protobuf serialization is not deterministic, the Cosmos SDK uses an additional ",(0,i.kt)("inlineCode",{parentName:"p"},"TxRaw")," type to denote the pinned bytes over which a transaction is signed. Any user can generate a valid ",(0,i.kt)("inlineCode",{parentName:"p"},"body")," and ",(0,i.kt)("inlineCode",{parentName:"p"},"auth_info")," for a transaction, and serialize these two messages using Protobuf. ",(0,i.kt)("inlineCode",{parentName:"p"},"TxRaw")," then pins the user's exact binary representation of ",(0,i.kt)("inlineCode",{parentName:"p"},"body")," and ",(0,i.kt)("inlineCode",{parentName:"p"},"auth_info"),", called respectively ",(0,i.kt)("inlineCode",{parentName:"p"},"body_bytes")," and ",(0,i.kt)("inlineCode",{parentName:"p"},"auth_info_bytes"),". The document that is signed by all signers of the transaction is ",(0,i.kt)("inlineCode",{parentName:"p"},"SignDoc")," (deterministically serialized using ",(0,i.kt)("a",{parentName:"p",href:"/main/architecture/adr-027-deterministic-protobuf-serialization"},"ADR-027"),"):"),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-protobuf",metastring:"reference",reference:!0},"https://github.com/cosmos/cosmos-sdk/blob/v0.47.0-rc1/proto/cosmos/tx/v1beta1/tx.proto#L48-L65\n")),(0,i.kt)("p",null,"Once signed by all signers, the ",(0,i.kt)("inlineCode",{parentName:"p"},"body_bytes"),", ",(0,i.kt)("inlineCode",{parentName:"p"},"auth_info_bytes")," and ",(0,i.kt)("inlineCode",{parentName:"p"},"signatures")," are gathered into ",(0,i.kt)("inlineCode",{parentName:"p"},"TxRaw"),", whose serialized bytes are broadcasted over the network."),(0,i.kt)("h4",{id:"sign_mode_legacy_amino_json"},(0,i.kt)("inlineCode",{parentName:"h4"},"SIGN_MODE_LEGACY_AMINO_JSON")),(0,i.kt)("p",null,"The legacy implementation of the ",(0,i.kt)("inlineCode",{parentName:"p"},"Tx")," interface is the ",(0,i.kt)("inlineCode",{parentName:"p"},"StdTx")," struct from ",(0,i.kt)("inlineCode",{parentName:"p"},"x/auth"),":"),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-go",metastring:"reference",reference:!0},"https://github.com/cosmos/cosmos-sdk/blob/v0.47.0-rc1/x/auth/migrations/legacytx/stdtx.go#L83-L93\n")),(0,i.kt)("p",null,"The document signed by all signers is ",(0,i.kt)("inlineCode",{parentName:"p"},"StdSignDoc"),":"),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-go",metastring:"reference",reference:!0},"https://github.com/cosmos/cosmos-sdk/blob/v0.47.0-rc1/x/auth/migrations/legacytx/stdsign.go#L38-L52\n")),(0,i.kt)("p",null,"which is encoded into bytes using Amino JSON. Once all signatures are gathered into ",(0,i.kt)("inlineCode",{parentName:"p"},"StdTx"),", ",(0,i.kt)("inlineCode",{parentName:"p"},"StdTx")," is serialized using Amino JSON, and these bytes are broadcasted over the network."),(0,i.kt)("h4",{id:"other-sign-modes"},"Other Sign Modes"),(0,i.kt)("p",null,"The Cosmos SDK also provides a couple of other sign modes for particular use cases."),(0,i.kt)("h4",{id:"sign_mode_direct_aux"},(0,i.kt)("inlineCode",{parentName:"h4"},"SIGN_MODE_DIRECT_AUX")),(0,i.kt)("p",null,(0,i.kt)("inlineCode",{parentName:"p"},"SIGN_MODE_DIRECT_AUX")," is a sign mode released in the Cosmos SDK v0.46 which targets transactions with multiple signers. Whereas ",(0,i.kt)("inlineCode",{parentName:"p"},"SIGN_MODE_DIRECT")," expects each signer to sign over both ",(0,i.kt)("inlineCode",{parentName:"p"},"TxBody")," and ",(0,i.kt)("inlineCode",{parentName:"p"},"AuthInfo")," (which includes all other signers' signer infos, i.e. their account sequence, public key and mode info), ",(0,i.kt)("inlineCode",{parentName:"p"},"SIGN_MODE_DIRECT_AUX")," allows N-1 signers to only sign over ",(0,i.kt)("inlineCode",{parentName:"p"},"TxBody")," and ",(0,i.kt)("em",{parentName:"p"},"their own")," signer info. Morever, each auxiliary signer (i.e. a signer using ",(0,i.kt)("inlineCode",{parentName:"p"},"SIGN_MODE_DIRECT_AUX"),") doesn't\nneed to sign over the fees:"),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-protobuf",metastring:"reference",reference:!0},"https://github.com/cosmos/cosmos-sdk/blob/v0.47.0-rc1/proto/cosmos/tx/v1beta1/tx.proto#L67-L97\n")),(0,i.kt)("p",null,"The use case is a multi-signer transaction, where one of the signers is appointed to gather all signatures, broadcast the signature and pay for fees, and the others only care about the transaction body. This generally allows for a better multi-signing UX. If Alice, Bob and Charlie are part of a 3-signer transaction, then Alice and Bob can both use ",(0,i.kt)("inlineCode",{parentName:"p"},"SIGN_MODE_DIRECT_AUX")," to sign over the ",(0,i.kt)("inlineCode",{parentName:"p"},"TxBody")," and their own signer info (no need an additional step to gather other signers' ones, like in ",(0,i.kt)("inlineCode",{parentName:"p"},"SIGN_MODE_DIRECT"),"), without specifying a fee in their SignDoc. Charlie can then gather both signatures from Alice and Bob, and\ncreate the final transaction by appending a fee. Note that the fee payer of the transaction (in our case Charlie) must sign over the fees, so must use ",(0,i.kt)("inlineCode",{parentName:"p"},"SIGN_MODE_DIRECT")," or ",(0,i.kt)("inlineCode",{parentName:"p"},"SIGN_MODE_LEGACY_AMINO_JSON"),"."),(0,i.kt)("p",null,"A concrete use case is implemented in ",(0,i.kt)("a",{parentName:"p",href:"/main/core/tips"},"transaction tips"),": the tipper may use ",(0,i.kt)("inlineCode",{parentName:"p"},"SIGN_MODE_DIRECT_AUX")," to specify a tip in the transaction, without signing over the actual transaction fees. Then, the fee payer appends fees inside the tipper's desired ",(0,i.kt)("inlineCode",{parentName:"p"},"TxBody"),", and as an exchange for paying the fees and broadcasting the transaction, receives the tipper's transaction tips as payment."),(0,i.kt)("h4",{id:"sign_mode_textual"},(0,i.kt)("inlineCode",{parentName:"h4"},"SIGN_MODE_TEXTUAL")),(0,i.kt)("p",null,(0,i.kt)("inlineCode",{parentName:"p"},"SIGN_MODE_TEXTUAL")," is a new sign mode for delivering a better signing experience on hardware wallets, it is currently still under implementation. If you wish to learn more, please refer to ",(0,i.kt)("a",{parentName:"p",href:"https://github.com/cosmos/cosmos-sdk/pull/10701"},"ADR-050"),"."),(0,i.kt)("h4",{id:"custom-sign-modes"},"Custom Sign modes"),(0,i.kt)("p",null,"There is the the opportunity to add your own custom sign mode to the Cosmos-SDK.  While we can not accept the implementation of the sign mode to the repository, we can accept a pull request to add the custom signmode to the SignMode enum located ",(0,i.kt)("a",{parentName:"p",href:"https://github.com/cosmos/cosmos-sdk/blob/v0.47.0-rc1/proto/cosmos/tx/signing/v1beta1/signing.proto#L17"},"here")),(0,i.kt)("h2",{id:"transaction-process"},"Transaction Process"),(0,i.kt)("p",null,"The process of an end-user sending a transaction is:"),(0,i.kt)("ul",null,(0,i.kt)("li",{parentName:"ul"},"decide on the messages to put into the transaction,"),(0,i.kt)("li",{parentName:"ul"},"generate the transaction using the Cosmos SDK's ",(0,i.kt)("inlineCode",{parentName:"li"},"TxBuilder"),","),(0,i.kt)("li",{parentName:"ul"},"broadcast the transaction using one of the available interfaces.")),(0,i.kt)("p",null,"The next paragraphs will describe each of these components, in this order."),(0,i.kt)("h3",{id:"messages"},"Messages"),(0,i.kt)("admonition",{type:"tip"},(0,i.kt)("p",{parentName:"admonition"},"Module ",(0,i.kt)("inlineCode",{parentName:"p"},"sdk.Msg"),"s are not to be confused with ",(0,i.kt)("a",{parentName:"p",href:"https://docs.tendermint.com/master/spec/abci/abci.html#messages"},"ABCI Messages")," which define interactions between the Tendermint and application layers.")),(0,i.kt)("p",null,(0,i.kt)("strong",{parentName:"p"},"Messages")," (or ",(0,i.kt)("inlineCode",{parentName:"p"},"sdk.Msg"),"s) are module-specific objects that trigger state transitions within the scope of the module they belong to. Module developers define the messages for their module by adding methods to the Protobuf ",(0,i.kt)("a",{parentName:"p",href:"/main/building-modules/msg-services"},(0,i.kt)("inlineCode",{parentName:"a"},"Msg")," service"),", and also implement the corresponding ",(0,i.kt)("inlineCode",{parentName:"p"},"MsgServer"),"."),(0,i.kt)("p",null,"Each ",(0,i.kt)("inlineCode",{parentName:"p"},"sdk.Msg"),"s is related to exactly one Protobuf ",(0,i.kt)("a",{parentName:"p",href:"/main/building-modules/msg-services"},(0,i.kt)("inlineCode",{parentName:"a"},"Msg")," service")," RPC, defined inside each module's ",(0,i.kt)("inlineCode",{parentName:"p"},"tx.proto")," file. A SDK app router automatically maps every ",(0,i.kt)("inlineCode",{parentName:"p"},"sdk.Msg")," to a corresponding RPC. Protobuf generates a ",(0,i.kt)("inlineCode",{parentName:"p"},"MsgServer")," interface for each module ",(0,i.kt)("inlineCode",{parentName:"p"},"Msg")," service, and the module developer needs to implement this interface.\nThis design puts more responsibility on module developers, allowing application developers to reuse common functionalities without having to implement state transition logic repetitively."),(0,i.kt)("p",null,"To learn more about Protobuf ",(0,i.kt)("inlineCode",{parentName:"p"},"Msg")," services and how to implement ",(0,i.kt)("inlineCode",{parentName:"p"},"MsgServer"),", click ",(0,i.kt)("a",{parentName:"p",href:"/main/building-modules/msg-services"},"here"),"."),(0,i.kt)("p",null,"While messages contain the information for state transition logic, a transaction's other metadata and relevant information are stored in the ",(0,i.kt)("inlineCode",{parentName:"p"},"TxBuilder")," and ",(0,i.kt)("inlineCode",{parentName:"p"},"Context"),"."),(0,i.kt)("h3",{id:"transaction-generation"},"Transaction Generation"),(0,i.kt)("p",null,"The ",(0,i.kt)("inlineCode",{parentName:"p"},"TxBuilder")," interface contains data closely related with the generation of transactions, which an end-user can freely set to generate the desired transaction:"),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-go",metastring:"reference",reference:!0},"https://github.com/cosmos/cosmos-sdk/blob/v0.47.0-rc1/client/tx_config.go#L33-L50\n")),(0,i.kt)("ul",null,(0,i.kt)("li",{parentName:"ul"},(0,i.kt)("inlineCode",{parentName:"li"},"Msg"),"s, the array of ",(0,i.kt)("a",{parentName:"li",href:"#messages"},"messages")," included in the transaction."),(0,i.kt)("li",{parentName:"ul"},(0,i.kt)("inlineCode",{parentName:"li"},"GasLimit"),", option chosen by the users for how to calculate how much gas they will need to pay."),(0,i.kt)("li",{parentName:"ul"},(0,i.kt)("inlineCode",{parentName:"li"},"Memo"),", a note or comment to send with the transaction."),(0,i.kt)("li",{parentName:"ul"},(0,i.kt)("inlineCode",{parentName:"li"},"FeeAmount"),", the maximum amount the user is willing to pay in fees."),(0,i.kt)("li",{parentName:"ul"},(0,i.kt)("inlineCode",{parentName:"li"},"TimeoutHeight"),", block height until which the transaction is valid."),(0,i.kt)("li",{parentName:"ul"},(0,i.kt)("inlineCode",{parentName:"li"},"Signatures"),", the array of signatures from all signers of the transaction.")),(0,i.kt)("p",null,"As there are currently two sign modes for signing transactions, there are also two implementations of ",(0,i.kt)("inlineCode",{parentName:"p"},"TxBuilder"),":"),(0,i.kt)("ul",null,(0,i.kt)("li",{parentName:"ul"},(0,i.kt)("a",{parentName:"li",href:"https://github.com/cosmos/cosmos-sdk/blob/v0.47.0-rc1/x/auth/tx/builder.go#L18-L34"},"wrapper")," for creating transactions for ",(0,i.kt)("inlineCode",{parentName:"li"},"SIGN_MODE_DIRECT"),","),(0,i.kt)("li",{parentName:"ul"},(0,i.kt)("a",{parentName:"li",href:"https://github.com/cosmos/cosmos-sdk/blob/v0.47.0-rc1/x/auth/migrations/legacytx/stdtx_builder.go#L15-L21"},"StdTxBuilder")," for ",(0,i.kt)("inlineCode",{parentName:"li"},"SIGN_MODE_LEGACY_AMINO_JSON"),".")),(0,i.kt)("p",null,"However, the two implementation of ",(0,i.kt)("inlineCode",{parentName:"p"},"TxBuilder")," should be hidden away from end-users, as they should prefer using the overarching ",(0,i.kt)("inlineCode",{parentName:"p"},"TxConfig")," interface:"),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-go",metastring:"reference",reference:!0},"https://github.com/cosmos/cosmos-sdk/blob/v0.47.0-rc1/client/tx_config.go#L22-L31\n")),(0,i.kt)("p",null,(0,i.kt)("inlineCode",{parentName:"p"},"TxConfig")," is an app-wide configuration for managing transactions. Most importantly, it holds the information about whether to sign each transaction with ",(0,i.kt)("inlineCode",{parentName:"p"},"SIGN_MODE_DIRECT")," or ",(0,i.kt)("inlineCode",{parentName:"p"},"SIGN_MODE_LEGACY_AMINO_JSON"),". By calling ",(0,i.kt)("inlineCode",{parentName:"p"},"txBuilder := txConfig.NewTxBuilder()"),", a new ",(0,i.kt)("inlineCode",{parentName:"p"},"TxBuilder")," will be created with the appropriate sign mode."),(0,i.kt)("p",null,"Once ",(0,i.kt)("inlineCode",{parentName:"p"},"TxBuilder")," is correctly populated with the setters exposed above, ",(0,i.kt)("inlineCode",{parentName:"p"},"TxConfig")," will also take care of correctly encoding the bytes (again, either using ",(0,i.kt)("inlineCode",{parentName:"p"},"SIGN_MODE_DIRECT")," or ",(0,i.kt)("inlineCode",{parentName:"p"},"SIGN_MODE_LEGACY_AMINO_JSON"),"). Here's a pseudo-code snippet of how to generate and encode a transaction, using the ",(0,i.kt)("inlineCode",{parentName:"p"},"TxEncoder()")," method:"),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-go"},"txBuilder := txConfig.NewTxBuilder()\ntxBuilder.SetMsgs(...) // and other setters on txBuilder\n\nbz, err := txConfig.TxEncoder()(txBuilder.GetTx())\n// bz are bytes to be broadcasted over the network\n")),(0,i.kt)("h3",{id:"broadcasting-the-transaction"},"Broadcasting the Transaction"),(0,i.kt)("p",null,"Once the transaction bytes are generated, there are currently three ways of broadcasting it."),(0,i.kt)("h4",{id:"cli"},"CLI"),(0,i.kt)("p",null,"Application developers create entry points to the application by creating a ",(0,i.kt)("a",{parentName:"p",href:"/main/core/cli"},"command-line interface"),", ",(0,i.kt)("a",{parentName:"p",href:"/main/core/grpc_rest"},"gRPC and/or REST interface"),", typically found in the application's ",(0,i.kt)("inlineCode",{parentName:"p"},"./cmd")," folder. These interfaces allow users to interact with the application through command-line."),(0,i.kt)("p",null,"For the ",(0,i.kt)("a",{parentName:"p",href:"/main/building-modules/module-interfaces#cli"},"command-line interface"),", module developers create subcommands to add as children to the application top-level transaction command ",(0,i.kt)("inlineCode",{parentName:"p"},"TxCmd"),". CLI commands actually bundle all the steps of transaction processing into one simple command: creating messages, generating transactions and broadcasting. For concrete examples, see the ",(0,i.kt)("a",{parentName:"p",href:"/main/run-node/interact-node"},"Interacting with a Node")," section. An example transaction made using CLI looks like:"),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-bash"},"simd tx send $MY_VALIDATOR_ADDRESS $RECIPIENT 1000stake\n")),(0,i.kt)("h4",{id:"grpc"},"gRPC"),(0,i.kt)("p",null,(0,i.kt)("a",{parentName:"p",href:"https://grpc.io"},"gRPC")," is the main component for the Cosmos SDK's RPC layer. Its principal usage is in the context of modules' ",(0,i.kt)("a",{parentName:"p",href:"/main/building-modules/query-services"},(0,i.kt)("inlineCode",{parentName:"a"},"Query")," services"),". However, the Cosmos SDK also exposes a few other module-agnostic gRPC services, one of them being the ",(0,i.kt)("inlineCode",{parentName:"p"},"Tx")," service:"),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-go",metastring:"reference",reference:!0},"https://github.com/cosmos/cosmos-sdk/blob/v0.47.0-rc1/proto/cosmos/tx/v1beta1/service.proto\n")),(0,i.kt)("p",null,"The ",(0,i.kt)("inlineCode",{parentName:"p"},"Tx")," service exposes a handful of utility functions, such as simulating a transaction or querying a transaction, and also one method to broadcast transactions."),(0,i.kt)("p",null,"Examples of broadcasting and simulating a transaction are shown ",(0,i.kt)("a",{parentName:"p",href:"/main/run-node/txs#programmatically-with-go"},"here"),"."),(0,i.kt)("h4",{id:"rest"},"REST"),(0,i.kt)("p",null,"Each gRPC method has its corresponding REST endpoint, generated using ",(0,i.kt)("a",{parentName:"p",href:"https://github.com/grpc-ecosystem/grpc-gateway"},"gRPC-gateway"),". Therefore, instead of using gRPC, you can also use HTTP to broadcast the same transaction, on the ",(0,i.kt)("inlineCode",{parentName:"p"},"POST /cosmos/tx/v1beta1/txs")," endpoint."),(0,i.kt)("p",null,"An example can be seen ",(0,i.kt)("a",{parentName:"p",href:"/main/run-node/txs#using-rest"},"here")),(0,i.kt)("h4",{id:"tendermint-rpc"},"Tendermint RPC"),(0,i.kt)("p",null,"The three methods presented above are actually higher abstractions over the Tendermint RPC ",(0,i.kt)("inlineCode",{parentName:"p"},"/broadcast_tx_{async,sync,commit}")," endpoints, documented ",(0,i.kt)("a",{parentName:"p",href:"https://docs.tendermint.com/master/rpc/#/Tx"},"here"),". This means that you can use the Tendermint RPC endpoints directly to broadcast the transaction, if you wish so."))}d.isMDXComponent=!0}}]);