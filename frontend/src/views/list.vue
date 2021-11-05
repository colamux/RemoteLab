<template>
  <div>
    <el-button type="primary" @click="test()"> test</el-button>
  </div>
</template>

<script>
export default {
  methods: {
    test() {
      this.$http({
        method: "get",
        url: "/admin/user",
        headers: {},
        data: {},
      }).then((res) => {
        console.log(res.status);
      });
    },
  },
};
</script>
