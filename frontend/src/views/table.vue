<template>
  <el-table :data="tableData" style="width: 100%" max-height="500">
    <el-table-column fixed prop="name" label="实验名称" width="450">
    </el-table-column>
    <el-table-column prop="score" label="分数" width="120"> </el-table-column>
    <el-table-column prop="data" label="时间" width="120">
    </el-table-column>
  </el-table>
</template>

<script>
export default {
  data() {
    return {
      tableData: [
        {
          date: "2016-05-07",
          name: "点亮LED灯",
          province: "上海",
          city: "普陀区",
          address: "上海市普陀区金沙江路 1518 弄",
          zip: 200333,
        },
        {
          date: "2016-05-07",
          name: "王小虎",
          province: "上海",
          city: "普陀区",
          address: "上海市普陀区金沙江路 1518 弄",
          zip: 200333,
        },
        {
          date: "2016-05-07",
          name: "王小虎",
          province: "上海",
          city: "普陀区",
          address: "上海市普陀区金沙江路 1518 弄",
          zip: 200333,
        },
      ],
    };
  },
};
</script>