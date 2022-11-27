## 基础
- [doc_link](https://c5tnppcmk4.feishu.cn/docx/UIw1dJSYBo4M9Zx1lKIc6xt2n9e)

- `net/http` 库的核心是 Server对象 
  - Server对象 接受一个addr参数以及一个Handler对象
  - Addr主要是指定Server的端口号
  - Handler 对象主要是一个处理请求的对象 可以直接是函数 也可以是ServerMux这类路由 Handler 本身只是个接口 总之最终要实现Handler的ServerHttp方法中完成工作
  - Server对象首先 建立会Listen端口 然后调用Accept等待请求进来。一旦请求进来 就使用Accept返回的conn对象 单独启动一个go程 然后通过参数判断实际使用的HandlerFucn
  - 最终通过HandlerFunc 处理返回结果
- context 构建 
  - 其中base context 以及 connContext 都可以在创建Server的时候自定义 传入一些自己希望传入的东西
  - 这类connContext 会在连接出现异常的时候 调用CancelFunc
  - ![img](https://static001.geekbang.org/resource/image/a1/fc/a1d0ece41997ac873a5292301ee988fc.jpg?wh=1920x1080)
  - 下面的一段代码 主要解释了 如何使用 context 在上层调用者处调用 同时等待两个信息 一个是context自己返回done 一个是调用者自己判定超时，但是不管哪一种 最终都要调用cancel 通知下游 上面已经不干了 你们自己cancel吧 
  - cancel是上游通知下游context done是下游被上游通知结束