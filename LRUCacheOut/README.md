# 简单LRU算法实现
* 完成功能
  1. 借用java的hashMap的散列算法实现的散列表
  2. 实现了数据访问后移动到链尾操作，热点数据访问时间复杂度接近O(1)
* 未完成功能
  1. 散列表动态扩容、散列因子、扩容后散列值迁移
  2. 链表的淘汰机制，实现了链表节点删除函数但未实现具体的淘汰策略