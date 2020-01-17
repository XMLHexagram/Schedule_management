<template>
  <div class="my_body">
    <form @submit="auth()">
      <div class="input-group">
        <input
          v-model="username"
          type="text"
          class="form-control"
          placeholder="用户名"
          aria-describedby="basic-addon1"
          required
        />
      </div>
      <div class="input-group">
        <input
          v-model="password"
          type="password"
          class="form-control"
          placeholder="密码"
          aria-describedby="basic-addon1"
          required
        />
      </div>
      <div class="btn-group">
        <button type="submit" class="btn btn-primary" aria-haspopup="true" aria-expanded="false">
          戳我登录
          <span class="caret"></span>
        </button>
      </div>
    </form>
  </div>
</template>

<script>
import axios from 'axios'
//import { baseURL } from "../main";

export default {
  name: 'login',
  data () {
    return {
        username: "",
        password: "",
    //   info: {}
    }
  },
  async created(){
      const token = localStorage.getItem("token");
      if (token != null) {
          this.$router.push("/home")
      }
  },
  methods: {
    async auth () {
      // localStorage 使用文档：https://developer.mozilla.org/zh-CN/docs/Web/API/Window/localStorage
      // 用 cookie 也可以
      const token = localStorage.getItem('token')
      if (token !== null) {
        // // 如果本地存在 token，那么就尝试从 https://api.hduhelp.com/student/info 获取个人信息
        // const studentInfoRes = await axios.get('/base/student/info', {
        //   // 固定写法，给 http 请求头添加 Authorization: token xxx，此时的 token 就代表着用户的用户名密码
        //   headers: {
        //     Authorization: `token ${token}`
        //   }
        // })
        // this.info = studentInfoRes.data.data
        this.$router.push("home");
      } else {
        // 如果本地没有 token，则直接跳转到数字杭电的授权地址
        // 文档参见：https://git.hduhelp.com/Lemon/api_gateway_new（GET /login/url/{GrantType} 获取登录链接）
        const casUrlRes = (await axios.get('/login/url/cas', {
          // 两个传入参数：
          params: {
            // clientID: 由后台部提供，测试可以填写 oa
            clientID: 'oa',
            //  redirect：获取 token 的地址，这里填写的 ${window.location.origin}/#/auth 是实际访问到 Auth.vue 的地址
            redirect: `${window.location.origin}/#/auth`
          }
        })).data
        const redirectURL = casUrlRes.data.url
        window.location.href = redirectURL
      }
    }
  },
//   async created () {
//     // 判断是否有登录权限
//     await this.auth()
//   }
}
</script>

<style>
.login {
    display: flex;
    margin: 0px;
    align-content: center;
}
</style>