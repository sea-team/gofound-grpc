# GoFound-grpc

[GoFound](https://github.com/sea-team/gofound)的 GRPC 版本,支持 http 和 grpc 调用

##### 为何要用 golang 实现一个全文检索引擎？

-  正如其名，`GoFound`去探索全文检索的世界，一个小巧精悍的全文检索引擎，支持持久化和单机亿级数据毫秒级查找。
-  传统的项目大多数会采用`ElasticSearch`来做全文检索，因为`ElasticSearch`够成熟，社区活跃、资料完善。缺点就是配置繁琐、基于 JVM 对内存消耗比较大。
-  所以我们需要一个更高效的搜索引擎，而又不会消耗太多的内存。 以最低的内存达到全文检索的目的，相比`ElasticSearch`，`gofound`是原生编译，会减少系统资源的消耗。而且对外无任何依赖。

## 普通启动 orDocker 启动

-  启动

   ```
   // 启动grpc
   go run main.go
   // 启动gateway
   cd gateway && go run main.go
   ```

-  docker 部署

   ```
   make build
   ```
