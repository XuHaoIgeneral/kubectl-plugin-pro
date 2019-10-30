## Kubectl-plugin-pro

kubectl-plugin-pro 是一个关于k8s命令行进行扩展的一个插件仓库。

---

#### 插件安装模板

- 编译需要的plugin文件，也可以直接在app目录下进行下载，命名格式为kubectl-xxx。
- 将目标文件放入`/usr/bin/`中，运行 `chmod +x ｛目标文件名｝`
- 运行 `kubectl plugin list` 查看可读取运行的插件是否安装成功

---

### 插件列表

##### namespace切换

- kubectl project <namespace> 
- 可以之间使用kubectl project 进入选择框对namespace进行手动选择