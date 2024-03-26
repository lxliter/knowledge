## Elastic Docs ›Elasticsearch Guide [8.12]

### What is Elasticsearch?
You know, for search (and analysis)
Elasticsearch is the distributed search and analytics engine at the heart of the Elastic Stack. 
Logstash and Beats facilitate collecting, aggregating, and enriching your data and storing it in Elasticsearch. 
Kibana enables you to interactively explore, visualize, 
and share insights into your data and manage and monitor the stack. 
Elasticsearch is where the indexing, search, and analysis magic happens.
```
你知道，用于搜索（和分析）
Elasticsearch 是 Elastic Stack 核心的分布式搜索和分析引擎。
Logstash 和 Beats 有助于收集、聚合和丰富您的数据并将其存储在 Elasticsearch 中。
Kibana 使您能够交互式地探索、可视化、
分享对数据的见解并管理和监控堆栈。
Elasticsearch 是索引、搜索和分析魔法发生的地方。
```

Elasticsearch provides near real-time search and analytics for all types of data. 
Whether you have structured or unstructured text, numerical data, or geospatial data, 
Elasticsearch can efficiently store and index it in a way that supports fast searches. 
You can go far beyond simple data retrieval and aggregate information to discover trends and patterns in your data. 
And as your data and query volume grows, 
the distributed nature of Elasticsearch enables your deployment to grow seamlessly right along with it.
```
Elasticsearch 为所有类型的数据提供近乎实时的搜索和分析。
无论您有结构化或非结构化文本、数字数据还是地理空间数据，
Elasticsearch 可以以支持快速搜索的方式有效地存储和索引它。
您不仅可以进行简单的数据检索和聚合信息，还可以发现数据中的趋势和模式。
随着您的数据和查询量的增长，
Elasticsearch 的分布式特性使您的部署能够随之无缝增长。
```

While not every problem is a search problem, 
Elasticsearch offers speed and flexibility to handle data in a wide variety of use cases:
```
虽然并非所有问题都是搜索问题，
Elasticsearch 提供速度和灵活性来处理各种用例中的数据：
```

- Add a search box to an app or website
- Store and analyze logs, metrics, and security event data
- Use machine learning to automatically model the behavior of your data in real time
- Use Elasticsearch as a vector database to create, store, and search vector embeddings
- Automate business workflows using Elasticsearch as a storage engine
- Manage, integrate, and analyze spatial information using Elasticsearch as a geographic information system (GIS)
- Store and process genetic data using Elasticsearch as a bioinformatics research tool

```
- 将搜索框添加到应用程序或网站
- 存储和分析日志、指标和安全事件数据
- 使用机器学习实时自动建模数据行为
- 使用 Elasticsearch 作为矢量数据库来创建、存储和搜索矢量嵌入
- 使用 Elasticsearch 作为存储引擎实现业务工作流程自动化
- 使用 Elasticsearch 作为地理信息系统 (GIS) 管理、集成和分析空间信息
- 使用 Elasticsearch 作为生物信息学研究工具来存储和处理遗传数据
```

We’re continually amazed by the novel[新颖的，珍奇的] ways people use search. 
But whether your use case is similar to one of these, or you’re using Elasticsearch to tackle a new problem, 
the way you work with your data, documents, and indices in Elasticsearch is the same.

```
人们使用搜索的新颖方式不断让我们感到惊讶。
但无论您的用例与其中之一类似，还是您正在使用 Elasticsearch 来解决新问题，
在 Elasticsearch 中处理数据、文档和索引的方式是相同的。
```

#### Data in: documents and indices
Elasticsearch is a distributed document store. 
Instead of storing information as rows of columnar data, 
Elasticsearch stores complex data structures that have been serialized as JSON documents. 
When you have multiple Elasticsearch nodes in a cluster, 
stored documents are distributed across the cluster and can be accessed immediately from any node.
```
Elasticsearch 是一个分布式文档存储。
不是将信息存储为柱状数据行，
Elasticsearch 存储已序列化为 JSON 文档的复杂数据结构。
当集群中有多个 Elasticsearch 节点时，
存储的文档分布在整个集群中，并且可以从任何节点立即访问。
```

