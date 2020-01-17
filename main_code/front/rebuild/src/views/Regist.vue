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
import {ApiInstance} from "../instances/index"
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
      try {
            const {data: res} = await ApiInstance.post(baseURL + '/auth/register',
                {username: this.username, password: this.password, code: this.code}
                )
                const token = res.data
                //console.log(token)
                localStorage.setItem("token",token)
                alert("注册成功")
                this.$router.push('/')
        } catch (e) {
        }
    }
  },
};
</script>