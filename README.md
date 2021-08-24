快速备份、创建、维护阿里云 SLS 日志服务索引。

应用场景：
- 将Logstore索引从项目A 拷贝至另一项目
- 本地维护SLS 索引（同一项目，仅首次备份）


满足功能：
- 备份索引
- 还原索引
- 索引追加
- 索引覆盖
- 索引创建


配置文件：
example-config.toml （需更名为config.toml)

```toml

dumpProject="project-prod"
dumpLogstore = ["test"]

receiveProject="project-dev"
receiveLogstore = ["test"]

```

### 如果已存在 logstore：
```shell
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

### 如果不存在 logstore：
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



## SLS 认证信息及Endpoint：
～/.config/sls-tools/config.toml

```
endpoint = "cn-beijing.log.aliyuncs.com"
accessKeyId = "Your-KeyId"
accessKeySecret = "Your-KeySecret"
```