When a document is stored, it is indexed and fully searchable in near real-time--within 1 second. 
Elasticsearch uses a data structure called an inverted index that supports very fast full-text searches. 
An inverted index lists every unique word that appears in any document and identifies all of the documents each word occurs in.
```
存储文档时，会在 1 秒内近乎实时地为其建立索引并完全可搜索。
Elasticsearch 使用称为倒排索引的数据结构，支持非常快速的全文搜索。
倒排索引列出了任何文档中出现的每个唯一单词，并标识了每个单词出现的所有文档。
```
```
Elasticsearch uses a data structure called an inverted index that supports very fast full-text searched.
```

An index can be thought of as an optimized collection of documents and each document is a collection of fields, 
which are the key-value pairs that contain your data. 
By default, Elasticsearch indexes all data in every field and each indexed field has a dedicated, optimized data structure. 
For example, text fields are stored in inverted indices, and numeric and geo fields are stored in BKD trees. 
The ability to use the per-field data structures to assemble and return search results is what makes Elasticsearch so fast.

```
索引可以被认为是文档的优化集合，每个文档都是字段的集合，
它们是包含您的数据的键值对。
默认情况下，Elasticsearch 会索引每个字段中的所有数据，并且每个索引字段都有一个专用的、优化的数据结构。
例如，文本字段存储在倒排索引中，数字和地理字段存储在 BKD 树中。
使用每个字段的数据结构来组合和返回搜索结果的能力使得 Elasticsearch 如此之快。
```

Elasticsearch also has the ability to be schema-less, 
which means that documents can be indexed without explicitly specifying how to handle each of the different fields that might occur in a document. 
When dynamic mapping is enabled, Elasticsearch automatically detects and adds new fields to the index. 
This default behavior makes it easy to index and explore your data—just start indexing documents and Elasticsearch will detect and map booleans, 
floating point and integer values, dates, and strings to the appropriate Elasticsearch data types.

```
Elasticsearch 还具有无模式的能力，
这意味着可以对文档建立索引，而无需明确指定如何处理文档中可能出现的每个不同字段。
启用动态映射后，Elasticsearch 会自动检测新字段并将其添加到索引中。
这种默认行为使索引和探索数据变得容易 — 只需开始索引文档，Elasticsearch 就会检测并映射布尔值，
浮点和整数值、日期和字符串转换为适当的 Elasticsearch 数据类型。
```

Ultimately, however, you know more about your data and how you want to use it than Elasticsearch can. 
You can define rules to control dynamic mapping and explicitly define mappings to take full control of how fields are stored and indexed.
```
但最终，您比 Elasticsearch 更了解您的数据以及您想要如何使用它。
您可以定义规则来控制动态映射并显式定义映射以完全控制字段的存储和索引方式。
```

Defining your own mappings enables you to:

- ***Distinguish between full-text string fields and exact value string fields***
- Perform language-specific text analysis
- Optimize fields for partial matching
- Use custom date formats
- Use data types such as geo_point and geo_shape that cannot be automatically detected
```
定义您自己的映射使您能够：
- 区分全文字符串字段和精确值字符串字段
- 执行特定语言的文本分析
- 优化部分匹配字段
- 使用自定义日期格式
- 使用无法自动检测的数据类型，例如geo_point和geo_shape
```

It’s often useful to index the same field in different ways for different purposes. 
For example, you might want to index a string field as both a text field for full-text search and as a keyword field for sorting or aggregating your data. 
Or, you might choose to use more than one language analyzer to process the contents of a string field that contains user input.
```
为了不同的目的以不同的方式对同一字段建立索引通常很有用。
例如，您可能希望将字符串字段索引为全文搜索的文本字段和排序或聚合数据的关键字字段。
或者，您可以选择使用多个语言分析器来处理包含用户输入的字符串字段的内容。
```

The analysis chain that is applied to a full-text field during indexing is also used at search time. 
When you query a full-text field, 
the query text undergoes the same analysis before the terms are looked up in the index.
```
在索引期间应用于全文字段的分析链也在搜索时使用。
当您查询全文字段时，
在索引中查找术语之前，查询文本会经历相同的分析。
```

