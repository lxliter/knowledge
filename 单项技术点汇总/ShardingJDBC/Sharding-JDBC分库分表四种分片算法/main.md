## Sharding-JDBC分库分表四种分片算法

- 1. 精确分片算法
精确分片算法（PreciseShardingAlgorithm）精确分片算法（***=与IN语句***），
用于处理使用单一键作为分片键的=与IN进行分片的场景。
需要配合StandardShardingStrategy使用

```
精确分片算法（PreciseShardingAlgorithm）精确分片算法（=与IN语句）
用于处理使用单一键作为分片键的=与IN进行分片的场景
```

- 2. 范围分片算法
范围分片算法（RangeShardingAlgorithm）用于处理使用单一键作为分片键的***BETWEEN AND***进行分片的场景。
需要配合StandardShardingStrategy使用。

```
范围分片算法（RangeShardingAlgorithm）用于处理使用单一键作为分片键的BETWEEN AND进行分片的场景。
需要配合StardardShardingStrategy使用。
```

- 3. 复合分片算法
复合分片算法（ComplexKeysShardingAlgorithm）用于***多个字段作为分片键的分片操作***，
***同时获取到多个分片健的值，根据多个字段处理业务逻辑***。
需要在复合分片策略（ComplexShardingStrategy）下使用。
```
复合分片算法（ComplexKeysShardingAlgorithm）用于多个字段作为分片键的分片操作，
同时获取到多个分片键的值，根据多个字段处理业务逻辑。
需要在复合分片策略（ComplexShardingStrategy）下使用。
```

- 4. Hint分片算法
Hint分片算法（HintShardingAlgorithm）稍有不同，上边的算法中我们都是解析语句提取分片键，
并设置分片策略进行分片。但有些时候我们并没有使用任何的分片键和分片策略，可还想将 SQL 路由到目标数据库和表，
就需要通过手动干预指定SQL的目标数据库和表信息，也叫强制路由。
```
Hint分片算法（HintShardingAlgorithm）稍有不同，
上边的算法中我们都是解析语句提取分片键，
并设置分片策略进行分片。
但有些时候我们并没有使用任何的分片键和分片策略，
可还想将SQL路由到目标数据库和表，
就需要通过手动干预指定SQL的目标数据库和表信息，也叫强制路由。
```

































