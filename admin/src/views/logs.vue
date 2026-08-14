<template>
  <div style="width:100%;">
    <el-main>
      <el-card class="box-card" shadow="always">
        <div slot="header">
          <span>操作日志</span>
        </div>
        <el-table :data="tableData" style="width: 100%;" height="calc(100vh - 131px)" :border="true" v-loading="visitorLoading">
          <el-table-column prop="ctime" label="时间" width="160"></el-table-column>
          <el-table-column label="操作记录">
            <template v-slot="{ row }">
              <el-tag class="wrap-tag" type="info"> {{ row.content }}</el-tag>
            </template>
          </el-table-column>
          <el-table-column prop="ip" label="IP" width="140"></el-table-column>
          <el-table-column prop="position" label="位置" width="220"></el-table-column>
        </el-table>
      </el-card>
    </el-main>
  </div>
</template>

<script>

import {getLogs} from "@/api/logs";

export default {
  data() {
    return {
      tableData: [],
      visitorLoading: false,
    }
  },
  mounted() {
    this.visitorLoading = true;
    getLogs(localStorage.getItem("ucode")).then((data) => {
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

.wrap-tag {
  white-space: normal !important; /* 关键：允许换行 */
  word-break: break-word !important; /* 按单词换行 */
  height: auto !important; /* 自动高度 */
  line-height: 1.5 !important; /* 行高 */
  padding: 4px 8px !important; /* 内边距 */
  max-width: 100%; /* 最大宽度 */
  font-size: 14px;
}
</style>