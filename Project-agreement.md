# 👀 相关项目约定 ✨
> `!一些通俗的约定在这里添加`

## 📔git相关

1. 先 `fork` 项目到自己的仓库

2. `git clone xxxx/fusu.git` 到自己的本地

3. 每次使用前先 `fetch&merge` ,实际上是用到 `git remote`

    ```bash
    #首次使用
    git remote add upstream https://github.com/DWDSec/fusu.git
    git remote -v
    # 这时候因该是有两个分支
    #以下为日常使用

    # 1 .抓取原仓库的更新
    git fetch upstream
    
    # 2 .合并远程的main分支
    git merge upstream/main

    # 3 .将本地更新后的更新到github的仓库(以上步骤是更新了本地的仓库与上游代码的更新，而我们自己的仓库的代码还未合并，简化每次更新要上github点击更新这个操作)

    git push 
    #或者是git push origin main,如果是差异性更新可以加 --force 参数

4. 这样每次准备进行开发的时候可以保持代码更新并不会合并冲突
 

## 🔈命名相关

### 全局变量
golang要求，首字母大写，之后保持驼峰书写

