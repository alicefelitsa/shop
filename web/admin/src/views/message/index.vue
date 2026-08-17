<template>
  <div class="content">
    <el-card shadow="always">
      <div slot="header">

        <!--搜索栏-->
        <div class="queryForm">
          <el-form :inline="true" :model="where" class="query-form-inline" size="small">
            <el-form-item label="姓名">
              <el-input v-model="where.name" placeholder="请输入" clearable class="queryElInput"></el-input>
            </el-form-item>
            <el-form-item label="邮箱">
              <el-input v-model="where.email" placeholder="请输入" clearable class="queryElInput"></el-input>
            </el-form-item>
            <el-form-item label="主题">
              <el-select v-model="where.subject" placeholder="请选择" clearable class="queryElInput">
                <el-option label="General Inquiry" value="general"></el-option>
                <el-option label="Order Inquiry" value="order"></el-option>
                <el-option label="Wholesale / Bulk Order" value="wholesale"></el-option>
                <el-option label="Technical Support" value="support"></el-option>
                <el-option label="Partnership" value="partnership"></el-option>
                <el-option label="Other" value="other"></el-option>
              </el-select>
            </el-form-item>
            <el-form-item>
              <el-button type="primary" @click="search">查询</el-button>
              <el-button icon="el-icon-refresh" @click="reset">重置</el-button>
            </el-form-item>
          </el-form>
        </div>
      </div>

      <!--工具栏-->
      <div class="toolbar">
        <el-button type="danger" size="small" icon="el-icon-delete" @click="del">删除</el-button>
        <el-button size="small" icon="el-icon-download" @click="handleExport">导出</el-button>
      </div>

      <!--数据表格-->
      <el-table class="tableData" :data="tableData" :highlight-selection-row="true" height="calc(100vh - 182px)"
                :border="true"
                v-loading="visitorLoading" @selection-change="handleSelectionChange">
        <el-table-column type="selection" width="55" align="center"></el-table-column>
        <el-table-column prop="id" label="ID" width="60px">
          <template v-slot="{row}">
            {{ row.id }}
          </template>
        </el-table-column>
        <el-table-column prop="name" label="姓名" min-width="100px">
          <template v-slot="{row}">
            {{ row.name }}
          </template>
        </el-table-column>
        <el-table-column prop="email" label="邮箱" min-width="160px">
          <template v-slot="{row}">
            {{ row.email }}
          </template>
        </el-table-column>
        <el-table-column prop="subject" label="主题" min-width="120px">
          <template v-slot="{row}">
            {{ subjectLabel(row.subject) }}
          </template>
        </el-table-column>
        <el-table-column prop="content" label="留言内容" min-width="220px" show-overflow-tooltip>
          <template v-slot="{row}">
            {{ row.content }}
          </template>
        </el-table-column>
        <el-table-column prop="ip" label="IP" min-width="110px">
          <template v-slot="{row}">
            {{ row.ip }}
          </template>
        </el-table-column>
        <el-table-column prop="ip_address" label="IP地址" min-width="140px" show-overflow-tooltip>
          <template v-slot="{row}">
            {{ row.ip_address }}
          </template>
        </el-table-column>
        <el-table-column prop="ctime" label="时间" min-width="130px">
          <template v-slot="{row}">
            {{ row.ctime }}
          </template>
        </el-table-column>
        <el-table-column label="操作" align="center" width="100px">
          <template v-slot="{row}">
            <el-button size="mini" @click="view(row)">查看</el-button>
          </template>
        </el-table-column>
      </el-table>

      <!--表格分页-->
      <div style="margin-top: 10px; text-align: center;" class="currentPage">
        <el-pagination
            @size-change="handleSizeChange"
            @current-change="handleCurrentChange"
            :current-page="where.page"
            :page-sizes="pageSizes"
            :page-size="where.limit"
            background
            layout="total, sizes, prev, pager, next, jumper"
            :total="totalData">
        </el-pagination>
      </div>

    </el-card>

    <!--查看留言详情-->
    <el-dialog title="留言详情" :visible.sync="dialogVisible" width="600px">
      <el-descriptions v-if="viewData.id" :column="2" border size="small">
        <el-descriptions-item label="姓名">{{ viewData.name }}</el-descriptions-item>
        <el-descriptions-item label="邮箱">{{ viewData.email }}</el-descriptions-item>
        <el-descriptions-item label="主题">{{ subjectLabel(viewData.subject) }}</el-descriptions-item>
        <el-descriptions-item label="时间">{{ viewData.ctime }}</el-descriptions-item>
        <el-descriptions-item label="IP">{{ viewData.ip }}</el-descriptions-item>
        <el-descriptions-item label="IP地址">{{ viewData.ip_address }}</el-descriptions-item>
        <el-descriptions-item label="留言内容" :span="2">
          <div style="white-space: pre-wrap;">{{ viewData.content }}</div>
        </el-descriptions-item>
      </el-descriptions>
      <div slot="footer">
        <el-button size="small" @click="dialogVisible=false">关闭</el-button>
      </div>
    </el-dialog>

  </div>
