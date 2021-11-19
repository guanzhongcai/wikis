# C++标准库之算法库

https://blog.csdn.net/qq_17044529/article/details/82646437



## 1.不修改序列的操作 

- all_of，any_of和none_of：判断一定范围内，是否全部，存在或不存在元素。
- for_each：将一个函数应用于某一范围的元素
- count和count_if：返回满足指定判别的元素数
- mismatch：查找两个范围第一个不同元素的位置
- find，find_if和find_if_not：查找满足特定条件的第一个元素
- find_end：查找一定范围内最后出现的元素序列
- find_first_of：查找元素集合中的任意元素
- adjacent_find：查找彼此相邻的两个相同（或其它的关系）的元素
- search：查找一个元素区间
- search_n：在区间中搜索连续一定数目次出现的元素

```c++
#include <iostream>
#include <algorithm>
#include <vector>
using namespace std;
int main()
{
    vector<int> ivec{ 1,2,2,4,5,6,7,8,9,10 };
    vector<int> v{ 5,6 ,7,8,9,10 };
    //判断容器内是否所有元素都满足条件
    if (all_of(ivec.begin(), ivec.end(), [](int i) {return i < 20; }))
        cout << "容器内所有元素都满足条件" << endl;
    //判断容器内是否有元素满足条件
    if(any_of(ivec.begin(),ivec.end(),[](int i){return i>9;}))
        cout << "容器内有任意一个元素满足条件" << endl;
    //判断容器内是否所有的元素都不满足条件
    if (none_of(ivec.begin(), ivec.end(), [](int i) {return i > 10; }))
        cout << "容器内所有元素都不满足条件" << endl;
    //遍历容器内元素
    for_each(ivec.begin(), ivec.end(), [](const int& n) { cout << n << " "; });
    cout << endl << "返回满足条件的元素个数：" << endl;
    cout << count(ivec.begin(), ivec.end(), 5) << endl;
    cout << count_if(ivec.begin(), ivec.end(), [](int i) {return i < 4; }) << endl;
    cout << "返回一个pair对象，first为第一个不匹配元素的迭代器，second为第一个匹配的元素的迭代器：" << endl;
    cout << *mismatch(ivec.begin(), ivec.end(), v.begin(), v.end()).first << endl;
    cout << *mismatch(ivec.begin(), ivec.end(), v.begin(), v.end()).second << endl;
    cout << "返回满足条件的第一个元素的迭代器：" << endl;
    cout << *find(ivec.begin(), ivec.end(), 5) << endl;
    cout << *find_if(ivec.begin(), ivec.end(), [](int i) {return i > 5; })<<endl;
    cout << "返回不满足条件的第一个元素的迭代器" << endl;
    cout << *find_if_not(ivec.begin(), ivec.end(), [](int i) {return i < 5; }) << endl;
    cout << "如果后一个区间是前一个区间的子区间，则返回后一个区间的起始迭代器，否则返回第一个区间的末尾迭代器：" << endl;
    cout << *find_end(ivec.begin(), ivec.end(), v.begin(),v.end())<<endl;
    cout << "返回两个区间首个相等元素的迭代器，如果没有则返回第一个区间的末端迭代器：" << endl;
    cout << *find_first_of(ivec.begin(), ivec.end(), v.begin(), v.end()) << endl;
    cout << "返回首个相邻元素相等的第一个元素的迭代器，如果没有则返回末端迭代器：" << endl;
    cout << *adjacent_find(ivec.begin(), ivec.end()) << endl;
    cout << "返回第一个相等的元素的迭代器，如果没有则返回第一个区间的末端迭代器：" << endl;
    cout << *search(ivec.begin(), ivec.end(), v.begin(), v.end()) << endl;
    cout << "返回满足条件(值为2，出现次数为2)的第一个元素的迭代器，如果没有则返回末端迭代器：" << endl;
    cout << *search_n(ivec.begin(), ivec.end(), 2, 2) << endl;
}
```



## 2.修改序列的操作 

- copy和copy_if ：将某一范围的元素复制到一个新的位置
- copy_n：复制一定数目的元素到新的位置
- copy_backward：按从后往前的顺序复制一个范围内的元素
- move：将某一范围的元素移动到一个新的位置
- move_backward：按从后往前的顺序移动某一范围的元素到新的位置
- fill：将一个值赋给一个范围内的元素
- fill_n：将一个值赋给一定数目的元素
- transform：将一个函数应用于某一范围的元素
- generate：赋值相继的函数调用结果给范围中的每个元素
- generate_n：赋值相继的函数调用结果给范围中的 N 个元素
- remove和remove_if：移除满足特定标准的元素
- remove_copy和remove_copy_if：复制一个范围内不满足特定条件的元素
- replace和replace_if：将所有满足特定条件的元素替换为另一个值
  replace_copy和replace_copy_if：复制一个范围内的元素，并将满足特定条件的元素替换为另一个值
  swap：交换两个对象的值
  swap_ranges：交换两个范围的元素
  iter_swap：交换两个迭代器所指向的元素
  reverse：将区间内的元素颠倒顺序
  reverse_copy：将区间内的元素颠倒顺序并复制
  rotate：将区间内的元素旋转
  rotate_copy：将区间内的元素旋转并复制
  random_shuffle：将范围内的元素随机重新排序
  unique：删除区间内连续重复的元素
  unique_copy：删除区间内连续重复的元素并复制