#### Information out: search and analyze
While you can use Elasticsearch as a document store and retrieve documents and their metadata, 
the real power comes from being able to easily access the full suite of search capabilities built on the Apache Lucene search engine library.

```
虽然您可以使用 Elasticsearch 作为文档存储并检索文档及其元数据，
真正的力量来自于能够轻松访问基于 Apache Lucene 搜索引擎库构建的全套搜索功能。
```

Elasticsearch provides a simple, coherent REST API for managing your cluster and indexing and searching your data. 
For testing purposes, you can easily submit requests directly from the command line or through the Developer Console in Kibana. 
From your applications, 
you can use the Elasticsearch client for your language of choice: Java, JavaScript, Go, .NET, PHP, Perl, Python or Ruby.
```
Elasticsearch 提供了一个简单、一致的 REST API，用于管理集群以及索引和搜索数据。
出于测试目的，您可以直接从命令行或通过 Kibana 中的开发者控制台轻松提交请求。
从您的应用程序中，
您可以使用适合您选择的语言的 Elasticsearch 客户端：Java、JavaScript、Go、.NET、PHP、Perl、Python 或 Ruby。
```

Searching your data

The Elasticsearch REST APIs support structured queries, full text queries, and complex queries that combine the two. 
Structured queries are similar to the types of queries you can construct in SQL. 
For example, you could search the gender and age fields in your employee index and sort the matches by the hire_date field. 
Full-text queries find all documents that match the query string and return them sorted by relevance—how good a match they are for your search terms.

```
搜索您的数据
Elasticsearch REST API 支持结构化查询、全文查询以及两者结合的复杂查询。
结构化查询类似于您可以在 SQL 中构建的查询类型。
例如，您可以搜索员工索引中的性别和年龄字段，并按雇佣日期字段对匹配项进行排序。
全文查询查找与查询字符串匹配的所有文档，并返回按相关性排序的文档 - 它们与您的搜索词的匹配程度。
```

In addition to searching for individual terms, 
you can perform phrase searches, similarity searches, and prefix searches, and get autocomplete suggestions.
```
除了搜索单个术语之外，
您可以执行短语搜索、相似性搜索和前缀搜索，并获得自动补全建议。
In addition to searching for individual terms,
you can perform phrase searches,similarity searched,and prefix searched,and get autocomplete suggestions.
```

Have geospatial or other numerical data that you want to search? 
Elasticsearch indexes non-textual data in optimized data structures that support high-performance geo and numerical queries.
```
您有要搜索的地理空间或其他数字数据吗？
Elasticsearch 在支持高性能地理和数字查询的优化数据结构中索引非文本数据。
Elasticsearch indexes non-textual data in optimized data structures that support high-performance geo and numerical queries.
```

You can access all of these search capabilities using Elasticsearch’s comprehensive JSON-style query language (Query DSL[domain-specific language]). 
You can also construct SQL-style queries to search and aggregate data natively inside Elasticsearch, 
and JDBC and ODBC drivers enable a broad range of third-party applications to interact with Elasticsearch via SQL.
```
您可以使用 Elasticsearch 的综合 JSON 式查询语言 (Query DSL) 访问所有这些搜索功能。
您还可以构建 SQL 样式的查询来在 Elasticsearch 内本地搜索和聚合数据，
JDBC 和 ODBC 驱动程序使广泛的第三方应用程序能够通过 SQL 与 Elasticsearch 交互。
You can access of these search capabilities using Elasticsearch's comprehensive JSON-style query language(Query DSL)
You can also construct SQL-style queries to search and aggregate data natively inside Elasticsearch,
and JDBC and ODBC drivers enable a broad range of third-party applications to interact with Elasticsearch via SQL.
```

Analyzing your data

Elasticsearch aggregations enable you to build complex summaries of your data and gain insight into key metrics, patterns, and trends. 
Instead of just finding the proverbial[prəˈvɜːrbiəl:谚语的，俗话所说的；众所周知的] “needle in a haystack”, aggregations enable you to answer questions like:

- How many needles are in the haystack?
- What is the average length of the needles?
- What is the median[ˈmiːdiən: 中间值的，中位数的；在中间的] length of the needles, broken down by manufacturer?
- How many needles were added to the haystack in each of the last six months?

