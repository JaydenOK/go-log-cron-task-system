### 日志存储分析系统
go-kafka-clickhouse 日常使用的后台、系统会产生大量的日志，直接存储mysql, mongo，数据量会很大，数据分析查询效率低。  
日志存储分析系统按模块存储，日志推送到elasticsearch或clickhouse存储，便于后面查询、分析。

#### 功能
- kafka中间件  
- go协程消费  
- 推送elasticsearch或clickhouse存储  
- api日志分析、查询  

#### 版本 
- go module
- go 1.19
- viper v1.15.0
- kafka
- elasticsearch  
- clickhouse  

#### 备注