## 3.划分操作 ^

is_partitioned：判断区间是否被给定的谓词划分
partition：把一个区间的元素分为两组
partition_copy：将区间内的元素分为两组复制到不同位置
stable_partition：将元素分为两组，同时保留其相对顺序
partition_point：定位已划分的区域的划分点
这里写代码片

## 4.排序操作 ^

is_sorted：检查区间元素是否按升序排列
is_sorted_until：找出最大的已排序子范围
sort：将区间按升序排序
partial_sort：将区间内较小的N个元素排序
partial_sort_copy：对区间内的元素进行复制并部分排序
stable_sort：将区间内的元素排序，同时保持相等的元素之间的顺序
nth_element：将给定的区间部分排序，确保区间被给定的元素划分

```c++
#include <iostream>
#include <algorithm>
#include <vector>
using namespace std;
int main()
{
    vector<int> ivec{ 3,2,2,4,5,6,7,8,9,10 };
    vector<int> v{ 1,5,6,7,4,9 };
    vector<int> v2{ 3,8,4,4,3,5,9 };
    vector<int> v3{ 5, 6, 4, 3, 2, 6, 7, 9, 3 };
    //默认升序
    cout << "是否按升序排序：" << is_sorted(ivec.begin(), ivec.end()) << endl;
    cout <<"输出最大的已排序子范围的第一个元素："<< *is_sorted_until(v.begin(), v.end()) << endl;
    cout << "将容器按降序排序：";
    sort(v.begin(), v.end(), [](int x, int y) {return x > y; });
    for (int& i : v)
        cout << i << " ";
    cout << endl;
    cout << "将容器前五个元素按降序排序，后面的元素按默认升序排序：";
    partial_sort(ivec.begin(), ivec.begin() + 5, ivec.end(), [](int x, int y) {return x > y; });
    for (int& i : ivec)
        cout << i << " ";
    cout << endl;
    cout << "将第一个区间按升序排序，然后复制给后一个区间：";
    partial_sort_copy(ivec.begin(), ivec.end(), v.begin(), v.end());
    for (int& i : v)
        cout << i << " ";
    cout << endl;
    cout << "将容器按降序排序，同时保持相等元素之间的顺序：";
    stable_sort(v2.begin(), v2.end());
    for (int& i : v2)
        cout << i << " ";
    cout << endl;

    nth_element(v3.begin(), v3.begin()+3, v3.end());
    for (int& i : v3)
        cout << i << " ";
    cout << endl;
  }
```


## 5.二分搜索操作(在已排序范围上) 

lower_bound：返回指向第一个不小于给定值的元素的迭代器
upper_bound：返回指向第一个大于给定值的元素的迭代器
binary_search：判断一个元素是否在区间内
equal_range：返回匹配特定键值的元素区间

```c++
#include <iostream>
#include <algorithm>
#include <vector>
using namespace std;
int main()
{
    vector<int> ivec{ 3,2,2,4,5,6,7,8,9,10 };
    sort(ivec.begin(), ivec.end());
    cout << "返回指向第一个不小于给定值的元素的迭代器：";
    cout<<*lower_bound(ivec.begin(), ivec.end(), 5)<<endl;
    cout << "返回指向第一个大于给定值的元素的迭代器：";
    cout << *upper_bound(ivec.begin(), ivec.end(), 5) << endl;
    cout << "判断一个给定值是否在区间内：";
    cout << binary_search(ivec.begin(), ivec.end(), 5) << endl;
    cout << "返回pair对象，first指向第一个不小于给定值的元素，second指向第一个大于给定值的元素，first相当于lower_bound()，second相当于upper_bound：";
    cout << *equal_range(ivec.begin(), ivec.end(), 5).first << "  ";
    cout << *equal_range(ivec.begin(), ivec.end(), 5).second << endl;
}
```



## 6.集合操作(在已排序范围上) 

