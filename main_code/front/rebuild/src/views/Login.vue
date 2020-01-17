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
        </button>
        <p>没有账号？</p>
        <button @click="goToRegist" class="btn btn-primary" aria-haspopup="true" aria-expanded="false">
          戳我注册
        </button>
      </div>
    </form>
  </div>
</template>

<script>
import { ApiInstance } from "../instances/index";
import { baseURL } from "../main";

export default {
  name: "login",
  data() {
    return {
      username: "",
      password: ""
    };
  },
  async created() {
    const token = localStorage.getItem("token");
    if (token != null) {
      this.$router.push("/home");
    }
  },
  methods: {
    async auth() {
      // localStorage 使用文档：https://developer.mozilla.org/zh-CN/docs/Web/API/Window/localStorage
      const token = localStorage.getItem("token");
      if (token != null) {
        this.$router.push("/home");
      } else {
        try {
          const { data: res } = await ApiInstance.post(
            baseURL + "/auth/login",
            { username: this.username, password: this.password }
          );
          const token1 = res.data;
          localStorage.setItem("token", token1);
          console.log(token1)
          this.$router.push("/home");
        } catch (e) {}
      }
    },
    goToRegist() {
      this.$router.push("/regist")
    }
  }
};
</script>

<style>
.login {
  display: flex;
  margin: 0px;
  align-content: center;
}
</style>