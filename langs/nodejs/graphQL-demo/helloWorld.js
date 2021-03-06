const express = require('express')
const { buildSchema } = require('graphql')
const graphqlHttp = require('express-graphql')

// 定义schema，查询和类型
const schema = buildSchema(`
    type Account {
        name: String
        age: Int
        sex: String
        department: String
    }
    type Query {
        hello: String
        accountName: String
        age: Int
        account: Account
    }
`)

// hello: String
// 其中的hello是方法名，String是返回值类型

// 定义查询对应的处理器resolver
const root = {
    hello: () => {
        return 'hello world'
    },
    accountName: () => {
        return '张三丰'
    },
    age: () => {
        return 18
    },
    account: () => {
        return {
            name: "jack",
            age: 22,
            sex: "male",
            department: "RD"
        }
    }
}

const app = express();

app.use('/graphql', graphqlHttp.graphqlHTTP({
    schema: schema,
    rootValue: root,
    graphiql: true
}))

app.listen(3000)
console.log(`please go to: http://localhost:3000/graphql`)

/**
 * 左侧输入：
 query {
  hello
  accountName
  account {
    name
    age
    sex
    department
  }
}
 */