<template>
  <div class="breadcrumb">
    <el-breadcrumb separator="/">
      <el-breadcrumb-item v-for="(item, i) in items" :item="item" :key="i">
        <span>{{ item.name }}</span>
      </el-breadcrumb-item>
    </el-breadcrumb>
  </div>
</template>

<script>
export default {
  data() {
    return {
      items: []
    }
  },
  created() {
    this.items = this.$route.matched
  },
  watch: {
    $route(to, from) {
      this.items = to.matched
    }
  }
}
</script>
