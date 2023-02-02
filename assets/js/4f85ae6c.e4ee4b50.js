"use strict";(self.webpackChunkcosmos_sdk_docs=self.webpackChunkcosmos_sdk_docs||[]).push([[2040],{3905:(e,t,n)=>{n.d(t,{Zo:()=>p,kt:()=>u});var r=n(7294);function i(e,t,n){return t in e?Object.defineProperty(e,t,{value:n,enumerable:!0,configurable:!0,writable:!0}):e[t]=n,e}function o(e,t){var n=Object.keys(e);if(Object.getOwnPropertySymbols){var r=Object.getOwnPropertySymbols(e);t&&(r=r.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),n.push.apply(n,r)}return n}function a(e){for(var t=1;t<arguments.length;t++){var n=null!=arguments[t]?arguments[t]:{};t%2?o(Object(n),!0).forEach((function(t){i(e,t,n[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(n)):o(Object(n)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(n,t))}))}return e}function s(e,t){if(null==e)return{};var n,r,i=function(e,t){if(null==e)return{};var n,r,i={},o=Object.keys(e);for(r=0;r<o.length;r++)n=o[r],t.indexOf(n)>=0||(i[n]=e[n]);return i}(e,t);if(Object.getOwnPropertySymbols){var o=Object.getOwnPropertySymbols(e);for(r=0;r<o.length;r++)n=o[r],t.indexOf(n)>=0||Object.prototype.propertyIsEnumerable.call(e,n)&&(i[n]=e[n])}return i}var l=r.createContext({}),c=function(e){var t=r.useContext(l),n=t;return e&&(n="function"==typeof e?e(t):a(a({},t),e)),n},p=function(e){var t=c(e.components);return r.createElement(l.Provider,{value:t},e.children)},d={inlineCode:"code",wrapper:function(e){var t=e.children;return r.createElement(r.Fragment,{},t)}},h=r.forwardRef((function(e,t){var n=e.components,i=e.mdxType,o=e.originalType,l=e.parentName,p=s(e,["components","mdxType","originalType","parentName"]),h=c(n),u=i,m=h["".concat(l,".").concat(u)]||h[u]||d[u]||o;return n?r.createElement(m,a(a({ref:t},p),{},{components:n})):r.createElement(m,a({ref:t},p))}));function u(e,t){var n=arguments,i=t&&t.mdxType;if("string"==typeof e||i){var o=n.length,a=new Array(o);a[0]=h;var s={};for(var l in t)hasOwnProperty.call(t,l)&&(s[l]=t[l]);s.originalType=e,s.mdxType="string"==typeof e?e:i,a[1]=s;for(var c=2;c<o;c++)a[c]=n[c];return r.createElement.apply(null,a)}return r.createElement.apply(null,n)}h.displayName="MDXCreateElement"},1581:(e,t,n)=>{n.r(t),n.d(t,{assets:()=>l,contentTitle:()=>a,default:()=>d,frontMatter:()=>o,metadata:()=>s,toc:()=>c});var r=n(7462),i=(n(7294),n(3905));const o={},a="ADR 035: Rosetta API Support",s={unversionedId:"architecture/adr-035-rosetta-api-support",id:"version-v0.47/architecture/adr-035-rosetta-api-support",title:"ADR 035: Rosetta API Support",description:"Authors",source:"@site/versioned_docs/version-v0.47/architecture/adr-035-rosetta-api-support.md",sourceDirName:"architecture",slug:"/architecture/adr-035-rosetta-api-support",permalink:"/v0.47/architecture/adr-035-rosetta-api-support",draft:!1,tags:[],version:"v0.47",frontMatter:{},sidebar:"tutorialSidebar",previous:{title:"ADR 034: Account Rekeying",permalink:"/v0.47/architecture/adr-034-account-rekeying"},next:{title:"ADR 036: Arbitrary Message Signature Specification",permalink:"/v0.47/architecture/adr-036-arbitrary-signature"}},l={},c=[{value:"Authors",id:"authors",level:2},{value:"Changelog",id:"changelog",level:2},{value:"Context",id:"context",level:2},{value:"Decision",id:"decision",level:2},{value:"Architecture",id:"architecture",level:2},{value:"The External Repo",id:"the-external-repo",level:3},{value:"Server",id:"server",level:4},{value:"Types",id:"types",level:4},{value:"Interfaces",id:"interfaces",level:5},{value:"2. Cosmos SDK Implementation",id:"2-cosmos-sdk-implementation",level:3},{value:"3. API service invocation",id:"3-api-service-invocation",level:3},{value:"Shared Process (Only Stargate)",id:"shared-process-only-stargate",level:4},{value:"Separate API service",id:"separate-api-service",level:4},{value:"Status",id:"status",level:2},{value:"Consequences",id:"consequences",level:2},{value:"Positive",id:"positive",level:3},{value:"References",id:"references",level:2}],p={toc:c};function d(e){let{components:t,...n}=e;return(0,i.kt)("wrapper",(0,r.Z)({},p,n,{components:t,mdxType:"MDXLayout"}),(0,i.kt)("h1",{id:"adr-035-rosetta-api-support"},"ADR 035: Rosetta API Support"),(0,i.kt)("h2",{id:"authors"},"Authors"),(0,i.kt)("ul",null,(0,i.kt)("li",{parentName:"ul"},"Jonathan Gimeno (@jgimeno)"),(0,i.kt)("li",{parentName:"ul"},"David Grierson (@senormonito)"),(0,i.kt)("li",{parentName:"ul"},"Alessio Treglia (@alessio)"),(0,i.kt)("li",{parentName:"ul"},"Frojdy Dymylja (@fdymylja)")),(0,i.kt)("h2",{id:"changelog"},"Changelog"),(0,i.kt)("ul",null,(0,i.kt)("li",{parentName:"ul"},"2021-05-12: the external library  ",(0,i.kt)("a",{parentName:"li",href:"https://github.com/tendermint/cosmos-rosetta-gateway"},"cosmos-rosetta-gateway")," has been moved within the Cosmos SDK.")),(0,i.kt)("h2",{id:"context"},"Context"),(0,i.kt)("p",null,(0,i.kt)("a",{parentName:"p",href:"https://www.rosetta-api.org/"},"Rosetta API")," is an open-source specification and set of tools developed by Coinbase to\nstandardise blockchain interactions."),(0,i.kt)("p",null,"Through the use of a standard API for integrating blockchain applications it will"),(0,i.kt)("ul",null,(0,i.kt)("li",{parentName:"ul"},"Be easier for a user to interact with a given blockchain"),(0,i.kt)("li",{parentName:"ul"},"Allow exchanges to integrate new blockchains quickly and easily"),(0,i.kt)("li",{parentName:"ul"},"Enable application developers to build cross-blockchain applications such as block explorers, wallets and dApps at\nconsiderably lower cost and effort.")),(0,i.kt)("h2",{id:"decision"},"Decision"),(0,i.kt)("p",null,"It is clear that adding Rosetta API support to the Cosmos SDK will bring value to all the developers and\nCosmos SDK based chains in the ecosystem. How it is implemented is key."),(0,i.kt)("p",null,"The driving principles of the proposed design are:"),(0,i.kt)("ol",null,(0,i.kt)("li",{parentName:"ol"},(0,i.kt)("strong",{parentName:"li"},"Extensibility:")," it must be as riskless and painless as possible for application developers to set-up network\nconfigurations to expose Rosetta API-compliant services."),(0,i.kt)("li",{parentName:"ol"},(0,i.kt)("strong",{parentName:"li"},"Long term support:")," This proposal aims to provide support for all the supported Cosmos SDK release series."),(0,i.kt)("li",{parentName:"ol"},(0,i.kt)("strong",{parentName:"li"},"Cost-efficiency:")," Backporting changes to Rosetta API specifications from ",(0,i.kt)("inlineCode",{parentName:"li"},"master")," to the various stable\nbranches of Cosmos SDK is a cost that needs to be reduced.")),(0,i.kt)("p",null,"We will achieve these delivering on these principles by the following:"),(0,i.kt)("ol",null,(0,i.kt)("li",{parentName:"ol"},"There will be a package ",(0,i.kt)("inlineCode",{parentName:"li"},"rosetta/lib"),"\nfor the implementation of the core Rosetta API features, particularly:\na. The types and interfaces (",(0,i.kt)("inlineCode",{parentName:"li"},"Client"),", ",(0,i.kt)("inlineCode",{parentName:"li"},"OfflineClient"),"...), this separates design from implementation detail.\nb. The ",(0,i.kt)("inlineCode",{parentName:"li"},"Server")," functionality as this is independent of the Cosmos SDK version.\nc. The ",(0,i.kt)("inlineCode",{parentName:"li"},"Online/OfflineNetwork"),", which is not exported, and implements the rosetta API using the ",(0,i.kt)("inlineCode",{parentName:"li"},"Client")," interface to query the node, build tx and so on.\nd. The ",(0,i.kt)("inlineCode",{parentName:"li"},"errors")," package to extend rosetta errors."),(0,i.kt)("li",{parentName:"ol"},"Due to differences between the Cosmos release series, each series will have its own specific implementation of ",(0,i.kt)("inlineCode",{parentName:"li"},"Client")," interface."),(0,i.kt)("li",{parentName:"ol"},"There will be two options for starting an API service in applications:\na. API shares the application process\nb. API-specific process.")),(0,i.kt)("h2",{id:"architecture"},"Architecture"),(0,i.kt)("h3",{id:"the-external-repo"},"The External Repo"),(0,i.kt)("p",null,"As section will describe the proposed external library, including the service implementation, plus the defined types and interfaces."),(0,i.kt)("h4",{id:"server"},"Server"),(0,i.kt)("p",null,(0,i.kt)("inlineCode",{parentName:"p"},"Server")," is a simple ",(0,i.kt)("inlineCode",{parentName:"p"},"struct")," that is started and listens to the port specified in the settings. This is meant to be used across all the Cosmos SDK versions that are actively supported."),(0,i.kt)("p",null,"The constructor follows:"),(0,i.kt)("p",null,(0,i.kt)("inlineCode",{parentName:"p"},"func NewServer(settings Settings) (Server, error)")),(0,i.kt)("p",null,(0,i.kt)("inlineCode",{parentName:"p"},"Settings"),", which are used to construct a new server, are the following:"),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-go"},"// Settings define the rosetta server settings\ntype Settings struct {\n    // Network contains the information regarding the network\n    Network *types.NetworkIdentifier\n    // Client is the online API handler\n    Client crgtypes.Client\n    // Listen is the address the handler will listen at\n    Listen string\n    // Offline defines if the rosetta service should be exposed in offline mode\n    Offline bool\n    // Retries is the number of readiness checks that will be attempted when instantiating the handler\n    // valid only for online API\n    Retries int\n    // RetryWait is the time that will be waited between retries\n    RetryWait time.Duration\n}\n")),(0,i.kt)("h4",{id:"types"},"Types"),(0,i.kt)("p",null,"Package types uses a mixture of rosetta types and custom defined type wrappers, that the client must parse and return while executing operations."),(0,i.kt)("h5",{id:"interfaces"},"Interfaces"),(0,i.kt)("p",null,"Every SDK version uses a different format to connect (rpc, gRPC, etc), query and build transactions, we have abstracted this in what is the ",(0,i.kt)("inlineCode",{parentName:"p"},"Client")," interface.\nThe client uses rosetta types, whilst the ",(0,i.kt)("inlineCode",{parentName:"p"},"Online/OfflineNetwork")," takes care of returning correctly parsed rosetta responses and errors."),(0,i.kt)("p",null,"Each Cosmos SDK release series will have their own ",(0,i.kt)("inlineCode",{parentName:"p"},"Client")," implementations.\nDevelopers can implement their own custom ",(0,i.kt)("inlineCode",{parentName:"p"},"Client"),"s as required."),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-go"},"// Client defines the API the client implementation should provide.\ntype Client interface {\n    // Needed if the client needs to perform some action before connecting.\n    Bootstrap() error\n    // Ready checks if the servicer constraints for queries are satisfied\n    // for example the node might still not be ready, it's useful in process\n    // when the rosetta instance might come up before the node itself\n    // the servicer must return nil if the node is ready\n    Ready() error\n\n    // Data API\n\n    // Balances fetches the balance of the given address\n    // if height is not nil, then the balance will be displayed\n    // at the provided height, otherwise last block balance will be returned\n    Balances(ctx context.Context, addr string, height *int64) ([]*types.Amount, error)\n    // BlockByHashAlt gets a block and its transaction at the provided height\n    BlockByHash(ctx context.Context, hash string) (BlockResponse, error)\n    // BlockByHeightAlt gets a block given its height, if height is nil then last block is returned\n    BlockByHeight(ctx context.Context, height *int64) (BlockResponse, error)\n    // BlockTransactionsByHash gets the block, parent block and transactions\n    // given the block hash.\n    BlockTransactionsByHash(ctx context.Context, hash string) (BlockTransactionsResponse, error)\n    // BlockTransactionsByHash gets the block, parent block and transactions\n    // given the block hash.\n    BlockTransactionsByHeight(ctx context.Context, height *int64) (BlockTransactionsResponse, error)\n    // GetTx gets a transaction given its hash\n    GetTx(ctx context.Context, hash string) (*types.Transaction, error)\n    // GetUnconfirmedTx gets an unconfirmed Tx given its hash\n    // NOTE(fdymylja): NOT IMPLEMENTED YET!\n    GetUnconfirmedTx(ctx context.Context, hash string) (*types.Transaction, error)\n    // Mempool returns the list of the current non confirmed transactions\n    Mempool(ctx context.Context) ([]*types.TransactionIdentifier, error)\n    // Peers gets the peers currently connected to the node\n    Peers(ctx context.Context) ([]*types.Peer, error)\n    // Status returns the node status, such as sync data, version etc\n    Status(ctx context.Context) (*types.SyncStatus, error)\n\n    // Construction API\n\n    // PostTx posts txBytes to the node and returns the transaction identifier plus metadata related\n    // to the transaction itself.\n    PostTx(txBytes []byte) (res *types.TransactionIdentifier, meta map[string]interface{}, err error)\n    // ConstructionMetadataFromOptions\n    ConstructionMetadataFromOptions(ctx context.Context, options map[string]interface{}) (meta map[string]interface{}, err error)\n    OfflineClient\n}\n\n// OfflineClient defines the functionalities supported without having access to the node\ntype OfflineClient interface {\n    NetworkInformationProvider\n    // SignedTx returns the signed transaction given the tx bytes (msgs) plus the signatures\n    SignedTx(ctx context.Context, txBytes []byte, sigs []*types.Signature) (signedTxBytes []byte, err error)\n    // TxOperationsAndSignersAccountIdentifiers returns the operations related to a transaction and the account\n    // identifiers if the transaction is signed\n    TxOperationsAndSignersAccountIdentifiers(signed bool, hexBytes []byte) (ops []*types.Operation, signers []*types.AccountIdentifier, err error)\n    // ConstructionPayload returns the construction payload given the request\n    ConstructionPayload(ctx context.Context, req *types.ConstructionPayloadsRequest) (resp *types.ConstructionPayloadsResponse, err error)\n    // PreprocessOperationsToOptions returns the options given the preprocess operations\n    PreprocessOperationsToOptions(ctx context.Context, req *types.ConstructionPreprocessRequest) (options map[string]interface{}, err error)\n    // AccountIdentifierFromPublicKey returns the account identifier given the public key\n    AccountIdentifierFromPublicKey(pubKey *types.PublicKey) (*types.AccountIdentifier, error)\n}\n")),(0,i.kt)("h3",{id:"2-cosmos-sdk-implementation"},"2. Cosmos SDK Implementation"),(0,i.kt)("p",null,"The Cosmos SDK implementation, based on version, takes care of satisfying the ",(0,i.kt)("inlineCode",{parentName:"p"},"Client")," interface.\nIn Stargate, Launchpad and 0.37, we have introduced the concept of rosetta.Msg, this message is not in the shared repository as the sdk.Msg type differs between Cosmos SDK versions."),(0,i.kt)("p",null,"The rosetta.Msg interface follows:"),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-go"},"// Msg represents a cosmos-sdk message that can be converted from and to a rosetta operation.\ntype Msg interface {\n    sdk.Msg\n    ToOperations(withStatus, hasError bool) []*types.Operation\n    FromOperations(ops []*types.Operation) (sdk.Msg, error)\n}\n")),(0,i.kt)("p",null,"Hence developers who want to extend the rosetta set of supported operations just need to extend their module's sdk.Msgs with the ",(0,i.kt)("inlineCode",{parentName:"p"},"ToOperations")," and ",(0,i.kt)("inlineCode",{parentName:"p"},"FromOperations")," methods."),(0,i.kt)("h3",{id:"3-api-service-invocation"},"3. API service invocation"),(0,i.kt)("p",null,"As stated at the start, application developers will have two methods for invocation of the Rosetta API service:"),(0,i.kt)("ol",null,(0,i.kt)("li",{parentName:"ol"},"Shared process for both application and API"),(0,i.kt)("li",{parentName:"ol"},"Standalone API service")),(0,i.kt)("h4",{id:"shared-process-only-stargate"},"Shared Process (Only Stargate)"),(0,i.kt)("p",null,"Rosetta API service could run within the same execution process as the application. This would be enabled via app.toml settings, and if gRPC is not enabled the rosetta instance would be spinned in offline mode (tx building capabilities only)."),(0,i.kt)("h4",{id:"separate-api-service"},"Separate API service"),(0,i.kt)("p",null,"Client application developers can write a new command to launch a Rosetta API server as a separate process too, using the rosetta command contained in the ",(0,i.kt)("inlineCode",{parentName:"p"},"/server/rosetta")," package. Construction of the command depends on Cosmos SDK version. Examples can be found inside ",(0,i.kt)("inlineCode",{parentName:"p"},"simd")," for stargate, and ",(0,i.kt)("inlineCode",{parentName:"p"},"contrib/rosetta/simapp")," for other release series."),(0,i.kt)("h2",{id:"status"},"Status"),(0,i.kt)("p",null,"Proposed"),(0,i.kt)("h2",{id:"consequences"},"Consequences"),(0,i.kt)("h3",{id:"positive"},"Positive"),(0,i.kt)("ul",null,(0,i.kt)("li",{parentName:"ul"},"Out-of-the-box Rosetta API support within Cosmos SDK."),(0,i.kt)("li",{parentName:"ul"},"Blockchain interface standardisation")),(0,i.kt)("h2",{id:"references"},"References"),(0,i.kt)("ul",null,(0,i.kt)("li",{parentName:"ul"},(0,i.kt)("a",{parentName:"li",href:"https://www.rosetta-api.org/"},"https://www.rosetta-api.org/"))))}d.isMDXComponent=!0}}]);