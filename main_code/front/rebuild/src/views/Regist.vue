<template>
  <div class="my_body">
    <form @submit="regist()">
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
      <div class="input-group">
        <input
          v-model="code"
          type="text"
          class="form-control"
          placeholder="邀请码"
          aria-describedby="basic-addon1"
          required
        />
      </div>
      <div class="btn-group">
        <button type="submit" class="btn btn-success" aria-haspopup="true" aria-expanded="false">
          戳我注册
          <span class="caret"></span>
        </button>
      </div>
    </form>
  </div>
</template>

<script>
import axios from "axios";
import { baseURL } from "../main";
export default {
  name: "regist",
  data() {
    return {
      username: "",
      password: "",
      code: "",
      token: ""
    };
  },
  methods: {
    async regist() {
      await axios
        .post(baseURL + "/auth/register", {
          username: this.username,
          password: this.password,
          code: this.code
        })
        .then(res => {
          this.token = res.data.data;
          localStorage.setItem("token", this.token);
          alert("注册成功！");
          this.$router.push("/home");
        })
        .catch(err => {
          this.token = err;
          console.log(1);
          console.log(this.token);
          alert("注册的时候好像有未知错误嘤？");
        });
    }
  },
};
</script>