<template>
  <div style="width:100%;">
    <el-main>
      <el-card class="box-card" shadow="always">
        <div slot="header">
          <span>访客统计数据<i style="color: #f56c6c">（相同IP算一个）</i></span>
        </div>
        <el-table :data="tableData" style="width: 100%; height: 383px" :border="true" v-loading="visitorLoading">
          <el-table-column prop="date" label="日期" align="center"></el-table-column>
          <el-table-column prop="num" label="访问数量" align="center"></el-table-column>
        </el-table>
      </el-card>
    </el-main>
  </div>
</template>

<script>

import {getStatistic} from "@/api/chat";

export default {
  data() {
    return {
      tableData: [],
      visitorLoading: false,
    }
  },
  mounted() {
    this.visitorLoading = true;
    getStatistic(localStorage.getItem("ucode")).then((data) => {
      //console.log(data);
      setTimeout(() => {
        this.visitorLoading = false;
        this.tableData = data;
      }, 200)
    }).catch((e) => {
      this.$message.error(e.message);
    });
  }
}
</script>


<style scoped>
.box-card {
  width: 98%;
  margin: 10px;
}

.el-main {
  padding: 0;
}
</style>