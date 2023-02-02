"use strict";(self.webpackChunkcosmos_sdk_docs=self.webpackChunkcosmos_sdk_docs||[]).push([[1932],{3905:(e,t,n)=>{n.d(t,{Zo:()=>d,kt:()=>h});var a=n(7294);function o(e,t,n){return t in e?Object.defineProperty(e,t,{value:n,enumerable:!0,configurable:!0,writable:!0}):e[t]=n,e}function i(e,t){var n=Object.keys(e);if(Object.getOwnPropertySymbols){var a=Object.getOwnPropertySymbols(e);t&&(a=a.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),n.push.apply(n,a)}return n}function r(e){for(var t=1;t<arguments.length;t++){var n=null!=arguments[t]?arguments[t]:{};t%2?i(Object(n),!0).forEach((function(t){o(e,t,n[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(n)):i(Object(n)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(n,t))}))}return e}function s(e,t){if(null==e)return{};var n,a,o=function(e,t){if(null==e)return{};var n,a,o={},i=Object.keys(e);for(a=0;a<i.length;a++)n=i[a],t.indexOf(n)>=0||(o[n]=e[n]);return o}(e,t);if(Object.getOwnPropertySymbols){var i=Object.getOwnPropertySymbols(e);for(a=0;a<i.length;a++)n=i[a],t.indexOf(n)>=0||Object.prototype.propertyIsEnumerable.call(e,n)&&(o[n]=e[n])}return o}var l=a.createContext({}),p=function(e){var t=a.useContext(l),n=t;return e&&(n="function"==typeof e?e(t):r(r({},t),e)),n},d=function(e){var t=p(e.components);return a.createElement(l.Provider,{value:t},e.children)},c={inlineCode:"code",wrapper:function(e){var t=e.children;return a.createElement(a.Fragment,{},t)}},m=a.forwardRef((function(e,t){var n=e.components,o=e.mdxType,i=e.originalType,l=e.parentName,d=s(e,["components","mdxType","originalType","parentName"]),m=p(n),h=o,u=m["".concat(l,".").concat(h)]||m[h]||c[h]||i;return n?a.createElement(u,r(r({ref:t},d),{},{components:n})):a.createElement(u,r({ref:t},d))}));function h(e,t){var n=arguments,o=t&&t.mdxType;if("string"==typeof e||o){var i=n.length,r=new Array(i);r[0]=m;var s={};for(var l in t)hasOwnProperty.call(t,l)&&(s[l]=t[l]);s.originalType=e,s.mdxType="string"==typeof e?e:o,r[1]=s;for(var p=2;p<i;p++)r[p]=n[p];return a.createElement.apply(null,r)}return a.createElement.apply(null,n)}m.displayName="MDXCreateElement"},5191:(e,t,n)=>{n.r(t),n.d(t,{assets:()=>l,contentTitle:()=>r,default:()=>c,frontMatter:()=>i,metadata:()=>s,toc:()=>p});var a=n(7462),o=(n(7294),n(3905));const i={sidebar_position:1},r="Node Client (Daemon)",s={unversionedId:"core/node",id:"version-v0.47/core/node",title:"Node Client (Daemon)",description:"The main endpoint of a Cosmos SDK application is the daemon client, otherwise known as the full-node client. The full-node runs the state-machine, starting from a genesis file. It connects to peers running the same client in order to receive and relay transactions, block proposals and signatures. The full-node is constituted of the application, defined with the Cosmos SDK, and of a consensus engine connected to the application via the ABCI.",source:"@site/versioned_docs/version-v0.47/core/03-node.md",sourceDirName:"core",slug:"/core/node",permalink:"/v0.47/core/node",draft:!1,tags:[],version:"v0.47",sidebarPosition:1,frontMatter:{sidebar_position:1},sidebar:"tutorialSidebar",previous:{title:"Context",permalink:"/v0.47/core/context"},next:{title:"Store",permalink:"/v0.47/core/store"}},l={},p=[{value:"<code>main</code> function",id:"main-function",level:2},{value:"<code>start</code> command",id:"start-command",level:2},{value:"Other commands",id:"other-commands",level:2}],d={toc:p};function c(e){let{components:t,...n}=e;return(0,o.kt)("wrapper",(0,a.Z)({},d,n,{components:t,mdxType:"MDXLayout"}),(0,o.kt)("h1",{id:"node-client-daemon"},"Node Client (Daemon)"),(0,o.kt)("admonition",{title:"Synopsis",type:"note"},(0,o.kt)("p",{parentName:"admonition"},"The main endpoint of a Cosmos SDK application is the daemon client, otherwise known as the full-node client. The full-node runs the state-machine, starting from a genesis file. It connects to peers running the same client in order to receive and relay transactions, block proposals and signatures. The full-node is constituted of the application, defined with the Cosmos SDK, and of a consensus engine connected to the application via the ABCI.")),(0,o.kt)("admonition",{type:"note"},(0,o.kt)("h3",{parentName:"admonition",id:"pre-requisite-readings"},"Pre-requisite Readings"),(0,o.kt)("ul",{parentName:"admonition"},(0,o.kt)("li",{parentName:"ul"},(0,o.kt)("a",{parentName:"li",href:"/v0.47/basics/app-anatomy"},"Anatomy of an SDK application")))),(0,o.kt)("h2",{id:"main-function"},(0,o.kt)("inlineCode",{parentName:"h2"},"main")," function"),(0,o.kt)("p",null,"The full-node client of any Cosmos SDK application is built by running a ",(0,o.kt)("inlineCode",{parentName:"p"},"main")," function. The client is generally named by appending the ",(0,o.kt)("inlineCode",{parentName:"p"},"-d")," suffix to the application name (e.g. ",(0,o.kt)("inlineCode",{parentName:"p"},"appd")," for an application named ",(0,o.kt)("inlineCode",{parentName:"p"},"app"),"), and the ",(0,o.kt)("inlineCode",{parentName:"p"},"main")," function is defined in a ",(0,o.kt)("inlineCode",{parentName:"p"},"./appd/cmd/main.go")," file. Running this function creates an executable ",(0,o.kt)("inlineCode",{parentName:"p"},"appd")," that comes with a set of commands. For an app named ",(0,o.kt)("inlineCode",{parentName:"p"},"app"),", the main command is ",(0,o.kt)("a",{parentName:"p",href:"#start-command"},(0,o.kt)("inlineCode",{parentName:"a"},"appd start")),", which starts the full-node."),(0,o.kt)("p",null,"In general, developers will implement the ",(0,o.kt)("inlineCode",{parentName:"p"},"main.go")," function with the following structure:"),(0,o.kt)("ul",null,(0,o.kt)("li",{parentName:"ul"},"First, an ",(0,o.kt)("a",{parentName:"li",href:"/v0.47/core/encoding"},(0,o.kt)("inlineCode",{parentName:"a"},"encodingCodec"))," is instantiated for the application."),(0,o.kt)("li",{parentName:"ul"},"Then, the ",(0,o.kt)("inlineCode",{parentName:"li"},"config")," is retrieved and config parameters are set. This mainly involves setting the Bech32 prefixes for ",(0,o.kt)("a",{parentName:"li",href:"/v0.47/basics/accounts#addresses"},"addresses"),".")),(0,o.kt)("pre",null,(0,o.kt)("code",{parentName:"pre",className:"language-go",metastring:"reference",reference:!0},"https://github.com/cosmos/cosmos-sdk/blob/v0.47.0-rc1/types/config.go#L14-L29\n")),(0,o.kt)("ul",null,(0,o.kt)("li",{parentName:"ul"},"Using ",(0,o.kt)("a",{parentName:"li",href:"https://github.com/spf13/cobra"},"cobra"),", the root command of the full-node client is created. After that, all the custom commands of the application are added using the ",(0,o.kt)("inlineCode",{parentName:"li"},"AddCommand()")," method of ",(0,o.kt)("inlineCode",{parentName:"li"},"rootCmd"),"."),(0,o.kt)("li",{parentName:"ul"},"Add default server commands to ",(0,o.kt)("inlineCode",{parentName:"li"},"rootCmd")," using the ",(0,o.kt)("inlineCode",{parentName:"li"},"server.AddCommands()")," method. These commands are separated from the ones added above since they are standard and defined at Cosmos SDK level. They should be shared by all Cosmos SDK-based applications. They include the most important command: the ",(0,o.kt)("a",{parentName:"li",href:"#start-command"},(0,o.kt)("inlineCode",{parentName:"a"},"start")," command"),"."),(0,o.kt)("li",{parentName:"ul"},"Prepare and execute the ",(0,o.kt)("inlineCode",{parentName:"li"},"executor"),".")),(0,o.kt)("pre",null,(0,o.kt)("code",{parentName:"pre",className:"language-go",metastring:"reference",reference:!0},"https://github.com/tendermint/tendermint/blob/v0.37.0-rc2/libs/cli/setup.go#L74-L78\n")),(0,o.kt)("p",null,"See an example of ",(0,o.kt)("inlineCode",{parentName:"p"},"main")," function from the ",(0,o.kt)("inlineCode",{parentName:"p"},"simapp")," application, the Cosmos SDK's application for demo purposes:"),(0,o.kt)("pre",null,(0,o.kt)("code",{parentName:"pre",className:"language-go",metastring:"reference",reference:!0},"https://github.com/cosmos/cosmos-sdk/blob/v0.47.0-rc1/simapp/simd/main.go\n")),(0,o.kt)("h2",{id:"start-command"},(0,o.kt)("inlineCode",{parentName:"h2"},"start")," command"),(0,o.kt)("p",null,"The ",(0,o.kt)("inlineCode",{parentName:"p"},"start")," command is defined in the ",(0,o.kt)("inlineCode",{parentName:"p"},"/server")," folder of the Cosmos SDK. It is added to the root command of the full-node client in the ",(0,o.kt)("a",{parentName:"p",href:"#main-function"},(0,o.kt)("inlineCode",{parentName:"a"},"main")," function")," and called by the end-user to start their node:"),(0,o.kt)("pre",null,(0,o.kt)("code",{parentName:"pre",className:"language-bash"},'# For an example app named "app", the following command starts the full-node.\nappd start\n\n# Using the Cosmos SDK\'s own simapp, the following commands start the simapp node.\nsimd start\n')),(0,o.kt)("p",null,"As a reminder, the full-node is composed of three conceptual layers: the networking layer, the consensus layer and the application layer. The first two are generally bundled together in an entity called the consensus engine (Tendermint Core by default), while the third is the state-machine defined with the help of the Cosmos SDK. Currently, the Cosmos SDK uses Tendermint as the default consensus engine, meaning the start command is implemented to boot up a Tendermint node."),(0,o.kt)("p",null,"The flow of the ",(0,o.kt)("inlineCode",{parentName:"p"},"start")," command is pretty straightforward. First, it retrieves the ",(0,o.kt)("inlineCode",{parentName:"p"},"config")," from the ",(0,o.kt)("inlineCode",{parentName:"p"},"context")," in order to open the ",(0,o.kt)("inlineCode",{parentName:"p"},"db")," (a ",(0,o.kt)("a",{parentName:"p",href:"https://github.com/syndtr/goleveldb"},(0,o.kt)("inlineCode",{parentName:"a"},"leveldb"))," instance by default). This ",(0,o.kt)("inlineCode",{parentName:"p"},"db")," contains the latest known state of the application (empty if the application is started from the first time."),(0,o.kt)("p",null,"With the ",(0,o.kt)("inlineCode",{parentName:"p"},"db"),", the ",(0,o.kt)("inlineCode",{parentName:"p"},"start")," command creates a new instance of the application using an ",(0,o.kt)("inlineCode",{parentName:"p"},"appCreator")," function:"),(0,o.kt)("pre",null,(0,o.kt)("code",{parentName:"pre",className:"language-go",metastring:"reference",reference:!0},"https://github.com/cosmos/cosmos-sdk/blob/v0.47.0-rc1/server/start.go#L220\n")),(0,o.kt)("p",null,"Note that an ",(0,o.kt)("inlineCode",{parentName:"p"},"appCreator")," is a function that fulfills the ",(0,o.kt)("inlineCode",{parentName:"p"},"AppCreator")," signature:"),(0,o.kt)("pre",null,(0,o.kt)("code",{parentName:"pre",className:"language-go",metastring:"reference",reference:!0},"https://github.com/cosmos/cosmos-sdk/blob/v0.47.0-rc1/server/types/app.go#L64-L66\n")),(0,o.kt)("p",null,"In practice, the ",(0,o.kt)("a",{parentName:"p",href:"/v0.47/basics/app-anatomy#constructor-function"},"constructor of the application")," is passed as the ",(0,o.kt)("inlineCode",{parentName:"p"},"appCreator"),"."),(0,o.kt)("pre",null,(0,o.kt)("code",{parentName:"pre",className:"language-go",metastring:"reference",reference:!0},"https://github.com/cosmos/cosmos-sdk/blob/v0.47.0-rc1/simapp/simd/cmd/root.go#L254-L268\n")),(0,o.kt)("p",null,"Then, the instance of ",(0,o.kt)("inlineCode",{parentName:"p"},"app")," is used to instantiate a new Tendermint node:"),(0,o.kt)("pre",null,(0,o.kt)("code",{parentName:"pre",className:"language-go",metastring:"reference",reference:!0},"https://github.com/cosmos/cosmos-sdk/blob/v0.47.0-rc1/server/start.go#L336-L348\n")),(0,o.kt)("p",null,"The Tendermint node can be created with ",(0,o.kt)("inlineCode",{parentName:"p"},"app")," because the latter satisfies the ",(0,o.kt)("a",{parentName:"p",href:"https://github.com/tendermint/tendermint/blob/v0.37.0-rc2/abci/types/application.go#L9-L35"},(0,o.kt)("inlineCode",{parentName:"a"},"abci.Application")," interface")," (given that ",(0,o.kt)("inlineCode",{parentName:"p"},"app")," extends ",(0,o.kt)("a",{parentName:"p",href:"/v0.47/core/baseapp"},(0,o.kt)("inlineCode",{parentName:"a"},"baseapp")),"). As part of the ",(0,o.kt)("inlineCode",{parentName:"p"},"node.New")," method, Tendermint makes sure that the height of the application (i.e. number of blocks since genesis) is equal to the height of the Tendermint node. The difference between these two heights should always be negative or null. If it is strictly negative, ",(0,o.kt)("inlineCode",{parentName:"p"},"node.New")," will replay blocks until the height of the application reaches the height of the Tendermint node. Finally, if the height of the application is ",(0,o.kt)("inlineCode",{parentName:"p"},"0"),", the Tendermint node will call ",(0,o.kt)("a",{parentName:"p",href:"/v0.47/core/baseapp#initchain"},(0,o.kt)("inlineCode",{parentName:"a"},"InitChain"))," on the application to initialize the state from the genesis file."),(0,o.kt)("p",null,"Once the Tendermint node is instantiated and in sync with the application, the node can be started:"),(0,o.kt)("pre",null,(0,o.kt)("code",{parentName:"pre",className:"language-go",metastring:"reference",reference:!0},"https://github.com/cosmos/cosmos-sdk/blob/v0.47.0-rc1/server/start.go#L350-L352\n")),(0,o.kt)("p",null,"Upon starting, the node will bootstrap its RPC and P2P server and start dialing peers. During handshake with its peers, if the node realizes they are ahead, it will query all the blocks sequentially in order to catch up. Then, it will wait for new block proposals and block signatures from validators in order to make progress."),(0,o.kt)("h2",{id:"other-commands"},"Other commands"),(0,o.kt)("p",null,"To discover how to concretely run a node and interact with it, please refer to our ",(0,o.kt)("a",{parentName:"p",href:"/v0.47/run-node/run-node"},"Running a Node, API and CLI")," guide."))}c.isMDXComponent=!0}}]);