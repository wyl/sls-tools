## 阿里云 SLS 日志工具
对阿里云日志索引进行备份、维护（追加、创建）的工具

### 目录相关
- sls-tools
- config.toml
### 执行命令
维护的时候记得更改配置文件中的 Logstore 的名称，支持全部索引格式。


./sls-tools

### 如果Logstore 在目标 Project 中已存在 
`备份的目录为 ./logstore`


```
wangyalong@wangyalongde-MacBook-Pro-for-Job sls-tools$ ./sls-tools
INFO 2021/08/24 17:01:40 project-prod -> test 备份完毕。
INFO 2021/08/24 16:59:51 正在执行 project-dev -> test。
Use the arrow keys to navigate: ↓ ↑ → ← 
? project-dev -> test 已存在索引，YES 追加索引，NO 覆盖索引。 [Yes/No]: 
 ▸ Yes
   No

✔ YES
INFO 2021/08/24 17:00:14 project-dev -> test 追加索引！
✔ NO
INFO 2021/08/24 17:00:53 project-dev -> test 覆盖索引！

INFO 2021/08/24 17:00:14 project-dev -> test 执行完毕。
```

### 如果Logstore 在目标 Project 中不存在
```
wangyalong@wangyalongde-MacBook-Pro-for-Job sls-tools$ ./sls-tools 
INFO 2021/08/24 17:04:46 project-prod -> test 备份完毕。
INFO 2021/08/24 17:04:46 正在执行 project-dev -> test。
Use the arrow keys to navigate: ↓ ↑ → ← 
? 【project-dev】 不存在 logstore 【test】，是否创建 logstore ？ [Yes/No]: 
 ▸ Yes
   No

✔ Yes
INFO 2021/08/24 17:04:57 project-dev -> test 创建索引！
INFO 2021/08/24 17:04:57 project-dev -> test 执行完毕。
```



### 配置文件描述
```toml 

dumpProject="project-prod" // 从该Project 中执行备份操作
dumpLogstore = ["test"]    // 备份Logstore 【"test"】

receiveProject="project-dev" // 从台Project 执行维护操作
receiveLogstore = ["test"]   // 维护Logstore 【"test"】

```



### 认证信息
```
~/.config/sls-tools/config.toml
   
endpoint = "cn-beijing.log.aliyuncs.com"
accessKeyId = "Your-KeyId"
accessKeySecret = "Your-KeySecret"
```

### 备份的索引文件 example
```
ProjectName = "project-dev"
LogStore = "test"

[Index]
  [Index.Keys]
    [Index.Keys.category]
      Token = [",", " ", "'", "\"", ";", "=", "(", ")", "[", "]", "{", "}", "?", "@", "&", "<", ">", "/", ":", "\n", "\t", "\r"]
      CaseSensitive = false
      Type = "text"
      DocValue = true
      Alias = ""
      Chn = false
    [Index.Keys.mac]
      Token = [",", " ", "'", "\"", ";", "=", "(", ")", "[", "]", "{", "}", "?", "@", "&", "<", ">", "/", ":", "\n", "\t", "\r"]
      CaseSensitive = false
      Type = "text"
      DocValue = true
      Alias = ""
      Chn = false
  [Index.Line]
    Token = [",", " ", "'", "\"", ";", "=", "(", ")", "[", "]", "{", "}", "?", "@", "&", "<", ">", "/", ":", "\n", "\t", "\r"]
    CaseSensitive = false
    Chn = false
```
### Jekyll Themes

Your Pages site will use the layout and styles from the Jekyll theme you have selected in your [repository settings](https://github.com/wyl/sls-tools/settings/pages). The name of this theme is saved in the Jekyll `_config.yml` configuration file.

### Support or Contact

Having trouble with Pages? Check out our [documentation](https://docs.github.com/categories/github-pages-basics/) or [contact support](https://support.github.com/contact) and we’ll help you sort it out.