</template>

<script>
import {DelMessage, GetMessage} from "@/api/message";
import {exportToExcel} from "@/utils/excel";

export default {
  name: "message",
  data() {
    return {
      tableData: [],
      totalData: 0,
      pageSizes: [30, 50, 100, 200],
      multipleSelection: [],
      where: {
        name: '',
        email: '',
        subject: '',
        page: 1,
        limit: 30,
      },
      visitorLoading: false,
      dialogVisible: false,
      viewData: {},
      // 主题编码与文案对应，与 h5 联系页下拉选项一致
      subjectMap: {
        'general': 'General Inquiry',
        'order': 'Order Inquiry',
        'wholesale': 'Wholesale / Bulk Order',
        'support': 'Technical Support',
        'partnership': 'Partnership',
        'other': 'Other'
      }
    }
  },
  mounted() {
    this.where.limit = this.pageSizes[0]
    this.getMessage()
  },
  methods: {
    //获取客户留言
    async getMessage() {
      this.visitorLoading = true;
      setTimeout(async () => {
        try {
          let data = await GetMessage({...this.where})
          this.tableData = data.data;
          this.totalData = data.count
        } catch (e) {
          this.$message.error(e.message);
        } finally {
          this.visitorLoading = false;
        }
      }, 200)
    },
    //主题编码转文案
    subjectLabel(subject) {
      return this.subjectMap[subject] || subject
    },
    //查询
    search() {
      this.where.page = 1
      this.getMessage()
    },
    //重置搜索条件
    reset() {
      this.where.name = ''
      this.where.email = ''
      this.where.subject = ''
      this.where.page = 1
      this.getMessage()
    },
    //页数
    handleSizeChange(val) {
      this.where.limit = val
      this.getMessage()
    },
    //页码
    handleCurrentChange(val) {
      this.where.page = val
      this.getMessage()
    },
    //多选
    handleSelectionChange(val) {
      this.multipleSelection = []
      val.forEach(item => {
        this.multipleSelection.push(item.id);
      })
    },
    //查看留言详情
    view(row) {
      this.viewData = {...row}
      this.dialogVisible = true;
    },
    //删除
    async del() {
      if (this.multipleSelection.length === 0) {
        this.$message.warning("请选择要删除的数据");
        return
      }
      this.$confirm('即将删除，是否继续?').then(async _ => {
        let ids = this.multipleSelection.join(',')
        try {
          let message = await DelMessage(ids)
          this.$message.success(message)
          await this.getMessage()
        } catch (e) {
          this.$message.error(e.message);
        }
      }).catch(_ => {
      });
    },
    //导出表格
    async handleExport() {
      const headers = ['姓名', '邮箱', '主题', '留言内容', 'IP', 'IP地址', '时间']
      // 将 tableData 转换为导出格式
      const rows = this.tableData.map(item => [
        item.name || '',
        item.email || '',
        this.subjectLabel(item.subject),
        item.content || '',
        item.ip || '',
        item.ip_address || '',
        item.ctime || ''
      ])
      await exportToExcel(headers, rows, '客户留言表')
    }
  }
}
</script>