```
分析您的数据
Elasticsearch 聚合使您能够构建复杂的数据摘要并深入了解关键指标、模式和趋势。
聚合不仅仅是寻找众所周知的“大海捞针”，还可以让您回答以下问题：

- 大海捞针有多少根？
- 针的平均长度是多少？
- 按制造商细分的针的平均长度是多少？
- 过去六个月每年大海捞针有多少？
Elasticsearch aggregations enable you to build complex summaries of your data and gain insight into key metrics,patterns,and trends.
Instead of just finding the proverbial "needle in a haystack", aggregations enable you to answer questions like:
How many needles are in the haystack?
What is the average length of the needles?
What is the median length of the needles,broken down by manufacturer?
How many needles were added to the haystack in each of the last six months?
```

You can also use aggregations to answer more subtle[不易察觉的，微妙的] questions, such as:

- What are your most popular needle manufacturers?
- Are there any unusual or anomalous[əˈnɑːmələs:异常的；不规则的；不恰当的] clumps[klʌmp:丛] of needles?
```
您还可以使用聚合来回答更微妙的问题，例如：

- 您最受欢迎的针制造商是哪些？
- 是否有任何不寻常或异常的针头团？
```

Because aggregations leverage[ˈlevərɪdʒ:影响力，手段；杠杆力，杠杆作用；<美>杠杆比率] the same data-structures used for search, they are also very fast. 
This enables you to analyze and visualize your data in real time. 
Your reports and dashboards update as your data changes so you can take action based on the latest information.
```
由于聚合利用与搜索相同的数据结构，因此它们也非常快。
这使您能够实时分析和可视化数据。
您的报告和仪表板会随着数据的变化而更新，以便您可以根据最新信息采取行动。

Because aggregations leverage the same data-structures used for search,
they are also very fast.
This enables you to analyze and visualize your data in real time.
Your reports and dashboards update as your data changes so you can take action based on the latest information.
```

What’s more, aggregations operate alongside search requests. 
You can search documents, filter results, and perform analytics at the same time, on the same data, in a single request. 
And because aggregations are calculated in the context of a particular search, you’re not just displaying a count of all size 70 needles, 
you’re displaying a count of the size 70 needles that match your users' search criteria—for example, 
all size 70 non-stick embroidery[ɪmˈbrɔɪdəri：刺绣技法，刺绣活儿；绣花] needles.
```
此外，聚合与搜索请求一起运行。
您可以在单个请求中同时对相同数据搜索文档、筛选结果和执行分析。
由于聚合是在特定搜索的上下文中计算的，因此您不仅仅显示所有尺寸 70 针的计数，
您正在显示符合用户搜索条件的 70 号针的数量，例如，
所有尺寸 70 不粘刺绣针。
```

But wait, there’s more
Want to automate the analysis of your time series data? 
You can use machine learning features to create accurate baselines of normal behavior in your data and identify anomalous[əˈnɑːmələs:异常的；不规则的；不恰当的] patterns. 
With machine learning, you can detect:

- Anomalies related to temporal[ˈtempərəl:] deviations in values, counts, or frequencies
- Statistical rarity[ˈrerəti:罕见的人（或物），珍品；稀有，罕见；]
- Unusual behaviors for a member of a population

And the best part? You can do this without having to specify algorithms, models, or other data science-related configurations.
```
但等等，还有更多
想要自动分析时间序列数据吗？ 您可以使用机器学习功能在数据中创建准确的正常行为基线并识别异常模式。
通过机器学习，您可以检测：
- 与值、计数或频率的时间偏差相关的异常
- 统计上的稀有性
- 群体成员的异常行为
最好的部分是什么？ 您无需指定算法、模型或其他与数据科学相关的配置即可执行此操作。
```

#### Scalability and resilience: clusters, nodes, and shards
```
可扩展性和弹性：集群、节点和分片
```

