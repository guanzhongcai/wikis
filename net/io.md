# IO复用

## **select/poll**

select 实现多路复⽤的⽅式是，将已连接的 Socket 都放到⼀个⽂件描述符集合，然后调⽤

select 函数将⽂件描述符集合拷⻉到内核⾥，让内核来检查是否有⽹络事件产⽣，检查的⽅

式很粗暴，就是通过遍历⽂件描述符集合的⽅式，当检查到有事件产⽣后，将此 Socket 标记

为可读或可写， 接着再把整个⽂件描述符集合拷⻉回⽤户态⾥，然后⽤户态还需要再通过遍 

历的⽅法找到可读或可写的 Socket，然后再对其处理。

所以，对于 select 这种⽅式，需要进⾏ **2** 次「遍历」⽂件描述符集合，⼀次是在内核态⾥，

⼀个次是在⽤户态⾥ ，⽽且还会发⽣ **2** 次「拷⻉」⽂件描述符集合，先从⽤户空间传⼊内核

空间，由内核修改后，再传出到⽤户空间中。

select 使⽤固定⻓度的 `BitsMap`，表示⽂件描述符集合，⽽且所⽀持的⽂件描述符的个数是

有限制的，在 Linux 系统中，由内核中的 `FD_SETSIZE` 限制， 默认最⼤值为 1024 ，只能监

听 0~1023 的⽂件描述符。

poll 不再⽤ BitsMap 来存储所关注的⽂件描述符，取⽽代之⽤动态数组，以链表形式来组

织，突破了 select 的⽂件描述符个数限制，当然还会受到系统⽂件描述符限制。

但是 poll 和 select 并没有太⼤的本质区别，都是使⽤「`线性结构`」存储进程关注的 **Socket**

集合，因此都**需要遍历⽂件描述符集合来找到可读或可写的** **Socket**，**时间复杂度为** **O(n)**，**⽽**

**且也需要在⽤户态与内核态之间拷⻉⽂件描述符集合**，这种⽅式随着并发数上来，性能的损

耗会呈指数级增⻓。
