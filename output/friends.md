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
## [『代码随想录』DAY8｜字符串：344.反转字符串 541.反转字符串 II](https://nickxu.me/post/programmercarl-day8.html)  by [NX](https://nickxu.me/), 2023-11-01

344. 反转字符串朴实无华1234567func reverseString(s []byte) {    n := len(s)    for i := 0; i < n/2; i++ {        s[n-i-1], s[i] = s[i], s[n-i-1]    }}541.反转字符串 II朴实无华的递归1234567891011121314151617181920func reve
## [『代码随想录』DAY7｜哈希表：454.四数相加 II 383. 赎金信](https://nickxu.me/post/programmercarl-day7.html)  by [NX](https://nickxu.me/), 2023-10-31

454.四数相加 II很顺理成章的思路123456789101112131415161718func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {    m := map[int]int{}    ans := 0    for _, a := range nums1 {        for _, b
## [浅谈 SAP 安全远程密码](https://homeboyc.cn/blog/%E6%B5%85%E8%B0%88-sap-%E5%AE%89%E5%85%A8%E8%BF%9C%E7%A8%8B%E5%AF%86%E7%A0%81/)  by [Atom](https://www.homeboyc.cn/), 2023-10-31

密码是一个让不少人头疼的问题，大多数人会立即承认密码很难记住和管理，特别是当密码要求变得越来越复杂时。幸运的是，有一些很棒的软件包和浏览器插
## [『代码随想录』DAY6｜哈希表：242.有效的字母异位词 349.两个数组的交集](https://nickxu.me/post/programmercarl-day6.html)  by [NX](https://nickxu.me/), 2023-10-30

242.有效的字母异位词这题没什么好说的12345678910111213141516171819func isAnagram(s string, t string) bool {    m := map[rune]int{}    for _, c := range s {        m[c]++    }    for _, c := range t {        m[c]--
## [『代码随想录』DAY4｜链表：19.删除链表的倒数第N个节点 142.环形链表II](https://nickxu.me/post/programmercarl-day4.html)  by [NX](https://nickxu.me/), 2023-10-28

本篇部分题目在 『算法拾遗』链表（Linked List） 中已经做过24.两两交换链表中的节点1234567891011121314151617181920func swapPairs(head *ListNode) *ListNode {    dummy:=&ListNode{        Next: head,    }        curr := dummy    for curr!
## [『代码随想录』DAY3｜链表：203.移除链表元素 707.设计链表 206.反转链表](https://nickxu.me/post/programmercarl-day3.html)  by [NX](https://nickxu.me/), 2023-10-27

本篇部分题目在 『算法拾遗』链表（Linked List） 中已经做过203.移除链表元素还有一个递归版，但是那个空间复杂度是 O(n)12345678910111213func removeElements(head *ListNode, val int) *ListNode {    dummy := &ListNode{Next: head}    curr := dummy    for
## [『代码随想录』DAY2｜数组：977.有序数组的平方 209.长度最小的子数组 59.螺旋矩阵II](https://nickxu.me/post/programmercarl-day2.html)  by [NX](https://nickxu.me/), 2023-10-26

本篇内容包含数组的另外两类经典题目：滑动窗口模拟同时复习一下双指针977.有序数组的平方相关链接代码随想录视频讲解想多了，居然想从中间入手往两边扫1234567891011121314151617181920212223242526272829303132func sortedSquares(nums []int) []int {    n := len(nums)    for i := 0;
## [『代码随想录』DAY1｜数组：704.二分查找 27.移除元素](https://nickxu.me/post/programmercarl-day1.html)  by [NX](https://nickxu.me/), 2023-10-25

本篇内容包含数组的两类经典题目：二分查找双指针704.二分查找相关链接代码随想录视频讲解边界条件注意Carl 哥的视频讲的很清楚，墙裂建议观看主流的就是两种写法，左闭右闭和左闭右开，如果你选择了一种写法，就应该全程保持这一种写法左闭右闭起始区间为 [0,n - 1] ，继续循环的条件是区间合法，也就是 left <= right更新左右边界值时，由于 mid 已经被排除，所以更新为 mid + 1
## [DNS 与 Pdns OpenWrt DNS 递归服务器搭建与 MosDNS 分流](https://xyxsw.ltd/2023/10/18/DNS%20%E4%B8%8E%20Pdns%20OpenWrt%20DNS%20%E9%80%92%E5%BD%92%E6%9C%8D%E5%8A%A1%E5%99%A8%E6%90%AD%E5%BB%BA%E4%B8%8E%20MosDNS%20%E5%88%86%E6%B5%81/)  by [xyxsw](https://xyxsw.ltd/), 2023-10-18

DNS什么是DNS域名系统（英语：Domain Name System，缩写：DNS）是互联网的一项服务。它作为将域名和IP地址相互映射的一个分布式数据库，能够使人更方便地访问互联网。简单来说就是用域名请求服务器换回来ip地址https://www.nslookup.io/这个网站可以看dns请求的结果、同理还有linux或者windows的nslookup命令DNS records for dn
## [『OSPP2023』我与 OSPP 的故事 —— 项目经验分享](https://nickxu.me/post/my-story-with-OSPP-project-experience-sharing.html)  by [NX](https://nickxu.me/), 2023-10-15

书接上文 『OSPP2023』我与 OSPP 的故事 —— 从听闻到中选 ，本文注重于描写项目开发的经历项目基本信息项目名称：为 Envoy Go 扩展建设插件市场项目导师：纪卓志项目描述：Envoy 是当前最流行的网络代理之一，Go 扩展是 MOSN 社区为 Envoy 增加的 Go 生态基础，也是 MOSN 社区 MoE 框架的基础。受益于Golang生态系统，研发可以轻松在 Envoy 实现
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
