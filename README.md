# GO-Agenda-CLI

 [toc]
 
## cobra工具应用
cobra 是一个很好的开发Linux命令行的工具。
### 如何安装cobra
- 使用命令 `
go get -v github.com/spf13/cobra/cobra`下载过程中，会出提示如下错误

> Fetching https://golang.org/x/sys/unix?go-get=1
 https fetch failed: Get https://golang.org/x/sys/unix?go-get=1: dial tcp 216.239.37.1:443: i/o timeout

这是熟悉的错误，请在 `$GOPATH/src/golang.org/x` 目录下用 `git clone` 下载 `sys` 和 `text` 项目，然后使用 `go install github.com/spf13/cobra/cobra`, 安装后在 `$GOBIN` 下出现了 `cobra` 可执行程序。

### Cobra 的简单使用

创建一个处理命令 `agenda register -uTestUser `或 `agenda register --user=TestUser` 的小程序。

简要步骤如下：

`cobra init`
`cobra add register`
需要的文件就产生了。 你需要阅读` main.go` 的 `main()` ; `root.go` 的 `Execute()`; 最后修改 `register.go`, `init()` 添加：

`registerCmd.Flags().StringP("user", "u", "Anonymous", "Help message for username")`
Run 匿名回调函数中添加：

``` go
username, _ := cmd.Flags().GetString("user")
fmt.Println("register called by " + username)
```
测试命令：

> $ go run main.go register --user=TestUser
register called by TestUser

参考文档：
 -  [官方文档](https://github.com/spf13/cobra#overview)
 -  [cobra简介](http://time-track.cn/cobra-brief-introduction.html)
## Agenda 命令设计

### 命令实现
>agenda help [option]

列出命令说明(可选是否列出具体功能的说明)。

>agenda register -uUserName --password pass -email=a@xxx.com -phone=phoneNum

用户注册，如果用户名已被使用，返回错误信息；如果注册成功，返回成功信息。

>agenda login -uUserName --password pass

用户登录，登录失败返回失败原因;登录成功，返回成功信息，并列出可选操作。

>agenda logout

用户退出登录，返回成功信息并列出可选操作。

>agenda -lUser

已登录用户查询已注册用户信息（用户名、邮箱、电话）

>agenda delete

已登录用户注销帐号，操作成功返回成功信息；否则，返回失败信息。若成功，删除一切与该用户的信息。

>agenda mkmeeting --title Name --participator user1 [user2 ....] --stime start --etime end

成功，则返回成功信息及注册信息；失败，则返回失败原因。

>agenda meetingadd --participator user1 [user2 ...]

成功，则返回新增参与者信息；失败，返回失败原因。

> agenda meetingdel --participator user1 [user2 ...]

成功，则返回成功信息（如果会议因为删除参与者而删除，返回额外信息）；失败，返回失败原因。

>agenda querymeeting -stime start -etime end

已登录用户查询自己的会议议程。

>agenda meetingdel --title Name

已登录用户删除会议。

>agenda meetingout --ttile Name

已登录用户退出自己参与的某一会议，若因此造成会议参与者为0,则会议被删除。

>agenda meetingclear

已登录用户清空自己发起的所有会议。

### 需求描述
- 业务需求：见后面附件
- 功能需求： 设计一组命令完成 agenda 的管理，例如：
	- agenda help ：列出命令说明
	- agenda register -uUserName --password pass -email=a@xxx.com ：注册用户
	- agenda help register ：列出 register 命令的描述
	- agenda cm ... : 创建一个会议
	- 原则上一个命令对应一个业务功能
- 持久化要求：
	- 使用 json 存储 User 和 Meeting 实体
	- 当前用户信息存储在 curUser.txt 中
- 开发需求
	- 团队：2-3人，一人作为 master 创建程序框架，其他人 fork 该项目，所有人同时开发。团队 不能少于 2 人
	- 时间：两周完成
- 项目目录
	- cmd ：存放命令实现代码
	- entity ：存放 User 和 Meeting 对象读写与处理逻辑
	- 其他目录 ： 自由添加
- 日志服务
	- 使用 log 包记录命令执行情况

## Agenda业务需求
**用户注册**

1. 注册新用户时，用户需设置一个唯一的用户名和一个密码。另外，还需登记邮箱及电话信息。
2. 如果注册时提供的用户名已由其他用户使用，应反馈一个适当的出错信息；成功注册后，亦应反馈一个成功注册的信息。

**用户登录**

1. 用户使用用户名和密码登录 Agenda 系统。
2. 用户名和密码同时正确则登录成功并反馈一个成功登录的信息。否则，登录失败并反馈一个失败登录的信息。

**用户登出**

1. 已登录的用户登出系统后，只能使用用户注册和用户登录功能。

**用户查询**

1. 已登录的用户可以查看已注册的所有用户的用户名、邮箱及电话信息。

**用户删除**

1. 已登录的用户可以删除本用户账户（即销号）。
2. 操作成功，需反馈一个成功注销的信息；否则，反馈一个失败注销的信息。
3. 删除成功则退出系统登录状态。删除后，该用户账户不再存在。
4. 用户账户删除以后：
	- 以该用户为 发起者 的会议将被删除
	- 以该用户为 参与者 的会议将从 参与者 列表中移除该用户。若因此造成会议 参与者 人数为0，则会议也将被删除。

**创建会议**

1. 已登录的用户可以添加一个新会议到其议程安排中。会议可以在多个已注册用户间举行，不允许包含未注册用户。添加会议时提供的信息应包括：
	- 会议主题(title)（在会议列表中具有唯一性）
	- 会议参与者(participator)
	- 会议起始时间(start time)
	- 会议结束时间(end time)
2. 注意，任何用户都无法分身参加多个会议。如果用户已有的会议安排（作为发起者或参与者）与将要创建的会议在时间上重叠 （允许仅有端点重叠的情况），则无法创建该会议。
3. 用户应获得适当的反馈信息，以便得知是成功地创建了新会议，还是在创建过程中出现了某些错误。

**增删会议参与者**

1. 已登录的用户可以向 自己发起的某一会议增加/删除 参与者 。
2. 增加参与者时需要做 时间重叠 判断（允许仅有端点重叠的情况）。
3. 删除会议参与者后，若因此造成会议 参与者 人数为0，则会议也将被删除。

**查询会议**

1. 已登录的用户可以查询自己的议程在某一时间段(time interval)内的所有会议安排。
2. 用户给出所关注时间段的起始时间和终止时间，返回该用户议程中在指定时间范围内找到的所有会议安排的列表。
3. 在列表中给出每一会议的起始时间、终止时间、主题、以及发起者和参与者。
4. 注意，查询会议的结果应包括用户作为 发起者或参与者 的会议。

**取消会议**

1. 已登录的用户可以取消 自己发起 的某一会议安排。
2. 取消会议时，需提供唯一标识：会议主题（title）。

**退出会议**

1. 已登录的用户可以退出 自己参与 的某一会议安排。
2. 退出会议时，需提供一个唯一标识：会议主题（title）。若因此造成会议 参与者 人数为0，则会议也将被删除。

**清空会议**

1. 已登录的用户可以清空 自己发起 的所有会议安排

## Code
> 代码放在了我的github上。[code](https://github.com/moandy/MygoLearn)
