# certbot-dns-manual-hook

配合certbot实现自动设置DNS解析插件

### 支持进度
- [x] 阿里云
- [ ] 腾通云
- [ ] 华为云

## 参数说明
|参数|说明|
|---|---|
|-action|可选[auth, clean]，默认为auth|
|-ak|AccessKey|
|-secret|AccessSecret|

可通过运行 ./manual-hook -h 查看帮助信息

## 使用示例

```
 certbot certonly --dry-run --manual --preferred-challenges dns \
  --manual-auth-hook "./manual-hook -ak '' -secret ''" \
  --manual-cleanup-hook "./manual-hook -ak '' -secret '' -action clean" \
  -d *.example.com 
```
