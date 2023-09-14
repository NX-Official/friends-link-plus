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
## [Markdown Editor CVE  - Marktext](https://www.ek1ng.com/MarkdownEditorCVE1.html)  by [ek1ng](https://ek1ng.com/), 2023-09-13

无CVE编号 XSS2RCEhttps://github.com/marktext/marktext/issues/2601https://github.com/marktext/marktext/commit/0dd09cc6842d260528c98151c394c5f63d733b62影响 <= 0.16.3 的marktext版本，点击链接触发。POC：123<a href="javasc
## [Zeabur CLI 从设计到实现](https://blog.aflybird.cn/2023/09/zeabur-cli/)  by [Bird](https://blog.aflybird.cn/), 2023-09-07

What is Zeabur? Why We use Zeabur? 也许你像我一样，热爱编程。用技术改变世界是多么酷的事情！ 但是，我们又不得不花大把时间在服务部署上，无论是配置基础设施，还是频繁
## [字节二面挂，还是人太菜了](https://nickxu.me/post/bytedance-interview-failed-2023-09.html)  by [NX](https://nickxu.me/), 2023-09-07

字节一面自我介绍简单介绍字节青训营项目是组队的吗项目耗时项目收获的点ELK 是你们搭建的吗ELK 的软件安装数据流大概是怎样的是通过什么写到 Logstash 里的Logstash 的功能你们用的的 fail2ban 是什么traceID 介绍微服务框架用的什么traceID 在框架中是怎么传递的对于异步的请求怎么处理的这个项目的挑战和难点Golang 的 Panic 关键字Panic 怎么恢复不
## [浅谈代码模块化设计](https://homeboyc.cn/blog/%E6%B5%85%E8%B0%88%E4%BB%A3%E7%A0%81%E6%A8%A1%E5%9D%97%E5%8C%96%E8%AE%BE%E8%AE%A1/)  by [Atom](https://www.homeboyc.cn/), 2023-09-07

# 前言 最近遇到了一些开发者使用团队内部框架时不太能够接受代码模块化设计理念的问题，因此特此开一篇文章浅谈在软件架构设计中，一些模块化的设计以及复用原则，以便于开发者更好的了解框架模块化设计的思想。# What‘s 模块？ 模块（module）也可以理解为组件，是系统中可服用的功能逻辑封装单位。在不同的架构设计中，模块的概念可能会有所不同。在Golang中，我们通常将一个package称为一个
## [Google Summer of Code & Chrome Extensions](https://daidr.me/archives/code-1086.html)  by [戴兜](https://im.daidr.me), 2023-09-04

Original PostI’m a sophomore from China passionate about web development. In my first year, I joined a technical club at our college. This club was my introduction to coding and open source. In th
## [渗透基本思路总结](https://www.ek1ng.com/Summary%20of%20penetration%20ideas.html)  by [ek1ng](https://ek1ng.com/), 2023-08-29

最近做了一阵子攻防相关的事，正好最近国护结束，做个总结，简单写一下渗透的基本思路（Check List）。不同的标题间内容并不完全独立，在实战中，比如先钓鱼获取到一台个人PC，但这台PC并不在办公网。而后通过收集个人PC的信息，能够登陆外网其他站点的后台，配合一个后台RCE进入办公网/生产网。这其中就有钓鱼，也有外网打点的部分。资产收集资产搜集通俗说就是“了解目标有什么东西”，讲究一个越全越好。路
## [🌟 你可能感兴趣的文章｜Posts you might be interested in](https://nickxu.me/posts-you-might-be-interested-in.html)  by [NX](https://nickxu.me/), 2023-08-21

Latest Resume｜最近在找实习哦Junior｜大三即将到来的生活，充满未知与期待Sophomore｜大二阿里云OSS被刷，我交了1000RMB学费！『OSPP2023』我与 OSPP 的故事 —— 从听闻到中选2023五一总结：近况与将来告别ELK！轻量级日志收集系统Grafana Loki初上手第五届字节跳动青训营项目总结写在大二下开学之初『CI/CD』结合GitHub Actions
## [阿里云OSS被刷，我交了1000RMB学费！](https://nickxu.me/post/aliyun-oss-brushed-1000rmb-fees.html)  by [NX](https://nickxu.me/), 2023-08-19

大致经过垂死病中惊坐起😱事情发生在 8 月 8 日凌晨，凌晨三点我突然看见手机上的消息我一开始是疑惑的，我的 OSS 是用来当做图床的，一个月也用不了几个钱账号里记得还有 20 多块钱，怎么会这么快用完然后我进阿里云一看，哇，我被人刷了？最后发现被刷了 3.57 TB，请求了 138 万次哇，我从没想到过这种事情会发生在我的身上而且我停机之后他还一直在刷，根本不带停的（我想，算了，300 块交学费
## [『Golang』并发编程之通道（Channel）](https://nickxu.me/post/golang-concurrent-programming-channel.html)  by [NX](https://nickxu.me/), 2023-08-17

通道（Channel）通道是什么，为什么使用通道「不要通过共享内存来通信，而应该通过通信来共享内存」通道可以在多个 goroutine 之间传递数据一个通道相当于一个先进先出（FIFO）的队列。也就是说，通道中的各个元素值都是严格地按照发送的顺序排列的，先被发送通道的元素值一定会先被接收。元素值的发送和接收都需要用到操作符 <-。我们也可以叫它接送操作符。一个左尖括号紧接着一个减号形象地代表了元素
## [组合优于继承](https://homeboyc.cn/blog/%E7%BB%84%E5%90%88%E4%BC%98%E4%BA%8E%E7%BB%A7%E6%89%BF/)  by [Atom](https://www.homeboyc.cn/), 2023-08-17

