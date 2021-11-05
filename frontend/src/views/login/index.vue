<template>
  <div class="login">
    <p class="login-title">嵌入式远程实验室平台</p>
    <div class="login-main">
      <el-form
        :model="ruleForm"
        :rules="rules"
        ref="ruleForm"
        label-width="100px"
        class="demo-ruleForm"
      >
        <el-form-item label="username" prop="name">
          <el-input v-model="ruleForm.name"></el-input>
        </el-form-item>

        <el-form-item label="password" prop="passwd">
          <el-input type="password" v-model="ruleForm.passwd"></el-input>
        </el-form-item>

        <el-form-item>
          <el-button type="primary" @click="login('ruleForm')">Login</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script>
export default {
  name: "login",
  data() {
    return {
      ruleForm: {
        name: "cola",
        passwd: "passwd",
      },
      rules: {
        name: [
          { required: true, message: "please enter name", trigger: "blur" },
          { min: 3, max: 20, message: "length error", trigger: "blur" },
        ],
        passwd: [
          { required: true, message: "please enter name", trigger: "blur" },
          { min: 3, max: 20, message: "length error", trigger: "blur" },
        ],
      },
    };
  },

  methods: {
    login() {
      let { name, passwd } = this.ruleForm;
      this.$http({
        method: "post",
        url: "/admin/add",
        data: {
          username: name,
          password: passwd,
        },
        headers: {
          //   'Content-Type':'application/json',
        },
        // data: JSON.stringify({
        //   username: name,
        //   password: passwd
        // }),
      }).then((res) => {
        if (res.status == 200) {
          // sessionStorage.setItem("username", name);
          localStorage.setItem("Authorization", res.data['data']['Authorization'])
          this.$router.push("/home");
        }
      });
    },
  },
};
</script>
