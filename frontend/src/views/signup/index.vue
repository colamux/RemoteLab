<template>
  <div class="signup">
    <p>注册</p>
    <!-- <el-button @click="go2login">登录</el-button> -->
    <div class="signup-main">
      <el-form ref="form" :model="form" label-width="80px">
        <el-form-item label="学号">
          <el-input v-model="form.username"></el-input>
        </el-form-item>

        <el-form-item label="电话">
            <el-input v-model="form.telephone"></el-input>
        </el-form-item>

        <el-form-item label="邮箱">
            <el-input v-model="form.email"></el-input>
        </el-form-item>

        <el-form-item label="密码">
            <el-input v-model="form.password"></el-input>
        </el-form-item>

        <el-form-item>
          <el-button type="primary" @click="signup">注册</el-button>
          <el-button @click="go2login">取消</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      form: {
        name: "",
        email:"xx@abc.com",
        telephone: "18712341234",
        passwd: "",
      },
    };
  },
  methods: {
    signup() {
        let {name,email,telephone, passwd} = this.form
    this.$http({
        method: "post",
        url: "/signup",
        data: {
          username: name,
          password: this.$encrypt(passwd),
          telephone: telephone,
          email: email
        },
        headers: {
          //   'Content-Type':'application/json',
        },
      }).then((res) => {
        if (res.status == 200) {
          // sessionStorage.setItem("username", name);
          alert("注册成功！")
          
        }
        // if (res.status == )
      });
    },
    go2login() {
        this.$router.push('/login')
    }
  },
};
</script>