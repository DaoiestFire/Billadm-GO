- [规范](#规范)
  - [参数和返回值](#参数和返回值)
- [账本](#账本)
  - [增加 POST](#增加-post)
    - [创建指定名称的账本](#创建指定名称的账本)
  - [删除 DELETE](#删除-delete)
    - [删除指定ID的账本](#删除指定id的账本)
  - [修改 PUT](#修改-put)
  - [查询 GET](#查询-get)
    - [查询所有账本](#查询所有账本)
    - [查询指定ID的账本](#查询指定id的账本)
- [消费记录](#消费记录)
  - [增加 POST](#增加-post-1)
    - [创建消费记录](#创建消费记录)
  - [删除 DELETE](#删除-delete-1)
    - [删除指定ID的消费记录](#删除指定id的消费记录)
  - [修改 PUT](#修改-put-1)
  - [查询 GET](#查询-get-1)
    - [查询指定账本中的所有消费记录](#查询指定账本中的所有消费记录)

## 规范

### 参数和返回值

* 端点：`http://127.0.0.1:31943`
* 均是 POST 方法
* 需要带参的接口，参数为 JSON 字符串，放置到 body 里，标头 Content-Type 为 `application/json`
* 返回值

  ```json
  {
    "code": 0,
    "msg": "",
    "data": {}
  }
  ```

  * `code`：非 0 为异常情况
  * `msg`：正常情况下是空字符串，异常情况下会返回错误文案
  * `data`：可能为 `{}`、`[]` 或者 `NULL`，根据不同接口而不同

## 账本

> /api/v1/ledger

### 增加 POST

#### 创建指定名称的账本

* 参数

  ```json
  {
    "name": "test-name"
  }
  ```
  * `name`：账本名称

* 返回值

  ```json
  {
    "code": 0,
    "msg": "",
    "data": "0197f526-8160-70e6-84f7-16f6f5c8417e"
  }
  ```
  * `data`：账本ID

### 删除 DELETE

#### 删除指定ID的账本

* 参数
  
  ```json
  {
    "ledger_id": "账本id"
  }
  ```

* 返回值

  ```json
  {
    "code": 0,
    "msg": "",
    "data": null
  }
  ```

### 修改 PUT

### 查询 GET

#### 查询所有账本

* 参数
  
  ```json
  {
    "ledger_id": "all"
  }
  ```
  * `ledger_id`：账本ID

* 返回值

  ```json
  {
    "code": 0,
    "msg": "",
    "data": [
        {
            "id": "0197fa7b-9b5b-73a4-b131-f022cabd1cf5",
            "name": "test-name",
            "created_at": "2025-07-12T01:14:54.1717438+08:00",
            "updated_at": "2025-07-12T01:14:54.1717438+08:00"
        }
    ]
  }
  ```
  * `data`：账本列表

#### 查询指定ID的账本

* 参数
  
  ```json
  {
    "ledger_id": "id1,id2"
  }
  ```
  * `ledger_id`：一个或多个账本ID，逗号分隔

* 返回值

  ```json
  {
    "code": 0,
    "msg": "",
    "data": [
        {
            "id": "0197fa7b-9b5b-73a4-b131-f022cabd1cf5",
            "name": "test-name",
            "created_at": "2025-07-12T01:14:54.1717438+08:00",
            "updated_at": "2025-07-12T01:14:54.1717438+08:00"
        }
    ]
  }
  ```
  * `data`：账本列表

## 消费记录

> /api/v1/tr

### 增加 POST

#### 创建消费记录

* 参数
  
  ```json
  {
    "ledger_id": "{{LEDGER_ID}}",
    "price": 3433,
    "transaction_type": "expense",
    "category": "",
    "description": "test des",
    "tags": "[\"tags1\",\"tags2\"]",
    "transaction_at": {{$date.millisecondsTimestamp}}
  }
  ```
  * `ledger_id`：账本ID
  * `price`：消费价格
  * `transaction_type`：消费类型，取值范围['expense','income','trasfer']，分别表示消费，收入，转账

* 返回值

  ```json
  {
    "code": 0,
    "msg": "",
    "data": null
  }
  ```

### 删除 DELETE

#### 删除指定ID的消费记录

* 参数
  
  ```json
  {
    "tr_id": "{{TR_ID}}"
  }
  ```
  * `tr_id`：消费记录ID

* 返回值

  ```json
  {
    "code": 0,
    "msg": "",
    "data": null
  }
  ```

### 修改 PUT

### 查询 GET

#### 查询指定账本中的所有消费记录

* 参数
  
  ```json
  {
    "ledger_id": "id1"
  }
  ```
  * `ledger_id`：账本ID

* 返回值

  ```json
  {
    "code": 0,
    "msg": "",
    "data": [
        {
            "transaction_id": "01981e51-c6df-716f-953f-dc8771370af8",
            "ledger_id": "01981e51-b02f-7730-9644-774491b459e4",
            "price": 4.5,
            "transaction_type": "expense",
            "category": "餐饮美食",
            "description": "早餐",
            "tags": [
                "三餐"
            ],
            "transaction_at": "2025-07-19T00:15:32.563+08:00",
            "created_at": "2025-07-19T00:15:32.5755985+08:00",
            "updated_at": "2025-07-19T00:15:32.5755985+08:00"
        },
        {
            "transaction_id": "01981e51-cf73-7789-a913-bcb1e2016b04",
            "ledger_id": "01981e51-b02f-7730-9644-774491b459e4",
            "price": 4.5,
            "transaction_type": "expense",
            "category": "餐饮美食",
            "description": "早餐",
            "tags": [
                "三餐"
            ],
            "transaction_at": "2025-07-19T00:15:34.742+08:00",
            "created_at": "2025-07-19T00:15:34.771494+08:00",
            "updated_at": "2025-07-19T00:15:34.771494+08:00"
        }
    ]
  }
  ```
  * `data`：消费记录列表