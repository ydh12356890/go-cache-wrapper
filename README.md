## 说明
基于开源otter封装一个缓存管理器
* 缓存的key和value的类型分别是string和[]byte
* 缓存的capacity和ttl可自定义
* 缓存中的entries上限为：capacity / ( len(key) + len(value) )
* 提供Get和Set接口
* 提供可使用prometheus抓取的缓存监控指标
* 使用go版本：1.20
## 安装
``go get github.com/ydh12356890/go-cache-wrapper``

## 导入
``import "github.com/ydh12356890/go-cache-wrapper/cache"``
