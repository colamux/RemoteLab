<template>
  <div>
    <el-menu-item v-if="!item.children" :index="item.url">
      <i :class="item.icon"></i>
      <span slot="title">{{ item.name }}</span>
    </el-menu-item>

    <el-submenu v-else index="item.url">
      <template slot="title">
        <i class="el-icon-menu"></i>
        <span>{{ item.name }}</span>
      </template>
      <NavItem v-for="(item, i) in item.children" :key="i" :item="item" />
    </el-submenu>
  </div>
</template>

<script>
export default {
  name: "NavItem",
  props: ["item"],
};
</script>