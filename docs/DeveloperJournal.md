# 开发日志

## Day 1

- 学习了 `curl` 指令的基本使用
- 学习了 Git Commit Messages 的规范，参考：
  - [How to Write Better Git Commit Messages – A Step-By-Step Guide (freecodecamp.org)](https://www.freecodecamp.org/news/how-to-write-better-git-commit-messages/)
  - [angular/CONTRIBUTING.md at main · angular/angular (github.com)](https://github.com/angular/angular/blob/main/CONTRIBUTING.md#-commit-message-format)
- 实现了对剪切板增删改查的操作

## Day 2

- 实现了基本的用户注册登录接口，还没有测试
- 小玩了一下，所以今天工作量很小

## Day 3

- 完成了用户注册登录
- 添加了 jwt 身份验证
- 发现 gin 的路由对大小写不敏感，改了 short-id 生成的逻辑
- 支持对剪切板保存作者信息，并且作者可以给剪切板设置访问控制
- 都用 postman 进行了测试
  - 对于合法的请求处理都没啥问题，对于一些特殊情况可能还存在问题，后续还需要再测测
  - 给前端返回结果的时候，成功的话返回啥、失败的话返回啥还需要再设计，现在写的比较随意
- 到目前为止主要是在完成业务层面的代码，一些细节的知识不太懂，后续还需要再补补，比如 jwt 的细节
