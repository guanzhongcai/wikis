# HTML

Hyper Text Markup Language 超文本标记语言，超文本表示它包括：文本、图片、音视频、动画等

浏览器的F12里，看到的左侧 `Elements` 表示html元素，右侧`Styles`表示css样式

html5和css3是未来的互联网方向！





## W3C

World Wide Web Consortium 万维网联盟，W3C标准包括：

- **结构**化标准语言：HTML、XML
- **表现**标准语言：CSS
- **行为**标准：DOM、ECMAScript



## HTML标签

```html
<!-- DOCTYPE：告诉浏览器，我们要使用什么规范 -->
<!DOCTYPE html>
<html lang="en">
  <!-- head标签代表网页头部 -->
  <head>
    <!-- meta描述性标签，用来描述网站的一些信息，譬如关键词keywords等 -->
    <meta charset="UTF-8">
    <meta name="keywords" content="编程,html,前端">
    <meta name="description" content="欢迎来学html的前端编程开发">
    <title>网页标题！</title>
  </head>
  <!-- body标签代表网页主体 -->
  <body>
    
  </body>
</html>
```



## 行内元素和块元素

块元素：无论内容多少，该元素独占一行，例如：p h1-h6

行原色：内容撑开长度，左右都是行内元素的可以，a strong em ...



## 页面结构分析

| 元素名  | 描述                                             |
| ------- | ------------------------------------------------ |
| header  | 标记头部区域的内容，用于页面或页面中的一块区域   |
| footer  | 标记脚部区域的内容，用于整个页面或页面的一块区域 |
| section | web页面的一块独立区域                            |
| article | 独立的文章内容                                   |
| aside   | 相关内容或应用，常用于侧边栏                     |
| nav     | 导航类辅助内容                                   |



## 表单的元素格式

| 属性      | 说明                                                         |
| --------- | ------------------------------------------------------------ |
| type      | 指定元素的类型。text、password，CheckBox、radio、submit、reset、file、hidden、image、buttone，默认为text |
| name      | 指定表单元素的名称                                           |
| value     | 元素的初始值。type为radio时必须指定一个值                    |
| size      | 指定表单元素的初始宽度。type为text、password是，size表示字符。其他类型，是像素单位 |
| maxlength | 输入的最大字符数                                             |
| checked   | 表示是否选中                                                 |



## 表单的应用

- 隐藏域 hidden
- 只读 readonly
- 禁用  disabled



## 表单初级验证

- placeholder：提示信息
- required：必须要填写的字段，否则提交不了！
- pattern：正则表达式，百度搜索“常用正则表达式”，[常用正则表达式大全](https://blog.csdn.net/qianhaohong/article/details/53435253)



## iframe内联框架

```html
<iframe src="//player.bilibili.com/player.html?aid=55631961&bvid=BV1x4411V75C&cid=97257967&page=11" scrolling="no" border="0" frameborder="no" framespacing="0" allowfullscreen="true"> 
</iframe>
```

使用a+iframe实现在本网页内加载网页

```html
<iframe src="" name="hello" frameborder="0" width="1000px" height="800px">
</iframe>
<a href="https://blog.kuangstudy.com" target="hello">点击跳转</a>
```

