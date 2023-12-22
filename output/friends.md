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
## [优雅地实现滚动容器遮罩](https://daidr.me/archives/code-1112.html)  by [戴兜](https://im.daidr.me), 2023-12-09

在设计前端页面时，常常会遇到这种情况：可滚动容器的边界并非父容器的边界，导致子元素溢出造成裁切，让页面产生比较怪异的视觉效果（左图）添加遮罩之后，效果自然了许多（右图）纯色遮罩以上图的这种情况举例，我们需要做的，是在可滚动容器的顶部和底部分别放置一个线性渐变的纯色遮罩，遮挡生硬的裁切线。创建两个元素 .top-mask 、.bottom-mask 来作为
## [Nginx Ingress & Grpc 端口复用配置](https://homeboyc.cn/blog/nginx-ingress-grpc-%E7%AB%AF%E5%8F%A3%E5%A4%8D%E7%94%A8%E9%85%8D%E7%BD%AE/)  by [Atom](https://www.homeboyc.cn/), 2023-12-07

# 前言 书接上回Caddy & GRPC 端口复用配置，新在集群上部署了一个带端口复用的服务，除了 Caddy 换成了 Nginx 之外，其他还是一样的配方。 看到有些教程给出的解
## [飞镖、骰子和硬币：从离散分布中抽样](https://nickxu.me/post/darts-dice-coins-chinese-version.html)  by [NX](https://nickxu.me/), 2023-12-04

本篇是原文 Darts, Dice, and Coins: Sampling from a Discrete Distribution 的全文中文翻译版本在被我的导师拉着参与 cloudwego/kitex 的一个负载均衡算法的改进时 #1184 ，他指出可能可以使用 Alias Method（别名方法）来实现更为高效的负载均衡。于是我仔细地研究了本文的英文原文（链接如上）， 起初我想寻求一个中文
## [『代码随想录』动态规划（Dynamic Programming）](https://nickxu.me/post/programmercarl-dynamic-programming.html)  by [NX](https://nickxu.me/), 2023-12-01

