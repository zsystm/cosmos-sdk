"use strict";(self.webpackChunkcosmos_sdk_docs=self.webpackChunkcosmos_sdk_docs||[]).push([[9270],{3905:(e,t,a)=>{a.d(t,{Zo:()=>d,kt:()=>p});var i=a(7294);function n(e,t,a){return t in e?Object.defineProperty(e,t,{value:a,enumerable:!0,configurable:!0,writable:!0}):e[t]=a,e}function o(e,t){var a=Object.keys(e);if(Object.getOwnPropertySymbols){var i=Object.getOwnPropertySymbols(e);t&&(i=i.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),a.push.apply(a,i)}return a}function s(e){for(var t=1;t<arguments.length;t++){var a=null!=arguments[t]?arguments[t]:{};t%2?o(Object(a),!0).forEach((function(t){n(e,t,a[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(a)):o(Object(a)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(a,t))}))}return e}function r(e,t){if(null==e)return{};var a,i,n=function(e,t){if(null==e)return{};var a,i,n={},o=Object.keys(e);for(i=0;i<o.length;i++)a=o[i],t.indexOf(a)>=0||(n[a]=e[a]);return n}(e,t);if(Object.getOwnPropertySymbols){var o=Object.getOwnPropertySymbols(e);for(i=0;i<o.length;i++)a=o[i],t.indexOf(a)>=0||Object.prototype.propertyIsEnumerable.call(e,a)&&(n[a]=e[a])}return n}var l=i.createContext({}),h=function(e){var t=i.useContext(l),a=t;return e&&(a="function"==typeof e?e(t):s(s({},t),e)),a},d=function(e){var t=h(e.components);return i.createElement(l.Provider,{value:t},e.children)},c={inlineCode:"code",wrapper:function(e){var t=e.children;return i.createElement(i.Fragment,{},t)}},u=i.forwardRef((function(e,t){var a=e.components,n=e.mdxType,o=e.originalType,l=e.parentName,d=r(e,["components","mdxType","originalType","parentName"]),u=h(a),p=n,g=u["".concat(l,".").concat(p)]||u[p]||c[p]||o;return a?i.createElement(g,s(s({ref:t},d),{},{components:a})):i.createElement(g,s({ref:t},d))}));function p(e,t){var a=arguments,n=t&&t.mdxType;if("string"==typeof e||n){var o=a.length,s=new Array(o);s[0]=u;var r={};for(var l in t)hasOwnProperty.call(t,l)&&(r[l]=t[l]);r.originalType=e,r.mdxType="string"==typeof e?e:n,s[1]=r;for(var h=2;h<o;h++)s[h]=a[h];return i.createElement.apply(null,s)}return i.createElement.apply(null,a)}u.displayName="MDXCreateElement"},3084:(e,t,a)=>{a.r(t),a.d(t,{assets:()=>l,contentTitle:()=>s,default:()=>c,frontMatter:()=>o,metadata:()=>r,toc:()=>h});var i=a(7462),n=(a(7294),a(3905));const o={},s="ADR 039: Epoched Staking",r={unversionedId:"architecture/adr-039-epoched-staking",id:"version-v0.47/architecture/adr-039-epoched-staking",title:"ADR 039: Epoched Staking",description:"Changelog",source:"@site/versioned_docs/version-v0.47/architecture/adr-039-epoched-staking.md",sourceDirName:"architecture",slug:"/architecture/adr-039-epoched-staking",permalink:"/v0.47/architecture/adr-039-epoched-staking",draft:!1,tags:[],version:"v0.47",frontMatter:{},sidebar:"tutorialSidebar",previous:{title:"ADR 038: KVStore state listening",permalink:"/v0.47/architecture/adr-038-state-listening"},next:{title:"ADR 040: Storage and SMT State Commitments",permalink:"/v0.47/architecture/adr-040-storage-and-smt-state-commitments"}},l={},h=[{value:"Changelog",id:"changelog",level:2},{value:"Authors",id:"authors",level:2},{value:"Status",id:"status",level:2},{value:"Abstract",id:"abstract",level:2},{value:"Context",id:"context",level:2},{value:"Design considerations",id:"design-considerations",level:2},{value:"Slashing",id:"slashing",level:3},{value:"Token lockup",id:"token-lockup",level:3},{value:"Pipelining the epochs",id:"pipelining-the-epochs",level:3},{value:"Rewards",id:"rewards",level:3},{value:"Parameterizing the epoch length",id:"parameterizing-the-epoch-length",level:3},{value:"Decision",id:"decision",level:2},{value:"Staking messages",id:"staking-messages",level:3},{value:"Slashing messages",id:"slashing-messages",level:3},{value:"Evidence Messages",id:"evidence-messages",level:3},{value:"Consequences",id:"consequences",level:2},{value:"Positive",id:"positive",level:3},{value:"Negative",id:"negative",level:3}],d={toc:h};function c(e){let{components:t,...a}=e;return(0,n.kt)("wrapper",(0,i.Z)({},d,a,{components:t,mdxType:"MDXLayout"}),(0,n.kt)("h1",{id:"adr-039-epoched-staking"},"ADR 039: Epoched Staking"),(0,n.kt)("h2",{id:"changelog"},"Changelog"),(0,n.kt)("ul",null,(0,n.kt)("li",{parentName:"ul"},"10-Feb-2021: Initial Draft")),(0,n.kt)("h2",{id:"authors"},"Authors"),(0,n.kt)("ul",null,(0,n.kt)("li",{parentName:"ul"},"Dev Ojha (@valardragon)"),(0,n.kt)("li",{parentName:"ul"},"Sunny Aggarwal (@sunnya97)")),(0,n.kt)("h2",{id:"status"},"Status"),(0,n.kt)("p",null,"Proposed"),(0,n.kt)("h2",{id:"abstract"},"Abstract"),(0,n.kt)("p",null,"This ADR updates the proof of stake module to buffer the staking weight updates for a number of blocks before updating the consensus' staking weights. The length of the buffer is dubbed an epoch. The prior functionality of the staking module is then a special case of the abstracted module, with the epoch being set to 1 block."),(0,n.kt)("h2",{id:"context"},"Context"),(0,n.kt)("p",null,"The current proof of stake module takes the design decision to apply staking weight changes to the consensus engine immediately. This means that delegations and unbonds get applied immediately to the validator set. This decision was primarily done as it was implementationally simplest, and because we at the time believed that this would lead to better UX for clients."),(0,n.kt)("p",null,"An alternative design choice is to allow buffering staking updates (delegations, unbonds, validators joining) for a number of blocks. This 'epoch'd proof of stake consensus provides the guarantee that the consensus weights for validators will not change mid-epoch, except in the event of a slash condition."),(0,n.kt)("p",null,"Additionally, the UX hurdle may not be as significant as was previously thought. This is because it is possible to provide users immediate acknowledgement that their bond was recorded and will be executed."),(0,n.kt)("p",null,"Furthermore, it has become clearer over time that immediate execution of staking events comes with limitations, such as:"),(0,n.kt)("ul",null,(0,n.kt)("li",{parentName:"ul"},(0,n.kt)("p",{parentName:"li"},"Threshold based cryptography. One of the main limitations is that because the validator set can change so regularly, it makes the running of multiparty computation by a fixed validator set difficult. Many threshold-based cryptographic features for blockchains such as randomness beacons and threshold decryption require a computationally-expensive DKG process (will take much longer than 1 block to create). To productively use these, we need to guarantee that the result of the DKG will be used for a reasonably long time. It wouldn't be feasible to rerun the DKG every block. By epoching staking, it guarantees we'll only need to run a new DKG once every epoch.")),(0,n.kt)("li",{parentName:"ul"},(0,n.kt)("p",{parentName:"li"},"Light client efficiency. This would lessen the overhead for IBC when there is high churn in the validator set. In the Tendermint light client bisection algorithm, the number of headers you need to verify is related to bounding the difference in validator sets between a trusted header and the latest header. If the difference is too great, you verify more header in between the two. By limiting the frequency of validator set changes, we can reduce the worst case size of IBC lite client proofs, which occurs when a validator set has high churn.")),(0,n.kt)("li",{parentName:"ul"},(0,n.kt)("p",{parentName:"li"},"Fairness of deterministic leader election. Currently we have no ways of reasoning of fairness of deterministic leader election in the presence of staking changes without epochs (tendermint/spec#217). Breaking fairness of leader election is profitable for validators, as they earn additional rewards from being the proposer. Adding epochs at least makes it easier for our deterministic leader election to match something we can prove secure. (Albeit, we still haven\u2019t proven if our current algorithm is fair with > 2 validators in the presence of stake changes)")),(0,n.kt)("li",{parentName:"ul"},(0,n.kt)("p",{parentName:"li"},"Staking derivative design. Currently, reward distribution is done lazily using the F1 fee distribution. While saving computational complexity, lazy accounting requires a more stateful staking implementation. Right now, each delegation entry has to track the time of last withdrawal. Handling this can be a challenge for some staking derivatives designs that seek to provide fungibility for all tokens staked to a single validator. Force-withdrawing rewards to users can help solve this, however it is infeasible to force-withdraw rewards to users on a per block basis. With epochs, a chain could more easily alter the design to have rewards be forcefully withdrawn (iterating over delegator accounts only once per-epoch), and can thus remove delegation timing from state. This may be useful for certain staking derivative designs."))),(0,n.kt)("h2",{id:"design-considerations"},"Design considerations"),(0,n.kt)("h3",{id:"slashing"},"Slashing"),(0,n.kt)("p",null,"There is a design consideration for whether to apply a slash immediately or at the end of an epoch. A slash event should apply to only members who are actually staked during the time of the infraction, namely during the epoch the slash event occured."),(0,n.kt)("p",null,"Applying it immediately can be viewed as offering greater consensus layer security, at potential costs to the aforementioned usecases. The benefits of immediate slashing for consensus layer security can be all be obtained by executing the validator jailing immediately (thus removing it from the validator set), and delaying the actual slash change to the validator's weight until the epoch boundary. For the use cases mentioned above, workarounds can be integrated to avoid problems, as follows:"),(0,n.kt)("ul",null,(0,n.kt)("li",{parentName:"ul"},"For threshold based cryptography, this setting will have the threshold cryptography use the original epoch weights, while consensus has an update that lets it more rapidly benefit from additional security. If the threshold based cryptography blocks liveness of the chain, then we have effectively raised the liveness threshold of the remaining validators for the rest of the epoch. (Alternatively, jailed nodes could still contribute shares) This plan will fail in the extreme case that more than 1/3rd of the validators have been jailed within a single epoch. For such an extreme scenario, the chain already have its own custom incident response plan, and defining how to handle the threshold cryptography should be a part of that."),(0,n.kt)("li",{parentName:"ul"},"For light client efficiency, there can be a bit included in the header indicating an intra-epoch slash (ala ",(0,n.kt)("a",{parentName:"li",href:"https://github.com/tendermint/spec/issues/199"},"https://github.com/tendermint/spec/issues/199"),")."),(0,n.kt)("li",{parentName:"ul"},"For fairness of deterministic leader election, applying a slash or jailing within an epoch would break the guarantee we were seeking to provide. This then re-introduces a new (but significantly simpler) problem for trying to provide fairness guarantees. Namely, that validators can adversarially elect to remove themself from the set of proposers. From a security perspective, this could potentially be handled by two different mechanisms (or prove to still be too difficult to achieve). One is making a security statement acknowledging the ability for an adversary to force an ahead-of-time fixed threshold of users to drop out of the proposer set within an epoch. The second method would be to  parameterize such that the cost of a slash within the epoch far outweights benefits due to being a proposer. However, this latter criterion is quite dubious, since being a proposer can have many advantageous side-effects in chains with complex state machines. (Namely, DeFi games such as Fomo3D)"),(0,n.kt)("li",{parentName:"ul"},"For staking derivative design, there is no issue introduced. This does not increase the state size of staking records, since whether a slash has occured is fully queryable given the validator address.")),(0,n.kt)("h3",{id:"token-lockup"},"Token lockup"),(0,n.kt)("p",null,"When someone makes a transaction to delegate, even though they are not immediately staked, their tokens should be moved into a pool managed by the staking module which will then be used at the end of an epoch. This prevents concerns where they stake, and then spend those tokens not realizing they were already allocated for staking, and thus having their staking tx fail."),(0,n.kt)("h3",{id:"pipelining-the-epochs"},"Pipelining the epochs"),(0,n.kt)("p",null,"For threshold based cryptography in particular, we need a pipeline for epoch changes. This is because when we are in epoch N, we want the epoch N+1 weights to be fixed so that the validator set can do the DKG accordingly. So if we are currently in epoch N, the stake weights for epoch N+1 should already be fixed, and new stake changes should be getting applied to epoch N + 2."),(0,n.kt)("p",null,"This can be handled by making a parameter for the epoch pipeline length. This parameter should not be alterable except during hard forks, to mitigate implementation complexity of switching the pipeline length."),(0,n.kt)("p",null,"With pipeline length 1, if I redelegate during epoch N, then my redelegation is applied prior to the beginning of epoch N+1.\nWith pipeline length 2, if I redelegate during epoch N, then my redelegation is applied prior to the beginning of epoch N+2."),(0,n.kt)("h3",{id:"rewards"},"Rewards"),(0,n.kt)("p",null,"Even though all staking updates are applied at epoch boundaries, rewards can still be distributed immediately when they are claimed. This is because they do not affect the current stake weights, as we do not implement auto-bonding of rewards. If such a feature were to be implemented, it would have to be setup so that rewards are auto-bonded at the epoch boundary."),(0,n.kt)("h3",{id:"parameterizing-the-epoch-length"},"Parameterizing the epoch length"),(0,n.kt)("p",null,"When choosing the epoch length, there is a trade-off queued state/computation buildup, and countering the previously discussed limitations of immediate execution if they apply to a given chain."),(0,n.kt)("p",null,"Until an ABCI mechanism for variable block times is introduced, it is ill-advised to be using high epoch lengths due to the computation buildup. This is because when a block's execution time is greater than the expected block time from Tendermint, rounds may increment."),(0,n.kt)("h2",{id:"decision"},"Decision"),(0,n.kt)("p",null,(0,n.kt)("strong",{parentName:"p"},"Step-1"),":  Implement buffering of all staking and slashing messages."),(0,n.kt)("p",null,"First we create a pool for storing tokens that are being bonded, but should be applied at the epoch boundary called the ",(0,n.kt)("inlineCode",{parentName:"p"},"EpochDelegationPool"),". Then, we have two separate queues, one for staking, one for slashing. We describe what happens on each message being delivered below:"),(0,n.kt)("h3",{id:"staking-messages"},"Staking messages"),(0,n.kt)("ul",null,(0,n.kt)("li",{parentName:"ul"},(0,n.kt)("strong",{parentName:"li"},"MsgCreateValidator"),": Move user's self-bond to ",(0,n.kt)("inlineCode",{parentName:"li"},"EpochDelegationPool")," immediately. Queue a message for the epoch boundary to handle the self-bond, taking the funds from the ",(0,n.kt)("inlineCode",{parentName:"li"},"EpochDelegationPool"),". If Epoch execution fail, return back funds from ",(0,n.kt)("inlineCode",{parentName:"li"},"EpochDelegationPool")," to user's account."),(0,n.kt)("li",{parentName:"ul"},(0,n.kt)("strong",{parentName:"li"},"MsgEditValidator"),": Validate message and if valid queue the message for execution at the end of the Epoch."),(0,n.kt)("li",{parentName:"ul"},(0,n.kt)("strong",{parentName:"li"},"MsgDelegate"),": Move user's funds to ",(0,n.kt)("inlineCode",{parentName:"li"},"EpochDelegationPool")," immediately. Queue a message for the epoch boundary to handle the delegation, taking the funds from the ",(0,n.kt)("inlineCode",{parentName:"li"},"EpochDelegationPool"),". If Epoch execution fail, return back funds from ",(0,n.kt)("inlineCode",{parentName:"li"},"EpochDelegationPool")," to user's account."),(0,n.kt)("li",{parentName:"ul"},(0,n.kt)("strong",{parentName:"li"},"MsgBeginRedelegate"),": Validate message and if valid queue the message for execution at the end of the Epoch."),(0,n.kt)("li",{parentName:"ul"},(0,n.kt)("strong",{parentName:"li"},"MsgUndelegate"),": Validate message and if valid queue the message for execution at the end of the Epoch.")),(0,n.kt)("h3",{id:"slashing-messages"},"Slashing messages"),(0,n.kt)("ul",null,(0,n.kt)("li",{parentName:"ul"},(0,n.kt)("strong",{parentName:"li"},"MsgUnjail"),": Validate message and if valid queue the message for execution at the end of the Epoch."),(0,n.kt)("li",{parentName:"ul"},(0,n.kt)("strong",{parentName:"li"},"Slash Event"),": Whenever a slash event is created, it gets queued in the slashing module to apply at the end of the epoch. The queues should be setup such that this slash applies immediately.")),(0,n.kt)("h3",{id:"evidence-messages"},"Evidence Messages"),(0,n.kt)("ul",null,(0,n.kt)("li",{parentName:"ul"},(0,n.kt)("strong",{parentName:"li"},"MsgSubmitEvidence"),": This gets executed immediately, and the validator gets jailed immediately. However in slashing, the actual slash event gets queued.")),(0,n.kt)("p",null,"Then we add methods to the end blockers, to ensure that at the epoch boundary the queues are cleared and delegation updates are applied."),(0,n.kt)("p",null,(0,n.kt)("strong",{parentName:"p"},"Step-2"),": Implement querying of queued staking txs."),(0,n.kt)("p",null,"When querying the staking activity of a given address, the status should return not only the amount of tokens staked, but also if there are any queued stake events for that address. This will require more work to be done in the querying logic, to trace the queued upcoming staking events."),(0,n.kt)("p",null,"As an initial implementation, this can be implemented as a linear search over all queued staking events. However, for chains that need long epochs, they should eventually build additional support for nodes that support querying to be able to produce results in constant time. (This is do-able by maintaining an auxilliary hashmap for indexing upcoming staking events by address)"),(0,n.kt)("p",null,(0,n.kt)("strong",{parentName:"p"},"Step-3"),": Adjust gas"),(0,n.kt)("p",null,"Currently gas represents the cost of executing a transaction when its done immediately. (Merging together costs of p2p overhead, state access overhead, and computational overhead) However, now a transaction can cause computation in a future block, namely at the epoch boundary."),(0,n.kt)("p",null,"To handle this, we should initially include parameters for estimating the amount of future computation (denominated in gas), and add that as a flat charge needed for the message.\nWe leave it as out of scope for how to weight future computation versus current computation in gas pricing, and have it set such that the are weighted equally for now."),(0,n.kt)("h2",{id:"consequences"},"Consequences"),(0,n.kt)("h3",{id:"positive"},"Positive"),(0,n.kt)("ul",null,(0,n.kt)("li",{parentName:"ul"},"Abstracts the proof of stake module that allows retaining the existing functionality"),(0,n.kt)("li",{parentName:"ul"},"Enables new features such as validator-set based threshold cryptography")),(0,n.kt)("h3",{id:"negative"},"Negative"),(0,n.kt)("ul",null,(0,n.kt)("li",{parentName:"ul"},"Increases complexity of integrating more complex gas pricing mechanisms, as they now have to consider future execution costs as well."),(0,n.kt)("li",{parentName:"ul"},"When epoch > 1, validators can no longer leave the network immediately, and must wait until an epoch boundary.")))}c.isMDXComponent=!0}}]);