merge：合并两个已排序的区间
inplace_merge：就地合并两个有序的区间
includes：如果一个集合是另外一个集合的子集则返回true
set_difference：计算两个集合的差集
set_intersection：计算两个集合的交集
set_symmetric_difference：计算两个集合的对称差
set_union：计算两个集合的并集
7.堆操作 ^
is_heap：检查给定的区间是否为一个堆
is_heap_until：查找区间中为堆的最大子区间
make_heap：根据区间内的元素创建出一个堆
push_heap：将元素加入到堆
pop_heap：将堆中的最大元素删除
sort_heap：将堆变成一个排好序的区间
#include <iostream>
#include <algorithm>
#include <vector>
using namespace std;
int main()
{
    vector<int> ivec{ 3,10,2,4,40,6,7,8,11,12 };
    vector<int> v{ 1,5,6,7,4,9 };
    vector<int> v2{ 3,8,4,4,3,5,9 };
    vector<int> v3{ 5, 6, 4, 3, 2, 6, 7, 9, 3 };

    make_heap(ivec.begin(), ivec.end());
    for (int& i : ivec)
        cout << i << ' ';
    cout << endl;
    
    //将堆中的最大元素删除，但元素还在容器中
    pop_heap(ivec.begin(), ivec.end()); 
    //将存在于容器中，不属于堆的元素添加到堆中去
    push_heap(ivec.begin(), ivec.end());
    cout << "检查给定的区间是否为一个堆："<<is_heap(ivec.begin(), ivec.end()) << endl;
    //返回堆中最后一个元素的后一个位置的迭代器
    auto heap_end = is_heap_until(ivec.begin(), ivec.end());
    
    for (auto i = ivec.begin(); i != heap_end; ++i)
        cout << *i << ' ';
    cout << endl;
    //将堆进行排序
    sort_heap(ivec.begin(), ivec.end());
    for (int& i : ivec)
        cout << i << ' ';
    cout << endl;

}
1
2
3
4
5
6
7
8
9
10
11
12
13
14
15
16
17
18
19
20
21
22
23
24
25
26
27
28
29
30
31
32
33
34
8.最小/最大操作 ^
max：返回两个元素中的较大者
max_element：返回区间内的最大元素
min：返回两个元素中的较小者
min_element：返回区间内的最小元素
minmax：返回两个元素中的的较大者和较小者
minmax_element：返回区间内的最小元素和最大元素
#include <iostream>
#include <algorithm>
#include <vector>
using namespace std;
int main()
{
    vector<int> ivec{ 3,10,2,4,20,6,7,8,11,12 };
    cout << "返回两个元素中的较大者：" << *max(ivec.begin(), ivec.end() - 1) << endl;
    cout << "返回区间内的最大元素：" << *max_element(ivec.begin(), ivec.end()) << endl;
    cout << "返回两个元素中的较小者：" << *min(ivec.begin(), ivec.end() - 1) << endl;
    cout << "返回区间内的最小元素：" << *min_element(ivec.begin(), ivec.end()) << endl;
    cout << "返回两个元素中的的较小者：" << *minmax(ivec.begin(), ivec.end() - 1).first << endl;
    cout << "返回两个元素中的的较大者：" << *minmax(ivec.begin(), ivec.end() - 1).second << endl;
    cout << "返回区间内的最小元素：" << *minmax_element(ivec.begin(), ivec.end()).first << endl;
    cout << "返回区间内的最大元素：" << *minmax_element(ivec.begin(), ivec.end()).second << endl;
}
1
2
3
4
5
6
7
8
9
10
11
12
13
14
15
16
9.比较操作 ^
equal：确定两个元素集合是否是相同的
lexicographical_compare：如果按字典顺序一个区间小于另一个区间，返回true
#include <iostream>
#include <algorithm>
#include <vector>
using namespace std;
int main()
{
    vector<int> ivec{ 3,2,2,4,5,6,7,8,9,10 };
    vector<int> v{ 1,5,6,7,4,9 };
    cout << "判断两个集合是否相等：" << equal(ivec.begin(), ivec.end(), v.begin(), v.end());
    cout << endl << "判断第一个区间按字典顺序是否小于另一个区间：" << lexicographical_compare(ivec.begin(), ivec.end(), v.begin(), v.end());
}
1
2
3
4
5
6
7
8
9
10
11
10.排列操作 ^
is_permutation：判断一个序列是否为另一个序列的排列组合
next_permutation：按字典顺序产生区间内元素下一个较大的排列组合
prev_permutation：按字典顺序产生区间内元素下一个较小的排列组合
#include <iostream>
#include <algorithm>
#include <vector>
using namespace std;
int main()
{
    vector<int> ivec{ 3,10,2,4,20,6,7,8,11,10 };
    cout<<"判断一个序列是否为另一个序列的排列组合："<<is_permutation(ivec.begin(), ivec.end(), v.begin(),v.end())<<endl;
    cout << "按字典顺序产生区间内元素下一个较大的排列组合：" << next_permutation(ivec.begin(), ivec.end()) << endl;
    for (int& i : ivec)
        cout << i << " ";
    cout << endl;
    cout << "按字典顺序产生区间内元素下一个较小的排列组合：" << prev_permutation(ivec.begin(), ivec.end()) << endl;
    for (int& i : ivec)
        cout << i << " ";
    cout << endl;
}
1
2
3
4
5
6
7
8
9
10
11
12
13
14
15
16
17
11.数值运算 ^
iota：用从起始值开始连续递增的值填充区间
accumulate：计算区间内元素的和
inner_product：计算两个区间元素的内积
adjacent_difference：计算区间内相邻元素之间的差
partial_sum：计算区间内元素的部分和
————————————————





![]()