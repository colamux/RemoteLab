<template>
  <el-menu router>
    <NavItem v-for="item in items" :key="item.url" :item="item" />
  </el-menu>
</template>

<script>
import NavItem from "./navItem";
const items = [
  {
    name: "首页",
    url: "/home",
    icon: "el-icon-document"
  },
  {
    name: "设备",
    url: "/equipment",
    children: [
      {
        name: "列表",
        url: "/equipment/list",
        icon: "el-icon-setting"
      },
      // {
      //   name: "列表",
      //   url: "/wms/list",
      // },
    ],
  },
];

export default {
  data() {
    return {
      items,
    };
  },
  components: {
    NavItem,
  },
  methods: {},
};
</script>