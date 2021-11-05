<template>
  <div class="module">
    <el-container>
      <!-- left -->
      <el-aside width="200px">
        <Nav />
      </el-aside>
      <el-container>
        <!-- header -->
        <el-header>
          <Breadcrumb />
          <el-button type="primary" class="exit" @click="goBack"
            >退出</el-button
          >
        </el-header>

        <!-- main -->
        <el-main>
          <router-view />
        </el-main>
      </el-container>
    </el-container>
  </div>
</template>

<script>
import Nav from "./nav";
import Breadcrumb from "./breadcrumb";
export default {
  data() {
    return {};
  },
  components: {
    Nav,
    Breadcrumb,
  },
  methods: {
    goBack() {
      localStorage.removeItem("Authorization");
      this.$router.push("/login");
    },
  },
};
</script>

