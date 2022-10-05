# spanner-benchmark
spanner benchmark for golang

## Spanner Benchmark Quickstart

### 测试环境：

-   Cloud Spanner 
-   Process Unit: 100 
-   Storage: 410GB

1.创建两个表, 分别带索引和不带索引(1M数据量)
<img src="https://lh6.googleusercontent.com/0POpXS92xoOcFykZQSvxnIrtTlOp1fASHakEvb2mqlfQtIWos2eslu7VCYERDnX3OqvyD1V2Ln_HgZ2dsnh372lLalDXNpStxySWfVAYXsed_o9bXmZAHIOhnoBHOiVfVCGIyNkF65AQ_yKbNijZddEEbp2deM0IewkDLX9yv6SONm0vJy4e32WK1gGoUOX_L-CCkcMAjPyJYiTJY9jmuoMD33PE7rcK_Q" alt="img" style="zoom:50%;" />

```sql
CREATE TABLE Users (
  UsersId INT64 NOT NULL,
  ServerId INT64,
  Power INT64,
  Charm INT64,
  Money INT64,
  intimacy INT64,
) PRIMARY KEY(UsersId);

CREATE TABLE Users_index (
  UsersId INT64 NOT NULL,
  ServerId INT64,
  Power INT64,
  Charm INT64,
  Money INT64,
  intimacy INT64,
) PRIMARY KEY(UsersId);

CREATE INDEX index_charm ON Users_index(Charm);

CREATE INDEX index_intimacy ON Users_index(intimacy);

CREATE INDEX index_money ON Users_index(Money);

CREATE INDEX index_power ON Users_index(Power);

CREATE INDEX index_serverid ON Users_index(ServerId);

CREATE INDEX serverId_power_index ON Users_index(ServerId, Power DESC);
```

### 2.性能测试场景和结果记录

历史数据100w条 in Spanner，每种测试测试100次, 取P50的测试数据[从小到大排序，取中位数] 数据插入性能比较(每次插入25条), 打印事务执行时间P50；所有操作执行100次, 按照P50统计结果

#### 2.1 DML批量写入方式测试：

##### 2.1.1 每次通过SQL 的方式批量写入25条记录，持续写入100次，产生2.5K 条记录，测试结果如下：

```bash
$ go run insertDML_bech.go 
----- writeBatchUsingDML Test: -----
Benchmark result:
  avg 97.1ms;  min 270ns;  p50 91.4ms;  max 211ms;
  p90 134ms;  p99 211ms;  p999 211ms;  p9999 211ms;
      270ns [  2] █
       50ms [ 63] ████████████████████████████████████████
      100ms [ 29] ██████████████████
      150ms [  5] ███
      200ms [  1] ▌
      250ms [  0] 
      300ms [  0] 
      350ms [  0] 
      400ms [  0] 
      450ms [  0] 

```

#### 2.2 通过mumation 方式同样每次插入25条记录，持续写入100次，产生2.5K 条记录，测试结果如下：

```bash
$ go run mumationInsert_bech.go 
Benchmark result:
  avg 18.6ms;  min 200ns;  p50 18.4ms;  max 45.1ms;
  p90 27.9ms;  p99 45.1ms;  p999 45.1ms;  p9999 45.1ms;
      200ns [  1] █
        5ms [  1] █
       10ms [ 28] █████████████████████████████████████
       15ms [ 30] ████████████████████████████████████████
       20ms [ 29] ██████████████████████████████████████▌
       25ms [  6] ████████
       30ms [  4] █████
       35ms [  0] 
       40ms [  0] 
       45ms [  1] █
```

#### 2.3 DML 方式批量删除

每次删除25条记录，连续测试100次，

```bash
$ go run deleteDML_bench.go 
Benchmark result:
  avg 31.9ms;  min 235ns;  p50 29ms;  max 67.2ms;
  p90 44.6ms;  p99 67.2ms;  p999 67.2ms;  p9999 67.2ms;
      235ns [  1] ▌
       10ms [  0] 
       20ms [ 57] ████████████████████████████████████████
       30ms [ 26] ██████████████████
       40ms [ 14] █████████▌
       50ms [  1] ▌
       60ms [  1] ▌
       70ms [  0] 
       80ms [  0] 
       90ms [  0] 
```

#### 2.4 使用mumation 方式批量更新：

由于mumation 方式对table 内不存的数据是进行insert 操作，对已存在的数据是进行更新操作，因此同样可以使用mumation方式来对update 性能测试：

```bash
$ go run mumationInsert_bech.go 
Benchmark result:
  avg 14.4ms;  min 146ns;  p50 12.8ms;  max 39.9ms;
  p90 23.2ms;  p99 39.9ms;  p999 39.9ms;  p9999 39.9ms;
      146ns [  1] 
        5ms [  0] 
       10ms [ 80] ████████████████████████████████████████
       15ms [  8] ████
       20ms [  4] ██
       25ms [  4] ██
       30ms [  2] █
       35ms [  1] 
       40ms [  0] 
       45ms [  0] 
```

#### 2.5 对users_index table 进行点查：

```bash
## sql := "SELECT UsersId,ServerId,Power,Charm,Money,intimacy FROM Users_index WHERE  ServerId=1"
$ go run queryDML_bench.go 
----- QueryDML Benchmark Test: -----
  avg 570ms;  min 434ms;  p50 490ms;  max 1.44s;
  p90 865ms;  p99 1.44s;  p999 1.44s;  p9999 1.44s;
      435ms [ 83] ████████████████████████████████████████
      600ms [  5] ██
      800ms [  6] ██▌
         1s [  2] ▌
       1.2s [  2] ▌
       1.4s [  2] ▌
       1.6s [  0] 
       1.8s [  0] 
         2s [  0] 
       2.2s [  0] 
```

#### 2.6 对users_index table 进行按照power 排序，查询榜单数据：

```bash
	## sql := "SELECT UsersId,Power FROM Users_index WHERE ServerId=1 ORDER BY Power DESC LIMIT 10"
$ go run queryDML_bench.go 
----- QueryDML Benchmark Test: -----
  avg 8.94ms;  min 3.77ms;  p50 5.97ms;  max 315ms;
  p90 7.16ms;  p99 315ms;  p999 315ms;  p9999 315ms;
     3.77ms [ 99] ████████████████████████████████████████
       50ms [  0] 
      100ms [  0] 
      150ms [  0] 
      200ms [  0] 
      250ms [  0] 
      300ms [  1] 
      350ms [  0] 
      400ms [  0] 
      450ms [  0] 
```