Elasticsearch is built to be always available and to scale with your needs. 
It does this by being distributed by nature. 
You can add servers (nodes) to a cluster to increase capacity and Elasticsearch automatically distributes your data and query load across all of the available nodes. 
No need to overhaul[彻底检修，全面改造；全面改革（制度或方法）] your application, Elasticsearch knows how to balance multi-node clusters to provide scale and high availability. 
The more nodes, the merrier.
```
Elasticsearch 旨在始终可用并根据您的需求进行扩展。
它通过自然分布来做到这一点。
您可以向集群添加服务器（节点）以增加容量，Elasticsearch 会自动在所有可用节点上分配数据和查询负载。
无需彻底修改您的应用程序，Elasticsearch 知道如何平衡多节点集群以提供规模和高可用性。
节点越多越好。
```

How does this work? Under the covers, 
an Elasticsearch index is really just a logical grouping of one or more physical shards, 
where each shard is actually a self-contained index. By distributing the documents in an index across multiple shards, 
and distributing those shards across multiple nodes, 
Elasticsearch can ensure redundancy, which both protects against hardware failures and increases query capacity as nodes are added to a cluster. 
As the cluster grows (or shrinks), Elasticsearch automatically migrates shards to rebalance the cluster.

```
这是如何运作的？ 在幕后，
Elasticsearch 索引实际上只是一个或多个物理分片的逻辑分组，
其中每个分片实际上是一个独立的索引。 通过将索引中的文档分布在多个分片上，
并将这些分片分布在多个节点上，
Elasticsearch 可以确保冗余，这既可以防止硬件故障，又可以在将节点添加到集群时提高查询能力。
随着集群的增长（或缩小），Elasticsearch 会自动迁移分片以重新平衡集群。
```

There are two types of shards: ***primaries and replicas***. 
Each document in an index belongs to one primary shard. 
A replica shard is a copy of a primary shard. 
Replicas provide redundant copies of your data to protect against hardware failure and increase capacity to serve read requests like searching or retrieving a document.
```
有两种类型的分片：主分片和副本分片。 索引中的每个文档都属于一个主分片。
副本分片是主分片的副本。
副本提供数据的冗余副本，以防止硬件故障并提高服务读取请求（例如搜索或检索文档）的能力。
```

The number of primary shards in an index is fixed at the time that an index is created, 
but the number of replica shards can be changed at any time, without interrupting indexing or query operations.
```
索引中主分片的数量在创建索引时是固定的，
但副本分片的数量可以随时更改，而无需中断索引或查询操作。
The number of primary shards in an index is fixed at the time that an index is created,
but the number of replica shards can be changed at any time,without interrupting indexing or query operations.
```

It depends...
There are a number of performance considerations and trade offs with respect to shard size and the number of primary shards configured for an index. 
The more shards, the more overhead[开销] there is simply in maintaining those indices. 
The larger the shard size, the longer it takes to move shards around when Elasticsearch needs to rebalance a cluster.
```
这取决于...
对于分片大小和为索引配置的主分片数量，存在许多性能考虑因素和权衡。
分片越多，维护这些索引的开销就越大。
分片大小越大，当 Elasticsearch 需要重新平衡集群时，移动分片所需的时间就越长。
```

Querying lots of small shards makes the processing per shard faster, but more queries means more overhead, 
so querying a smaller number of larger shards might be faster. In short…it depends.

As a starting point:
- Aim to keep the average shard size between a few GB and a few tens of GB. For use cases with time-based data, it is common to see shards in the 20GB to 40GB range.
- Avoid the gazillion[ɡəˈzɪljən:极大量] shards problem. The number of shards a node can hold is proportional to the available heap space. 
As a general rule, the number of shards per GB of heap space should be less than 20.

The best way to determine the optimal configuration for your use case is through testing with your own data and queries.
```
查询大量小分片可以使每个分片的处理速度更快，但更多查询意味着更多开销，
因此查询较少数量的较大分片可能会更快。 简而言之……这取决于。

作为起点：
- 目标是将平均分片大小保持在几 GB 到几十 GB 之间。 对于具有基于时间的数据的用例，20GB 到 40GB 范围内的分片很常见。
- 避免大量碎片问题。 节点可以容纳的分片数量与可用堆空间成正比。作为一般规则，每 GB 堆空间的分片数量应小于 20。

确定适合您的用例的最佳配置的最佳方法是使用您自己的数据和查询进行测试。
```

