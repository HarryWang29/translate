# 项目的由来
最近倒是进了一些群，发现很多人对规则的转换都不是很了解，之前我都是自己写个脚本自己运行一下，就删除了，后来看到了[这个项目](https://github.com/ne1llee/v2ray2clash)，感觉很不错，不过这个老哥支持的不多，我个人喜欢使用命令行生成配置文件直接导入的方式，于是这个版本暂时只支持了命令行方式，web订阅后期再进行开发

这个项目也有一些借鉴了[@ne1llee](https://github.com/ne1llee)老哥的代码，其中makefile我就直接抄了过来，要是老哥觉得不合适，可与我联系
# 功能
- [ ] 命令行版本
- [ ] web版本
- [X] 将`v2ray`加入到`神机规则`中（clash）
- [ ] 将`ssr`加入到`神机规则`中（clash）
- [ ] 将`ss`加入到`神机规则`中（clash）
- [ ] 将`v2ray`加入到`神机规则`中（surge）
- [ ] 将`ssr`加入到`神机规则`中（surge）
- [ ] 将`ss`加入到`神机规则`中（surge）

.

.

.

# 使用方法
* 从release下载二进制文件
* 解压响应版本的zip
## mac/linux
* 执行如下命令修改可执行权限
    ```bash
    sudo chmod 766 translate-darwin-amd64
    ```
* 执行如下命令将v2ray转为clash(保留双引号)
    ```bash
    ./translate-darwin-amd64 v2ray clash --subLink="订阅链接"
    ```

## windows
### v2ray --> clash
* 用文本编辑器打开`v2ray2clash.bat`
* 修改订阅链接
* 保存退出
* 双击`v2ray2clash.bat`

# 感谢
* [神机规则](https://github.com/ConnersHua/Profiles/tree/master)
* [ne1lle](https://github.com/ne1llee/v2ray2clash)