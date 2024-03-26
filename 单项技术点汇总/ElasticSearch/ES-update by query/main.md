## ES: update by query

```
POST delivery_store_coupon/delivery_store_coupon/_update_by_query
{
  "script": {
    "lang": "painless",
    "source": "if (ctx._source.soureType == 0) {ctx._source.soureType = [0]} else if (ctx._source.soureType == 1) {ctx._source.soureType = [1]} else if (ctx._source.soureType == 2) {ctx._source.soureType = [2]}"
  }
}
```

### 文章目录
- _update_by_query 的应用场景
- 造数据
  - 1、修改一个字段的值
  - 2、***给es里某个字段增加一个子类型，要求之前的数据也能被查询到***

es 版本为7.9.3

### 造数据
```
POST test
{
  "mappings" : {
      "properties" : {
        "name" : {
          "type" : "text",
          "fields" : {
            "keyword" : {
              "type" : "keyword",
              "ignore_above" : 256
            }
          }
        }
      }
    }
}

POST test/_doc/1
{
  "name": "chb",
  "age": "20"
}

POST test/_doc/2
{
  "name": "ling",
  "age": 18
}

POST test/_doc/3
{
  "name": "旺仔",
  "age": 1
}

POST test/_doc/4
{
  "name": "李四"
}
```

### 1、修改一个字段的值
```
# 修改李四的年龄为44
POST test/_update_by_query
{
  "script": {
    "source": "ctx._source.age = 44",
    "lang": "painless" // 无痛的
  },
  "query": {
    "bool": {
      "must_not": [
        {
          "exists": {
              "field": "age"
          }
        }
      ]
    }
  }
}
```

### 2、 给es里某个字段增加一个子类型，要求之前的数据也能被查询到
修改mapping，添加一个子字段
```
POST test/_mapping
{
  "properties": {
    "name": {
      "type": "text",
      "fields": {
        "keyword": {
          "type": "keyword",
          "ignore_above": 256
        },
        "ik_smart": {
          "type": "text",
          "analyzer": "ik_smart"
        }
      }
    }
  }
}
```
插入一条新的数据
```
PUT test/_doc/5
{
  "name": "王五",
  "age": 35
}
```
查询 李四，王五，发现查不到李四
```
GET test/_search
{
  "query": {
    "match": {
      "name.ik_smart": "李四"
    }
  }
}

GET test/_search
{
  "query": {
    "match": {
      "name.ik_smart": "王五"
    }
  }
}
```
因为李四是 更改mapping之前插入，新增字段没有在老数据上生效，导致查询不出
为了之前的数据也能被查询到，我们通过 _update_by_query
```
POST test/_update_by_query
```

### ES. _update_by_query
场景： 给es里某个字段增加一个子类型，要求之前的数据也能被查询到
- 如上场景，我们可以使用es里的_update_by_query
- 例如
```
POST class/_update_by_query
```
直接对加完类型的索引使用即可。

下面是一个例子
```
PUT class
{
  "mappings" : {
      "properties" : {
        "student" : {
          "type" : "nested",
          "properties" : {
            "name" : {
              "type" : "keyword",
              "fields" : {
                "ik" : {
                  "type" : "text",
                  "analyzer" : "ik_max_word"
                }
              }
            }
          }
        }
      }
    }
}
```
创建一个索引，设置mapping
```
{
  "acknowledged" : true,
  "shards_acknowledged" : true,
  "index" : "class"
}
```
索引创建完毕，开始写入数据
```
PUT class/_doc/1
{
  "student" : [
    {
        "name" : "GO"
    }
  ]
}
```
写入一条数据，然后修改mapping，增加一个子字段
```

PUT class/_mapping
{
 "properties" : {
        "student" : {
          "type" : "nested",
          "properties" : {
            "name" : {
              "type" : "keyword",
              "fields" : {
                "ik" : {
                  "type" : "text",
                  "analyzer" : "ik_max_word"
                },
                "ik_smart" : {
                  "type" : "text",
                  "analyzer" : "ik_smart"
                }
              }
            }
          }
        }
      }
}
```
加号字段后我们再插入一条数据
```
PUT class/_doc/2
{
    "student" : [
        {
          "name" : "GO"
        }
    ]
}
```
然后我们分别使用最大颗粒度查询这两个name
```
GET class/_search
{
  "query": {
    "nested": {
      "path": "student",
      "query": {
        "bool": {
          "filter": {
           "term": {
             "student.name.ik_smart": "go"
           }
          }
        }
      }
    }
  }
}
 
 
GET class/_search
{
  "query": {
    "nested": {
      "path": "student",
      "query": {
        "bool": {
          "filter": {
           "term": {
             "student.name.ik_smart": "php"
           }
          }
        }
      }
    }
  }
}
```
结果go查不到，而php能查到，因php是在更改mapping后插入的，新增加的子字段在php上生效，所以可以查询到，
而go是在修改mapping之前插入的，新字段没有生效，所以查询不到，如果我们要让旧数据也生效，该怎么做呢

```
POST class/_update_by_query
```

我们可以直接对索引适用update_by_query，这样旧的数据也会增加子字段，
再次查询的时候我们就会发现，go也能搜索到了