In case of disaster
A cluster’s nodes need good, reliable connections to each other. To provide better connections, 
you typically co-locate[放在同一] the nodes in the same data center or nearby data centers. 
However, to maintain high availability, you also need to avoid any single point of failure. 
In the event of a major outage[中断] in one location, servers in another location need to be able to take over. 
The answer? Cross-cluster replication (CCR).

CCR provides a way to automatically synchronize indices from your primary cluster to a secondary remote cluster that can serve as a hot backup. 
If the primary cluster fails, the secondary cluster can take over. You can also use CCR to create secondary clusters to serve read requests in geo-proximity[地理位置接近] to your users.

Cross-cluster replication is active-passive. The index on the primary cluster is the active leader index and handles all write requests. 
Indices replicated to secondary clusters are read-only followers.

Care and feeding
As with any enterprise system, you need tools to secure, manage, and monitor your Elasticsearch clusters. 
Security, monitoring, and administrative features that are integrated into Elasticsearch enable you to use Kibana as a control center for managing a cluster. 
Features like downsampling and index lifecycle management help you intelligently manage your data over time.
```
发生灾害时
集群的节点之间需要良好、可靠的连接。 为了提供更好的连接，您通常将节点放在同一数据中心或附近的数据中心中。 
但是，为了保持高可用性，您还需要避免任何单点故障。 如果一个位置发生重大中断，另一位置的服务器需要能够接管。 答案？ 跨集群复制 (CCR)。

CCR 提供了一种自动将索引从主集群同步到可用作热备份的辅助远程集群的方法。 
如果主集群发生故障，辅助集群可以接管。 您还可以使用 CCR 创建辅助集群来处理与用户地理位置接近的读取请求。

跨集群复制是主动-被动的。 主集群上的索引是活动领导索引并处理所有写入请求。 复制到辅助集群的索引是只读追随者。

护理和喂养
与任何企业系统一样，您需要工具来保护、管理和监控您的 Elasticsearch 集群。 
集成到 Elasticsearch 中的安全、监控和管理功能使您能够使用 Kibana 作为管理集群的控制中心。 
下采样和索引生命周期管理等功能可帮助您随着时间的推移智能地管理数据。
```

### What’s new in 8.12
Here are the highlights of what’s new and improved in Elasticsearch 8.12! 
For detailed information about this release, see the Release notes and Migration guide.
```
以下是 Elasticsearch 8.12 中的新增功能和改进的亮点！
有关此版本的详细信息，请参阅发行说明和迁移指南。
```

Other versions:
8.11 | 8.10 | 8.9 | 8.8 | 8.7 | 8.6 | 8.5 | 8.4 | 8.3 | 8.2 | 8.1 | 8.0

Enable query phase parallelism within a single shard
Activate inter-segment search concurrency by default in the query phase, 
in order to enable parallelizing search execution across segments that a single shard is made of.
```
在单个分片内启用查询阶段并行性
查询阶段默认激活段间搜索并发，
为了能够跨单个分片组成的段并行搜索执行。
```

Add new int8_hsnw index type for int8 quantization for HNSW
This commit adds a new index type called int8_hnsw. This index will automatically quantized float32 values into int8 byte values. 
While this increases disk usage by 25%, it reduces memory required for fast HNSW search by 75%. 
Dramatically reducing the resource overhead required for dense vector search.
```
为 HNSW 的 int8 量化添加新的 int8_hsnw 索引类型
此提交添加了一个名为 int8_hnsw 的新索引类型。 该索引会自动将 float32 值量化为 int8 字节值。
虽然这会使磁盘使用量增加 25%，但快速 HNSW 搜索所需的内存减少了 75%。
大幅减少密集向量搜索所需的资源开销。
```

### Quick start
This guide helps you learn how to:
- install and run Elasticsearch and Kibana (using Elastic Cloud or Docker),
- add simple (non-timestamped) dataset to Elasticsearch,
- run basic searches.
```
本指南可帮助您学习如何：
- 安装并运行 Elasticsearch 和 Kibana（使用 Elastic Cloud 或 Docker），
- 将简单（无时间戳）数据集添加到 Elasticsearch，
- 运行基本搜索。
```

