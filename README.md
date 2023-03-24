### 日志存储及定时任务系统
go-log-cron-task-system  
Log storage and analysis system, 日志存储及定时任务系统。  
- 日常使用的后台、系统会产生大量的日志，直接存储mysql, mongo，数据量会很大，数据分析查询效率低。日志存储分析系统按模块存储，日志推送到elasticsearch或clickhouse存储，便于后面查询、分析。  
- 定时任务批量并发处理任务。  

#### 功能
- kafka-rabbitmq中间件  
- go协程消费  
- 推送elasticsearch或clickhouse存储  
- api日志分析、查询  

#### 版本
- kafka-rabbitmq-elasticsearch-clickhouse 
- go module
- go 1.19
- viper v1.15.0
- kafka
- rabbitmq
- elasticsearch  
- clickhouse  

#### 备注