# 前言 你可能听过“组合优于继承”这个观点，但是这个说法可能有点笼统，所以我想通过这篇文章详细解释一下：组合和继承分别是什么？为什么前者优于后者？# 继承和组合分别是什么 组合和继承都是为了解决同一个问题——“代码复用”。# 继承是什么 当一个类中有你想要复用的功能时，继承便会发生。我们通常会创建一个子类来扩张基类的功能，然后通过注入新的方法来拓展或重写基类的部件。// 基类 abstra
## [自我觉察与追寻内心宁静的碎碎念](https://blog.aflybird.cn/2023/08/the-pursuit-of-self-awareness-and-inner-peace/)  by [Bird](https://blog.aflybird.cn/), 2023-08-16

很久没有认真地去做一次自我觉察了。本篇文章更多的是我与自己的对话，如果读者也能从中获得一些启发，无论是增添了几分决心、还是悟到了一些觉察、改变的方式，我倍感荣幸。
## [『LeetCode-HOT-100』T41～T50](https://nickxu.me/post/leetcode-hot-100-t41-t50.html)  by [NX](https://nickxu.me/), 2023-08-13

二叉树的层序遍历简单的 BFS 练习12345678910111213141516171819202122232425262728293031func levelOrder(root *TreeNode) [][]int {    ans := [][]int{}    if root == nil {        return ans    }    queue := []TreeNode{}
## [源码分析——Go语言依赖注入库 samber/do](https://blog.aflybird.cn/2023/08/read-open-source-go-dependency-injection-library-samber-do/)  by [Bird](https://blog.aflybird.cn/), 2023-08-05

琢磨设计模式与抽象，可以说是我的最爱之一了。刚学 Go 的时候，我就陶醉于其的 interface 设计。这次，我们来聊聊 Go 语言的依赖注入（DI）库 samber/do。本文不是一行行分析源码，而是尝试一步步复现作者的设计思路。
## [『LeetCode-HOT-100』T31～T40](https://nickxu.me/post/leetcode-hot-100-t31-t40.html)  by [NX](https://nickxu.me/), 2023-08-05

颜色分类这真的是 Medium 吗，哈哈哈😂123456789101112131415161718192021222324func sortColors(nums []int) {    var red, white, blue int    for i := 0; i < len(nums); i++ {        switch nums[i] {        case 0:
## [『LeetCode-HOT-100』T21～T30](https://nickxu.me/post/leetcode-hot-100-t21-t30.html)  by [NX](https://nickxu.me/), 2023-08-03

全排列板子题，不解释123456789101112131415161718func permute(nums []int) (ans [][]int) {    var dfs func(begain, end int)    dfs = func(begain, end int) {        if begain == end {            // 切片是引用类型，需要深拷贝一下
## [『LeetCode-HOT-100』T11～T20](https://nickxu.me/post/leetcode-hot-100-t11-t20.html)  by [NX](https://nickxu.me/), 2023-07-30

有效的括号栈的经典题目了属于是12345678910111213141516171819202122232425262728293031323334353637383940func isValid(s string) bool {    stack := ""    for k := 0; k < len(s); k++ {        i := s[k]        switch i {
## [解决 gRPC 中 oneof 类型未导出的问题](https://homeboyc.cn/blog/%E8%A7%A3%E5%86%B3-grpc-%E4%B8%AD-oneof-%E7%B1%BB%E5%9E%8B%E6%9C%AA%E5%AF%BC%E5%87%BA%E7%9A%84%E9%97%AE%E9%A2%98/)  by [Atom](https://www.homeboyc.cn/), 2023-07-30

# 场景 proto 中有个项是 oneof，但是最新的 gRPC generator并没有将此字段的接口导出。Proto:message Setting { string item_id = 1; // 设置项的内部唯一id string item_label = 2; // 设置项的名称 SettingItemType item_type = 3; // 设置项的类型 oneof attr
## [LLM Agent之结构化输出](https://blog.marlene.top/index.php/develop/88.html)  by [Marlene](https://blog.marlene.top/), 2023-07-26

引言自去年年底以来，GPT的迅速发展诞生了一系列大模型。出现了更新、更大、更强的GPT-4。OpenAI不断推出GPT-4，ChatGPT Plugins，代码解释器，Function calling,图片处理等等。7月的WAIC上，笔者也有幸见到了国内一众企业相继展示自家的大模型。在这段时间里，LLM从最初的PE工程走向智能体交互。而笔者从最开始考虑LLM能不能多人协作，思考”一个专家完成所有任
## [Java RMI 攻击梳理总结](https://www.ek1ng.com/java-rmi-attack.html)  by [ek1ng](https://ek1ng.com/), 2023-07-26

RMI 是什么定义RMI（Remote Method Invocation）是远程方法调用，类似RPC（Remote Procedure Calls）。RPC是打包和传送数据结构，而在Java中，通常传递一个完整的对象，包含数据和操作数据的方法。通过RMI，能够让客户端JVM上的对象，像调用本地对象一样调用服务端JVM上的对象。RMI引入了 Stubs（客户端存根）和 Skeletons（服务端骨
## [重学 Java 反射机制](https://www.ek1ng.com/java-reflect-learning.html)  by [ek1ng](https://ek1ng.com/), 2023-07-25

近期跟一些java的最新漏洞，发现自己的语言基础太差了，跟着p牛的java安全漫谈重新学一下反射，p牛的文章确实是讲复杂的东西讲的浅显易懂。反射的定义对象可以通过反射获取对应的类，类可以通过反射获取所有方法，拿到的方法可以调用，这种机制就是反射。反射机制在安全方面的意义例如我们要完成RCE，但代码中绝大多数时候并没有Runtime，ProcessBuilder等常见的用于命令执行的类来让我们调用。
