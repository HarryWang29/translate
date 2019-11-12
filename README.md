# 项目的由来
最近倒是进了一些群，发现很多人对规则的转换都不是很了解，之前我都是自己写个脚本自己运行一下，就删除了，后来看到了[这个项目](https://github.com/ne1llee/v2ray2clash)，感觉很不错，不过这个老哥支持的不多，我个人喜欢使用命令行生成配置文件直接导入的方式，于是这个版本暂时只支持了命令行方式，web订阅后期再进行开发

这个项目也有一些借鉴了[@ne1llee](https://github.com/ne1llee)老哥的代码，其中makefile我就直接抄了过来，要是老哥觉得不合适，可与我联系
# 功能
- [X] 命令行版本
- [ ] web版本 (关于web订阅一直没有开始：web订阅很多app都不允许修改，很不爽啊，我还是喜欢用文本导入，还可以自定义规则)
- [X] 支持多订阅链接
- [X] 支持喵帕斯订阅提取v3节点
- [X] 将`v2ray`加入到`神机规则`中（clash）
- [ ] 将`ssr`加入到`神机规则`中（clash）
- [ ] 将`ss`加入到`神机规则`中（clash）
- [ ] 将`v2ray`加入到`神机规则`中（surge）
- [ ] ~~将`ssr`加入到`神机规则`中（surge）~~(写readme的时候没动脑子，哈哈)
- [ ] 将`ss`加入到`神机规则`中（surge）

.

.

.

# 使用方法
* 从[release](https://github.com/HarryWang29/translate/releases)下载二进制文件
* 解压相应版本的压缩包
* 多订阅链接直接提供多`--subLink=""`标签即可，样例如下
    ```bash
    ./translate-darwin-amd64 vmess clash --subLink="订阅链接1" --subLink="订阅链接2"
    or
    ./translate-darwin-amd64 vmess clash --subLink="订阅链接1,订阅链接2"
    ```
 * 喵帕斯节点过滤功能，现在只支持vmess协议的订阅，增加`--npsboost`标签，样例如下：
    * 获取喵帕斯v2ray订阅链接为`link1`
    * 获取喵帕斯ssr订阅链接为`link2`
    * 将`link2`中`mu=1`更改为`mu=0`为`link3`
    * 组成命令为
        ```bash
        ./translate-darwin-amd64 vmess clash --subLink="link1" --npsboost="link3"
        ```
    * 注意，喵帕斯过滤功能无法和多订阅一起使用，此功能会按照ss订阅来过滤节点，会将第二个机场的所有节点全部过滤
## mac/linux
* 执行如下命令修改可执行权限
    ```bash
    sudo chmod 766 translate-darwin-amd64
    ```
* 执行如下命令将v2ray转为clash(保留双引号)
    ```bash
    ./translate-darwin-amd64 vmess clash --subLink="订阅链接"
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