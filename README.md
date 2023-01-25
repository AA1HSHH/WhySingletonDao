# WhySingletonDao
## env deploy

```shell
sudo docker compose up -d
```

```shell
sudo docker compose down -v
```
## 疑问点
对之前发帖子的demo有点疑问，为什么在数据库操作的时候需要建一个单例userDao呢？

我实际尝试了下多协程插入和读取也没出现什么问题，这样做有什么好处吗？测试的代码见XXX_test。