DAY 38斐波那契数123456789101112func fib(n int) int {    if n == 0 || n == 1 {        return n    }    dp := make([]int, n+1)    dp[0], dp[1] = 0, 1    for i := 2; i <= n; i++ {        dp[i] = dp[i-1] + dp[
## [『代码随想录』贪心（Greedy）](https://nickxu.me/post/programmercarl-greedy.html)  by [NX](https://nickxu.me/), 2023-11-28

DAY 31分发饼干遍历每块饼干，尝试将其分配给胃口最小的那个尚未满足的孩子1234567891011121314func findContentChildren(g []int, s []int) int {    sort.Ints(g)    sort.Ints(s)    n, m := len(g), len(s)    i, j := 0, 0    for i < n && j <
## [『代码随想录』回溯（Backtracking）](https://nickxu.me/post/programmercarl-backtracking.html)  by [NX](https://nickxu.me/), 2023-11-17

DAY 2477.组合很经典的回溯算法123456789101112131415161718func combine(n int, k int) [][]int {    ans := [][]int{}    curr := []int{}    var dfs func(s int)    dfs = func(s int) {        if len(curr) == k {
## [『代码随想录』二叉树（Binary Tree）](https://nickxu.me/post/programmercarl-binary-tree.html)  by [NX](https://nickxu.me/), 2023-11-07

DAY 14稍微又复习了一遍二叉树的基础144.二叉树的前序遍历123456789func preorderTraversal(root *TreeNode) []int {if root == nil {return []int{}}ans := []int{root.Val}ans = append(ans, preorderTraversal(root.Left)...)ans = appe
## [Termium 如何保护您的信息安全？](https://homeboyc.cn/blog/termium-%E5%A6%82%E4%BD%95%E4%BF%9D%E6%8A%A4%E6%82%A8%E7%9A%84%E4%BF%A1%E6%81%AF%E5%AE%89%E5%85%A8/)  by [Atom](https://www.homeboyc.cn/), 2023-11-06

在 Termium，我们非常注重保护您的信息安全、以负责任的方式处理您的信息，并确保您拥有控制权。说得容易做的难，我们之所以有信心保护您的信息
## [浅谈 OPAQUE 安全远程密码](https://homeboyc.cn/blog/%E6%B5%85%E8%B0%88-opaque-%E5%AE%89%E5%85%A8%E8%BF%9C%E7%A8%8B%E5%AF%86%E7%A0%81/)  by [Atom](https://www.homeboyc.cn/), 2023-11-06

继前一篇谈 SAP 协议之后，我们来浅了解一下 OPAQUE。（一个已经被WhatsAPP用于生产的新兴协议） 和 SAP 的建立于“反复打磨”上的安全性不同，
## [Mac Windows 网线直连文件传输](https://nickxu.me/post/mac-windows-file-copy-by-directly-LAN.html)  by [NX](https://nickxu.me/), 2023-11-05

今天又学到一个小技巧，使用一条网线直连两台 PC 进行文件传输我是把 Mac 上的文件传到 Windows 上，但其实两台 Windows 也大同小异首先当然是网线插上下一步是自己手动改一下 IP 到同一网段，之后 ping 一下，能 ping 通就算没问题传输文件使用 HTTP 其实最方便，直接用 py 起一个 http server12cd /path/to/dirpython3 -m htt
## [『代码随想录』栈与队列（Stack & Queue）](https://nickxu.me/post/programmercarl-stack-and-queue.html)  by [NX](https://nickxu.me/), 2023-11-03

DAY 10232.用栈实现队列因为栈和队列的出队顺序是反的，所以再来个栈倒腾一下就是正的了123456789101112131415161718192021222324252627282930313233343536373839404142type MyQueue struct {    in  []int    out []int}func Constructor() MyQueue {
## [『代码随想录』字符串（String）](https://nickxu.me/post/programmercarl-string.html)  by [NX](https://nickxu.me/), 2023-11-01

DAY 8344. 反转字符串朴实无华1234567func reverseString(s []byte) {    n := len(s)    for i := 0; i < n/2; i++ {        s[n-i-1], s[i] = s[i], s[n-i-1]    }}541.反转字符串 II朴实无华的递归1234567891011121314151617181920func
## [浅谈 SAP 安全远程密码](https://homeboyc.cn/blog/%E6%B5%85%E8%B0%88-sap-%E5%AE%89%E5%85%A8%E8%BF%9C%E7%A8%8B%E5%AF%86%E7%A0%81/)  by [Atom](https://www.homeboyc.cn/), 2023-10-31

密码是一个让不少人头疼的问题，大多数人会立即承认密码很难记住和管理，特别是当密码要求变得越来越复杂时。幸运的是，有一些很棒的软件包和浏览器插
## [『代码随想录』哈希表（Hash Map）](https://nickxu.me/post/programmercarl-hash-map.html)  by [NX](https://nickxu.me/), 2023-10-30

DAY 6242.有效的字母异位词这题没什么好说的12345678910111213141516171819func isAnagram(s string, t string) bool {    m := map[rune]int{}    for _, c := range s {        m[c]++    }    for _, c := range t {        m[c]-
## [『代码随想录』链表（Linked List）](https://nickxu.me/post/programmercarl-linked-list.html)  by [NX](https://nickxu.me/), 2023-10-27

DAY 3本篇部分题目在 『算法拾遗』链表（Linked List） 中已经做过203.移除链表元素还有一个递归版，但是那个空间复杂度是 O(n)12345678910111213func removeElements(head *ListNode, val int) *ListNode {    dummy := &ListNode{Next: head}    curr := dummy
## [DNS 与 Pdns OpenWrt DNS 递归服务器搭建与 MosDNS 分流](https://xyxsw.ltd/2023/10/18/DNS%20%E4%B8%8E%20Pdns%20OpenWrt%20DNS%20%E9%80%92%E5%BD%92%E6%9C%8D%E5%8A%A1%E5%99%A8%E6%90%AD%E5%BB%BA%E4%B8%8E%20MosDNS%20%E5%88%86%E6%B5%81/)  by [xyxsw](https://xyxsw.ltd/), 2023-10-18

DNS什么是DNS域名系统（英语：Domain Name System，缩写：DNS）是互联网的一项服务。它作为将域名和IP地址相互映射的一个分布式数据库，能够使人更方便地访问互联网。简单来说就是用域名请求服务器换回来ip地址https://www.nslookup.io/这个网站可以看dns请求的结果、同理还有linux或者windows的nslookup命令DNS records for dn
## [在实验网络中使用 step-ca 签发 SSL 证书](https://baimeow.cn/posts/dn11/acme/)  by [柏喵Sakura](https://baimeow.cn/), 2023-10-08

在真实的互联网中，我们常常会通过校验 SSL 证书的方式创建加密 socket 来访问互联网上的资源，这能够有效阻止他人看到通信的内容。而在 DN11 中，由于每个人都相当于是 ISP，流量可能途径每一个人，随时可能被人抓取，因此使用 SSL 证书这件事就更为重要了。本文将在使用 PDNS 搭建好实验性网络 DNS 的基础上使用 step-ca 搭建 ACME Server 并使用 acme.s
## [[简中] Google Summer of Code & Chrome Extensions](https://daidr.me/archives/code-1097.html)  by [戴兜](https://im.daidr.me), 2023-09-25

这篇博客是 https://im.daidr.me/blog/1086 的简中翻译版本。原文地址：https://developer.chrome.com/blog/google-summer-of-code-and-chrome-extensions/我是一名来自中国的大二学生，对Web开发非常感兴趣。在我大一的时候，我加入了一个技术社团（杭电助手），这可以说是让我接触编程和
## [人生好比修仙](https://blog.aflybird.cn/2023/09/life-is-cultivating-immortality/)  by [Bird](https://blog.aflybird.cn/), 2023-09-24

最近看《凡人修仙传》魔怔了，自认为「悟」出了点人生与修仙的关联。 声明：本人是纯正的唯物主义者，信仰科学。以下对于修仙的描述皆为文学手法。 凡人
## [IBGP FullMesh 实现多节点自治域](https://baimeow.cn/posts/dn11/configureibgp/)  by [柏喵Sakura](https://baimeow.cn/), 2023-09-23

最近不少群友有把自己的其他服务器也接入自己的 AS 的需求，那这篇教程也是时候该出了。总览 这个事情比较复杂，我们先不急着开始配，先了解一下我们需要一个什么效果。DN11 作为一个发展了有一段时间的实验性网络，我们目前有很大一批人，他们手上都已经有了一个节点，并且在这个节点上配置了 BGP 从而联通到其他节点。那这个情景下，他们的自治域就是由那一个节点组成的单节点自治域。我们现在需要添加一个新
