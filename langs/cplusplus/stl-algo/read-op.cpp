#include <iostream>
#include <algorithm>
#include <vector>
using namespace std;

int main()
{
    vector<int> ivec{1, 2, 2, 4, 5, 6, 7, 8, 9, 10};
    vector<int> v{5, 6, 7, 8, 9, 10};
    //判断容器内是否所有元素都满足条件
    if (all_of(ivec.begin(), ivec.end(), [](int i)
               { return i < 20; }))
        cout << "容器内所有元素都满足条件" << endl;
    //判断容器内是否有元素满足条件
    if (any_of(ivec.begin(), ivec.end(), [](int i)
               { return i > 9; }))
        cout << "容器内有任意一个元素满足条件" << endl;
    //判断容器内是否所有的元素都不满足条件
    if (none_of(ivec.begin(), ivec.end(), [](int i)
                { return i > 10; }))
        cout << "容器内所有元素都不满足条件" << endl;
    //遍历容器内元素
    for_each(ivec.begin(), ivec.end(), [](const int &n)
             { cout << n << " "; });
    cout << endl
         << "返回满足条件的元素个数：" << endl;
    cout << count(ivec.begin(), ivec.end(), 5) << endl;
    cout << count_if(ivec.begin(), ivec.end(), [](int i)
                     { return i < 4; })
         << endl;
    cout << "返回一个pair对象，first为第一个不匹配元素的迭代器，second为第一个匹配的元素的迭代器：" << endl;
    cout << *mismatch(ivec.begin(), ivec.end(), v.begin(), v.end()).first << endl;
    cout << *mismatch(ivec.begin(), ivec.end(), v.begin(), v.end()).second << endl;
    cout << "返回满足条件的第一个元素的迭代器：" << endl;
    cout << *find(ivec.begin(), ivec.end(), 5) << endl;
    cout << *find_if(ivec.begin(), ivec.end(), [](int i)
                     { return i > 5; })
         << endl;
    cout << "返回不满足条件的第一个元素的迭代器" << endl;
    cout << *find_if_not(ivec.begin(), ivec.end(), [](int i)
                         { return i < 5; })
         << endl;
    cout << "如果后一个区间是前一个区间的子区间，则返回后一个区间的起始迭代器，否则返回第一个区间的末尾迭代器：" << endl;
    cout << *find_end(ivec.begin(), ivec.end(), v.begin(), v.end()) << endl;
    cout << "返回两个区间首个相等元素的迭代器，如果没有则返回第一个区间的末端迭代器：" << endl;
    cout << *find_first_of(ivec.begin(), ivec.end(), v.begin(), v.end()) << endl;
    cout << "返回首个相邻元素相等的第一个元素的迭代器，如果没有则返回末端迭代器：" << endl;
    cout << *adjacent_find(ivec.begin(), ivec.end()) << endl;
    cout << "返回第一个相等的元素的迭代器，如果没有则返回第一个区间的末端迭代器：" << endl;
    cout << *search(ivec.begin(), ivec.end(), v.begin(), v.end()) << endl;
    cout << "返回满足条件(值为2，出现次数为2)的第一个元素的迭代器，如果没有则返回末端迭代器：" << endl;
    cout << *search_n(ivec.begin(), ivec.end(), 2, 2) << endl;
}

// g++ read-op.cpp -std=c++14