If you’re interested in using Elasticsearch with Python, check out Elastic Search Labs. 
This is the best place to explore AI-powered search use cases[搜索用例], such as working with embeddings, vector search, and retrieval augmented[增加，增大；加强，补充] generation (RAG).
- Tutorial: this walks you through building a complete search solution with Elasticsearch, from the ground up.
- elasticsearch-labs repository: it contains a range of Python notebooks and example apps.
```
如果您有兴趣通过 Python 使用 Elasticsearch，请查看 Elastic Search Labs。
这是探索人工智能驱动的搜索用例的最佳场所，例如使用嵌入、矢量搜索和检索增强生成 (RAG)。
- 教程：这将引导您从头开始使用 Elasticsearch 构建完整的搜索解决方案。
- elasticsearch-labs 存储库：它包含一系列 Python 笔记本和示例应用程序。
```

#### Run Elasticsearch
The simplest way to set up Elasticsearch is to create a managed[托管的] deployment with Elasticsearch Service on Elastic Cloud. 
If you prefer to manage your own test environment, install and run Elasticsearch using Docker.
```
设置 Elasticsearch 的最简单方法是使用 Elastic Cloud 上的 Elasticsearch 服务创建托管部署。
如果您更喜欢管理自己的测试环境，请使用 Docker 安装并运行 Elasticsearch。
The simplest way to set up Elasticsearch is to create managed deployment with Elasticsearch Service on Elastic Cloud.
If you prefer to manage your own test environment, install and run Elasticsearch using Docker.
```

Elasticsearch Service
- 1.Get a free trial.
- 2.Log into Elastic Cloud.
- 3.Click Create deployment.
- 4.Give your deployment a name.
- 5.Click Create deployment and download the password for the elastic user.
- 6.Click Continue to open Kibana, the user interface for Elastic Cloud.
- 7.Click Explore on my own.
```
Elasticsearch服务
- 1.获得免费试用。
- 2.登录弹性云。
- 3.单击创建部署。
- 4.为您的部署命名。
- 5.单击创建部署并下载弹性用户的密码。
- 6.单击继续打开 Kibana，Elastic Cloud 的用户界面。
- 7.点击“我自己探索”。
```

Self-managed
Start a single-node cluster
We’ll use a single-node Elasticsearch cluster in this quick start, which makes sense for testing and development. 
Refer to Install Elasticsearch with Docker for advanced Docker documentation.
```
自我管理
启动单节点集群
在本快速入门中，我们将使用单节点 Elasticsearch 集群，这对于测试和开发很有意义。
有关高级 Docker 文档，请参阅使用 Docker 安装 Elasticsearch。
Self-managed
Start a single-node cluster
We'll use a single-node Elasticsearch cluster in this quick start,
which makes sense for testing and development.
Refer to install Elasticsearch with Docker for advanced Docker documentation.
```

- 1.Run the following Docker commands:
```bash
docker network create elastic
docker pull docker.elastic.co/elasticsearch/elasticsearch:8.12.2
docker run --name es01 --net elastic -p 9200:9200 -p 9300:9300 -e "discovery.type=single-node" -t docker.elastic.co/elasticsearch/elasticsearch:8.12.2
```
```
Run the following Docker commands:
docker network create elastic
docker pull docker.elastic.co/elasticsearch/elasticsearch:8.12.2
docker run --name es01 --net elastic -p 9200:9200 -p 9300:9300 -e "discovery.type=single-node" -t docker.elastic.co/elasticsearch/elasticsearch:8.12.2
```

- 2.Copy the generated elastic password and enrollment[登记，注册] token, which are output to your terminal. 
You’ll use these to enroll Kibana with your Elasticsearch cluster and log in. 
These credentials are only shown when you start Elasticsearch for the first time.
We recommend storing the elastic password as an environment variable in your shell. Example:
```
- 2.复制生成的elastic密码和注册令牌，这些内容将输出到您的终端。
您将使用它们在 Elasticsearch 集群中注册 Kibana 并登录。
这些凭据仅在您首次启动 Elasticsearch 时显示。
我们建议将弹性密码存储为 shell 中的环境变量。 例子：
Copy the generated elastic password adn enrollment token,which are output to your terminal.
You'll use these to enroll Kibana with your Elasticsearch cluster and log in.
These credentials are only shown when you start Elasticsearch for the first time.
We recommend storing the elastic password as an environemnt variable in your shell.Example:
```
```
export ELASTIC_PASSWORD="your_password"
```

