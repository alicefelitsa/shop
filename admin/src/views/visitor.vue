<template>
  <div style="width: 100%;">
    <el-main style="margin:0 5px;">
      <el-form :inline="true" style="margin: 10px 0 -20px 10px;">
        <el-form-item>
          <el-button type="danger" @click="del">删除访客</el-button>
        </el-form-item>
        <el-form-item style="margin-top: 2px">
          <el-input placeholder="请输入IP进行查询" v-model="search" clearable>
            <template slot="prepend">IP</template>
          </el-input>
        </el-form-item>
      </el-form>
      <el-table v-loading="visitorLoading"
                :data="tableData.filter(data => !search || data.ip.toLowerCase().includes(search.toLowerCase()))"
                height="calc(100vh - 65px)" style="width: 100%;" @selection-change="handleSelectionChange">
        <el-table-column type="selection" width="55" align="center"></el-table-column>
        <el-table-column prop="uuid" label="昵称" width="100"></el-table-column>
        <el-table-column prop="ip" label="IP" width="130"></el-table-column>
        <el-table-column prop="position" label="位置" width="230"></el-table-column>
        <el-table-column prop="ctime" label="注册时间" width="160"></el-table-column>
        <el-table-column prop="online_time" label="在线时间" width="160"></el-table-column>
        <el-table-column prop="offline_time" label="离线时间" width="160"></el-table-column>
        <el-table-column prop="scan_num" label="扫码次数" width="80"></el-table-column>
        <el-table-column prop="device" label="设备" :show-overflow-tooltip="true" min-width="200"></el-table-column>
      </el-table>
    </el-main>
  </div>
</template>


<script>
import {delClient, getClientList} from "@/api/chat";

export default {
  data() {
    return {
      tableData: [],
      multipleSelection: [],
      search: '',
      visitorLoading: false,
    }
  },
  mounted() {
    this.getVisitorList()
  },
  methods: {
    //获取访客列表
    getVisitorList() {
      this.visitorLoading = true;
      getClientList(localStorage.getItem("ucode")).then((data) => {
        //console.log(data.client);
        setTimeout(() => {
          this.visitorLoading = false;
          this.tableData = data.client;
        }, 200)
      }).catch((e) => {
        this.$message.error(e.message);
      });
    },
    handleSelectionChange(val) {
      this.multipleSelection = []
      val.forEach(item => {
        this.multipleSelection.push(item.id);
      })
    },
    //删除访客
    del() {
      if (this.multipleSelection.length > 0) {
        let ids = this.multipleSelection.join(',')
        //console.log(ids)
        delClient(localStorage.getItem("ucode"), ids).then((msg) => {
          this.$message.success(msg)
          this.getVisitorList()
        }).catch((e) => {
          this.$message.error(e.message);
        });
      } else {
        this.$message.warning("请选择要删除的访客");
      }
    }
  }
}
</script>
<style scoped>
.el-main {
  padding: 0;
}
</style>