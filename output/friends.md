# My Friends
- [戴兜｜Coding the world.](https://im.daidr.me)
- [Bird｜aFlyBird0](https://blog.aflybird.cn/)
- [NX｜这家伙真勤奋，什么都留下了](https://nickxu.me/)
- [柏喵Sakura｜过气 emo 师傅](https://baimeow.cn/)
- [Marlene｜去不了异世界也要拿出真本事~](https://blog.marlene.top/)
- [ek1ng｜Hidden Gem](https://ek1ng.com/)
- [Atom｜aka asjdf](https://www.homeboyc.cn/)
- [xyxsw｜大相的问题](https://xyxsw.ltd/)
- [hakuya｜恋恋真可爱，嘿嘿嘿。。。🤤🤤🤤。。。](https://www.hakuya.work)

# Recent Blogs
## [你还惦记着你那二面呢｜Allow Everything to Happen](https://nickxu.me/post/allow-everything-to-happen.html)  by [NX](https://nickxu.me/), 2023-10-12

「你还惦记着你那二面呢」我真的好累，今天还吐了，感觉病了我消失了很长一段时间，本来我想写一篇长文，记录我从 8 月 22 号返校到现在发生的所有事情但是我真的好累，简单地说，就是诸事不顺，面试、亚运、学业…我不想写我总是想在最后来一波升华，但是我其实并没有走出来，架构不出，我写不出来我太想成功了，但这正是我没法成功的原因就像你找东西，你越找就越找不到，经常是后面不经意的时候突然就冒出来了就我现在这
## [浅析 JWT Refresh Token](https://nickxu.me/post/jwt-refresh-token.html)  by [NX](https://nickxu.me/), 2023-10-10

最近逛 V 站又看见有人在讨论 JWT，感觉很多人讲的很乱，我想简单记录下我印象中的理解Session最开始应该是 session 方案，就是用户登陆后服务器返回客户端一个存根（token）来标识当前的会话服务器在缓存中保存这个 token，用户请求时需要传递这个 token（不管你是使用 cookie header 还是 query）每次请求服务器都在缓存中查找这个 token，并且找到当前会话
## [在实验网络中使用 step-ca 签发 SSL 证书](https://baimeow.cn/posts/dn11/acme/)  by [柏喵Sakura](https://baimeow.cn/), 2023-10-08

在真实的互联网中，我们常常会通过校验 SSL 证书的方式创建加密 socket 来访问互联网上的资源，这能够有效阻止他人看到通信的内容。而在 DN11 中，由于每个人都相当于是 ISP，流量可能途径每一个人，随时可能被人抓取，因此使用 SSL 证书这件事就更为重要了。本文将在使用 PDNS 搭建好实验性网络 DNS 的基础上使用 step-ca 搭建 ACME Server 并使用 acme.s
## [[简中] Google Summer of Code & Chrome Extensions](https://daidr.me/archives/code-1097.html)  by [戴兜](https://im.daidr.me), 2023-09-25

这篇博客是 https://im.daidr.me/blog/1086 的简中翻译版本。原文地址：https://developer.chrome.com/blog/google-summer-of-code-and-chrome-extensions/我是一名来自中国的大二学生，对Web开发非常感兴趣。在我大一的时候，我加入了一个技术社团（杭电助手），这可以说是让我接触编程和
## [人生好比修仙](https://blog.aflybird.cn/2023/09/life-is-cultivating-immortality/)  by [Bird](https://blog.aflybird.cn/), 2023-09-24

最近看《凡人修仙传》魔怔了，自认为「悟」出了点人生与修仙的关联。 声明：本人是纯正的唯物主义者，信仰科学。以下对于修仙的描述皆为文学手法。 凡人
## [IBGP FullMesh 实现多节点自治域](https://baimeow.cn/posts/dn11/configureibgp/)  by [柏喵Sakura](https://baimeow.cn/), 2023-09-23

最近不少群友有把自己的其他服务器也接入自己的 AS 的需求，那这篇教程也是时候该出了。总览 这个事情比较复杂，我们先不急着开始配，先了解一下我们需要一个什么效果。DN11 作为一个发展了有一段时间的实验性网络，我们目前有很大一批人，他们手上都已经有了一个节点，并且在这个节点上配置了 BGP 从而联通到其他节点。那这个情景下，他们的自治域就是由那一个节点组成的单节点自治域。我们现在需要添加一个新
## [DN11 使用 EBGP 重分发 IBGP OSPF（同as内部peer）](https://xyxsw.ltd/2023/09/20/DN11%20%E4%BD%BF%E7%94%A8%20EBGP%20%E9%87%8D%E5%88%86%E5%8F%91%20IBGP%20OSPF%EF%BC%88%E5%90%8Cas%E5%86%85%E9%83%A8peer%EF%BC%89/)  by [xyxsw](https://xyxsw.ltd/), 2023-09-19

DN11 使用 EBGP 重分发 IBGP OSPF（同as内部peer）连接 wg peer1.本地机器# wg_aws.conf[Interface]# 你的私钥PrivateKey = <PRIVATEKAY># 你开的端口号（注意防火墙）ListenPort = <PORT>PostUp = /sbin/ip addr add dev %i <your ip> peer <peer ip>
## [校园网打洞纪实，绕个普普通通的防火墙](https://baimeow.cn/posts/dn11/travelihdu/)  by [柏喵Sakura](https://baimeow.cn/), 2023-09-16

WireGuard 在 i-HDU 之死 突然集体下线了 前几天接到悲报，说是 Vidar-Team 300b 节点从 DN11 下线了校园网需要走web认证登陆，所以起初我以为是校园网的登陆脚本炸了，找个人把自己账号登上去救一下急就好了后来派煎包去看看，他回来告诉我，Vidar 的 WiFi 有网，网很好校园网也登着这就很让人迷惑了，按理说有网的情况下，万万不该这么多隧道一起炸了由于是暑
## [2023OSPP大作戦、完結です!](https://blog.marlene.top/index.php/develop/93.html)  by [Marlene](https://blog.marlene.top/), 2023-09-16

项目信息项目名称：AliceBot 插件商店实现项目产出要求：编写一个插件商店用于用户分享自己编写插件和适配器使用 GitHub Workflow 和 GitHub App 自动获取用户提交的信息并更新商店页面界面美观，易于使用。时间规划：2023.7-2023.9项目进度根据项目计划书内容和项目导师要求，本项目已全部完成，功能内容测试完成，已做相关优化。插件商店页面部分要求描述AliceBot使
## [Markdown editor CVE of Marktext](https://www.ek1ng.com/MarkdownEditorCVE1.html)  by [ek1ng](https://ek1ng.com/), 2023-09-13

无CVE编号 XSS2RCEhttps://github.com/marktext/marktext/issues/2601https://github.com/marktext/marktext/commit/0dd09cc6842d260528c98151c394c5f63d733b62影响 <= 0.16.3 的marktext版本，点击链接触发。POC：123<a href="javasc
## [『LeetCode-HOT-100』T61～T70](https://nickxu.me/post/leetcode-hot-100-t61-t70.html)  by [NX](https://nickxu.me/), 2023-09-12

207. 课程表拓扑排序板子1234567891011121314151617181920212223242526272829303132333435363738394041func canFinish(numCourses int, prerequisites [][]int) bool {    n:=len(prerequisites)    rd:=make([]int,numCourse
## [『LeetCode-HOT-100』T51～T60](https://nickxu.me/post/leetcode-hot-100-t51-t60.html)  by [NX](https://nickxu.me/), 2023-09-12

142. 环形链表 II请见 『算法拾遗』链表（Linked List）146. LRU 缓存标准写法12345678910111213141516171819202122232425262728293031323334353637383940414243444546474849505152535455565758596061626364656667686970717273747576777879
## [Zeabur CLI 从设计到实现](https://blog.aflybird.cn/2023/09/zeabur-cli/)  by [Bird](https://blog.aflybird.cn/), 2023-09-07

What is Zeabur? Why We use Zeabur? 也许你像我一样，热爱编程。用技术改变世界是多么酷的事情！ 但是，我们又不得不花大把时间在服务部署上，无论是配置基础设施，还是频繁
## [字节二面挂，还是人太菜了](https://nickxu.me/post/bytedance-interview-failed-2023-09.html)  by [NX](https://nickxu.me/), 2023-09-07

字节一面自我介绍简单介绍字节青训营项目是组队的吗项目耗时项目收获的点ELK 是你们搭建的吗ELK 的软件安装数据流大概是怎样的是通过什么写到 Logstash 里的Logstash 的功能你们用的的 fail2ban 是什么traceID 介绍微服务框架用的什么traceID 在框架中是怎么传递的对于异步的请求怎么处理的这个项目的挑战和难点Golang 的 Panic 关键字Panic 怎么恢复不
## [浅谈代码模块化设计](https://homeboyc.cn/blog/%E6%B5%85%E8%B0%88%E4%BB%A3%E7%A0%81%E6%A8%A1%E5%9D%97%E5%8C%96%E8%AE%BE%E8%AE%A1/)  by [Atom](https://www.homeboyc.cn/), 2023-09-07

# 前言 最近遇到了一些开发者使用团队内部框架时不太能够接受代码模块化设计理念的问题，因此特此开一篇文章浅谈在软件架构设计中，一些模块化的设计以及复用原则，以便于开发者更好的了解框架模块化设计的思想。# What‘s 模块？ 模块（module）也可以理解为组件，是系统中可服用的功能逻辑封装单位。在不同的架构设计中，模块的概念可能会有所不同。在Golang中，我们通常将一个package称为一个
## [Google Summer of Code & Chrome Extensions](https://daidr.me/archives/code-1086.html)  by [戴兜](https://im.daidr.me), 2023-09-04

Original PostI’m a sophomore from China passionate about web development. In my first year, I joined a technical club at our college. This club was my introduction to coding and open source. In th
## [go-zero api文件生成项目框架 ｜ 青训营笔记](https://xyxsw.ltd/2023/08/30/go-zero%20api%E6%96%87%E4%BB%B6%20%EF%BD%9C%20%E9%9D%92%E8%AE%AD%E8%90%A5%E7%AC%94%E8%AE%B0/)  by [xyxsw](https://xyxsw.ltd/), 2023-08-30

go-zero api文件生成项目框架数据结构.api文件是go-zero自创的文件格式，和protobuf的语法有很大不同，但好在不难理解。具体语法介绍可看官网文档：go-zero.dev/docs/tasks/…先编写接口定义，然后直接一次性使用 goctl 生成代码。api文件支持将一个数据结构嵌入另一个结构中，便于编写统一的响应格式。首先定义空请求结构和基础的响应结构：syntax = "
## [gorm 进阶 ｜ 青训营笔记](https://xyxsw.ltd/2023/08/30/gorm%E8%BF%9B%E9%98%B6%20%EF%BD%9C%20%E9%9D%92%E8%AE%AD%E8%90%A5%E7%AC%94%E8%AE%B0/)  by [xyxsw](https://xyxsw.ltd/), 2023-08-30

gorm关于简单的 gorm 总结 可以看上一篇文章  gorm 初体验 ｜ 青训营笔记软删除Gorm提供了软删除的能力，需要在结构体中定义一个Deleted字段，此时再调用Delete删除函数，则会生成update语句，并将deleted字段赋值为当前删除时间。type Product struct {    ID      uint    Code    string    Price   u
## [渗透测试思路总结](https://www.ek1ng.com/Summary%20of%20penetration%20ideas.html)  by [ek1ng](https://ek1ng.com/), 2023-08-29

最近做了一阵子攻防相关的事，正好最近国护结束，做个总结，简单写一下渗透的基本思路（Check List）。不同的标题间内容并不完全独立，在实战中，比如先钓鱼获取到一台个人PC，但这台PC并不在办公网。而后通过收集个人PC的信息，能够登陆外网其他站点的后台，配合一个后台RCE进入办公网/生产网。这其中就有钓鱼，也有外网打点的部分。资产收集资产搜集通俗说就是“了解目标有什么东西”，讲究一个越全越好。路
## [go字节三件套 ｜ 青训营笔记](https://xyxsw.ltd/2023/08/28/go%E5%AD%97%E8%8A%82%E4%B8%89%E4%BB%B6%E5%A5%97%20%EF%BD%9C%20%E9%9D%92%E8%AE%AD%E8%90%A5%E7%AC%94%E8%AE%B0/)  by [xyxsw](https://xyxsw.ltd/), 2023-08-28

Gorm、Kitex、Hertz：三件套介绍与基本用法GormGorm是Golang中广受欢迎的ORM（对象关系映射）框架，已经发展数十年，具有强大的功能和出色的性能。ORM框架用于将面向对象的概念与数据库中的表相对应，简化数据操作过程。在Golang中，自定义的结构体与数据库表一一对应，结构体的实例对应表中的一条记录。基本用法定义结构体：在Gorm中，定义结构体来映射数据库表。type User