- 3.Copy the http_ca.crt SSL certificate from the container to your local machine.
```
docker cp es01:/usr/share/elasticsearch/config/certs/http_ca.crt .
```

- 4.Make a REST API call to Elasticsearch to ensure the Elasticsearch container is running.
```
curl --cacert http_ca.crt -u elastic:$ELASTIC_PASSWORD https://localhost:9200
```

Run Kibana
Kibana is the user interface for Elastic. 
It’s great[非常适合] for getting started with Elasticsearch and exploring your data. 
We’ll be using the Dev Tools Console in Kibana to make REST API calls to Elasticsearch.
In a new terminal session, start Kibana and connect it to your Elasticsearch container:

```
docker pull docker.elastic.co/kibana/kibana:8.12.2
docker run --name kibana --net elastic -p 5601:5601 docker.elastic.co/kibana/kibana:8.12.2
```

```
运行Kibana
Kibana 是 Elastic 的用户界面。
它非常适合开始使用 Elasticsearch 和探索数据。
我们将使用 Kibana 中的开发工具控制台对 Elasticsearch 进行 REST API 调用。
在新的终端会话中，启动 Kibana 并将其连接到您的 Elasticsearch 容器：
docker pull docker.elastic.co/kibana/kibana:8.12.2
docker run --name kibana --net elastic -p 5601:5601 docker.elastic.co/kibana/kibana:8.12.2
```

When you start Kibana, a unique URL is output to your terminal. To access Kibana:
- 1.Open the generated URL in your browser.
- 2.Paste the enrollment token that you copied earlier, to connect your Kibana instance with Elasticsearch.
- 3.Log in to Kibana as the elastic user with the password that was generated when you started Elasticsearch.
```
当您启动 Kibana 时，一个唯一的 URL 会输出到您的终端。 访问 Kibana：
- 1.在浏览器中打开生成的 URL。
- 2.粘贴您之前复制的注册令牌，以将您的 Kibana 实例与 Elasticsearch 连接。
- 3.使用启动Elasticsearch时生成的密码以elastic用户身份登录Kibana。
When you start Kibana, a unique URL is output to your terminal. To access Kibana:
- 1.Open the generated URL in your browser.
- 2.Paste the enrollment token that you copied earlier, to connect your Kiaban instance with Elasticsearch.
- 3.Log in to Kibana as the elastic user with the password that was generated when you started Elasticsearch.
```

#### Send requests to Elasticsearch
You send data and other requests to Elasticsearch using REST APIs. 
This lets you interact with Elasticsearch using any client that sends HTTP requests, such as curl. 
You can also use Kibana’s Console to send requests to Elasticsearch.
```
您使用 REST API 向 Elasticsearch 发送数据和其他请求。
这使您可以使用任何发送 HTTP 请求的客户端（例如curl）与 Elasticsearch 交互。
您还可以使用 Kibana 的控制台向 Elasticsearch 发送请求。
You send data and other requests to Elasticsearch using REST APIs.
This lets you interact with Elasticsearch using any client that sends HTTP requests,such as curl.
You can also use Kibana's Console to sned requests to Elasticsearch.
```

Elasticsearch Service
Use Kibana
1.Open Kibana’s main menu ("☰" near Elastic logo) and go to Dev Tools > Console.
```
History Settings Help
GET _search
{
    "query":{
        "match_all":{}
    }
}
```
2.Run the following test API request in Console:
```
GET /
```
Use curl
To communicate with Elasticsearch using curl or another client, you need your cluster’s endpoint.
- 1.Open Kibana’s main menu and click Manage this deployment.
- 2.From your deployment menu, go to the Elasticsearch page. Click Copy endpoint.
- 3.To submit an example API request, run the following curl command in a new terminal session. 
Replace <password> with the password for the elastic user. Replace <elasticsearch_endpoint> with your endpoint.


















