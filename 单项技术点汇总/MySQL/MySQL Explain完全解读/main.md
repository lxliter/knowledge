### MySQL Explain完全解读
***EXPLAIN作为MySQL的性能分析神器，读懂其结果是很有必要的***，然而我在各种搜索引擎上竟然找不到特别完整的解读。<br>
都是只有重点，没有细节（例如***type***的取值不全、***Extra***缺乏完整的介绍等）。<br>

所以，我肝了将近一个星期，整理了一下。<br>
这应该是全网最全面、最细致的EXPLAIN解读文章了，下面是全文。<br>

文章比较长，建议收藏。<br>

#### TIPS
本文基于MySQL 8.0编写，理论支持MySQL 5.0及更高版本。<br>
EXPLAIN使用<br>
explain可用来分析SQL的执行计划。格式如下：<br>
```
{EXPLAIN | DESCRIBE | DESC}
    tbl_name [col_name | wild]
    
{EXPLAIN | DESCRIBE | DESC}
    [explain_type]
    {explainable_stmt | FOR CONNECTION connection_id}

{EXPLAIN | DESCRIBE | DESC} ANALYZE select_statement

explain_type: {
    FORMAT = format_name
}

format_name: {
    TRADITIONAL
  | JSON
  | TREE
}

explainable_stmt: {
    SELECT statement
  | TABLE statement
  | DELETE statement
  | INSERT statement
  | REPLACE statement
  | UPDATE statement
}    
```

示例：

```
EXPLAIN format = TRADITIONAL json SELECT tt.TicketNumber, tt.TimeIn,
               tt.ProjectReference, tt.EstimatedShipDate,
               tt.ActualShipDate, tt.ClientID,
               tt.ServiceCodes, tt.RepetitiveID,
               tt.CurrentProcess, tt.CurrentDPPerson,
               tt.RecordVolume, tt.DPPrinted, et.COUNTRY,
               et_1.COUNTRY, do.CUSTNAME
        FROM tt, et, et AS et_1, do
        WHERE tt.SubmitTime IS NULL
          AND tt.ActualPC = et.EMPLOYID
          AND tt.AssignedPC = et_1.EMPLOYID
          AND tt.ClientID = do.CUSTNMBR;
```

```
EXPLAIN format = TRADITIONAL json ***
```

结果输出展示：

| 字段            | format=json时的名称 | 含义             |
|---------------|-----------------|----------------|
| id            | select_id       | 该语句的唯一标识       |
| select_type   | 无               | 查询类型           |
| table         | table_name      | 表名             |
| partitions    | partitions      | 匹配的分区          |
| type          | access_type     | 联接类型           |
| possible_keys | possible_keys   | 可能得索引选择        |
| key           | key             | 实际选择的索引        |
| key_len       | key_length      | 索引长度           |
| ref           | ref             | 索引的哪一列被引用了     |
| rows          | rows            | 估计要扫描的行        |
| filtered      | filtered        | 表示符合查询条件的数据百分比 |
| Extra         | 没有              | 附加信息           |


##### 结果解读
###### id
该语句的唯一标识。如果***explain的结果包括多个id值***，***则数字越大越先执行***；***而对于相同id的行，则表示从上往下依次执行***。
- explain的结果包括多个id值
- 数字越大越先执行
- 对于相同id的行，则表示从上往下依次执行

###### select_type
查询类型，有如下几种取值：

| 查询类型                 | 作用                                                                                              |
|----------------------|-------------------------------------------------------------------------------------------------|
| SIMPLE               | 简单查询（未使用UNION或子查询）                                                                              |
| PRIMARY              | 最外层的查询                                                                                          |
| UNION                | 在UNION中的第二个和随后的SELECT被标记为UNION。如果UNION被FROM子句中的子查询包含，那么它的第一个SELECT会被标记为DERIVED                  |
| DEPENDENT UNION      | UNION中的第二个或后面的查询，依赖了外面的查询                                                                       |
| UNION RESULT         | UNION的结果                                                                                        |
| SUBQUERY             | 子查询中的第一个SELECT                                                                                  |
| DEPENDENT SUBQUERY   | 子查询中的第一个SELECT，依赖了外面的查询                                                                         |
| DERIVED              | 用来表示包含在FROM子句的子查询中的SELECT，MySQL会递归执行并将结果放到一个临时表中，MySQL内部将其称为Derived table（派生表），因为该临时表是从子查询派生出来的 |
| MATERIALIZED         | 物化子查询                                                                                           |
| UNCACHEABLE SUBQUERY | 子查询，结果无法缓存，必须针对外部查询的每一行重新评估                                                                     |
| UNCACHEABLE UNION    | UNION属于UNCACHEABLE SUBQUERY的第二个或后面的查询                                                           |




















