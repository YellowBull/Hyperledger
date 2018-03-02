# Hyperledger Fabric 是什么？
  Linux基金会在2015年创立了Hyperledger Fabric，以推进跨行业的区块链技术。他没有宣传单一的区块链标准，而是鼓励一种合作的方式，通过社区进程开发区块链技术，知识产权鼓励开放开发，并随着时间的推移采用关键标准。<br/>
  Hyperledger Fabric是Hyperledger上的区块链项目之一，就如同其他区块链技术一样，他有一个账本，使用智能合约，并且是一个由参与者管理他们交易的系统。<br/>
  与其他区块链系统最大的不同点在于Hyperledger Fabric是私有的，而且是被许可的，不是一个允许未知身份参与网络的开放的无许可的系统（要求协议验证事务并确保网络的安全），Hyperledger Fabric组织的成员可以通过一个Membership Service Provider（成员服务提供者即MSP）来注册。<br/>
  Hyperledger Fabric还提供了几个可插拔的组件。账本数据可以以多种格式存储，一致的机制可以被转换和输出，并且支持不同的MSPs<br/>
  Hyperledger Fabric也提供了创建通道的能力，允许一组参与者与创建一个单独的共同维护的交易账本。对于有些参与者可能是竞争对手的网络来说，这是一个特别重要的选择，他们不希望自己的每笔交易都能获得--例如，他们向一些参与者提供了一个特别的价格，而不是其他参与者。如果两个参与者形成一个通道，那么这些参与者--以及其他参与者--都有该渠道的分类账本。<br/> 
## 共享账本
  Hyperledger Fabirc有一个分类子系统，包括两个组成部分：世界状态（world state）和事务日志（transaction log）。每个参与
  者都有一份账本的副本刀他们所属的每一个Hyperledger Fabric的网络上。<br/>
  在给定的时间点上，世界状态（world state）组件描述了总账的状态。他是账本的数据库。事务日志（transaction log）组件记录所有
  导致当前世界状态值得事务。这是世界状态（world state）的更新历史。那么账本是世界状态（world state）数据库和事务日志（transaction log）
  历史组合。<br/>
  该账本为世界状态（world state）提供了可替换的数据库方案。默认情况下，这是一个LeveDB简直存储数据库。事务日志（transaction log）
  不需要是可插拔的，他只是记录了区块链网络使用的账本数据库之前和之后的值。
## 智能合约
  Hyperledger Fabric的智能契约是用Chaincode编写的，并且当应用程序需要与账本进行交互时，被应用程序外部的应用程序调用。在大多数情况下，
  Chaincode只与总账的数据库组件交互，例如世界状态（例如，查询他），而不是事务日志。<br/>
  Chaincode可疑用几种编程语言实现，目前支持的chaincode编写的是Go语言，在今后的发行版中将会逐步新增java和其他语言的支持。<br/>
## 隐私
  根据网络的需要，企业对企业（B2B）网络的参与者可能对他们所共享的信息非常敏感。对其他网络来说，隐私不会成为首要关注的问题。<br/>
  Hyperledger Fabric支持需要将隐私（使用通道）作为关键操作需求的网络，同时也是相对开放的网络。<br/>
## 共识
  事务必须按照他们发生的顺序写在账本上，即使他们可能是网络中不同的参与者生成的。要做到这一点，必须建立事务的顺序，并且必须
  在账本中建立一种拒绝错误事务（或恶意的）的方法。<br/>
  这是一个经过彻底研究的计算机科学领域，有很多方法可以实现它，每一个都有不同的权衡。例如，PBFT（拜占庭容错）可以为文件副本
  提供一种机制，使其能够相互通信，从而保持每个副本的一致性，即使是在出现腐化的情况下。或者，在比特币中，排序是通过一个名为“
  挖矿”的过程来实现的，在这个过程中，竞争的计算机竞相解决一个加密难题，该难题定义了所有流程随后构建的顺序。<br/>
  Hyperledger Fabric的设计使得网络启动者可以选择一种最能代表参与者之间关系的共识机制，就像隐私一样，需要有一系列的需求，从人际
  关系高度机构化的网络到更加对等的网络。<br/>
  关于Hyperledger Fabric共识机制，他目前包括SOLO和Kafka，并将很快拓展到SBFT（简化的拜占庭容错）。<br/>
## 身份管理
  为了支持许可的网络，Hyperledger Fabric提供了一个成员服务（membership identify service），它管理用户id并对网络上的所有
  参与者进行身份验证。访问控制列表可以通过特定网络操作的授权来提供额外的权限。例如，一个特定的用户ID可疑被允许调用一个链代码应用
  程序，但是阻止了部署新的链代码。关于Hyperledger Fabric网络的一个真理是，成员互相了解（身份），但他们不知道彼此在做什么（
  隐私和机密性）。<br/>